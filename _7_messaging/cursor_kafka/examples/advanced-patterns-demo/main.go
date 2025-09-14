package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("🏗️ Advanced Patterns and Anti-patterns Demo")
	fmt.Println(strings.Repeat("=", 55))
	fmt.Println("Mastering advanced Kafka patterns and avoiding common pitfalls...")
	fmt.Println()

	// Run different advanced pattern demonstrations
	demonstrateSagaPattern()
	demonstrateOutboxPattern()
	demonstrateCQRSPattern()
	demonstrateAntiPatterns()
}

func demonstrateSagaPattern() {
	fmt.Println("🔄 Saga Pattern Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Create saga orchestrator
	orchestrator := NewSagaOrchestrator()

	// Define order processing saga steps
	steps := []SagaStep{
		{
			ID:   "validate-order",
			Name: "Validate Order",
			Execute: func(ctx context.Context) error {
				fmt.Println("   ✅ Validating order...")
				time.Sleep(100 * time.Millisecond)
				return nil
			},
			Compensate: func(ctx context.Context) error {
				fmt.Println("   🔄 Compensating order validation...")
				return nil
			},
			MaxRetries: 3,
		},
		{
			ID:   "reserve-inventory",
			Name: "Reserve Inventory",
			Execute: func(ctx context.Context) error {
				fmt.Println("   ✅ Reserving inventory...")
				time.Sleep(150 * time.Millisecond)
				return nil
			},
			Compensate: func(ctx context.Context) error {
				fmt.Println("   🔄 Releasing inventory reservation...")
				return nil
			},
			MaxRetries: 3,
		},
		{
			ID:   "process-payment",
			Name: "Process Payment",
			Execute: func(ctx context.Context) error {
				fmt.Println("   ✅ Processing payment...")
				time.Sleep(200 * time.Millisecond)
				// Simulate occasional failure
				if rand.Float64() < 0.3 {
					return fmt.Errorf("payment processing failed")
				}
				return nil
			},
			Compensate: func(ctx context.Context) error {
				fmt.Println("   🔄 Refunding payment...")
				return nil
			},
			MaxRetries: 3,
		},
		{
			ID:   "send-confirmation",
			Name: "Send Confirmation",
			Execute: func(ctx context.Context) error {
				fmt.Println("   ✅ Sending confirmation email...")
				time.Sleep(100 * time.Millisecond)
				return nil
			},
			Compensate: func(ctx context.Context) error {
				fmt.Println("   🔄 Sending cancellation email...")
				return nil
			},
			MaxRetries: 3,
		},
	}

	orchestrator.SetSteps(steps)

	// Execute saga
	fmt.Println("🔧 Executing order processing saga...")
	ctx := context.Background()
	
	err := orchestrator.ExecuteSaga(ctx)
	if err != nil {
		fmt.Printf("   ❌ Saga execution failed: %v\n", err)
		fmt.Println("   🔄 Compensation steps executed")
	} else {
		fmt.Println("   ✅ Saga executed successfully")
	}

	// Show saga state
	state := orchestrator.GetState()
	fmt.Printf("   📊 Saga State: %s (Steps completed: %d/%d)\n", 
		state.Status, state.CurrentStep, len(steps))
	fmt.Println()
}

func demonstrateOutboxPattern() {
	fmt.Println("📦 Outbox Pattern Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Create outbox processor
	processor := NewOutboxProcessor()

	// Simulate business operations that generate events
	fmt.Println("🔧 Simulating business operations with outbox pattern...")
	
	operations := []BusinessOperation{
		{
			ID:   "op1",
			Type: "CreateUser",
			Data: map[string]interface{}{
				"user_id": "user123",
				"email":   "user@example.com",
				"name":    "John Doe",
			},
		},
		{
			ID:   "op2",
			Type: "UpdateUser",
			Data: map[string]interface{}{
				"user_id": "user123",
				"email":   "john.doe@example.com",
			},
		},
		{
			ID:   "op3",
			Type: "CreateOrder",
			Data: map[string]interface{}{
				"order_id": "order456",
				"user_id":  "user123",
				"amount":   99.99,
			},
		},
	}

	// Process operations
	for i, op := range operations {
		fmt.Printf("   Processing operation %d: %s\n", i+1, op.Type)
		
		// Store in outbox
		event := OutboxEvent{
			ID:         fmt.Sprintf("evt_%d", i+1),
			AggregateID: op.Data["user_id"].(string),
			EventType:   op.Type,
			EventData:   []byte(fmt.Sprintf(`{"operation_id":"%s","data":%s}`, op.ID, op.Data)),
			CreatedAt:   time.Now(),
		}
		
		if err := processor.StoreEvent(event); err != nil {
			fmt.Printf("   ❌ Failed to store event: %v\n", err)
		} else {
			fmt.Printf("   ✅ Event stored in outbox\n")
		}
	}

	// Process outbox events
	fmt.Println("\n🔧 Processing outbox events...")
	processed, err := processor.ProcessEvents()
	if err != nil {
		fmt.Printf("   ❌ Failed to process events: %v\n", err)
	} else {
		fmt.Printf("   ✅ Processed %d events from outbox\n", processed)
	}
	fmt.Println()
}

func demonstrateCQRSPattern() {
	fmt.Println("🔄 CQRS Pattern Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Create CQRS manager
	cqrsManager := NewCQRSManager()

	// Register projections
	cqrsManager.RegisterProjection("UserProjection", &UserProjection{})
	cqrsManager.RegisterProjection("OrderProjection", &OrderProjection{})

	// Simulate commands
	fmt.Println("🔧 Processing commands...")
	commands := []Command{
		{
			ID:   "cmd1",
			Type: "CreateUser",
			Data: map[string]interface{}{
				"user_id": "user123",
				"email":   "user@example.com",
				"name":    "John Doe",
			},
		},
		{
			ID:   "cmd2",
			Type: "UpdateUser",
			Data: map[string]interface{}{
				"user_id": "user123",
				"email":   "john.doe@example.com",
			},
		},
		{
			ID:   "cmd3",
			Type: "CreateOrder",
			Data: map[string]interface{}{
				"order_id": "order456",
				"user_id":  "user123",
				"amount":   99.99,
			},
		},
	}

	for i, cmd := range commands {
		fmt.Printf("   Processing command %d: %s\n", i+1, cmd.Type)
		
		ctx := context.Background()
		if err := cqrsManager.HandleCommand(ctx, cmd); err != nil {
			fmt.Printf("   ❌ Command failed: %v\n", err)
		} else {
			fmt.Printf("   ✅ Command processed successfully\n")
		}
	}

	// Simulate queries
	fmt.Println("\n🔧 Processing queries...")
	queries := []Query{
		{
			ID:   "qry1",
			Type: "GetUser",
			Criteria: map[string]interface{}{
				"user_id": "user123",
			},
		},
		{
			ID:   "qry2",
			Type: "GetUserOrders",
			Criteria: map[string]interface{}{
				"user_id": "user123",
			},
		},
	}

	for i, qry := range queries {
		fmt.Printf("   Processing query %d: %s\n", i+1, qry.Type)
		
		ctx := context.Background()
		result, err := cqrsManager.HandleQuery(ctx, qry)
		if err != nil {
			fmt.Printf("   ❌ Query failed: %v\n", err)
		} else {
			fmt.Printf("   ✅ Query result: %v\n", result)
		}
	}
	fmt.Println()
}

func demonstrateAntiPatterns() {
	fmt.Println("⚠️ Anti-patterns Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Demonstrate tight coupling anti-pattern
	fmt.Println("🔧 Demonstrating tight coupling anti-pattern...")
	tightlyCoupledService := &TightlyCoupledService{}
	err := tightlyCoupledService.ProcessOrder(Order{
		ID:     "order123",
		Amount: 99.99,
		Items:  []string{"item1", "item2"},
	})
	if err != nil {
		fmt.Printf("   ❌ Tightly coupled service failed: %v\n", err)
	} else {
		fmt.Println("   ✅ Tightly coupled service completed (but fragile)")
	}

	// Demonstrate loose coupling solution
	fmt.Println("\n🔧 Demonstrating loose coupling solution...")
	looselyCoupledService := &LooselyCoupledService{}
	err = looselyCoupledService.ProcessOrder(Order{
		ID:     "order456",
		Amount: 149.99,
		Items:  []string{"item3", "item4"},
	})
	if err != nil {
		fmt.Printf("   ❌ Loosely coupled service failed: %v\n", err)
	} else {
		fmt.Println("   ✅ Loosely coupled service completed (resilient)")
	}

	// Demonstrate error handling anti-pattern
	fmt.Println("\n🔧 Demonstrating poor error handling...")
	poorErrorHandler := &PoorErrorHandler{}
	poorErrorHandler.ProcessMessages([]string{"msg1", "msg2", "msg3"})

	// Demonstrate proper error handling
	fmt.Println("\n🔧 Demonstrating proper error handling...")
	properErrorHandler := &ProperErrorHandler{}
	properErrorHandler.ProcessMessages([]string{"msg1", "msg2", "msg3"})
	fmt.Println()
}

// Saga Pattern Implementation
type SagaOrchestrator struct {
	steps []SagaStep
	state SagaState
	mutex sync.RWMutex
}

type SagaStep struct {
	ID         string
	Name       string
	Execute    func(context.Context) error
	Compensate func(context.Context) error
	Status     StepStatus
	RetryCount int
	MaxRetries int
}

type StepStatus string

const (
	STEP_PENDING   StepStatus = "PENDING"
	STEP_RUNNING   StepStatus = "RUNNING"
	STEP_COMPLETED StepStatus = "COMPLETED"
	STEP_FAILED    StepStatus = "FAILED"
)

type SagaState struct {
	SagaID      string
	Status      SagaStatus
	CurrentStep int
	Data        map[string]interface{}
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SagaStatus string

const (
	SAGA_PENDING  SagaStatus = "PENDING"
	SAGA_RUNNING  SagaStatus = "RUNNING"
	SAGA_COMPLETED SagaStatus = "COMPLETED"
	SAGA_FAILED   SagaStatus = "FAILED"
)

func NewSagaOrchestrator() *SagaOrchestrator {
	return &SagaOrchestrator{
		state: SagaState{
			SagaID:      fmt.Sprintf("saga_%d", time.Now().Unix()),
			Status:      SAGA_PENDING,
			CurrentStep: 0,
			Data:        make(map[string]interface{}),
			CreatedAt:   time.Now(),
		},
	}
}

func (so *SagaOrchestrator) SetSteps(steps []SagaStep) {
	so.mutex.Lock()
	defer so.mutex.Unlock()
	
	so.steps = steps
}

func (so *SagaOrchestrator) ExecuteSaga(ctx context.Context) error {
	so.mutex.Lock()
	defer so.mutex.Unlock()
	
	so.state.Status = SAGA_RUNNING
	so.state.UpdatedAt = time.Now()
	
	for i, step := range so.steps {
		so.state.CurrentStep = i
		step.Status = STEP_RUNNING
		
		// Execute step with retry logic
		if err := so.executeStepWithRetry(ctx, step); err != nil {
			// Compensate for completed steps
			so.compensateCompletedSteps(ctx, i)
			so.state.Status = SAGA_FAILED
			so.state.UpdatedAt = time.Now()
			return fmt.Errorf("saga execution failed at step %d: %w", i, err)
		}
		
		step.Status = STEP_COMPLETED
		so.steps[i] = step
	}
	
	so.state.Status = SAGA_COMPLETED
	so.state.UpdatedAt = time.Now()
	return nil
}

func (so *SagaOrchestrator) executeStepWithRetry(ctx context.Context, step SagaStep) error {
	for step.RetryCount <= step.MaxRetries {
		if err := step.Execute(ctx); err != nil {
			step.RetryCount++
			fmt.Printf("   ⚠️ Step %s failed, retry %d/%d: %v\n", 
				step.Name, step.RetryCount, step.MaxRetries, err)
			
			if step.RetryCount <= step.MaxRetries {
				time.Sleep(time.Duration(step.RetryCount) * time.Second)
				continue
			}
			return err
		}
		return nil
	}
	return fmt.Errorf("max retries exceeded for step %s", step.Name)
}

func (so *SagaOrchestrator) compensateCompletedSteps(ctx context.Context, failedStepIndex int) {
	fmt.Println("   🔄 Executing compensation steps...")
	for i := failedStepIndex - 1; i >= 0; i-- {
		step := so.steps[i]
		if step.Status == STEP_COMPLETED {
			fmt.Printf("   🔄 Compensating step: %s\n", step.Name)
			step.Compensate(ctx)
		}
	}
}

func (so *SagaOrchestrator) GetState() SagaState {
	so.mutex.RLock()
	defer so.mutex.RUnlock()
	return so.state
}

// Outbox Pattern Implementation
type OutboxProcessor struct {
	events []OutboxEvent
	mutex  sync.RWMutex
}

type OutboxEvent struct {
	ID          string
	AggregateID string
	EventType   string
	EventData   []byte
	CreatedAt   time.Time
	Processed   bool
	ProcessedAt *time.Time
}

type BusinessOperation struct {
	ID   string
	Type string
	Data map[string]interface{}
}

func NewOutboxProcessor() *OutboxProcessor {
	return &OutboxProcessor{
		events: make([]OutboxEvent, 0),
	}
}

func (op *OutboxProcessor) StoreEvent(event OutboxEvent) error {
	op.mutex.Lock()
	defer op.mutex.Unlock()
	
	op.events = append(op.events, event)
	return nil
}

func (op *OutboxProcessor) ProcessEvents() (int, error) {
	op.mutex.Lock()
	defer op.mutex.Unlock()
	
	processed := 0
	for i, event := range op.events {
		if event.Processed {
			continue
		}
		
		// Simulate publishing to Kafka
		fmt.Printf("   📤 Publishing event: %s (ID: %s)\n", event.EventType, event.ID)
		time.Sleep(50 * time.Millisecond) // Simulate network delay
		
		// Mark as processed
		now := time.Now()
		op.events[i].Processed = true
		op.events[i].ProcessedAt = &now
		processed++
	}
	
	return processed, nil
}

// CQRS Pattern Implementation
type CQRSManager struct {
	commandStore  map[string]Command
	queryStore    map[string]interface{}
	eventStore    []Event
	projections   map[string]Projection
	mutex         sync.RWMutex
}

type Command struct {
	ID        string
	Type      string
	Data      map[string]interface{}
	UserID    string
	Timestamp time.Time
}

type Query struct {
	ID       string
	Type     string
	Criteria map[string]interface{}
	UserID   string
}

type Event struct {
	ID        string
	Type      string
	Data      map[string]interface{}
	Timestamp time.Time
	Version   int
}

type Projection interface {
	Handle(event Event) error
	Query(criteria map[string]interface{}) (interface{}, error)
}

type UserProjection struct {
	users map[string]map[string]interface{}
}

func (up *UserProjection) Handle(event Event) error {
	if up.users == nil {
		up.users = make(map[string]map[string]interface{})
	}
	
	userID := event.Data["user_id"].(string)
	
	switch event.Type {
	case "CreateUser":
		up.users[userID] = event.Data
	case "UpdateUser":
		if user, exists := up.users[userID]; exists {
			for key, value := range event.Data {
				if key != "user_id" {
					user[key] = value
				}
			}
		}
	}
	
	return nil
}

func (up *UserProjection) Query(criteria map[string]interface{}) (interface{}, error) {
	userID := criteria["user_id"].(string)
	if user, exists := up.users[userID]; exists {
		return user, nil
	}
	return nil, fmt.Errorf("user not found")
}

type OrderProjection struct {
	orders map[string][]map[string]interface{}
}

func (op *OrderProjection) Handle(event Event) error {
	if op.orders == nil {
		op.orders = make(map[string][]map[string]interface{})
	}
	
	userID := event.Data["user_id"].(string)
	
	switch event.Type {
	case "CreateOrder":
		if op.orders[userID] == nil {
			op.orders[userID] = make([]map[string]interface{}, 0)
		}
		op.orders[userID] = append(op.orders[userID], event.Data)
	}
	
	return nil
}

func (op *OrderProjection) Query(criteria map[string]interface{}) (interface{}, error) {
	userID := criteria["user_id"].(string)
	if orders, exists := op.orders[userID]; exists {
		return orders, nil
	}
	return []map[string]interface{}{}, nil
}

func NewCQRSManager() *CQRSManager {
	return &CQRSManager{
		commandStore: make(map[string]Command),
		queryStore:   make(map[string]interface{}),
		eventStore:   make([]Event, 0),
		projections:  make(map[string]Projection),
	}
}

func (cqrs *CQRSManager) RegisterProjection(name string, projection Projection) {
	cqrs.mutex.Lock()
	defer cqrs.mutex.Unlock()
	
	cqrs.projections[name] = projection
}

func (cqrs *CQRSManager) HandleCommand(ctx context.Context, cmd Command) error {
	cqrs.mutex.Lock()
	defer cqrs.mutex.Unlock()
	
	// Store command
	cmd.Timestamp = time.Now()
	cqrs.commandStore[cmd.ID] = cmd
	
	// Create event
	event := Event{
		ID:        fmt.Sprintf("evt_%d", time.Now().UnixNano()),
		Type:      cmd.Type,
		Data:      cmd.Data,
		Timestamp: time.Now(),
		Version:   len(cqrs.eventStore) + 1,
	}
	
	// Store event
	cqrs.eventStore = append(cqrs.eventStore, event)
	
	// Update projections
	for _, projection := range cqrs.projections {
		projection.Handle(event)
	}
	
	return nil
}

func (cqrs *CQRSManager) HandleQuery(ctx context.Context, query Query) (interface{}, error) {
	cqrs.mutex.RLock()
	defer cqrs.mutex.RUnlock()
	
	// Get projection
	projection, exists := cqrs.projections[query.Type]
	if !exists {
		return nil, fmt.Errorf("projection not found: %s", query.Type)
	}
	
	// Query projection
	return projection.Query(query.Criteria)
}

// Anti-pattern Demonstrations
type Order struct {
	ID     string
	Amount float64
	Items  []string
}

type TightlyCoupledService struct{}

func (tcs *TightlyCoupledService) ProcessOrder(order Order) error {
	fmt.Println("   🔗 Tightly coupled: Direct service calls")
	
	// Simulate direct service calls
	fmt.Println("     → Calling payment service directly")
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("     → Calling inventory service directly")
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("     → Calling notification service directly")
	time.Sleep(100 * time.Millisecond)
	
	return nil
}

type LooselyCoupledService struct{}

func (lcs *LooselyCoupledService) ProcessOrder(order Order) error {
	fmt.Println("   🔗 Loosely coupled: Event-driven approach")
	
	// Simulate event publishing
	events := []string{"PaymentRequested", "InventoryReservationRequested", "OrderConfirmed"}
	for _, event := range events {
		fmt.Printf("     → Publishing event: %s\n", event)
		time.Sleep(50 * time.Millisecond)
	}
	
	return nil
}

type PoorErrorHandler struct{}

func (peh *PoorErrorHandler) ProcessMessages(messages []string) {
	fmt.Println("   ⚠️ Poor error handling: Ignoring errors")
	for _, msg := range messages {
		fmt.Printf("     Processing: %s\n", msg)
		// Simulate occasional failure
		if rand.Float64() < 0.3 {
			fmt.Printf("     ❌ Failed to process %s (error ignored)\n", msg)
		} else {
			fmt.Printf("     ✅ Processed %s\n", msg)
		}
	}
}

type ProperErrorHandler struct{}

func (peh *ProperErrorHandler) ProcessMessages(messages []string) {
	fmt.Println("   ✅ Proper error handling: Handling errors gracefully")
	for _, msg := range messages {
		fmt.Printf("     Processing: %s\n", msg)
		// Simulate occasional failure
		if rand.Float64() < 0.3 {
			fmt.Printf("     ❌ Failed to process %s\n", msg)
			fmt.Printf("     🔄 Sending to dead letter queue\n")
		} else {
			fmt.Printf("     ✅ Processed %s\n", msg)
		}
	}
}
