package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// CircuitBreaker represents a circuit breaker pattern
type CircuitBreaker struct {
	Name          string
	MaxRequests   int32
	Interval      time.Duration
	Timeout       time.Duration
	FailureCount  int32
	SuccessCount  int32
	LastFailTime  time.Time
	State         string // "closed", "open", "half-open"
	mutex         sync.RWMutex
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(name string, maxRequests int32, interval, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		Name:        name,
		MaxRequests: maxRequests,
		Interval:    interval,
		Timeout:     timeout,
		State:       "closed",
	}
}

// Call executes a function with circuit breaker protection
func (cb *CircuitBreaker) Call(fn func() (interface{}, error)) (interface{}, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	// Check if circuit should be opened
	if cb.State == "closed" && cb.FailureCount >= cb.MaxRequests {
		if time.Since(cb.LastFailTime) > cb.Interval {
			cb.State = "open"
			cb.FailureCount = 0
		}
	}
	
	// Check if circuit should be half-opened
	if cb.State == "open" && time.Since(cb.LastFailTime) > cb.Timeout {
		cb.State = "half-open"
		cb.SuccessCount = 0
	}
	
	// Execute function based on state
	switch cb.State {
	case "closed", "half-open":
		result, err := fn()
		if err != nil {
			cb.FailureCount++
			cb.LastFailTime = time.Now()
			if cb.State == "half-open" {
				cb.State = "open"
			}
			return nil, err
		}
		
		cb.SuccessCount++
		if cb.State == "half-open" && cb.SuccessCount >= cb.MaxRequests {
			cb.State = "closed"
			cb.FailureCount = 0
		}
		
		return result, nil
	case "open":
		return nil, fmt.Errorf("circuit breaker is open")
	default:
		return nil, fmt.Errorf("unknown circuit breaker state")
	}
}

// GetState returns the current state of the circuit breaker
func (cb *CircuitBreaker) GetState() string {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.State
}

// Bulkhead represents a bulkhead pattern for resource isolation
type Bulkhead struct {
	Name        string
	MaxWorkers  int
	MaxQueue    int
	Workers     chan struct{}
	Queue       chan func()
	IsRunning   bool
	StopChan    chan struct{}
	mutex       sync.RWMutex
}

// NewBulkhead creates a new bulkhead
func NewBulkhead(name string, maxWorkers, maxQueue int) *Bulkhead {
	return &Bulkhead{
		Name:     name,
		MaxWorkers: maxWorkers,
		MaxQueue:   maxQueue,
		Workers:   make(chan struct{}, maxWorkers),
		Queue:     make(chan func(), maxQueue),
		StopChan:  make(chan struct{}),
	}
}

// Start starts the bulkhead workers
func (b *Bulkhead) Start() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	
	if b.IsRunning {
		return
	}
	
	b.IsRunning = true
	
	// Start workers
	for i := 0; i < b.MaxWorkers; i++ {
		go b.worker()
	}
}

// Stop stops the bulkhead workers
func (b *Bulkhead) Stop() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	
	if !b.IsRunning {
		return
	}
	
	b.IsRunning = false
	close(b.StopChan)
}

// Execute executes a function in the bulkhead
func (b *Bulkhead) Execute(fn func()) error {
	select {
	case b.Queue <- fn:
		return nil
	default:
		return fmt.Errorf("bulkhead queue is full")
	}
}

// worker processes functions from the queue
func (b *Bulkhead) worker() {
	for {
		select {
		case <-b.StopChan:
			return
		case fn := <-b.Queue:
			b.Workers <- struct{}{} // Acquire worker
			fn()
			<-b.Workers // Release worker
		}
	}
}

// GetStats returns bulkhead statistics
func (b *Bulkhead) GetStats() map[string]interface{} {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	
	return map[string]interface{}{
		"name":         b.Name,
		"max_workers":  b.MaxWorkers,
		"max_queue":    b.MaxQueue,
		"active_workers": len(b.Workers),
		"queue_length":  len(b.Queue),
		"is_running":   b.IsRunning,
	}
}

// Saga represents a saga pattern for distributed transactions
type Saga struct {
	ID          string
	Steps       []SagaStep
	CurrentStep int
	State       string // "running", "completed", "failed", "compensating"
	Result      interface{}
	Error       error
	mutex       sync.RWMutex
}

// SagaStep represents a step in a saga
type SagaStep struct {
	ID           string
	Action       func() (interface{}, error)
	Compensation func() error
	Timeout      time.Duration
}

// NewSaga creates a new saga
func NewSaga(id string) *Saga {
	return &Saga{
		ID:          id,
		Steps:       make([]SagaStep, 0),
		CurrentStep: 0,
		State:       "running",
	}
}

// AddStep adds a step to the saga
func (s *Saga) AddStep(step SagaStep) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	s.Steps = append(s.Steps, step)
}

// Execute executes the saga
func (s *Saga) Execute() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	s.State = "running"
	
	for i, step := range s.Steps {
		s.CurrentStep = i
		
		// Execute step with timeout
		ctx, cancel := context.WithTimeout(context.Background(), step.Timeout)
		done := make(chan struct{})
		var result interface{}
		var err error
		
		go func() {
			defer close(done)
			result, err = step.Action()
		}()
		
		select {
		case <-done:
			cancel()
			if err != nil {
				s.State = "failed"
				s.Error = err
				// Compensate for completed steps
				s.compensate(i - 1)
				return err
			}
		case <-ctx.Done():
			cancel()
			s.State = "failed"
			s.Error = fmt.Errorf("step %s timed out", step.ID)
			// Compensate for completed steps
			s.compensate(i - 1)
			return s.Error
		}
	}
	
	s.State = "completed"
	s.Result = result
	return nil
}

// compensate compensates for completed steps
func (s *Saga) compensate(lastStep int) {
	s.State = "compensating"
	
	for i := lastStep; i >= 0; i-- {
		if s.Steps[i].Compensation != nil {
			if err := s.Steps[i].Compensation(); err != nil {
				log.Printf("Compensation failed for step %s: %v", s.Steps[i].ID, err)
			}
		}
	}
}

// GetState returns the current state of the saga
func (s *Saga) GetState() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.State
}

// EventStore represents an event store for event sourcing
type EventStore struct {
	Events map[string][]Event
	mutex  sync.RWMutex
}

// Event represents an event in the event store
type Event struct {
	ID        string
	Type      string
	Data      interface{}
	Timestamp time.Time
	Version   int64
}

// NewEventStore creates a new event store
func NewEventStore() *EventStore {
	return &EventStore{
		Events: make(map[string][]Event),
	}
}

// Append appends an event to the event store
func (es *EventStore) Append(streamID string, eventType string, data interface{}) error {
	es.mutex.Lock()
	defer es.mutex.Unlock()
	
	event := Event{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		Type:      eventType,
		Data:      data,
		Timestamp: time.Now(),
		Version:   int64(len(es.Events[streamID])),
	}
	
	es.Events[streamID] = append(es.Events[streamID], event)
	return nil
}

// GetEvents gets events for a stream
func (es *EventStore) GetEvents(streamID string) ([]Event, error) {
	es.mutex.RLock()
	defer es.mutex.RUnlock()
	
	events, exists := es.Events[streamID]
	if !exists {
		return nil, fmt.Errorf("stream not found: %s", streamID)
	}
	
	// Return a copy
	result := make([]Event, len(events))
	copy(result, events)
	return result, nil
}

// GetEventsFromVersion gets events from a specific version
func (es *EventStore) GetEventsFromVersion(streamID string, fromVersion int64) ([]Event, error) {
	es.mutex.RLock()
	defer es.mutex.RUnlock()
	
	events, exists := es.Events[streamID]
	if !exists {
		return nil, fmt.Errorf("stream not found: %s", streamID)
	}
	
	if fromVersion >= int64(len(events)) {
		return []Event{}, nil
	}
	
	result := make([]Event, len(events)-int(fromVersion))
	copy(result, events[fromVersion:])
	return result, nil
}

// CQRS represents a CQRS pattern implementation
type CQRS struct {
	CommandStore *EventStore
	QueryStore   map[string]interface{}
	mutex        sync.RWMutex
}

// NewCQRS creates a new CQRS instance
func NewCQRS() *CQRS {
	return &CQRS{
		CommandStore: NewEventStore(),
		QueryStore:   make(map[string]interface{}),
	}
}

// ExecuteCommand executes a command
func (cqrs *CQRS) ExecuteCommand(streamID string, commandType string, data interface{}) error {
	// Append command to event store
	if err := cqrs.CommandStore.Append(streamID, commandType, data); err != nil {
		return err
	}
	
	// Update query store based on command
	cqrs.mutex.Lock()
	defer cqrs.mutex.Unlock()
	
	switch commandType {
	case "CreateUser":
		userData := data.(map[string]interface{})
		cqrs.QueryStore[streamID] = map[string]interface{}{
			"id":    streamID,
			"name":  userData["name"],
			"email": userData["email"],
			"state": "active",
		}
	case "UpdateUser":
		userData := data.(map[string]interface{})
		if existing, exists := cqrs.QueryStore[streamID]; exists {
			existingUser := existing.(map[string]interface{})
			if name, ok := userData["name"]; ok {
				existingUser["name"] = name
			}
			if email, ok := userData["email"]; ok {
				existingUser["email"] = email
			}
		}
	case "DeleteUser":
		if existing, exists := cqrs.QueryStore[streamID]; exists {
			existingUser := existing.(map[string]interface{})
			existingUser["state"] = "deleted"
		}
	}
	
	return nil
}

// Query queries the read store
func (cqrs *CQRS) Query(streamID string) (interface{}, error) {
	cqrs.mutex.RLock()
	defer cqrs.mutex.RUnlock()
	
	result, exists := cqrs.QueryStore[streamID]
	if !exists {
		return nil, fmt.Errorf("stream not found: %s", streamID)
	}
	
	return result, nil
}

// GetAllQueries gets all queries
func (cqrs *CQRS) GetAllQueries() map[string]interface{} {
	cqrs.mutex.RLock()
	defer cqrs.mutex.RUnlock()
	
	result := make(map[string]interface{})
	for k, v := range cqrs.QueryStore {
		result[k] = v
	}
	return result
}

// RetryPolicy represents a retry policy
type RetryPolicy struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
	Multiplier  float64
	Jitter      bool
}

// NewRetryPolicy creates a new retry policy
func NewRetryPolicy(maxAttempts int, baseDelay, maxDelay time.Duration, multiplier float64, jitter bool) *RetryPolicy {
	return &RetryPolicy{
		MaxAttempts: maxAttempts,
		BaseDelay:   baseDelay,
		MaxDelay:    maxDelay,
		Multiplier:  multiplier,
		Jitter:      jitter,
	}
}

// Execute executes a function with retry logic
func (rp *RetryPolicy) Execute(fn func() (interface{}, error)) (interface{}, error) {
	var lastErr error
	delay := rp.BaseDelay
	
	for attempt := 0; attempt < rp.MaxAttempts; attempt++ {
		result, err := fn()
		if err == nil {
			return result, nil
		}
		
		lastErr = err
		
		if attempt == rp.MaxAttempts-1 {
			break
		}
		
		// Calculate delay
		actualDelay := delay
		if rp.Jitter {
			actualDelay = time.Duration(float64(delay) * (0.5 + rand.Float64()*0.5))
		}
		
		time.Sleep(actualDelay)
		
		// Calculate next delay
		delay = time.Duration(float64(delay) * rp.Multiplier)
		if delay > rp.MaxDelay {
			delay = rp.MaxDelay
		}
	}
	
	return nil, fmt.Errorf("max attempts reached: %v", lastErr)
}

// TimeoutPolicy represents a timeout policy
type TimeoutPolicy struct {
	Timeout time.Duration
}

// NewTimeoutPolicy creates a new timeout policy
func NewTimeoutPolicy(timeout time.Duration) *TimeoutPolicy {
	return &TimeoutPolicy{
		Timeout: timeout,
	}
}

// Execute executes a function with timeout
func (tp *TimeoutPolicy) Execute(fn func() (interface{}, error)) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), tp.Timeout)
	defer cancel()
	
	done := make(chan struct{})
	var result interface{}
	var err error
	
	go func() {
		defer close(done)
		result, err = fn()
	}()
	
	select {
	case <-done:
		return result, err
	case <-ctx.Done():
		return nil, fmt.Errorf("operation timed out after %v", tp.Timeout)
	}
}

// Demonstrate circuit breaker pattern
func demonstrateCircuitBreaker() {
	fmt.Println("=== Circuit Breaker Pattern Demo ===\n")
	
	// Create circuit breaker
	cb := NewCircuitBreaker("api-service", 3, 5*time.Second, 10*time.Second)
	
	// Simulate API calls
	fmt.Println("Testing Circuit Breaker:")
	for i := 0; i < 10; i++ {
		result, err := cb.Call(func() (interface{}, error) {
			// Simulate API call with 70% failure rate
			if rand.Float64() < 0.7 {
				return nil, fmt.Errorf("API call failed")
			}
			return fmt.Sprintf("Success %d", i), nil
		})
		
		if err != nil {
			fmt.Printf("  Call %d: Failed - %v (State: %s)\n", i+1, err, cb.GetState())
		} else {
			fmt.Printf("  Call %d: Success - %v (State: %s)\n", i+1, result, cb.GetState())
		}
		
		time.Sleep(500 * time.Millisecond)
	}
}

// Demonstrate bulkhead pattern
func demonstrateBulkhead() {
	fmt.Println("=== Bulkhead Pattern Demo ===\n")
	
	// Create bulkhead
	bh := NewBulkhead("database", 3, 10)
	bh.Start()
	defer bh.Stop()
	
	// Simulate work
	fmt.Println("Testing Bulkhead:")
	for i := 0; i < 15; i++ {
		err := bh.Execute(func() {
			fmt.Printf("  Executing task %d\n", i+1)
			time.Sleep(100 * time.Millisecond)
		})
		
		if err != nil {
			fmt.Printf("  Task %d: Failed - %v\n", i+1, err)
		}
	}
	
	// Display stats
	stats := bh.GetStats()
	fmt.Printf("\nBulkhead Stats: %+v\n", stats)
}

// Demonstrate saga pattern
func demonstrateSaga() {
	fmt.Println("=== Saga Pattern Demo ===\n")
	
	// Create saga
	saga := NewSaga("order-processing")
	
	// Add steps
	saga.AddStep(SagaStep{
		ID: "reserve-inventory",
		Action: func() (interface{}, error) {
			fmt.Println("  Reserving inventory...")
			time.Sleep(100 * time.Millisecond)
			if rand.Float64() < 0.3 {
				return nil, fmt.Errorf("inventory unavailable")
			}
			return "inventory-reserved", nil
		},
		Compensation: func() error {
			fmt.Println("  Releasing inventory...")
			return nil
		},
		Timeout: 5 * time.Second,
	})
	
	saga.AddStep(SagaStep{
		ID: "charge-payment",
		Action: func() (interface{}, error) {
			fmt.Println("  Charging payment...")
			time.Sleep(100 * time.Millisecond)
			if rand.Float64() < 0.2 {
				return nil, fmt.Errorf("payment failed")
			}
			return "payment-charged", nil
		},
		Compensation: func() error {
			fmt.Println("  Refunding payment...")
			return nil
		},
		Timeout: 5 * time.Second,
	})
	
	saga.AddStep(SagaStep{
		ID: "create-shipment",
		Action: func() (interface{}, error) {
			fmt.Println("  Creating shipment...")
			time.Sleep(100 * time.Millisecond)
			return "shipment-created", nil
		},
		Compensation: func() error {
			fmt.Println("  Canceling shipment...")
			return nil
		},
		Timeout: 5 * time.Second,
	})
	
	// Execute saga
	fmt.Println("Executing Saga:")
	err := saga.Execute()
	if err != nil {
		fmt.Printf("Saga failed: %v\n", err)
	} else {
		fmt.Printf("Saga completed successfully: %v\n", saga.Result)
	}
	
	fmt.Printf("Final state: %s\n", saga.GetState())
}

// Demonstrate event sourcing
func demonstrateEventSourcing() {
	fmt.Println("=== Event Sourcing Demo ===\n")
	
	// Create event store
	eventStore := NewEventStore()
	
	// Append events
	fmt.Println("Appending Events:")
	eventStore.Append("user-123", "UserCreated", map[string]interface{}{
		"name":  "John Doe",
		"email": "john@example.com",
	})
	
	eventStore.Append("user-123", "UserUpdated", map[string]interface{}{
		"name": "John Smith",
	})
	
	eventStore.Append("user-123", "UserDeleted", map[string]interface{}{
		"reason": "account closure",
	})
	
	// Get all events
	events, err := eventStore.GetEvents("user-123")
	if err != nil {
		fmt.Printf("Error getting events: %v\n", err)
		return
	}
	
	fmt.Printf("Events for user-123:\n")
	for _, event := range events {
		fmt.Printf("  %s: %+v (version: %d)\n", event.Type, event.Data, event.Version)
	}
	
	// Get events from version 1
	eventsFromV1, err := eventStore.GetEventsFromVersion("user-123", 1)
	if err != nil {
		fmt.Printf("Error getting events from version 1: %v\n", err)
		return
	}
	
	fmt.Printf("\nEvents from version 1:\n")
	for _, event := range eventsFromV1 {
		fmt.Printf("  %s: %+v (version: %d)\n", event.Type, event.Data, event.Version)
	}
}

// Demonstrate CQRS pattern
func demonstrateCQRS() {
	fmt.Println("=== CQRS Pattern Demo ===\n")
	
	// Create CQRS instance
	cqrs := NewCQRS()
	
	// Execute commands
	fmt.Println("Executing Commands:")
	cqrs.ExecuteCommand("user-123", "CreateUser", map[string]interface{}{
		"name":  "John Doe",
		"email": "john@example.com",
	})
	
	cqrs.ExecuteCommand("user-123", "UpdateUser", map[string]interface{}{
		"name": "John Smith",
	})
	
	cqrs.ExecuteCommand("user-456", "CreateUser", map[string]interface{}{
		"name":  "Jane Doe",
		"email": "jane@example.com",
	})
	
	// Query data
	fmt.Println("\nQuerying Data:")
	user123, err := cqrs.Query("user-123")
	if err != nil {
		fmt.Printf("Error querying user-123: %v\n", err)
	} else {
		fmt.Printf("  user-123: %+v\n", user123)
	}
	
	user456, err := cqrs.Query("user-456")
	if err != nil {
		fmt.Printf("Error querying user-456: %v\n", err)
	} else {
		fmt.Printf("  user-456: %+v\n", user456)
	}
	
	// Get all queries
	allQueries := cqrs.GetAllQueries()
	fmt.Printf("\nAll Queries: %+v\n", allQueries)
}

// Demonstrate retry policy
func demonstrateRetryPolicy() {
	fmt.Println("=== Retry Policy Demo ===\n")
	
	// Create retry policy
	retryPolicy := NewRetryPolicy(3, 100*time.Millisecond, 1*time.Second, 2.0, true)
	
	// Test retry
	fmt.Println("Testing Retry Policy:")
	result, err := retryPolicy.Execute(func() (interface{}, error) {
		// Simulate operation with 80% failure rate
		if rand.Float64() < 0.8 {
			return nil, fmt.Errorf("operation failed")
		}
		return "operation succeeded", nil
	})
	
	if err != nil {
		fmt.Printf("  Final result: Failed - %v\n", err)
	} else {
		fmt.Printf("  Final result: Success - %v\n", result)
	}
}

// Demonstrate timeout policy
func demonstrateTimeoutPolicy() {
	fmt.Println("=== Timeout Policy Demo ===\n")
	
	// Create timeout policy
	timeoutPolicy := NewTimeoutPolicy(2 * time.Second)
	
	// Test timeout
	fmt.Println("Testing Timeout Policy:")
	result, err := timeoutPolicy.Execute(func() (interface{}, error) {
		// Simulate long-running operation
		time.Sleep(3 * time.Second)
		return "operation completed", nil
	})
	
	if err != nil {
		fmt.Printf("  Result: Failed - %v\n", err)
	} else {
		fmt.Printf("  Result: Success - %v\n", result)
	}
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 11: Distributed Systems Patterns")
	fmt.Println("========================================================================\n")
	
	// Run all demonstrations
	demonstrateCircuitBreaker()
	fmt.Println()
	demonstrateBulkhead()
	fmt.Println()
	demonstrateSaga()
	fmt.Println()
	demonstrateEventSourcing()
	fmt.Println()
	demonstrateCQRS()
	fmt.Println()
	demonstrateRetryPolicy()
	fmt.Println()
	demonstrateTimeoutPolicy()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Circuit breaker prevents cascading failures")
	fmt.Println("2. Bulkhead isolates resources and limits failures")
	fmt.Println("3. Saga manages distributed transactions with compensation")
	fmt.Println("4. Event sourcing stores events instead of state")
	fmt.Println("5. CQRS separates command and query responsibilities")
	fmt.Println("6. Retry policies handle transient failures")
	fmt.Println("7. Timeout policies prevent hanging operations")
	fmt.Println("8. Go provides excellent tools for implementing patterns")
	
	fmt.Println("\n📚 Next Topic: Cloud Networking & Containerization")
}
