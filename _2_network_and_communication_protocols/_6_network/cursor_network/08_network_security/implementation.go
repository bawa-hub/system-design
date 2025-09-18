package main

import (
	"bytes"
	"context"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"strings"
	"time"
)

// SecurityConfig represents security configuration
type SecurityConfig struct {
	MinTLSVersion     uint16
	CipherSuites      []uint16
	CertFile          string
	KeyFile           string
	InsecureSkipVerify bool
	ClientAuth        tls.ClientAuthType
}

// DefaultSecurityConfig returns a secure default configuration
func DefaultSecurityConfig() *SecurityConfig {
	return &SecurityConfig{
		MinTLSVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		},
		InsecureSkipVerify: false,
		ClientAuth:         tls.NoClientCert,
	}
}

// SecureServer represents a secure HTTP server
type SecureServer struct {
	Server   *http.Server
	Config   *SecurityConfig
	Cert     []byte
	Key      []byte
	Running  bool
}

// NewSecureServer creates a new secure server
func NewSecureServer(addr string, config *SecurityConfig) *SecureServer {
	if config == nil {
		config = DefaultSecurityConfig()
	}
	
	// Generate self-signed certificate if not provided
	cert, key, err := generateSelfSignedCert()
	if err != nil {
		log.Printf("Error generating certificate: %v", err)
		return nil
	}
	
	// Create TLS configuration
	tlsConfig := &tls.Config{
		MinVersion:               config.MinTLSVersion,
		CipherSuites:             config.CipherSuites,
		PreferServerCipherSuites: true,
		InsecureSkipVerify:       config.InsecureSkipVerify,
		ClientAuth:               config.ClientAuth,
	}
	
	// Load certificate
	certificate, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Printf("Error loading certificate: %v", err)
		return nil
	}
	
	tlsConfig.Certificates = []tls.Certificate{certificate}
	
	server := &http.Server{
		Addr:      addr,
		TLSConfig: tlsConfig,
	}
	
	return &SecureServer{
		Server:  server,
		Config:  config,
		Cert:    cert,
		Key:     key,
		Running: false,
	}
}

// Start starts the secure server
func (s *SecureServer) Start() error {
	s.Running = true
	fmt.Printf("Secure server starting on %s\n", s.Server.Addr)
	return s.Server.ListenAndServeTLS("", "")
}

// Stop stops the secure server
func (s *SecureServer) Stop() error {
	s.Running = false
	return s.Server.Close()
}

// generateSelfSignedCert generates a self-signed certificate
func generateSelfSignedCert() ([]byte, []byte, error) {
	// Generate private key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	
	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:  []string{"Test Company"},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		DNSNames:     []string{"localhost"},
	}
	
	// Create certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}
	
	// Encode certificate
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	
	// Encode private key
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	
	return certPEM, keyPEM, nil
}

// SecureClient represents a secure HTTP client
type SecureClient struct {
	Client *http.Client
	Config *SecurityConfig
}

// NewSecureClient creates a new secure client
func NewSecureClient(config *SecurityConfig) *SecureClient {
	if config == nil {
		config = DefaultSecurityConfig()
	}
	
	// Create TLS configuration
	tlsConfig := &tls.Config{
		MinVersion:         config.MinTLSVersion,
		CipherSuites:       config.CipherSuites,
		InsecureSkipVerify: config.InsecureSkipVerify,
	}
	
	// Create HTTP client with custom transport
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		Timeout: 30 * time.Second,
	}
	
	return &SecureClient{
		Client: client,
		Config: config,
	}
}

// Get makes a secure GET request
func (c *SecureClient) Get(url string) (*http.Response, error) {
	return c.Client.Get(url)
}

// Post makes a secure POST request
func (c *SecureClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return c.Client.Post(url, contentType, body)
}

// Do makes a custom secure request
func (c *SecureClient) Do(req *http.Request) (*http.Response, error) {
	return c.Client.Do(req)
}

// EncryptionService provides encryption/decryption services
type EncryptionService struct {
	AESKey []byte
	RSAPriv *rsa.PrivateKey
	RSAPub  *rsa.PublicKey
}

// NewEncryptionService creates a new encryption service
func NewEncryptionService() (*EncryptionService, error) {
	// Generate AES key
	aesKey := make([]byte, 32) // 256-bit key
	if _, err := rand.Read(aesKey); err != nil {
		return nil, err
	}
	
	// Generate RSA key pair
	rsaPriv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	
	return &EncryptionService{
		AESKey:  aesKey,
		RSAPriv: rsaPriv,
		RSAPub:  &rsaPriv.PublicKey,
	}, nil
}

// EncryptAES encrypts data using AES-GCM
func (es *EncryptionService) EncryptAES(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(es.AESKey)
	if err != nil {
		return nil, err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// DecryptAES decrypts data using AES-GCM
func (es *EncryptionService) DecryptAES(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(es.AESKey)
	if err != nil {
		return nil, err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	
	return plaintext, nil
}

// EncryptRSA encrypts data using RSA
func (es *EncryptionService) EncryptRSA(plaintext []byte) ([]byte, error) {
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, es.RSAPub, plaintext, nil)
	if err != nil {
		return nil, err
	}
	
	return ciphertext, nil
}

// DecryptRSA decrypts data using RSA
func (es *EncryptionService) DecryptRSA(ciphertext []byte) ([]byte, error) {
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, es.RSAPriv, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	
	return plaintext, nil
}

// SignRSA signs data using RSA
func (es *EncryptionService) SignRSA(data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, es.RSAPriv, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}
	
	return signature, nil
}

// VerifyRSA verifies RSA signature
func (es *EncryptionService) VerifyRSA(data, signature []byte) error {
	hash := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(es.RSAPub, crypto.SHA256, hash[:], signature)
}

// HashService provides hashing services
type HashService struct{}

// NewHashService creates a new hash service
func NewHashService() *HashService {
	return &HashService{}
}

// HashSHA256 hashes data using SHA-256
func (hs *HashService) HashSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// HashSHA256String hashes a string using SHA-256
func (hs *HashService) HashSHA256String(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// HashWithSalt hashes data with salt
func (hs *HashService) HashWithSalt(data, salt []byte) []byte {
	combined := append(data, salt...)
	hash := sha256.Sum256(combined)
	return hash[:]
}

// PasswordService provides password hashing and verification
type PasswordService struct {
	HashService *HashService
}

// NewPasswordService creates a new password service
func NewPasswordService() *PasswordService {
	return &PasswordService{
		HashService: NewHashService(),
	}
}

// HashPassword hashes a password with salt
func (ps *PasswordService) HashPassword(password string) (string, error) {
	// Generate random salt
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	
	// Hash password with salt
	hash := ps.HashService.HashWithSalt([]byte(password), salt)
	
	// Combine hash and salt
	combined := append(hash, salt...)
	
	// Encode as base64
	return base64.StdEncoding.EncodeToString(combined), nil
}

// VerifyPassword verifies a password against its hash
func (ps *PasswordService) VerifyPassword(password, hashedPassword string) bool {
	// Decode base64
	data, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false
	}
	
	// Check length
	if len(data) < 32 {
		return false
	}
	
	// Extract hash and salt
	hash := data[:32]
	salt := data[32:]
	
	// Hash password with salt
	expectedHash := ps.HashService.HashWithSalt([]byte(password), salt)
	
	// Compare hashes
	return bytes.Equal(hash, expectedHash)
}

// JWTService provides JWT token services
type JWTService struct {
	SecretKey []byte
}

// NewJWTService creates a new JWT service
func NewJWTService(secretKey []byte) *JWTService {
	if secretKey == nil {
		secretKey = make([]byte, 32)
		rand.Read(secretKey)
	}
	
	return &JWTService{
		SecretKey: secretKey,
	}
}

// JWTClaims represents JWT claims
type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
	Iat      int64  `json:"iat"`
}

// GenerateToken generates a JWT token
func (js *JWTService) GenerateToken(userID, username string, expirationTime time.Duration) (string, error) {
	now := time.Now()
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		Exp:      now.Add(expirationTime).Unix(),
		Iat:      now.Unix(),
	}
	
	// Create header
	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}
	
	// Encode header
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJSON)
	
	// Encode claims
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)
	
	// Create signature
	message := headerEncoded + "." + claimsEncoded
	signature := js.HashService.HashSHA256([]byte(message + string(js.SecretKey)))
	signatureEncoded := base64.RawURLEncoding.EncodeToString(signature)
	
	// Create JWT
	token := message + "." + signatureEncoded
	
	return token, nil
}

// VerifyToken verifies a JWT token
func (js *JWTService) VerifyToken(token string) (*JWTClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}
	
	// Verify signature
	message := parts[0] + "." + parts[1]
	signature := js.HashService.HashSHA256([]byte(message + string(js.SecretKey)))
	signatureEncoded := base64.RawURLEncoding.EncodeToString(signature)
	
	if parts[2] != signatureEncoded {
		return nil, fmt.Errorf("invalid signature")
	}
	
	// Decode claims
	claimsJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	
	var claims JWTClaims
	if err := json.Unmarshal(claimsJSON, &claims); err != nil {
		return nil, err
	}
	
	// Check expiration
	if time.Now().Unix() > claims.Exp {
		return nil, fmt.Errorf("token expired")
	}
	
	return &claims, nil
}

// SecurityMiddleware provides security middleware for HTTP
type SecurityMiddleware struct {
	JWTService      *JWTService
	PasswordService *PasswordService
}

// NewSecurityMiddleware creates a new security middleware
func NewSecurityMiddleware() *SecurityMiddleware {
	return &SecurityMiddleware{
		JWTService:      NewJWTService(nil),
		PasswordService: NewPasswordService(),
	}
}

// AuthMiddleware provides authentication middleware
func (sm *SecurityMiddleware) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		
		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}
		
		// Verify token
		claims, err := sm.JWTService.VerifyToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		
		// Add user info to request context
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "username", claims.Username)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// CORSMiddleware provides CORS middleware
func (sm *SecurityMiddleware) CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	}
}

// SecurityHeadersMiddleware adds security headers
func (sm *SecurityMiddleware) SecurityHeadersMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		
		next.ServeHTTP(w, r)
	}
}

// Demonstrate secure server
func demonstrateSecureServer() {
	fmt.Println("=== Secure Server Demo ===\n")
	
	// Create secure server
	server := NewSecureServer(":8443", nil)
	if server == nil {
		log.Println("Failed to create secure server")
		return
	}
	
	// Set up routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message": "Secure server response",
			"time":    time.Now().Format(time.RFC3339),
			"secure":  true,
		}
		json.NewEncoder(w).Encode(response)
	})
	
	// Start server in goroutine
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Test with secure client
	client := NewSecureClient(&SecurityConfig{
		InsecureSkipVerify: true, // For self-signed certificate
	})
	
	resp, err := client.Get("https://localhost:8443/")
	if err != nil {
		log.Printf("Client error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Secure response: %s\n", string(body))
	
	// Stop server
	server.Stop()
}

// Demonstrate encryption
func demonstrateEncryption() {
	fmt.Println("=== Encryption Demo ===\n")
	
	// Create encryption service
	es, err := NewEncryptionService()
	if err != nil {
		log.Printf("Error creating encryption service: %v", err)
		return
	}
	
	// Test AES encryption
	plaintext := "Hello, World! This is a secret message."
	fmt.Printf("Original: %s\n", plaintext)
	
	ciphertext, err := es.EncryptAES([]byte(plaintext))
	if err != nil {
		log.Printf("AES encryption error: %v", err)
		return
	}
	fmt.Printf("AES Encrypted: %x\n", ciphertext)
	
	decrypted, err := es.DecryptAES(ciphertext)
	if err != nil {
		log.Printf("AES decryption error: %v", err)
		return
	}
	fmt.Printf("AES Decrypted: %s\n", string(decrypted))
	
	// Test RSA encryption
	rsaCiphertext, err := es.EncryptRSA([]byte(plaintext))
	if err != nil {
		log.Printf("RSA encryption error: %v", err)
		return
	}
	fmt.Printf("RSA Encrypted: %x\n", rsaCiphertext)
	
	rsaDecrypted, err := es.DecryptRSA(rsaCiphertext)
	if err != nil {
		log.Printf("RSA decryption error: %v", err)
		return
	}
	fmt.Printf("RSA Decrypted: %s\n", string(rsaDecrypted))
	
	// Test RSA signing
	signature, err := es.SignRSA([]byte(plaintext))
	if err != nil {
		log.Printf("RSA signing error: %v", err)
		return
	}
	fmt.Printf("RSA Signature: %x\n", signature)
	
	err = es.VerifyRSA([]byte(plaintext), signature)
	if err != nil {
		log.Printf("RSA verification error: %v", err)
		return
	}
	fmt.Println("RSA Signature verified successfully")
}

// Demonstrate password hashing
func demonstratePasswordHashing() {
	fmt.Println("=== Password Hashing Demo ===\n")
	
	// Create password service
	ps := NewPasswordService()
	
	// Test password hashing
	password := "mySecretPassword123"
	fmt.Printf("Original password: %s\n", password)
	
	hashedPassword, err := ps.HashPassword(password)
	if err != nil {
		log.Printf("Password hashing error: %v", err)
		return
	}
	fmt.Printf("Hashed password: %s\n", hashedPassword)
	
	// Test password verification
	valid := ps.VerifyPassword(password, hashedPassword)
	fmt.Printf("Password verification: %t\n", valid)
	
	// Test with wrong password
	wrongPassword := "wrongPassword"
	valid = ps.VerifyPassword(wrongPassword, hashedPassword)
	fmt.Printf("Wrong password verification: %t\n", valid)
}

// Demonstrate JWT tokens
func demonstrateJWT() {
	fmt.Println("=== JWT Token Demo ===\n")
	
	// Create JWT service
	js := NewJWTService(nil)
	
	// Generate token
	userID := "user123"
	username := "john_doe"
	expirationTime := 24 * time.Hour
	
	token, err := js.GenerateToken(userID, username, expirationTime)
	if err != nil {
		log.Printf("Token generation error: %v", err)
		return
	}
	fmt.Printf("Generated token: %s\n", token)
	
	// Verify token
	claims, err := js.VerifyToken(token)
	if err != nil {
		log.Printf("Token verification error: %v", err)
		return
	}
	fmt.Printf("Token claims: UserID=%s, Username=%s, Exp=%d\n", 
		claims.UserID, claims.Username, claims.Exp)
	
	// Test with invalid token
	invalidToken := "invalid.token.here"
	_, err = js.VerifyToken(invalidToken)
	if err != nil {
		fmt.Printf("Invalid token verification (expected): %v\n", err)
	}
}

// Demonstrate security middleware
func demonstrateSecurityMiddleware() {
	fmt.Println("=== Security Middleware Demo ===\n")
	
	// Create security middleware
	sm := NewSecurityMiddleware()
	
	// Create test server
	mux := http.NewServeMux()
	
	// Public endpoint
	mux.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message": "This is a public endpoint",
			"time":    time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(response)
	})
	
	// Protected endpoint
	mux.HandleFunc("/protected", sm.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(string)
		username := r.Context().Value("username").(string)
		
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message":  "This is a protected endpoint",
			"user_id":  userID,
			"username": username,
			"time":     time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(response)
	}))
	
	// Apply security middleware
	handler := sm.SecurityHeadersMiddleware(sm.CORSMiddleware(mux.ServeHTTP))
	
	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	
	go func() {
		fmt.Println("Security middleware server starting on :8080")
		log.Fatal(server.ListenAndServe())
	}()
	
	// Give server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Test public endpoint
	resp, err := http.Get("http://localhost:8080/public")
	if err != nil {
		log.Printf("Public endpoint error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Public endpoint response: %s\n", string(body))
	
	// Test protected endpoint without token
	resp, err = http.Get("http://localhost:8080/protected")
	if err != nil {
		log.Printf("Protected endpoint error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Protected endpoint response (no token): %s\n", string(body))
	
	// Generate token for testing
	js := NewJWTService(nil)
	token, err := js.GenerateToken("user123", "john_doe", time.Hour)
	if err != nil {
		log.Printf("Token generation error: %v", err)
		return
	}
	
	// Test protected endpoint with token
	req, err := http.NewRequest("GET", "http://localhost:8080/protected", nil)
	if err != nil {
		log.Printf("Request creation error: %v", err)
		return
	}
	
	req.Header.Set("Authorization", "Bearer "+token)
	
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Printf("Protected endpoint error: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	
	fmt.Printf("Protected endpoint response (with token): %s\n", string(body))
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 8: Network Security & Encryption")
	fmt.Println("====================================================================\n")
	
	// Run all demonstrations
	demonstrateSecureServer()
	fmt.Println()
	demonstrateEncryption()
	fmt.Println()
	demonstratePasswordHashing()
	fmt.Println()
	demonstrateJWT()
	fmt.Println()
	demonstrateSecurityMiddleware()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Network security protects data and communications")
	fmt.Println("2. Encryption ensures confidentiality and integrity")
	fmt.Println("3. TLS/SSL provides secure communication over networks")
	fmt.Println("4. Authentication verifies user and system identity")
	fmt.Println("5. Proper key management is crucial for security")
	fmt.Println("6. Security middleware protects web applications")
	
	fmt.Println("\n📚 Next Topic: Routing & Switching")
}
