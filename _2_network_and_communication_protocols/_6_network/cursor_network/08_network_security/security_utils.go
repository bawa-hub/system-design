package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

// SecurityUtils provides various security utility functions
type SecurityUtils struct{}

// NewSecurityUtils creates a new security utils instance
func NewSecurityUtils() *SecurityUtils {
	return &SecurityUtils{}
}

// GenerateRandomString generates a random string of specified length
func (su *SecurityUtils) GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// GenerateSecureRandomString generates a cryptographically secure random string
func (su *SecurityUtils) GenerateSecureRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	
	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b), nil
}

// GenerateAPIKey generates a secure API key
func (su *SecurityUtils) GenerateAPIKey() (string, error) {
	return su.GenerateSecureRandomString(32)
}

// GenerateSessionID generates a secure session ID
func (su *SecurityUtils) GenerateSessionID() (string, error) {
	return su.GenerateSecureRandomString(16)
}

// HashMD5 hashes data using MD5 (not recommended for security)
func (su *SecurityUtils) HashMD5(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// HashSHA1 hashes data using SHA-1 (not recommended for security)
func (su *SecurityUtils) HashSHA1(data string) string {
	hash := sha1.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// HashSHA256 hashes data using SHA-256
func (su *SecurityUtils) HashSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// HashSHA512 hashes data using SHA-512
func (su *SecurityUtils) HashSHA512(data string) string {
	hash := sha512.Sum512([]byte(data))
	return hex.EncodeToString(hash[:])
}

// GenerateHMAC generates HMAC for data with given key
func (su *SecurityUtils) GenerateHMAC(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyHMAC verifies HMAC for data with given key
func (su *SecurityUtils) VerifyHMAC(data, key, expectedMAC string) bool {
	actualMAC := su.GenerateHMAC(data, key)
	return hmac.Equal([]byte(actualMAC), []byte(expectedMAC))
}

// ValidatePasswordStrength validates password strength
func (su *SecurityUtils) ValidatePasswordStrength(password string) (bool, []string) {
	var issues []string
	
	if len(password) < 8 {
		issues = append(issues, "Password must be at least 8 characters long")
	}
	
	if len(password) > 128 {
		issues = append(issues, "Password must be no more than 128 characters long")
	}
	
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	
	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasDigit = true
		case strings.ContainsRune("!@#$%^&*()_+-=[]{}|;:,.<>?", char):
			hasSpecial = true
		}
	}
	
	if !hasUpper {
		issues = append(issues, "Password must contain at least one uppercase letter")
	}
	
	if !hasLower {
		issues = append(issues, "Password must contain at least one lowercase letter")
	}
	
	if !hasDigit {
		issues = append(issues, "Password must contain at least one digit")
	}
	
	if !hasSpecial {
		issues = append(issues, "Password must contain at least one special character")
	}
	
	return len(issues) == 0, issues
}

// SanitizeInput sanitizes user input to prevent XSS
func (su *SecurityUtils) SanitizeInput(input string) string {
	// Replace potentially dangerous characters
	input = strings.ReplaceAll(input, "<", "&lt;")
	input = strings.ReplaceAll(input, ">", "&gt;")
	input = strings.ReplaceAll(input, "\"", "&quot;")
	input = strings.ReplaceAll(input, "'", "&#x27;")
	input = strings.ReplaceAll(input, "&", "&amp;")
	input = strings.ReplaceAll(input, "/", "&#x2F;")
	
	return input
}

// ValidateEmail validates email format
func (su *SecurityUtils) ValidateEmail(email string) bool {
	// Simple email validation
	if len(email) < 5 {
		return false
	}
	
	if !strings.Contains(email, "@") {
		return false
	}
	
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	
	if len(parts[0]) == 0 || len(parts[1]) == 0 {
		return false
	}
	
	if !strings.Contains(parts[1], ".") {
		return false
	}
	
	return true
}

// ValidateIPAddress validates IP address format
func (su *SecurityUtils) ValidateIPAddress(ip string) bool {
	// Check if it's a valid IPv4 or IPv6 address
	if net.ParseIP(ip) == nil {
		return false
	}
	return true
}

// IsPrivateIP checks if an IP address is private
func (su *SecurityUtils) IsPrivateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	
	// Check private IP ranges
	privateRanges := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
		"169.254.0.0/16",
	}
	
	for _, rangeStr := range privateRanges {
		_, network, _ := net.ParseCIDR(rangeStr)
		if network.Contains(parsedIP) {
			return true
		}
	}
	
	return false
}

// GenerateCSRFToken generates a CSRF token
func (su *SecurityUtils) GenerateCSRFToken() (string, error) {
	return su.GenerateSecureRandomString(32)
}

// ValidateCSRFToken validates a CSRF token
func (su *SecurityUtils) ValidateCSRFToken(token, sessionID string) bool {
	// In a real implementation, you would store and validate CSRF tokens
	// This is a simplified version
	return len(token) == 32 && len(sessionID) > 0
}

// RateLimiter provides rate limiting functionality
type RateLimiter struct {
	requests map[string][]time.Time
	mutex    sync.RWMutex
	limit    int
	window   time.Duration
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

// Allow checks if a request is allowed for the given key
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

// GetRemainingRequests returns the number of remaining requests
func (rl *RateLimiter) GetRemainingRequests(key string) int {
	rl.mutex.RLock()
	defer rl.mutex.RUnlock()
	
	now := time.Now()
	cutoff := now.Add(-rl.window)
	
	requests := rl.requests[key]
	var validRequests []time.Time
	for _, req := range requests {
		if req.After(cutoff) {
			validRequests = append(validRequests, req)
		}
	}
	
	return rl.limit - len(validRequests)
}

// SecurityHeaders provides security header management
type SecurityHeaders struct {
	Headers map[string]string
}

// NewSecurityHeaders creates a new security headers instance
func NewSecurityHeaders() *SecurityHeaders {
	return &SecurityHeaders{
		Headers: map[string]string{
			"X-Content-Type-Options": "nosniff",
			"X-Frame-Options":        "DENY",
			"X-XSS-Protection":       "1; mode=block",
			"Strict-Transport-Security": "max-age=31536000; includeSubDomains",
			"Content-Security-Policy":   "default-src 'self'",
			"Referrer-Policy":           "strict-origin-when-cross-origin",
			"Permissions-Policy":        "geolocation=(), microphone=(), camera=()",
		},
	}
}

// SetHeader sets a security header
func (sh *SecurityHeaders) SetHeader(name, value string) {
	sh.Headers[name] = value
}

// ApplyHeaders applies security headers to HTTP response
func (sh *SecurityHeaders) ApplyHeaders(w http.ResponseWriter) {
	for name, value := range sh.Headers {
		w.Header().Set(name, value)
	}
}

// SecurityAudit provides security auditing functionality
type SecurityAudit struct {
	Events []SecurityEvent
	mutex  sync.RWMutex
}

// SecurityEvent represents a security event
type SecurityEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
	Severity  string    `json:"severity"`
	Message   string    `json:"message"`
	IP        string    `json:"ip"`
	UserID    string    `json:"user_id,omitempty"`
}

// NewSecurityAudit creates a new security audit instance
func NewSecurityAudit() *SecurityAudit {
	return &SecurityAudit{
		Events: make([]SecurityEvent, 0),
	}
}

// LogEvent logs a security event
func (sa *SecurityAudit) LogEvent(eventType, severity, message, ip, userID string) {
	sa.mutex.Lock()
	defer sa.mutex.Unlock()
	
	event := SecurityEvent{
		Timestamp: time.Now(),
		Type:      eventType,
		Severity:  severity,
		Message:   message,
		IP:        ip,
		UserID:    userID,
	}
	
	sa.Events = append(sa.Events, event)
	
	// Keep only last 1000 events
	if len(sa.Events) > 1000 {
		sa.Events = sa.Events[len(sa.Events)-1000:]
	}
}

// GetEvents returns security events
func (sa *SecurityAudit) GetEvents() []SecurityEvent {
	sa.mutex.RLock()
	defer sa.mutex.RUnlock()
	
	events := make([]SecurityEvent, len(sa.Events))
	copy(events, sa.Events)
	return events
}

// GetEventsByType returns security events of a specific type
func (sa *SecurityAudit) GetEventsByType(eventType string) []SecurityEvent {
	sa.mutex.RLock()
	defer sa.mutex.RUnlock()
	
	var filtered []SecurityEvent
	for _, event := range sa.Events {
		if event.Type == eventType {
			filtered = append(filtered, event)
		}
	}
	return filtered
}

// Demonstrate security utilities
func demonstrateSecurityUtils() {
	fmt.Println("=== Security Utils Demo ===\n")
	
	// Create security utils
	su := NewSecurityUtils()
	
	// Test random string generation
	randomStr, err := su.GenerateSecureRandomString(16)
	if err != nil {
		log.Printf("Error generating random string: %v", err)
		return
	}
	fmt.Printf("Random string: %s\n", randomStr)
	
	// Test API key generation
	apiKey, err := su.GenerateAPIKey()
	if err != nil {
		log.Printf("Error generating API key: %v", err)
		return
	}
	fmt.Printf("API key: %s\n", apiKey)
	
	// Test password validation
	passwords := []string{
		"weak",
		"password123",
		"StrongP@ssw0rd!",
		"VeryStrongP@ssw0rd123!",
	}
	
	for _, password := range passwords {
		valid, issues := su.ValidatePasswordStrength(password)
		fmt.Printf("Password: %s\n", password)
		fmt.Printf("Valid: %t\n", valid)
		if !valid {
			fmt.Printf("Issues: %v\n", issues)
		}
		fmt.Println()
	}
	
	// Test input sanitization
	maliciousInput := "<script>alert('XSS')</script>"
	sanitized := su.SanitizeInput(maliciousInput)
	fmt.Printf("Original: %s\n", maliciousInput)
	fmt.Printf("Sanitized: %s\n", sanitized)
	
	// Test email validation
	emails := []string{
		"user@example.com",
		"invalid-email",
		"test@domain.co.uk",
		"notanemail",
	}
	
	for _, email := range emails {
		valid := su.ValidateEmail(email)
		fmt.Printf("Email: %s - Valid: %t\n", email, valid)
	}
	
	// Test HMAC
	data := "Hello, World!"
	key := "secret-key"
	hmac := su.GenerateHMAC(data, key)
	fmt.Printf("Data: %s\n", data)
	fmt.Printf("HMAC: %s\n", hmac)
	
	valid := su.VerifyHMAC(data, key, hmac)
	fmt.Printf("HMAC valid: %t\n", valid)
}

// Demonstrate rate limiting
func demonstrateRateLimiting() {
	fmt.Println("=== Rate Limiting Demo ===\n")
	
	// Create rate limiter (5 requests per minute)
	rl := NewRateLimiter(5, time.Minute)
	
	// Test rate limiting
	clientIP := "192.168.1.100"
	
	for i := 0; i < 7; i++ {
		allowed := rl.Allow(clientIP)
		remaining := rl.GetRemainingRequests(clientIP)
		fmt.Printf("Request %d: Allowed=%t, Remaining=%d\n", i+1, allowed, remaining)
	}
}

// Demonstrate security audit
func demonstrateSecurityAudit() {
	fmt.Println("=== Security Audit Demo ===\n")
	
	// Create security audit
	audit := NewSecurityAudit()
	
	// Log some events
	audit.LogEvent("login", "info", "User logged in successfully", "192.168.1.100", "user123")
	audit.LogEvent("login", "warning", "Failed login attempt", "192.168.1.101", "")
	audit.LogEvent("access", "info", "User accessed protected resource", "192.168.1.100", "user123")
	audit.LogEvent("error", "error", "Database connection failed", "192.168.1.100", "user123")
	
	// Get all events
	events := audit.GetEvents()
	fmt.Printf("Total events: %d\n", len(events))
	
	// Get events by type
	loginEvents := audit.GetEventsByType("login")
	fmt.Printf("Login events: %d\n", len(loginEvents))
	
	// Display events
	for _, event := range events {
		fmt.Printf("[%s] %s: %s (IP: %s, User: %s)\n", 
			event.Timestamp.Format("15:04:05"), 
			event.Severity, 
			event.Message, 
			event.IP, 
			event.UserID)
	}
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 8: Security Utils")
	fmt.Println("=====================================================\n")
	
	// Run all demonstrations
	demonstrateSecurityUtils()
	fmt.Println()
	demonstrateRateLimiting()
	fmt.Println()
	demonstrateSecurityAudit()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Security utilities help implement common security patterns")
	fmt.Println("2. Input validation prevents many security vulnerabilities")
	fmt.Println("3. Rate limiting protects against abuse and DoS attacks")
	fmt.Println("4. Security auditing helps track and investigate incidents")
	fmt.Println("5. Proper random number generation is crucial for security")
	fmt.Println("6. Security headers protect against common web vulnerabilities")
}
