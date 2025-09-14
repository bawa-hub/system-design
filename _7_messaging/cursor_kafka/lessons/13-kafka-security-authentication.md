# Lesson 13: Kafka Security and Authentication

## 🎯 Learning Objectives
By the end of this lesson, you will understand:
- Kafka security architecture and components
- Authentication mechanisms (SASL, SSL/TLS, OAuth2)
- Authorization and access control (ACLs, RBAC)
- Encryption and data protection
- Security best practices and compliance
- Monitoring and auditing security events
- Production security configurations

## 📚 Theory Section

### 1. **Kafka Security Overview**

#### **Security Layers:**
```
┌─────────────────────────────────────┐
│           Application Layer         │
├─────────────────────────────────────┤
│         Authorization (ACLs)        │
├─────────────────────────────────────┤
│        Authentication (SASL)        │
├─────────────────────────────────────┤
│         Encryption (SSL/TLS)        │
├─────────────────────────────────────┤
│         Network Security            │
└─────────────────────────────────────┘
```

#### **Security Components:**
- **Authentication**: Verifying user identity
- **Authorization**: Controlling access to resources
- **Encryption**: Protecting data in transit and at rest
- **Auditing**: Logging security events
- **Monitoring**: Detecting security threats

### 2. **Authentication Mechanisms**

#### **SASL (Simple Authentication and Security Layer):**
```go
type SASLConfig struct {
    Mechanism   string // PLAIN, SCRAM-SHA-256, SCRAM-SHA-512, GSSAPI, OAUTHBEARER
    Username    string
    Password    string
    Token       string // For OAuth2
    Realm       string // For GSSAPI
}

func (sc *SASLConfig) ConfigureSASL(config *sarama.Config) {
    config.Net.SASL.Enable = true
    config.Net.SASL.Mechanism = sarama.SASLMechanism(sc.Mechanism)
    config.Net.SASL.User = sc.Username
    config.Net.SASL.Password = sc.Password
    
    switch sc.Mechanism {
    case "SCRAM-SHA-256":
        config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
            return &XDGSCRAMClient{HashGeneratorFcn: SHA256}
        }
    case "SCRAM-SHA-512":
        config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
            return &XDGSCRAMClient{HashGeneratorFcn: SHA512}
        }
    case "OAUTHBEARER":
        config.Net.SASL.TokenProvider = &OAuthTokenProvider{
            Token: sc.Token,
        }
    }
}
```

#### **SASL Mechanisms:**
```go
type SASLMechanism string

const (
    SASL_PLAIN        SASLMechanism = "PLAIN"
    SASL_SCRAM_SHA256 SASLMechanism = "SCRAM-SHA-256"
    SASL_SCRAM_SHA512 SASLMechanism = "SCRAM-SHA-512"
    SASL_GSSAPI       SASLMechanism = "GSSAPI"
    SASL_OAUTHBEARER  SASLMechanism = "OAUTHBEARER"
)

func (sm SASLMechanism) IsSecure() bool {
    switch sm {
    case SASL_SCRAM_SHA256, SASL_SCRAM_SHA512, SASL_GSSAPI, SASL_OAUTHBEARER:
        return true
    default:
        return false
    }
}
```

#### **OAuth2 Authentication:**
```go
type OAuth2Config struct {
    ClientID     string
    ClientSecret string
    TokenURL     string
    Scope        []string
    GrantType    string // client_credentials, password, authorization_code
}

type OAuth2TokenProvider struct {
    config     *OAuth2Config
    token      *oauth2.Token
    tokenMutex sync.RWMutex
    httpClient *http.Client
}

func (otp *OAuth2TokenProvider) Token() (*oauth2.Token, error) {
    otp.tokenMutex.RLock()
    if otp.token != nil && otp.token.Valid() {
        otp.tokenMutex.RUnlock()
        return otp.token, nil
    }
    otp.tokenMutex.RUnlock()
    
    // Refresh token
    otp.tokenMutex.Lock()
    defer otp.tokenMutex.Unlock()
    
    ctx := context.Background()
    conf := &oauth2.Config{
        ClientID:     otp.config.ClientID,
        ClientSecret: otp.config.ClientSecret,
        TokenURL:     otp.config.TokenURL,
        Scopes:       otp.config.Scope,
    }
    
    token, err := conf.ClientCredentialsTokenSource(ctx).Token()
    if err != nil {
        return nil, err
    }
    
    otp.token = token
    return token, nil
}
```

### 3. **SSL/TLS Encryption**

#### **TLS Configuration:**
```go
type TLSConfig struct {
    EnableTLS     bool
    CertFile      string
    KeyFile       string
    CAFile        string
    InsecureSkip bool
    MinVersion    uint16
    MaxVersion    uint16
    CipherSuites  []uint16
}

func (tc *TLSConfig) ConfigureTLS(config *sarama.Config) error {
    if !tc.EnableTLS {
        return nil
    }
    
    tlsConfig := &tls.Config{
        InsecureSkipVerify: tc.InsecureSkip,
        MinVersion:         tc.MinVersion,
        MaxVersion:         tc.MaxVersion,
        CipherSuites:       tc.CipherSuites,
    }
    
    // Load client certificate
    if tc.CertFile != "" && tc.KeyFile != "" {
        cert, err := tls.LoadX509KeyPair(tc.CertFile, tc.KeyFile)
        if err != nil {
            return err
        }
        tlsConfig.Certificates = []tls.Certificate{cert}
    }
    
    // Load CA certificate
    if tc.CAFile != "" {
        caCert, err := ioutil.ReadFile(tc.CAFile)
        if err != nil {
            return err
        }
        
        caCertPool := x509.NewCertPool()
        if !caCertPool.AppendCertsFromPEM(caCert) {
            return fmt.Errorf("failed to parse CA certificate")
        }
        tlsConfig.RootCAs = caCertPool
    }
    
    config.Net.TLS.Enable = true
    config.Net.TLS.Config = tlsConfig
    
    return nil
}
```

#### **Certificate Management:**
```go
type CertificateManager struct {
    certDir    string
    keySize    int
    validFor   time.Duration
    caCert     *x509.Certificate
    caKey      *rsa.PrivateKey
    mutex      sync.RWMutex
}

func (cm *CertificateManager) GenerateCACertificate() error {
    // Generate CA private key
    caKey, err := rsa.GenerateKey(rand.Reader, cm.keySize)
    if err != nil {
        return err
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
        NotAfter:              time.Now().Add(cm.validFor),
        KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
        ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
        BasicConstraintsValid: true,
        IsCA:                  true,
    }
    
    // Create CA certificate
    caCertDER, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caKey.PublicKey, caKey)
    if err != nil {
        return err
    }
    
    // Parse CA certificate
    cm.caCert, err = x509.ParseCertificate(caCertDER)
    if err != nil {
        return err
    }
    
    cm.caKey = caKey
    
    // Save CA certificate and key
    return cm.saveCACertificate()
}
```

### 4. **Authorization and Access Control**

#### **ACL (Access Control List) Management:**
```go
type ACLManager struct {
    kafkaClient sarama.Client
    mutex       sync.RWMutex
}

type ACLRule struct {
    Principal    string
    Host         string
    Operation    string
    Permission   string
    ResourceType string
    ResourceName string
}

func (am *ACLManager) CreateACL(rule ACLRule) error {
    // Create ACL using Kafka Admin API
    admin, err := sarama.NewClusterAdmin([]string{"localhost:9092"}, nil)
    if err != nil {
        return err
    }
    defer admin.Close()
    
    acl := sarama.Acl{
        Principal:    rule.Principal,
        Host:         rule.Host,
        Operation:    sarama.AclOperation(rule.Operation),
        Permission:   sarama.AclPermission(rule.Permission),
        ResourceType: sarama.AclResourceType(rule.ResourceType),
        ResourceName: rule.ResourceName,
    }
    
    return admin.CreateACLs([]sarama.Acl{acl})
}

func (am *ACLManager) ListACLs() ([]ACLRule, error) {
    admin, err := sarama.NewClusterAdmin([]string{"localhost:9092"}, nil)
    if err != nil {
        return nil, err
    }
    defer admin.Close()
    
    acls, err := admin.ListAcls(sarama.AclFilter{})
    if err != nil {
        return nil, err
    }
    
    var rules []ACLRule
    for _, acl := range acls {
        rules = append(rules, ACLRule{
            Principal:    acl.Principal,
            Host:         acl.Host,
            Operation:    string(acl.Operation),
            Permission:   string(acl.Permission),
            ResourceType: string(acl.ResourceType),
            ResourceName: acl.ResourceName,
        })
    }
    
    return rules, nil
}
```

#### **RBAC (Role-Based Access Control):**
```go
type RBACManager struct {
    roles       map[string]*Role
    users       map[string]*User
    permissions map[string][]Permission
    mutex       sync.RWMutex
}

type Role struct {
    Name        string
    Description string
    Permissions []Permission
}

type User struct {
    Username string
    Roles    []string
    Groups   []string
}

type Permission struct {
    Resource string
    Action   string
    Effect   string // Allow, Deny
}

func (rm *RBACManager) CreateRole(name, description string, permissions []Permission) error {
    rm.mutex.Lock()
    defer rm.mutex.Unlock()
    
    role := &Role{
        Name:        name,
        Description: description,
        Permissions: permissions,
    }
    
    rm.roles[name] = role
    return nil
}

func (rm *RBACManager) AssignRoleToUser(username, roleName string) error {
    rm.mutex.Lock()
    defer rm.mutex.Unlock()
    
    if _, exists := rm.roles[roleName]; !exists {
        return fmt.Errorf("role %s does not exist", roleName)
    }
    
    if rm.users[username] == nil {
        rm.users[username] = &User{Username: username}
    }
    
    rm.users[username].Roles = append(rm.users[username].Roles, roleName)
    return nil
}

func (rm *RBACManager) CheckPermission(username, resource, action string) bool {
    rm.mutex.RLock()
    defer rm.mutex.RUnlock()
    
    user, exists := rm.users[username]
    if !exists {
        return false
    }
    
    for _, roleName := range user.Roles {
        role, exists := rm.roles[roleName]
        if !exists {
            continue
        }
        
        for _, permission := range role.Permissions {
            if permission.Resource == resource && permission.Action == action {
                return permission.Effect == "Allow"
            }
        }
    }
    
    return false
}
```

### 5. **Security Monitoring and Auditing**

#### **Security Event Logger:**
```go
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

func (sel *SecurityEventLogger) LogEvent(event SecurityEvent) {
    sel.mutex.Lock()
    defer sel.mutex.Unlock()
    
    event.Timestamp = time.Now()
    sel.events = append(sel.events, event)
    
    // Keep only last maxEvents
    if len(sel.events) > sel.maxEvents {
        sel.events = sel.events[1:]
    }
    
    // Log to external system (e.g., SIEM)
    sel.logToExternalSystem(event)
}

func (sel *SecurityEventLogger) GetEvents(filter SecurityEventFilter) []SecurityEvent {
    sel.mutex.RLock()
    defer sel.mutex.RUnlock()
    
    var filteredEvents []SecurityEvent
    for _, event := range sel.events {
        if sel.matchesFilter(event, filter) {
            filteredEvents = append(filteredEvents, event)
        }
    }
    
    return filteredEvents
}
```

#### **Security Metrics:**
```go
type SecurityMetrics struct {
    AuthenticationFailures int64
    AuthorizationFailures  int64
    TLSHandshakeFailures   int64
    SuspiciousActivities   int64
    FailedLogins          int64
    SuccessfulLogins      int64
    mutex                 sync.RWMutex
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
```

### 6. **Security Best Practices**

#### **Security Configuration:**
```go
type SecurityConfig struct {
    // Authentication
    SASLConfig SASLConfig
    OAuth2Config OAuth2Config
    
    // Encryption
    TLSConfig TLSConfig
    
    // Authorization
    ACLConfig ACLConfig
    RBACConfig RBACConfig
    
    // Monitoring
    AuditConfig AuditConfig
    MetricsConfig MetricsConfig
    
    // Compliance
    ComplianceConfig ComplianceConfig
}

type ComplianceConfig struct {
    EnableAuditLogging    bool
    EnableDataEncryption  bool
    EnableAccessLogging   bool
    RetentionPeriod       time.Duration
    ComplianceStandard    string // SOC2, GDPR, HIPAA, PCI-DSS
}
```

#### **Security Hardening:**
```go
type SecurityHardening struct {
    // Network Security
    EnableFirewall        bool
    RestrictPorts         []int
    EnableVPN             bool
    
    // System Security
    DisableUnusedServices bool
    EnableSELinux         bool
    EnableAppArmor        bool
    
    // Kafka Security
    EnableSASL            bool
    EnableSSL             bool
    EnableACLs            bool
    DisablePlaintext      bool
    
    // Monitoring
    EnableIntrusionDetection bool
    EnableLogMonitoring      bool
    EnableAnomalyDetection   bool
}
```

## 🧪 Hands-on Experiments

### Experiment 1: SASL Authentication Setup

**Goal**: Configure SASL authentication for Kafka

**Steps**:
1. Configure SASL mechanisms (PLAIN, SCRAM-SHA-256)
2. Create user credentials
3. Test authentication with different mechanisms
4. Verify security configuration

**Expected Results**:
- SASL authentication working
- User credentials validated
- Secure communication established

### Experiment 2: SSL/TLS Encryption

**Goal**: Implement SSL/TLS encryption for Kafka

**Steps**:
1. Generate certificates (CA, server, client)
2. Configure TLS on Kafka brokers
3. Configure TLS on clients
4. Test encrypted communication

**Expected Results**:
- TLS encryption working
- Certificates validated
- Encrypted data transmission

### Experiment 3: ACL Authorization

**Goal**: Set up ACL-based authorization

**Steps**:
1. Create ACL rules for different users
2. Test access control for topics and operations
3. Verify permission enforcement
4. Test role-based access

**Expected Results**:
- ACL rules enforced
- Access control working
- Permission validation successful

### Experiment 4: Security Monitoring

**Goal**: Implement security monitoring and auditing

**Steps**:
1. Set up security event logging
2. Configure audit trails
3. Implement security metrics
4. Test monitoring alerts

**Expected Results**:
- Security events logged
- Audit trails maintained
- Monitoring alerts working

## 📊 Security Metrics and Monitoring

### **Key Security Metrics:**
- **Authentication Success Rate**: Percentage of successful authentications
- **Authorization Success Rate**: Percentage of successful authorizations
- **TLS Handshake Success Rate**: Percentage of successful TLS handshakes
- **Failed Login Attempts**: Number of failed authentication attempts
- **Suspicious Activities**: Number of detected suspicious activities
- **Security Events**: Total number of security events

### **Security Monitoring Dashboard:**
```go
type SecurityDashboard struct {
    metrics        *SecurityMetrics
    eventLogger    *SecurityEventLogger
    alertManager   *SecurityAlertManager
}

func (sd *SecurityDashboard) GenerateSecurityReport() SecurityReport {
    return SecurityReport{
        Metrics:      sd.metrics.GetSecurityReport(),
        RecentEvents: sd.eventLogger.GetRecentEvents(24 * time.Hour),
        Alerts:       sd.alertManager.GetActiveAlerts(),
        Timestamp:    time.Now(),
    }
}
```

## 🎯 Key Takeaways

1. **Security is multi-layered** with authentication, authorization, and encryption
2. **SASL provides authentication** with multiple mechanisms
3. **SSL/TLS ensures encryption** for data in transit
4. **ACLs control access** to Kafka resources
5. **RBAC provides role-based** access control
6. **Monitoring and auditing** are essential for security
7. **Compliance requirements** must be considered

## 📝 Lesson 13 Assessment Questions

1. **What are the different layers of Kafka security?**
2. **What SASL mechanisms are available and which are secure?**
3. **How do you configure SSL/TLS encryption for Kafka?**
4. **What is the difference between ACLs and RBAC?**
5. **How do you implement security monitoring and auditing?**
6. **What are the best practices for Kafka security?**
7. **How do you handle certificate management?**
8. **What compliance requirements should you consider?**

---

## 🔄 Next Lesson Preview: Multi-cluster and Cross-Datacenter

**What we'll learn**:
- Multi-cluster Kafka architectures
- Cross-datacenter replication
- Disaster recovery strategies
- Cluster federation
- Data synchronization
- Network partitioning
- Global data distribution

**Hands-on experiments**:
- Set up multi-cluster Kafka
- Configure cross-datacenter replication
- Test disaster recovery procedures
- Implement cluster federation
