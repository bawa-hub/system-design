# 📚 MICROSERVICES TERMINOLOGY - COMPLETE REFERENCE

## 🎯 Essential Terms You Must Know

### 1. Service
**Definition**: An independent business capability that can be developed, deployed, and scaled independently.

**Examples**:
- User Service (handles user management)
- Order Service (handles order processing)
- Payment Service (handles payment processing)
- Notification Service (handles notifications)

**Key Characteristics**:
- Single responsibility
- Autonomous
- Loosely coupled
- Highly cohesive

### 2. API Gateway
**Definition**: Single entry point for all client requests that routes requests to appropriate services.

**Responsibilities**:
- Request routing
- Load balancing
- Authentication and authorization
- Rate limiting and throttling
- Request/response transformation
- Logging and monitoring

**Examples**: Kong, Zuul, Ambassador, Traefik

### 3. Service Discovery
**Definition**: Mechanism for services to find and communicate with each other in a distributed system.

**Types**:
- **Client-side discovery**: Client queries service registry
- **Server-side discovery**: Load balancer queries service registry

**Examples**: Consul, etcd, Eureka, Zookeeper

### 4. Load Balancing
**Definition**: Distributing incoming requests across multiple service instances to improve performance and availability.

**Algorithms**:
- **Round-robin**: Requests distributed in circular order
- **Least connections**: Route to service with fewest active connections
- **Weighted**: Route based on assigned weights
- **IP hash**: Route based on client IP hash

**Examples**: HAProxy, NGINX, AWS ELB, Kubernetes Service

### 5. Circuit Breaker
**Definition**: Pattern to prevent cascade failures by opening the circuit when a service is failing.

**States**:
- **Closed**: Normal operation, requests pass through
- **Open**: Service is failing, requests are rejected immediately
- **Half-open**: Testing if service has recovered

**Examples**: Hystrix, Resilience4j, Go resilience libraries

### 6. Bulkhead Pattern
**Definition**: Isolating resources to prevent failures from spreading across the system.

**Examples**:
- Separate thread pools for different operations
- Isolated database connections
- Separate service instances for different clients

**Benefits**:
- Prevents cascade failures
- Improves system resilience
- Better resource utilization

### 7. Saga Pattern
**Definition**: Managing distributed transactions across multiple services without using traditional ACID transactions.

**Types**:
- **Choreography**: Each service knows what to do next
- **Orchestration**: Central coordinator manages the flow

**Example**: E-commerce order processing
1. Create order
2. Reserve inventory
3. Process payment
4. Send confirmation

### 8. Event Sourcing
**Definition**: Storing state changes as a sequence of events rather than storing current state.

**Benefits**:
- Complete audit trail
- Time travel debugging
- Event replay capability
- Better scalability

**Example**: Bank account balance
- Instead of storing balance: 1000
- Store events: +500, -200, +300

### 9. CQRS (Command Query Responsibility Segregation)
**Definition**: Separating read and write operations into different models.

**Benefits**:
- Independent scaling of read/write
- Different data models for different needs
- Better performance optimization
- Simplified complex domains

**Example**:
- Write model: Optimized for transactions
- Read model: Optimized for queries

### 10. Domain Events
**Definition**: Business events that occur in the domain and are published when something significant happens.

**Examples**:
- UserRegistered
- OrderPlaced
- PaymentProcessed
- InventoryReserved

**Benefits**:
- Loose coupling between services
- Eventual consistency
- Better scalability

## 🔧 Technical Terms

### 11. Container
**Definition**: Lightweight, portable unit that packages application code and dependencies.

**Benefits**:
- Consistent environment
- Easy deployment
- Resource isolation
- Scalability

**Examples**: Docker, Podman, containerd

### 12. Orchestration
**Definition**: Automated management of containerized applications.

**Responsibilities**:
- Container lifecycle management
- Service discovery
- Load balancing
- Health monitoring
- Scaling

**Examples**: Kubernetes, Docker Swarm, Mesos

### 13. Service Mesh
**Definition**: Infrastructure layer that handles service-to-service communication.

**Features**:
- Traffic management
- Security
- Observability
- Policy enforcement

**Examples**: Istio, Linkerd, Consul Connect

### 14. API Versioning
**Definition**: Managing different versions of APIs to maintain backward compatibility.

**Strategies**:
- **URL versioning**: /api/v1/users
- **Header versioning**: Accept: application/vnd.api+json;version=1
- **Content versioning**: Different JSON schemas

### 15. Health Check
**Definition**: Endpoint that reports the health status of a service.

**Types**:
- **Liveness**: Is the service running?
- **Readiness**: Is the service ready to accept requests?

**Example**:
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T00:00:00Z",
  "checks": {
    "database": "ok",
    "redis": "ok"
  }
}
```

## 📊 Data Terms

### 16. Database per Service
**Definition**: Each service has its own database, no shared data.

**Benefits**:
- Data isolation
- Independent scaling
- Technology diversity
- Team autonomy

**Challenges**:
- Data consistency
- Cross-service queries
- Transaction management

### 17. Eventual Consistency
**Definition**: System will become consistent over time, but not immediately.

**Example**: User updates profile, notification service eventually sees the change.

**Trade-offs**:
- Better availability
- Better performance
- Temporary inconsistency

### 18. Strong Consistency
**Definition**: All nodes see the same data at the same time.

**Example**: Bank account balance must be immediately consistent.

**Trade-offs**:
- Lower availability
- Lower performance
- Immediate consistency

### 19. CAP Theorem
**Definition**: In a distributed system, you can only guarantee two of three properties:
- **Consistency**: All nodes see the same data
- **Availability**: System remains operational
- **Partition Tolerance**: System continues despite network failures

### 20. BASE Properties
**Definition**: Alternative to ACID for distributed systems:
- **Basically Available**: System is available most of the time
- **Soft state**: System state may change over time
- **Eventual consistency**: System will become consistent

## 🔒 Security Terms

### 21. Zero Trust Architecture
**Definition**: Never trust, always verify - every request is authenticated and authorized.

**Principles**:
- Verify explicitly
- Use least privilege access
- Assume breach

### 22. mTLS (Mutual TLS)
**Definition**: Both client and server authenticate each other using certificates.

**Benefits**:
- Strong authentication
- Encrypted communication
- Service-to-service security

### 23. JWT (JSON Web Token)
**Definition**: Compact, URL-safe token for securely transmitting information between parties.

**Structure**:
- Header: Algorithm and token type
- Payload: Claims and data
- Signature: Verification

### 24. OAuth 2.0
**Definition**: Authorization framework for third-party applications.

**Flows**:
- Authorization Code
- Client Credentials
- Resource Owner Password
- Implicit

### 25. RBAC (Role-Based Access Control)
**Definition**: Access control based on user roles.

**Example**:
- Admin: Full access
- User: Limited access
- Guest: Read-only access

## 📈 Monitoring Terms

### 26. Observability
**Definition**: Ability to understand system behavior through external outputs.

**Three Pillars**:
- **Metrics**: Numerical data over time
- **Logs**: Discrete events
- **Traces**: Request flow through services

### 27. SLI (Service Level Indicator)
**Definition**: Measurable aspect of service performance.

**Examples**:
- Request latency
- Error rate
- Availability

### 28. SLO (Service Level Objective)
**Definition**: Target value for an SLI.

**Example**: 99.9% availability

### 29. SLA (Service Level Agreement)
**Definition**: Contract between service provider and consumer.

**Example**: 99.9% uptime or refund

### 30. Distributed Tracing
**Definition**: Tracking requests across multiple services.

**Benefits**:
- Request flow visualization
- Performance bottleneck identification
- Error root cause analysis

## 🚀 Deployment Terms

### 31. Blue-Green Deployment
**Definition**: Two identical production environments, switch traffic between them.

**Benefits**:
- Zero downtime
- Quick rollback
- Risk mitigation

### 32. Canary Deployment
**Definition**: Gradual rollout to a subset of users.

**Benefits**:
- Risk reduction
- Real-world testing
- Gradual validation

### 33. Rolling Update
**Definition**: Gradually replace old instances with new ones.

**Benefits**:
- Continuous availability
- Gradual deployment
- Risk mitigation

### 34. Feature Flags
**Definition**: Runtime configuration to enable/disable features.

**Benefits**:
- A/B testing
- Gradual rollouts
- Quick rollbacks

### 35. GitOps
**Definition**: Git as the single source of truth for declarative infrastructure and applications.

**Benefits**:
- Version control
- Audit trail
- Automated deployments

## 🎯 Key Takeaways

1. **Terminology is crucial** - understand these terms to communicate effectively
2. **Context matters** - same term can mean different things in different contexts
3. **Keep learning** - new terms emerge as the field evolves
4. **Practice** - use these terms in your implementations

## 🔗 Next Steps

Now that you understand the terminology, we'll move to:
1. **Architecture Patterns** - Common patterns and their use cases
2. **When to Use What** - Decision-making framework
3. **Implementation** - Building your first microservice

**Ready to continue? Let's learn about architecture patterns!** 🚀
