# Topic 8: Network Security & Encryption - Secure Communication Mastery

## Network Security Overview

Network security is the practice of protecting networks and data from unauthorized access, misuse, modification, or destruction. It encompasses multiple layers of security controls and technologies.

### Security Principles
- **Confidentiality**: Data is only accessible to authorized users
- **Integrity**: Data has not been modified or corrupted
- **Availability**: Data and services are available when needed
- **Authentication**: Verifying the identity of users/systems
- **Authorization**: Controlling access to resources
- **Non-repudiation**: Preventing denial of actions

### Security Threats
- **Eavesdropping**: Unauthorized listening to communications
- **Man-in-the-middle**: Intercepting and modifying communications
- **Spoofing**: Impersonating legitimate entities
- **Denial of Service**: Overwhelming systems to make them unavailable
- **Data tampering**: Modifying data in transit or storage
- **Replay attacks**: Replaying captured communications

## Cryptography Fundamentals

### Symmetric Encryption
- **Same key**: Encryption and decryption use the same key
- **Fast**: Efficient for large amounts of data
- **Key distribution**: Challenge of sharing keys securely
- **Examples**: AES, DES, 3DES, ChaCha20

### Asymmetric Encryption
- **Key pairs**: Public and private keys
- **Public key**: Can be shared openly
- **Private key**: Must be kept secret
- **Examples**: RSA, ECC, DSA

### Hash Functions
- **One-way**: Cannot be reversed
- **Fixed length**: Output is always the same size
- **Collision resistant**: Hard to find two inputs with same output
- **Examples**: SHA-256, SHA-3, BLAKE2

### Digital Signatures
- **Authentication**: Proves message came from sender
- **Integrity**: Proves message hasn't been modified
- **Non-repudiation**: Sender cannot deny sending
- **Process**: Hash message, encrypt hash with private key

## TLS/SSL Protocol

### TLS Overview
Transport Layer Security (TLS) is a cryptographic protocol that provides secure communication over a network. It's the successor to SSL (Secure Sockets Layer).

### TLS Versions
- **SSL 1.0**: Never released publicly
- **SSL 2.0**: Deprecated, security vulnerabilities
- **SSL 3.0**: Deprecated, POODLE vulnerability
- **TLS 1.0**: Deprecated, security issues
- **TLS 1.1**: Deprecated, CBC attacks
- **TLS 1.2**: Widely used, secure
- **TLS 1.3**: Latest version, improved security and performance

### TLS Handshake Process
1. **Client Hello**: Supported cipher suites, random data
2. **Server Hello**: Chosen cipher suite, server certificate
3. **Certificate Verification**: Client validates server certificate
4. **Key Exchange**: Generate shared secret
5. **Finished**: Handshake complete, encrypted communication

### TLS Record Protocol
- **Fragmentation**: Split data into manageable chunks
- **Compression**: Optional data compression
- **Encryption**: Encrypt data with agreed cipher
- **MAC**: Message authentication code for integrity

## Certificate Management

### X.509 Certificates
- **Standard format**: Widely used certificate format
- **Contains**: Public key, identity, validity period, signature
- **Issued by**: Certificate Authority (CA)
- **Chain of trust**: Root CA → Intermediate CA → End entity

### Certificate Authority (CA)
- **Trusted entity**: Issues and validates certificates
- **Root CA**: Self-signed, trusted by default
- **Intermediate CA**: Signed by root CA
- **Public CA**: Commercial certificate providers
- **Private CA**: Internal certificate authority

### Certificate Validation
- **Chain validation**: Verify certificate chain
- **Expiration check**: Certificate not expired
- **Revocation check**: Certificate not revoked (CRL/OCSP)
- **Hostname verification**: Certificate matches hostname
- **Purpose validation**: Certificate used for intended purpose

## Authentication Methods

### Password-based Authentication
- **Username/password**: Most common method
- **Hashed passwords**: Store password hashes, not plaintext
- **Salt**: Random data added to password before hashing
- **Strong passwords**: Complexity requirements, length

### Multi-Factor Authentication (MFA)
- **Something you know**: Password, PIN
- **Something you have**: Token, smart card, phone
- **Something you are**: Biometric (fingerprint, face)
- **Time-based**: TOTP (Time-based One-Time Password)

### Certificate-based Authentication
- **Client certificates**: X.509 certificates for clients
- **Mutual TLS**: Both client and server authenticate
- **PKI**: Public Key Infrastructure for certificate management
- **Smart cards**: Hardware tokens with certificates

### Token-based Authentication
- **JWT**: JSON Web Tokens
- **OAuth**: Authorization framework
- **SAML**: Security Assertion Markup Language
- **API keys**: Simple authentication tokens

## Network Security Protocols

### IPsec (IP Security)
- **Network layer**: Security at IP layer
- **AH**: Authentication Header for integrity
- **ESP**: Encapsulating Security Payload for confidentiality
- **IKE**: Internet Key Exchange for key management
- **Tunnel mode**: Encrypts entire IP packet
- **Transport mode**: Encrypts only payload

### VPN (Virtual Private Network)
- **Secure tunnel**: Encrypted connection over public network
- **Remote access**: Connect to corporate network remotely
- **Site-to-site**: Connect multiple networks securely
- **Types**: IPsec VPN, SSL VPN, OpenVPN, WireGuard

### SSH (Secure Shell)
- **Remote access**: Secure command-line access
- **File transfer**: SCP, SFTP for secure file transfer
- **Port forwarding**: Secure tunneling of other protocols
- **Key-based auth**: Public/private key authentication

### DNSSEC (DNS Security Extensions)
- **DNS integrity**: Prevents DNS spoofing
- **Digital signatures**: Signs DNS records
- **Chain of trust**: Validates DNS responses
- **Prevents**: Cache poisoning, man-in-the-middle

## Application Security

### HTTPS (HTTP Secure)
- **TLS over HTTP**: Encrypts HTTP communication
- **Port 443**: Default HTTPS port
- **Certificate validation**: Verify server identity
- **HSTS**: HTTP Strict Transport Security

### Secure WebSocket (WSS)
- **TLS over WebSocket**: Encrypts WebSocket communication
- **wss://**: Secure WebSocket protocol
- **Same security**: Benefits as HTTPS
- **Real-time security**: Secure real-time communication

### Email Security
- **S/MIME**: Secure/Multipurpose Internet Mail Extensions
- **PGP**: Pretty Good Privacy
- **DKIM**: DomainKeys Identified Mail
- **SPF**: Sender Policy Framework
- **DMARC**: Domain-based Message Authentication

## Security Best Practices

### Secure Coding
- **Input validation**: Validate all user input
- **Output encoding**: Encode output to prevent XSS
- **SQL injection**: Use parameterized queries
- **Buffer overflow**: Bounds checking, safe functions
- **Memory management**: Proper allocation/deallocation

### Network Security
- **Firewalls**: Control network traffic
- **Intrusion Detection**: Monitor for attacks
- **Network segmentation**: Isolate network segments
- **Access control**: Limit network access
- **Monitoring**: Log and monitor network activity

### Key Management
- **Key generation**: Use cryptographically secure random
- **Key storage**: Secure storage of private keys
- **Key rotation**: Regularly change keys
- **Key escrow**: Backup and recovery procedures
- **Key destruction**: Secure deletion of old keys

### Incident Response
- **Detection**: Identify security incidents
- **Containment**: Limit damage and spread
- **Eradication**: Remove threat and vulnerabilities
- **Recovery**: Restore normal operations
- **Lessons learned**: Improve security posture

## Go Security Implementation

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

### Certificate Generation
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

### Password Hashing
```go
func hashPassword(password string) (string, error) {
    salt := make([]byte, 32)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }
    
    hash := sha256.Sum256(append([]byte(password), salt...))
    return base64.StdEncoding.EncodeToString(append(hash[:], salt...)), nil
}

func verifyPassword(password, hashedPassword string) bool {
    data, err := base64.StdEncoding.DecodeString(hashedPassword)
    if err != nil {
        return false
    }
    
    if len(data) < 32 {
        return false
    }
    
    hash := data[:32]
    salt := data[32:]
    
    expectedHash := sha256.Sum256(append([]byte(password), salt...))
    return bytes.Equal(hash, expectedHash[:])
}
```

## Security Monitoring

### Logging
- **Security events**: Authentication, authorization, errors
- **Network activity**: Connections, traffic patterns
- **System events**: File access, process execution
- **Application logs**: Error logs, access logs

### Intrusion Detection
- **Signature-based**: Known attack patterns
- **Anomaly-based**: Unusual behavior detection
- **Network-based**: Monitor network traffic
- **Host-based**: Monitor individual systems

### Vulnerability Management
- **Vulnerability scanning**: Regular security scans
- **Patch management**: Keep systems updated
- **Risk assessment**: Evaluate security risks
- **Remediation**: Fix identified vulnerabilities

## Compliance and Standards

### Security Standards
- **ISO 27001**: Information security management
- **NIST**: National Institute of Standards and Technology
- **PCI DSS**: Payment Card Industry Data Security Standard
- **HIPAA**: Health Insurance Portability and Accountability Act
- **GDPR**: General Data Protection Regulation

### Security Frameworks
- **OWASP**: Open Web Application Security Project
- **CIS Controls**: Center for Internet Security Controls
- **NIST Cybersecurity Framework**: Risk management framework
- **SANS**: Security training and resources

## Common Security Vulnerabilities

### OWASP Top 10
1. **Injection**: SQL, NoSQL, OS command injection
2. **Broken Authentication**: Weak authentication mechanisms
3. **Sensitive Data Exposure**: Inadequate protection of sensitive data
4. **XML External Entities**: XXE attacks
5. **Broken Access Control**: Inadequate access controls
6. **Security Misconfiguration**: Insecure default configurations
7. **Cross-Site Scripting**: XSS attacks
8. **Insecure Deserialization**: Unsafe deserialization
9. **Known Vulnerabilities**: Using components with known vulnerabilities
10. **Insufficient Logging**: Inadequate logging and monitoring

### Network Vulnerabilities
- **Man-in-the-middle**: Intercepting communications
- **DNS spoofing**: Redirecting traffic to malicious sites
- **ARP spoofing**: Redirecting traffic at layer 2
- **Port scanning**: Discovering open ports and services
- **Denial of Service**: Overwhelming systems

## Interview Questions

### Basic Questions
1. What is the difference between symmetric and asymmetric encryption?
2. Explain the TLS handshake process.
3. What are the main security principles?

### Intermediate Questions
1. How do you implement secure authentication in a web application?
2. Explain certificate validation and the chain of trust.
3. What are the differences between TLS 1.2 and TLS 1.3?

### Advanced Questions
1. Design a secure communication system for a distributed application.
2. How would you implement mutual TLS authentication?
3. Explain the security implications of different encryption algorithms.

## Next Steps
After mastering network security, proceed to:
- **Topic 9**: Routing & Switching
- **Topic 10**: Network Performance & Optimization
- **Topic 11**: Distributed Systems Networking

Master network security, and you'll understand how to protect networks and data! 🚀
