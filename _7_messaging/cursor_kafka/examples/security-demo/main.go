package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("🔐 Kafka Security and Authentication Demo")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Securing Kafka clusters with authentication and authorization...")
	fmt.Println()

	// Run different security demonstrations
	demonstrateSASLAuthentication()
	demonstrateTLSEncryption()
	demonstrateACLAuthorization()
	demonstrateSecurityMonitoring()
}

func demonstrateSASLAuthentication() {
	fmt.Println("🔑 SASL Authentication Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create SASL configurations
	saslConfigs := []SASLConfig{
		{
			Mechanism: "PLAIN",
			Username:  "admin",
			Password:  "admin-secret",
		},
		{
			Mechanism: "SCRAM-SHA-256",
			Username:  "user1",
			Password:  "user1-secret",
		},
		{
			Mechanism: "SCRAM-SHA-512",
			Username:  "user2",
			Password:  "user2-secret",
		},
	}

	fmt.Println("🔧 Testing SASL authentication mechanisms...")
	for _, config := range saslConfigs {
		fmt.Printf("   Testing %s authentication...\n", config.Mechanism)
		
		// Simulate authentication
		success := simulateSASLAuthentication(config)
		if success {
			fmt.Printf("   ✅ Authentication successful for user: %s\n", config.Username)
		} else {
			fmt.Printf("   ❌ Authentication failed for user: %s\n", config.Username)
		}
	}
	fmt.Println()
}

func demonstrateTLSEncryption() {
	fmt.Println("🔒 TLS Encryption Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create TLS configuration
	tlsConfig := TLSConfig{
		EnableTLS:     true,
		MinVersion:    tls.VersionTLS12,
		MaxVersion:    tls.VersionTLS13,
		InsecureSkip:  false,
	}

	fmt.Println("🔧 Testing TLS encryption...")
	
	// Generate certificates
	fmt.Println("   Generating certificates...")
	caCert, caKey, err := generateCACertificate()
	if err != nil {
		fmt.Printf("   ❌ Error generating CA certificate: %v\n", err)
	} else {
		fmt.Println("   ✅ CA certificate generated successfully")
	}

	// Generate server certificate
	serverCert, serverKey, err := generateServerCertificate(caCert, caKey)
	if err != nil {
		fmt.Printf("   ❌ Error generating server certificate: %v\n", err)
	} else {
		fmt.Println("   ✅ Server certificate generated successfully")
	}

	// Generate client certificate
	clientCert, clientKey, err := generateClientCertificate(caCert, caKey)
	if err != nil {
		fmt.Printf("   ❌ Error generating client certificate: %v\n", err)
	} else {
		fmt.Println("   ✅ Client certificate generated successfully")
	}

	// Test TLS handshake
	fmt.Println("   Testing TLS handshake...")
	success := simulateTLSHandshake(tlsConfig, serverCert, serverKey, clientCert, clientKey)
	if success {
		fmt.Println("   ✅ TLS handshake successful")
	} else {
		fmt.Println("   ❌ TLS handshake failed")
	}
	fmt.Println()
}

func demonstrateACLAuthorization() {
	fmt.Println("🛡️ ACL Authorization Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create ACL manager
	aclManager := NewACLManager()

	// Define ACL rules
	aclRules := []ACLRule{
		{
			Principal:    "User:admin",
			Host:         "*",
			Operation:    "All",
			Permission:   "Allow",
			ResourceType: "Topic",
			ResourceName: "user-events",
		},
		{
			Principal:    "User:user1",
			Host:         "*",
			Operation:    "Read",
			Permission:   "Allow",
			ResourceType: "Topic",
			ResourceName: "user-events",
		},
		{
			Principal:    "User:user2",
			Host:         "*",
			Operation:    "Write",
			Permission:   "Allow",
			ResourceType: "Topic",
			ResourceName: "order-events",
		},
	}

	fmt.Println("🔧 Testing ACL authorization...")
	
	// Create ACL rules
	for i, rule := range aclRules {
		fmt.Printf("   Creating ACL rule %d: %s %s on %s\n", 
			i+1, rule.Principal, rule.Operation, rule.ResourceName)
		
		err := aclManager.CreateACL(rule)
		if err != nil {
			fmt.Printf("   ❌ Error creating ACL: %v\n", err)
		} else {
			fmt.Printf("   ✅ ACL rule created successfully\n")
		}
	}

	// Test authorization
	fmt.Println("\n   Testing authorization...")
	testCases := []struct {
		user     string
		resource string
		action   string
		allowed  bool
	}{
		{"admin", "user-events", "Read", true},
		{"admin", "user-events", "Write", true},
		{"user1", "user-events", "Read", true},
		{"user1", "user-events", "Write", false},
		{"user2", "order-events", "Write", true},
		{"user2", "user-events", "Read", false},
	}

	for _, tc := range testCases {
		allowed := aclManager.CheckPermission(tc.user, tc.resource, tc.action)
		if allowed == tc.allowed {
			fmt.Printf("   ✅ %s %s on %s: %v (expected: %v)\n", 
				tc.user, tc.action, tc.resource, allowed, tc.allowed)
		} else {
			fmt.Printf("   ❌ %s %s on %s: %v (expected: %v)\n", 
				tc.user, tc.action, tc.resource, allowed, tc.allowed)
		}
	}
	fmt.Println()
}

func demonstrateSecurityMonitoring() {
	fmt.Println("📊 Security Monitoring Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create security monitoring components
	eventLogger := NewSecurityEventLogger()
	metrics := NewSecurityMetrics()
	alertManager := NewSecurityAlertManager()

	fmt.Println("🔧 Testing security monitoring...")
	
	// Simulate security events
	events := []SecurityEvent{
		{
			EventType: "AuthenticationSuccess",
			User:      "admin",
			Resource:  "kafka-cluster",
			Action:    "Login",
			Result:    "Success",
			IPAddress: "192.168.1.100",
		},
		{
			EventType: "AuthenticationFailure",
			User:      "hacker",
			Resource:  "kafka-cluster",
			Action:    "Login",
			Result:    "Failure",
			IPAddress: "192.168.1.200",
		},
		{
			EventType: "AuthorizationFailure",
			User:      "user1",
			Resource:  "admin-topic",
			Action:    "Write",
			Result:    "Denied",
			IPAddress: "192.168.1.101",
		},
		{
			EventType: "SuspiciousActivity",
			User:      "unknown",
			Resource:  "kafka-cluster",
			Action:    "MultipleFailedLogins",
			Result:    "Detected",
			IPAddress: "192.168.1.300",
		},
	}

	// Log events
	for i, event := range events {
		fmt.Printf("   Logging event %d: %s\n", i+1, event.EventType)
		eventLogger.LogEvent(event)
		
		// Update metrics
		switch event.EventType {
		case "AuthenticationSuccess":
			metrics.RecordSuccessfulLogin()
		case "AuthenticationFailure":
			metrics.RecordAuthenticationFailure()
		case "AuthorizationFailure":
			metrics.RecordAuthorizationFailure()
		case "SuspiciousActivity":
			metrics.RecordSuspiciousActivity()
		}
	}

	// Generate security report
	fmt.Println("\n   Generating security report...")
	report := metrics.GetSecurityReport()
	fmt.Printf("   📊 Security Metrics:\n")
	fmt.Printf("      Successful Logins: %d\n", report.SuccessfulLogins)
	fmt.Printf("      Failed Logins: %d\n", report.FailedLogins)
	fmt.Printf("      Authentication Failures: %d\n", report.AuthenticationFailures)
	fmt.Printf("      Authorization Failures: %d\n", report.AuthorizationFailures)
	fmt.Printf("      Suspicious Activities: %d\n", report.SuspiciousActivities)

	// Check for alerts
	fmt.Println("\n   Checking for security alerts...")
	alerts := alertManager.CheckAlerts(report)
	if len(alerts) > 0 {
		fmt.Printf("   🚨 %d security alerts detected:\n", len(alerts))
		for i, alert := range alerts {
			fmt.Printf("      Alert %d: %s - %s\n", i+1, alert.Severity, alert.Message)
		}
	} else {
		fmt.Println("   ✅ No security alerts detected")
	}
	fmt.Println()
}

// SASL Configuration
type SASLConfig struct {
	Mechanism string
	Username  string
	Password  string
	Token     string
	Realm     string
}

func simulateSASLAuthentication(config SASLConfig) bool {
	// Simulate authentication based on mechanism
	switch config.Mechanism {
	case "PLAIN":
		return config.Username == "admin" && config.Password == "admin-secret"
	case "SCRAM-SHA-256":
		return config.Username == "user1" && config.Password == "user1-secret"
	case "SCRAM-SHA-512":
		return config.Username == "user2" && config.Password == "user2-secret"
	default:
		return false
	}
}

// TLS Configuration
type TLSConfig struct {
	EnableTLS     bool
	MinVersion    uint16
	MaxVersion    uint16
	InsecureSkip  bool
}

func generateCACertificate() (*x509.Certificate, *rsa.PrivateKey, error) {
	// Generate CA private key
	caKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Create CA certificate template
	caTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:  []string{"Kafka CA"},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// Create CA certificate
	caCertDER, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caKey.PublicKey, caKey)
	if err != nil {
		return nil, nil, err
	}

	// Parse CA certificate
	caCert, err := x509.ParseCertificate(caCertDER)
	if err != nil {
		return nil, nil, err
	}

	return caCert, caKey, nil
}

func generateServerCertificate(caCert *x509.Certificate, caKey *rsa.PrivateKey) (*x509.Certificate, *rsa.PrivateKey, error) {
	// Generate server private key
	serverKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Create server certificate template
	serverTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Organization:  []string{"Kafka Server"},
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
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1"), net.ParseIP("192.168.1.1")},
		DNSNames:     []string{"localhost", "kafka-server"},
	}

	// Create server certificate
	serverCertDER, err := x509.CreateCertificate(rand.Reader, &serverTemplate, caCert, &serverKey.PublicKey, caKey)
	if err != nil {
		return nil, nil, err
	}

	// Parse server certificate
	serverCert, err := x509.ParseCertificate(serverCertDER)
	if err != nil {
		return nil, nil, err
	}

	return serverCert, serverKey, nil
}

func generateClientCertificate(caCert *x509.Certificate, caKey *rsa.PrivateKey) (*x509.Certificate, *rsa.PrivateKey, error) {
	// Generate client private key
	clientKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Create client certificate template
	clientTemplate := x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject: pkix.Name{
			Organization:  []string{"Kafka Client"},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}

	// Create client certificate
	clientCertDER, err := x509.CreateCertificate(rand.Reader, &clientTemplate, caCert, &clientKey.PublicKey, caKey)
	if err != nil {
		return nil, nil, err
	}

	// Parse client certificate
	clientCert, err := x509.ParseCertificate(clientCertDER)
	if err != nil {
		return nil, nil, err
	}

	return clientCert, clientKey, nil
}

func simulateTLSHandshake(config TLSConfig, serverCert *x509.Certificate, serverKey *rsa.PrivateKey, clientCert *x509.Certificate, clientKey *rsa.PrivateKey) bool {
	// Simulate TLS handshake
	// In a real implementation, this would involve actual TLS handshake
	return config.EnableTLS && serverCert != nil && clientCert != nil
}

// ACL Manager
type ACLManager struct {
	rules map[string]ACLRule
	mutex sync.RWMutex
}

type ACLRule struct {
	Principal    string
	Host         string
	Operation    string
	Permission   string
	ResourceType string
	ResourceName string
}

func NewACLManager() *ACLManager {
	return &ACLManager{
		rules: make(map[string]ACLRule),
	}
}

func (am *ACLManager) CreateACL(rule ACLRule) error {
	am.mutex.Lock()
	defer am.mutex.Unlock()
	
	key := fmt.Sprintf("%s:%s:%s:%s", rule.Principal, rule.ResourceType, rule.ResourceName, rule.Operation)
	am.rules[key] = rule
	return nil
}

func (am *ACLManager) CheckPermission(user, resource, action string) bool {
	am.mutex.RLock()
	defer am.mutex.RUnlock()
	
	// Check for exact match
	key := fmt.Sprintf("User:%s:Topic:%s:%s", user, resource, action)
	if rule, exists := am.rules[key]; exists {
		return rule.Permission == "Allow"
	}
	
	// Check for wildcard match
	key = fmt.Sprintf("User:%s:Topic:%s:All", user, resource)
	if rule, exists := am.rules[key]; exists {
		return rule.Permission == "Allow"
	}
	
	return false
}

// Security Event Logger
type SecurityEventLogger struct {
	events    []SecurityEvent
	mutex     sync.RWMutex
	maxEvents int
}

type SecurityEvent struct {
	Timestamp   time.Time
	EventType   string
	User        string
	Resource    string
	Action      string
	Result      string
	IPAddress   string
	UserAgent   string
	Details     map[string]interface{}
}

func NewSecurityEventLogger() *SecurityEventLogger {
	return &SecurityEventLogger{
		events:    make([]SecurityEvent, 0),
		maxEvents: 1000,
	}
}

func (sel *SecurityEventLogger) LogEvent(event SecurityEvent) {
	sel.mutex.Lock()
	defer sel.mutex.Unlock()
	
	event.Timestamp = time.Now()
	sel.events = append(sel.events, event)
	
	// Keep only last maxEvents
	if len(sel.events) > sel.maxEvents {
		sel.events = sel.events[1:]
	}
}

// Security Metrics
type SecurityMetrics struct {
	AuthenticationFailures int64
	AuthorizationFailures  int64
	TLSHandshakeFailures   int64
	SuspiciousActivities   int64
	FailedLogins          int64
	SuccessfulLogins      int64
}

func NewSecurityMetrics() *SecurityMetrics {
	return &SecurityMetrics{}
}

func (sm *SecurityMetrics) RecordAuthenticationFailure() {
	atomic.AddInt64(&sm.AuthenticationFailures, 1)
}

func (sm *SecurityMetrics) RecordAuthorizationFailure() {
	atomic.AddInt64(&sm.AuthorizationFailures, 1)
}

func (sm *SecurityMetrics) RecordTLSHandshakeFailure() {
	atomic.AddInt64(&sm.TLSHandshakeFailures, 1)
}

func (sm *SecurityMetrics) RecordSuspiciousActivity() {
	atomic.AddInt64(&sm.SuspiciousActivities, 1)
}

func (sm *SecurityMetrics) RecordFailedLogin() {
	atomic.AddInt64(&sm.FailedLogins, 1)
}

func (sm *SecurityMetrics) RecordSuccessfulLogin() {
	atomic.AddInt64(&sm.SuccessfulLogins, 1)
}

func (sm *SecurityMetrics) GetSecurityReport() SecurityReport {
	return SecurityReport{
		AuthenticationFailures: atomic.LoadInt64(&sm.AuthenticationFailures),
		AuthorizationFailures:  atomic.LoadInt64(&sm.AuthorizationFailures),
		TLSHandshakeFailures:   atomic.LoadInt64(&sm.TLSHandshakeFailures),
		SuspiciousActivities:   atomic.LoadInt64(&sm.SuspiciousActivities),
		FailedLogins:          atomic.LoadInt64(&sm.FailedLogins),
		SuccessfulLogins:      atomic.LoadInt64(&sm.SuccessfulLogins),
		Timestamp:             time.Now(),
	}
}

type SecurityReport struct {
	AuthenticationFailures int64
	AuthorizationFailures  int64
	TLSHandshakeFailures   int64
	SuspiciousActivities   int64
	FailedLogins          int64
	SuccessfulLogins      int64
	Timestamp             time.Time
}

// Security Alert Manager
type SecurityAlertManager struct {
	alerts []SecurityAlert
	mutex  sync.RWMutex
}

type SecurityAlert struct {
	Severity    string
	Message     string
	Timestamp   time.Time
	Threshold   int64
	Current     int64
}

func NewSecurityAlertManager() *SecurityAlertManager {
	return &SecurityAlertManager{
		alerts: make([]SecurityAlert, 0),
	}
}

func (sam *SecurityAlertManager) CheckAlerts(report SecurityReport) []SecurityAlert {
	sam.mutex.Lock()
	defer sam.mutex.Unlock()
	
	var newAlerts []SecurityAlert
	
	// Check for failed login threshold
	if report.FailedLogins > 10 {
		newAlerts = append(newAlerts, SecurityAlert{
			Severity:  "HIGH",
			Message:   "High number of failed login attempts detected",
			Timestamp: time.Now(),
			Threshold: 10,
			Current:   report.FailedLogins,
		})
	}
	
	// Check for suspicious activities
	if report.SuspiciousActivities > 5 {
		newAlerts = append(newAlerts, SecurityAlert{
			Severity:  "CRITICAL",
			Message:   "Multiple suspicious activities detected",
			Timestamp: time.Now(),
			Threshold: 5,
			Current:   report.SuspiciousActivities,
		})
	}
	
	// Check for authentication failures
	if report.AuthenticationFailures > 20 {
		newAlerts = append(newAlerts, SecurityAlert{
			Severity:  "MEDIUM",
			Message:   "High number of authentication failures",
			Timestamp: time.Now(),
			Threshold: 20,
			Current:   report.AuthenticationFailures,
		})
	}
	
	sam.alerts = append(sam.alerts, newAlerts...)
	return newAlerts
}
