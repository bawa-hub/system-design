package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// =====================
// Utilities
// =====================

func b64url(data []byte) string { return base64.RawURLEncoding.EncodeToString(data) }

func newID(n int) (string, error) {
	// cryptographically strong random ID (n bytes -> 2n hex-ish using base64url)
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return b64url(buf), nil
}

// =====================
// JWT / JWKS types
// =====================

type AccessClaims struct {
	jwt.RegisteredClaims
	Roles    []string `json:"roles,omitempty"`
	TokenUse string   `json:"token_use,omitempty"` // "access" or "refresh"
}

type jwkKey struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use,omitempty"`
	Alg string `json:"alg,omitempty"`
	N   string `json:"n,omitempty"`
	E   string `json:"e,omitempty"`
}

type jwks struct {
	Keys []jwkKey `json:"keys"`
}

// =====================
// In-memory stores (for demo only!)
// =====================

type refreshRecord struct {
	Username  string
	ExpiresAt time.Time
	Revoked   bool
}

type server struct {
	mu           sync.RWMutex
	issuer       string
	audience     string
	activeKID    string
	privKey      *rsa.PrivateKey
	pubKeys      map[string]*rsa.PublicKey // kid -> pub
	users        map[string]string         // username -> bcrypt hash
	refreshStore map[string]refreshRecord  // jti -> record
}

// =====================
// Key management
// =====================

func generateKey() (*rsa.PrivateKey, string, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, "", err
	}
	// kid = first 10 chars of sha256(modulus)
	h := sha256.Sum256(priv.PublicKey.N.Bytes())
	kid := b64url(h[:])[:10]
	return priv, kid, nil
}

func (s *server) rotateKeys() error {
	priv, kid, err := generateKey()
	if err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	// keep old pubkeys for verifying older tokens
	if s.pubKeys == nil {
		s.pubKeys = map[string]*rsa.PublicKey{}
	}
	if s.privKey != nil {
		s.pubKeys[s.activeKID] = &s.privKey.PublicKey
	}
	s.privKey = priv
	s.activeKID = kid
	return nil
}

func (s *server) currentPublicJWK() jwkKey {
	s.mu.RLock()
	defer s.mu.RUnlock()
	pub := &s.privKey.PublicKey
	// exponent usually 65537
	e := big.NewInt(int64(pub.E)).Bytes()
	return jwkKey{
		Kty: "RSA",
		Kid: s.activeKID,
		Use: "sig",
		Alg: "RS256",
		N:   b64url(pub.N.Bytes()),
		E:   b64url(e),
	}
}

func (s *server) allPublicJWKS() jwks {
	s.mu.RLock()
	defer s.mu.RUnlock()
	keys := []jwkKey{}
	// current
	{
		pub := &s.privKey.PublicKey
		e := big.NewInt(int64(pub.E)).Bytes()
		keys = append(keys, jwkKey{Kty: "RSA", Kid: s.activeKID, Use: "sig", Alg: "RS256", N: b64url(pub.N.Bytes()), E: b64url(e)})
	}
	// historical
	for kid, pub := range s.pubKeys {
		e := big.NewInt(int64(pub.E)).Bytes()
		keys = append(keys, jwkKey{Kty: "RSA", Kid: kid, Use: "sig", Alg: "RS256", N: b64url(pub.N.Bytes()), E: b64url(e)})
	}
	return jwks{Keys: keys}
}

// =====================
// JWT helpers
// =====================

func (s *server) signToken(username, tokenUse string, ttl time.Duration, extra map[string]any) (string, *AccessClaims, error) {
	if tokenUse != "access" && tokenUse != "refresh" {
		return "", nil, errors.New("tokenUse must be access or refresh")
	}
	jti, err := newID(16)
	if err != nil {
		return "", nil, err
	}
	now := time.Now()
	claims := &AccessClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   username,
			Audience:  jwt.ClaimStrings{s.audience},
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			NotBefore: jwt.NewNumericDate(now.Add(-30 * time.Second)), // small clock skew
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        jti,
		},
		Roles:    []string{"user"},
		TokenUse: tokenUse,
	}
	for k, v := range extra {
		switch k {
		case "roles":
			if r, ok := v.([]string); ok {
				claims.Roles = r
			}
		default:
			// attach custom claims into a map by using WithClaims below
		}
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tok.Header["kid"] = s.activeKID
	s.mu.RLock()
	priv := s.privKey
	s.mu.RUnlock()
	signed, err := tok.SignedString(priv)
	if err != nil {
		return "", nil, err
	}
	return signed, claims, nil
}

func (s *server) verifyToken(tokenString, expectedUse string) (*AccessClaims, error) {
	parser := jwt.NewParser(jwt.WithValidMethods([]string{"RS256"}))
	claims := &AccessClaims{}
	tok, err := parser.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		kid, _ := t.Header["kid"].(string)
		if kid == "" {
			return nil, errors.New("missing kid")
		}
		s.mu.RLock()
		defer s.mu.RUnlock()
		if kid == s.activeKID {
			return &s.privKey.PublicKey, nil
		}
		if k, ok := s.pubKeys[kid]; ok {
			return k, nil
		}
		return nil, fmt.Errorf("unknown kid: %s", kid)
	})
	if err != nil || !tok.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	// Validate iss, aud, exp, nbf handled by jwt library's Valid()
	if claims.Issuer != s.issuer {
		return nil, errors.New("bad issuer")
	}
if !containsAudience(claims.Audience, s.audience) {
    return nil, errors.New("bad audience")
}
	if expectedUse != "" && claims.TokenUse != expectedUse {
		return nil, errors.New("wrong token use")
	}
	return claims, nil
}

// =====================
// HTTP helpers / middleware
// =====================

func jsonWrite(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (s *server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			jsonWrite(w, http.StatusUnauthorized, map[string]string{"error": "missing bearer token"})
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		claims, err := s.verifyToken(token, "access")
		if err != nil {
			jsonWrite(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
			return
		}
		// attach username to context via header for demo simplicity
		w.Header().Set("X-User", claims.Subject)
		next.ServeHTTP(w, r)
	})
}

// =====================
// Handlers
// =====================

func (s *server) handleJWKS(w http.ResponseWriter, r *http.Request) {
	jsonWrite(w, http.StatusOK, s.allPublicJWKS())
}

func (s *server) handleLogin(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		jsonWrite(w, http.StatusBadRequest, map[string]string{"error": "bad json"})
		return
	}
	// validate user
	s.mu.RLock()
	hash, ok := s.users[in.Username]
	s.mu.RUnlock()
	if !ok || bcrypt.CompareHashAndPassword([]byte(hash), []byte(in.Password)) != nil {
		jsonWrite(w, http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
		return
	}
	accessTTL := 5 * time.Minute
	refreshTTL := 24 * time.Hour * 7

	accessToken, _, err := s.signToken(in.Username, "access", accessTTL, nil)
	if err != nil {
		jsonWrite(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	refreshToken, refreshClaims, err := s.signToken(in.Username, "refresh", refreshTTL, nil)
	if err != nil {
		jsonWrite(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	// store refresh jti
	s.mu.Lock()
	s.refreshStore[refreshClaims.ID] = refreshRecord{Username: in.Username, ExpiresAt: refreshClaims.ExpiresAt.Time}
	s.mu.Unlock()

	// Set refresh as HttpOnly cookie scoped to /refresh (demo; add Secure in prod)
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Path:     "/refresh",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(refreshTTL.Seconds()),
		// Secure: true, // enable behind HTTPS
	}
	http.SetCookie(w, cookie)

	jsonWrite(w, http.StatusOK, map[string]any{
		"access_token": accessToken,
		"token_type":   "Bearer",
		"expires_in":   int(accessTTL.Seconds()),
		"user":         in.Username,
		"kid":          s.activeKID,
	})
}

func (s *server) handleRefresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("refresh_token")
	if err != nil {
		jsonWrite(w, http.StatusUnauthorized, map[string]string{"error": "missing refresh cookie"})
		return
	}
	claims, err := s.verifyToken(c.Value, "refresh")
	if err != nil {
		jsonWrite(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}
	// check store
	s.mu.RLock()
	rec, ok := s.refreshStore[claims.ID]
	s.mu.RUnlock()
	if !ok || rec.Revoked || time.Now().After(rec.ExpiresAt) || rec.Username != claims.Subject {
		jsonWrite(w, http.StatusUnauthorized, map[string]string{"error": "refresh invalid"})
		return
	}
	// rotate: revoke old, issue new refresh & access
	accessTTL := 5 * time.Minute
	refreshTTL := 24 * time.Hour * 7

	accessToken, _, err := s.signToken(claims.Subject, "access", accessTTL, nil)
	if err != nil {
		jsonWrite(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	newRefresh, newRC, err := s.signToken(claims.Subject, "refresh", refreshTTL, nil)
	if err != nil {
		jsonWrite(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	s.mu.Lock()
	rec.Revoked = true
	s.refreshStore[claims.ID] = rec
	s.refreshStore[newRC.ID] = refreshRecord{Username: claims.Subject, ExpiresAt: newRC.ExpiresAt.Time}
	s.mu.Unlock()

	http.SetCookie(w, &http.Cookie{Name: "refresh_token", Value: newRefresh, HttpOnly: true, Path: "/refresh", SameSite: http.SameSiteStrictMode, MaxAge: int(refreshTTL.Seconds())})
	jsonWrite(w, http.StatusOK, map[string]any{"access_token": accessToken, "token_type": "Bearer", "expires_in": int(accessTTL.Seconds())})
}

func (s *server) handleLogout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("refresh_token")
	if err == nil {
		if claims, err2 := s.verifyToken(c.Value, "refresh"); err2 == nil {
			s.mu.Lock()
			rec := s.refreshStore[claims.ID]
			rec.Revoked = true
			s.refreshStore[claims.ID] = rec
			s.mu.Unlock()
		}
	}
	http.SetCookie(w, &http.Cookie{Name: "refresh_token", Value: "", Path: "/refresh", Expires: time.Unix(0, 0)})
	jsonWrite(w, http.StatusOK, map[string]string{"status": "logged out"})
}

func (s *server) handleProtected(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("X-User")
	jsonWrite(w, http.StatusOK, map[string]any{"hello": user, "time": time.Now().Format(time.RFC3339)})
}

func (s *server) handleRotateKeys(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if err := s.rotateKeys(); err != nil {
		jsonWrite(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	jsonWrite(w, http.StatusOK, map[string]string{"kid": s.activeKID})
}

func containsAudience(aud []string, target string) bool {
    for _, v := range aud {
        if v == target {
            return true
        }
    }
    return false
}

// =====================
// Boot
// =====================

func newServer() *server {
	s := &server{
		issuer:       "http://localhost:8080/",
		audience:     "example-api",
		pubKeys:      map[string]*rsa.PublicKey{},
		users:        map[string]string{},
		refreshStore: map[string]refreshRecord{},
	}
	if err := s.rotateKeys(); err != nil {
		log.Fatalf("keygen failed: %v", err)
	}
	// demo user: alice / password123
	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	s.users["alice"] = string(hash)
	return s
}

func main() {
	s := newServer()

	mux := http.NewServeMux()
	mux.HandleFunc("/.well-known/jwks.json", s.handleJWKS)
	mux.HandleFunc("/login", s.handleLogin)
	mux.HandleFunc("/refresh", s.handleRefresh)
	mux.HandleFunc("/logout", s.handleLogout)
	mux.HandleFunc("/rotate-keys", s.handleRotateKeys)
	mux.Handle("/me", s.authMiddleware(http.HandlerFunc(s.handleProtected)))

	addr := ":8090"
	log.Printf("JWT demo server listening on %s\n", addr)
	log.Printf("Login with: curl -s -X POST localhost:8080/login -d '{"+"\"username\":\"alice\",\"password\":\"password123\""+"}' -H 'Content-Type: application/json' | jq\n")
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}


