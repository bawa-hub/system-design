# 🧪 LESSON 1 PRACTICE EXERCISES

## 🎯 Learning Objectives
By completing these exercises, you will:
- Understand microservices fundamentals
- Practice building Go microservices
- Learn API design best practices
- Experience containerization
- Test your understanding

## 📚 Exercise 1: Theory Understanding

### Question 1.1: Microservices vs Monoliths
**Scenario**: You're building an e-commerce platform with the following requirements:
- 50 developers across 8 teams
- Need to scale user management and order processing independently
- Different teams prefer different technologies
- High availability requirements (99.9% uptime)

**Question**: Should you choose microservices or monolith? Justify your answer.

**Answer Space**:
```
Your answer here...
```

### Question 1.2: Service Boundaries
**Scenario**: You have a user management system that handles:
- User registration and authentication
- User profile management
- User preferences and settings
- User activity logging
- Email notifications

**Question**: How would you split this into microservices? What would be the service boundaries?

**Answer Space**:
```
Your answer here...
```

### Question 1.3: Data Management
**Scenario**: You have an order processing system with:
- Orders table
- Order items table
- Customer information
- Payment information
- Inventory information

**Question**: How would you handle data in a microservices architecture? What are the challenges?

**Answer Space**:
```
Your answer here...
```

## 🛠️ Exercise 2: API Design

### Task 2.1: Design User Service API
Design a complete REST API for the User Service with the following requirements:
- Create user
- Get user by ID
- Get all users (with pagination)
- Update user
- Delete user
- Search users

**Requirements**:
- Use proper HTTP methods
- Include appropriate status codes
- Design request/response schemas
- Consider error handling

**Your API Design**:
```
# Create User
POST /api/v1/users
Request Body: { ... }
Response: { ... }

# Get User
GET /api/v1/users/{id}
Response: { ... }

# Get All Users
GET /api/v1/users?page=1&limit=10&search=john
Response: { ... }

# Update User
PUT /api/v1/users/{id}
Request Body: { ... }
Response: { ... }

# Delete User
DELETE /api/v1/users/{id}
Response: { ... }

# Search Users
GET /api/v1/users/search?q=john&page=1&limit=10
Response: { ... }
```

### Task 2.2: Error Handling
Design error responses for the following scenarios:
- User not found
- Invalid input data
- Email already exists
- Database connection error
- Service unavailable

**Your Error Response Design**:
```
# User Not Found
Status: 404
Response: { ... }

# Invalid Input
Status: 400
Response: { ... }

# Email Already Exists
Status: 409
Response: { ... }

# Database Error
Status: 500
Response: { ... }

# Service Unavailable
Status: 503
Response: { ... }
```

## 🔧 Exercise 3: Implementation Practice

### Task 3.1: Test the User Service
1. Start the services using the run script
2. Test all API endpoints
3. Verify error handling
4. Check health endpoints

**Commands to Run**:
```bash
cd lessons/lesson-1/implementation
./run.sh
```

**Test Cases**:
```bash
# 1. Health Check
curl http://localhost:8080/health

# 2. Create User
curl -X POST http://localhost:8080/api/v1/users \
  -H 'Content-Type: application/json' \
  -d '{"email":"john@example.com","first_name":"John","last_name":"Doe"}'

# 3. Get All Users
curl http://localhost:8080/api/v1/users

# 4. Get User by ID (use ID from step 2)
curl http://localhost:8080/api/v1/users/{user_id}

# 5. Update User
curl -X PUT http://localhost:8080/api/v1/users/{user_id} \
  -H 'Content-Type: application/json' \
  -d '{"first_name":"Jane"}'

# 6. Search Users
curl "http://localhost:8080/api/v1/users?search=john&page=1&limit=10"

# 7. Delete User
curl -X DELETE http://localhost:8080/api/v1/users/{user_id}

# 8. Test Error Cases
curl -X POST http://localhost:8080/api/v1/users \
  -H 'Content-Type: application/json' \
  -d '{"email":"invalid-email","first_name":"","last_name":"Doe"}'
```

**Expected Results**:
- All successful operations should return appropriate success responses
- Error cases should return proper error messages
- Health check should return service status
- All responses should be valid JSON

### Task 3.2: Modify the Service
Add a new endpoint to get user statistics:
- Total number of users
- Users created in the last 24 hours
- Most common first names

**Implementation Steps**:
1. Add new model for statistics
2. Add new handler method
3. Add new route
4. Test the new endpoint

**Your Implementation**:
```go
// Add to models/user.go
type UserStats struct {
    TotalUsers     int    `json:"total_users"`
    RecentUsers    int    `json:"recent_users"`
    CommonNames    []string `json:"common_names"`
}

// Add to handlers/user_handler.go
func (h *UserHandler) GetUserStats(w http.ResponseWriter, r *http.Request) {
    // Your implementation here
}

// Add to main.go
api.HandleFunc("/users/stats", userHandler.GetUserStats).Methods("GET")
```

### Task 3.3: Add Validation
Enhance the validation in the CreateUser and UpdateUser endpoints:
- Email format validation
- Name length validation (2-50 characters)
- Required field validation
- Custom error messages

**Your Enhanced Validation**:
```go
// Add custom validation rules
type CreateUserRequest struct {
    Email     string `json:"email" validate:"required,email,min=5,max=100"`
    FirstName string `json:"first_name" validate:"required,min=2,max=50,alpha"`
    LastName  string `json:"last_name" validate:"required,min=2,max=50,alpha"`
}

// Add custom error messages
func (h *UserHandler) validateRequest(req interface{}) error {
    // Your validation logic here
}
```

## 🐳 Exercise 4: Containerization

### Task 4.1: Docker Commands
Practice using Docker commands to manage your services:

```bash
# 1. Build the user service image
docker build -t user-service ./user-service

# 2. Run the user service container
docker run -d -p 8080:8080 --name user-service user-service

# 3. Check container status
docker ps

# 4. View container logs
docker logs user-service

# 5. Execute commands in container
docker exec -it user-service sh

# 6. Stop and remove container
docker stop user-service
docker rm user-service
```

### Task 4.2: Docker Compose
Practice using Docker Compose:

```bash
# 1. Start all services
docker-compose up -d

# 2. View service status
docker-compose ps

# 3. View logs for specific service
docker-compose logs user-service

# 4. Scale a service
docker-compose up -d --scale user-service=3

# 5. Stop all services
docker-compose down

# 6. Clean up volumes
docker-compose down -v
```

## 🧪 Exercise 5: Testing

### Task 5.1: Manual Testing
Create a comprehensive test plan for the User Service:

**Test Plan**:
```
1. Functional Tests
   - Create user with valid data
   - Create user with invalid data
   - Get user by valid ID
   - Get user by invalid ID
   - Update user with valid data
   - Update user with invalid data
   - Delete existing user
   - Delete non-existing user
   - Get all users with pagination
   - Search users

2. Error Handling Tests
   - Invalid JSON
   - Missing required fields
   - Invalid email format
   - Duplicate email
   - Invalid user ID format
   - Database connection error

3. Performance Tests
   - Create 100 users
   - Get all users with large dataset
   - Concurrent requests

4. Health Check Tests
   - Service health endpoint
   - Database health check
   - Metrics endpoint
```

### Task 5.2: Load Testing
Use a tool like `hey` or `ab` to perform load testing:

```bash
# Install hey (if not installed)
go install github.com/rakyll/hey@latest

# Load test the health endpoint
hey -n 1000 -c 10 http://localhost:8080/health

# Load test the create user endpoint
hey -n 100 -c 5 -m POST -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","first_name":"Test","last_name":"User"}' \
  http://localhost:8080/api/v1/users

# Load test the get users endpoint
hey -n 1000 -c 10 http://localhost:8080/api/v1/users
```

## 📊 Exercise 6: Monitoring

### Task 6.1: Health Monitoring
Monitor the health of your services:

```bash
# Check service health
curl http://localhost:8080/health

# Check metrics
curl http://localhost:8080/metrics

# Monitor database connections
curl http://localhost:8080/metrics | grep database
```

### Task 6.2: Log Analysis
Analyze the logs to understand service behavior:

```bash
# View real-time logs
docker-compose logs -f user-service

# View logs with timestamps
docker-compose logs -t user-service

# Filter logs by level
docker-compose logs user-service | grep ERROR
```

## 🎯 Exercise 7: Troubleshooting

### Task 7.1: Common Issues
Practice troubleshooting common microservices issues:

**Issue 1**: Service won't start
- Check Docker logs
- Verify environment variables
- Check port conflicts
- Verify database connection

**Issue 2**: Database connection errors
- Check database status
- Verify connection string
- Check network connectivity
- Verify credentials

**Issue 3**: API returning errors
- Check request format
- Verify endpoint URLs
- Check service logs
- Verify database data

### Task 7.2: Debugging Techniques
Practice debugging techniques:

```bash
# 1. Check container status
docker ps -a

# 2. View container logs
docker logs user-service

# 3. Execute commands in container
docker exec -it user-service sh

# 4. Check network connectivity
docker exec -it user-service ping postgres

# 5. Check database connectivity
docker exec -it user-service psql -h postgres -U postgres -d microservices
```

## ✅ Completion Checklist

- [ ] Read all theory notes
- [ ] Understand microservices concepts
- [ ] Complete all exercises
- [ ] Test the implementation
- [ ] Practice Docker commands
- [ ] Understand error handling
- [ ] Know how to troubleshoot
- [ ] Ready for Lesson 2

## 🎯 Next Steps

After completing these exercises, you should:
1. Have a solid understanding of microservices fundamentals
2. Know how to build a basic microservice in Go
3. Understand API design best practices
4. Be comfortable with Docker and containerization
5. Know how to test and troubleshoot microservices

**Ready for Lesson 2? Let me know when you've completed these exercises!** 🚀
