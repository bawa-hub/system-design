# Network Security & Encryption Quick Reference Guide

## 🚀 Essential Security Concepts

### Security Principles
- **Confidentiality**: Data only accessible to authorized users
- **Integrity**: Data has not been modified or corrupted
- **Availability**: Data and services available when needed
- **Authentication**: Verifying identity of users/systems
- **Authorization**: Controlling access to resources
- **Non-repudiation**: Preventing denial of actions

### Encryption Types
- **Symmetric**: Same key for encryption/decryption (AES, ChaCha20)
- **Asymmetric**: Public/private key pairs (RSA, ECC)
- **Hash Functions**: One-way functions (SHA-256, SHA-3)
- **Digital Signatures**: Authentication and integrity (RSA, ECDSA)

## 🔧 Go Security Patterns

### TLS Configuration
```go
config := &tls.Config{
    MinVersion:         tls.VersionTLS12,
    CipherSuites:       []uint16{
        tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
        tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
    },
    PreferServerCipherSuites: true,
    InsecureSkipVerify: false,
}
```

### Secure HTTP Server
```go
func startSecureServer() {
    server := &http.Server{
        Addr:      ":8443",
        TLSConfig: config,
    }
    
    log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
```

### Secure HTTP Client
```go
func createSecureClient() *http.Client {
    return &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: false,
            },
        },
        Timeout: 30 * time.Second,
    }
}
```

## 🔐 Encryption Implementation

### AES Encryption
```go
func encryptAES(plaintext, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
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

func decryptAES(ciphertext, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
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
    return gcm.Open(nil, nonce, ciphertext, nil)
}
```

### RSA Encryption
```go
func encryptRSA(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
    hash := sha256.New()
    return rsa.EncryptOAEP(hash, rand.Reader, publicKey, plaintext, nil)
}

func decryptRSA(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
    hash := sha256.New()
    return rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, nil)
}
```

### Digital Signatures
```go
func signRSA(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
    hash := sha256.Sum256(data)
    return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
}

func verifyRSA(data, signature []byte, publicKey *rsa.PublicKey) error {
    hash := sha256.Sum256(data)
    return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
}
```

## 🔑 Password Security

### Password Hashing
```go
func hashPassword(password string) (string, error) {
    salt := make([]byte, 32)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }
    
    hash := sha256.Sum256(append([]byte(password), salt...))
    combined := append(hash[:], salt...)
    
    return base64.StdEncoding.EncodeToString(combined), nil
}

func verifyPassword(password, hashedPassword string) bool {
    data, err := base64.StdEncoding.DecodeString(hashedPassword)
    if err != nil || len(data) < 32 {
        return false
    }
    
    hash := data[:32]
    salt := data[32:]
    
    expectedHash := sha256.Sum256(append([]byte(password), salt...))
    return bytes.Equal(hash, expectedHash[:])
}
```

### Password Validation
```go
func validatePassword(password string) (bool, []string) {
    var issues []string
    
    if len(password) < 8 {
        issues = append(issues, "Password must be at least 8 characters")
    }
    
    hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
    hasDigit := strings.ContainsAny(password, "0123456789")
    hasSpecial := strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:,.<>?")
    
    if !hasUpper {
        issues = append(issues, "Password must contain uppercase letter")
    }
    if !hasLower {
        issues = append(issues, "Password must contain lowercase letter")
    }
    if !hasDigit {
        issues = append(issues, "Password must contain digit")
    }
    if !hasSpecial {
        issues = append(issues, "Password must contain special character")
    }
    
    return len(issues) == 0, issues
}
```

## 🎫 JWT Tokens

### JWT Implementation
```go
type JWTClaims struct {
    UserID   string `json:"user_id"`
    Username string `json:"username"`
    Exp      int64  `json:"exp"`
    Iat      int64  `json:"iat"`
}

func generateJWT(userID, username string, secretKey []byte) (string, error) {
    claims := JWTClaims{
        UserID:   userID,
        Username: username,
        Exp:      time.Now().Add(24 * time.Hour).Unix(),
        Iat:      time.Now().Unix(),
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

func verifyJWT(tokenString string, secretKey []byte) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    
    if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, err
}
```

## 🛡️ Security Middleware

### Authentication Middleware
```go
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Authorization required", http.StatusUnauthorized)
            return
        }
        
        token = strings.TrimPrefix(token, "Bearer ")
        claims, err := verifyJWT(token, secretKey)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}
```

### Security Headers Middleware
```go
func securityHeadersMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("Strict-Transport-Security", "max-age=31536000")
        w.Header().Set("Content-Security-Policy", "default-src 'self'")
        
        next.ServeHTTP(w, r)
    }
}
```

### CORS Middleware
```go
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    }
}
```

## 🚦 Rate Limiting

### Rate Limiter Implementation
```go
type RateLimiter struct {
    requests map[string][]time.Time
    mutex    sync.RWMutex
    limit    int
    window   time.Duration
}

func (rl *RateLimiter) Allow(key string) bool {
    rl.mutex.Lock()
    defer rl.mutex.Unlock()
    
    now := time.Now()
    cutoff := now.Add(-rl.window)
    
    // Clean old requests
    requests := rl.requests[key]
    var validRequests []time.Time
    for _, req := range requests {
        if req.After(cutoff) {
            validRequests = append(validRequests, req)
        }
    }
    
    if len(validRequests) >= rl.limit {
        return false
    }
    
    validRequests = append(validRequests, now)
    rl.requests[key] = validRequests
    return true
}
```

## 🔍 Input Validation

### Input Sanitization
```go
func sanitizeInput(input string) string {
    input = strings.ReplaceAll(input, "<", "&lt;")
    input = strings.ReplaceAll(input, ">", "&gt;")
    input = strings.ReplaceAll(input, "\"", "&quot;")
    input = strings.ReplaceAll(input, "'", "&#x27;")
    input = strings.ReplaceAll(input, "&", "&amp;")
    return input
}
```

### Email Validation
```go
func validateEmail(email string) bool {
    if len(email) < 5 || !strings.Contains(email, "@") {
        return false
    }
    
    parts := strings.Split(email, "@")
    if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
        return false
    }
    
    return strings.Contains(parts[1], ".")
}
```

## 📊 Security Monitoring

### Security Event Logging
```go
type SecurityEvent struct {
    Timestamp time.Time `json:"timestamp"`
    Type      string    `json:"type"`
    Severity  string    `json:"severity"`
    Message   string    `json:"message"`
    IP        string    `json:"ip"`
    UserID    string    `json:"user_id,omitempty"`
}

func logSecurityEvent(eventType, severity, message, ip, userID string) {
    event := SecurityEvent{
        Timestamp: time.Now(),
        Type:      eventType,
        Severity:  severity,
        Message:   message,
        IP:        ip,
        UserID:    userID,
    }
    
    // Log to file, database, or monitoring system
    log.Printf("Security Event: %+v", event)
}
```

## 🔒 Certificate Management

### Self-Signed Certificate Generation
```go
func generateSelfSignedCert() ([]byte, []byte, error) {
    priv, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, nil, err
    }
    
    template := x509.Certificate{
        SerialNumber: big.NewInt(1),
        Subject: pkix.Name{
            Organization: []string{"Test Company"},
        },
        NotBefore:    time.Now(),
        NotAfter:     time.Now().Add(365 * 24 * time.Hour),
        KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
        ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
        IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1)},
        DNSNames:     []string{"localhost"},
    }
    
    certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
    if err != nil {
        return nil, nil, err
    }
    
    certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
    keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
    
    return certPEM, keyPEM, nil
}
```

## 🎯 Best Practices

1. **Use strong encryption**: AES-256, RSA-2048+, SHA-256+
2. **Implement proper key management**: Secure storage, rotation
3. **Validate all input**: Sanitize, validate, escape
4. **Use HTTPS everywhere**: TLS 1.2+ for all communications
5. **Implement proper authentication**: Multi-factor when possible
6. **Log security events**: Monitor and audit
7. **Keep systems updated**: Patch vulnerabilities
8. **Use secure defaults**: Strong configurations
9. **Implement rate limiting**: Prevent abuse
10. **Follow OWASP guidelines**: Security best practices

## 🚨 Common Vulnerabilities

### OWASP Top 10
1. **Injection**: SQL, NoSQL, OS command injection
2. **Broken Authentication**: Weak authentication mechanisms
3. **Sensitive Data Exposure**: Inadequate protection
4. **XML External Entities**: XXE attacks
5. **Broken Access Control**: Inadequate access controls
6. **Security Misconfiguration**: Insecure defaults
7. **Cross-Site Scripting**: XSS attacks
8. **Insecure Deserialization**: Unsafe deserialization
9. **Known Vulnerabilities**: Outdated components
10. **Insufficient Logging**: Inadequate monitoring

---

**Remember**: Security is not optional. Implement these patterns, and you'll build secure, robust applications! 🚀
