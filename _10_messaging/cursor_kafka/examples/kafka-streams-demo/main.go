package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("🌊 Kafka Streams Fundamentals Demo")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Real-time stream processing with Kafka...")
	fmt.Println()

	// Run different stream processing demonstrations
	demonstrateBasicStreamProcessing()
	demonstrateStatefulAggregations()
	demonstrateEventSourcing()
	demonstrateCQRS()
}

func demonstrateBasicStreamProcessing() {
	fmt.Println("🔄 Basic Stream Processing Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create stream processor
	processor := NewStreamProcessor("user-events", "processed-events")
	
	// Add transformations
	processor.Map(func(data []byte) []byte {
		var event map[string]interface{}
		if err := json.Unmarshal(data, &event); err != nil {
			return nil
		}
		
		// Transform: add processing timestamp
		event["processed_at"] = time.Now().Unix()
		event["status"] = "processed"
		
		transformed, _ := json.Marshal(event)
		return transformed
	})
	
	processor.Filter(func(data []byte) bool {
		var event map[string]interface{}
		if err := json.Unmarshal(data, &event); err != nil {
			return false
		}
		
		// Filter: only process events with user_id
		_, hasUserID := event["user_id"]
		return hasUserID
	})
	
	// Simulate stream processing
	fmt.Println("🔧 Processing stream events...")
	events := generateSampleEvents(5)
	
	for i, event := range events {
		fmt.Printf("   Processing event %d: %s\n", i+1, event["type"])
		
		// Transform event
		transformed := processor.ProcessEvent(event)
		if transformed != nil {
			fmt.Printf("   ✅ Transformed: %s\n", string(transformed))
		} else {
			fmt.Printf("   ❌ Filtered out\n")
		}
	}
	fmt.Println()
}

func demonstrateStatefulAggregations() {
	fmt.Println("📊 Stateful Aggregations Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create aggregation processor
	processor := NewAggregationProcessor(5 * time.Second)
	
	// Simulate windowed aggregations
	fmt.Println("🔧 Processing windowed aggregations...")
	
	events := []AggregationEvent{
		{UserID: "user1", Amount: 100.0, Timestamp: time.Now()},
		{UserID: "user2", Amount: 200.0, Timestamp: time.Now()},
		{UserID: "user1", Amount: 150.0, Timestamp: time.Now()},
		{UserID: "user3", Amount: 300.0, Timestamp: time.Now()},
		{UserID: "user2", Amount: 250.0, Timestamp: time.Now()},
	}
	
	for i, event := range events {
		fmt.Printf("   Processing event %d: User %s, Amount $%.2f\n", i+1, event.UserID, event.Amount)
		
		processor.ProcessEvent(event)
		
		// Show current aggregations
		aggregations := processor.GetCurrentAggregations()
		fmt.Printf("   📊 Current aggregations:\n")
		for userID, agg := range aggregations {
			fmt.Printf("      %s: Count=%d, Sum=$%.2f, Avg=$%.2f\n", 
				userID, agg["count"], agg["sum"], agg["avg"])
		}
		fmt.Println()
	}
}

func demonstrateEventSourcing() {
	fmt.Println("📝 Event Sourcing Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create event sourcing processor
	processor := NewEventSourcingProcessor()
	
	// Simulate event sourcing
	fmt.Println("🔧 Processing events for user stream...")
	
	userID := "user123"
	events := []Event{
		{ID: "evt1", Type: "UserCreated", Data: map[string]interface{}{"name": "John Doe", "email": "john@example.com"}},
		{ID: "evt2", Type: "UserUpdated", Data: map[string]interface{}{"name": "John Smith"}},
		{ID: "evt3", Type: "UserActivated", Data: map[string]interface{}{"status": "active"}},
		{ID: "evt4", Type: "UserUpdated", Data: map[string]interface{}{"email": "john.smith@example.com"}},
	}
	
	for i, event := range events {
		fmt.Printf("   Appending event %d: %s\n", i+1, event.Type)
		processor.AppendEvent(userID, event)
	}
	
	// Replay events
	fmt.Println("\n🔧 Replaying events...")
	stream := processor.GetStream(userID)
	for i, event := range stream {
		fmt.Printf("   Event %d: %s (Version %d)\n", i+1, event.Type, event.Version)
	}
	
	// Show current state
	fmt.Println("\n📊 Current state:")
	state := processor.GetCurrentState(userID)
	for key, value := range state {
		fmt.Printf("   %s: %v\n", key, value)
	}
	fmt.Println()
}

func demonstrateCQRS() {
	fmt.Println("🔄 CQRS (Command Query Responsibility Segregation) Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create CQRS processor
	processor := NewCQRSProcessor()
	
	// Simulate commands and queries
	fmt.Println("🔧 Processing commands...")
	
	commands := []Command{
		{ID: "cmd1", Type: "CreateUser", Data: map[string]interface{}{"name": "Alice", "email": "alice@example.com"}},
		{ID: "cmd2", Type: "UpdateUser", Data: map[string]interface{}{"id": "user1", "name": "Alice Smith"}},
		{ID: "cmd3", Type: "DeleteUser", Data: map[string]interface{}{"id": "user2"}},
	}
	
	for i, cmd := range commands {
		fmt.Printf("   Processing command %d: %s\n", i+1, cmd.Type)
		err := processor.HandleCommand(cmd)
		if err != nil {
			fmt.Printf("   ❌ Error: %v\n", err)
		} else {
			fmt.Printf("   ✅ Command processed successfully\n")
		}
	}
	
	// Simulate queries
	fmt.Println("\n🔧 Processing queries...")
	
	queries := []Query{
		{ID: "qry1", Type: "GetUser", Data: map[string]interface{}{"id": "user1"}},
		{ID: "qry2", Type: "ListUsers", Data: map[string]interface{}{}},
	}
	
	for i, qry := range queries {
		fmt.Printf("   Processing query %d: %s\n", i+1, qry.Type)
		result, err := processor.HandleQuery(qry)
		if err != nil {
			fmt.Printf("   ❌ Error: %v\n", err)
		} else {
			fmt.Printf("   ✅ Query result: %v\n", result)
		}
	}
	fmt.Println()
}

// Stream Processor Implementation
type StreamProcessor struct {
	inputTopic     string
	outputTopic    string
	transformations []Transform
	mutex          sync.RWMutex
}

type Transform struct {
	Type string
	Func func([]byte) []byte
}

func NewStreamProcessor(inputTopic, outputTopic string) *StreamProcessor {
	return &StreamProcessor{
		inputTopic:      inputTopic,
		outputTopic:     outputTopic,
		transformations: make([]Transform, 0),
	}
}

func (sp *StreamProcessor) Map(transform func([]byte) []byte) *StreamProcessor {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	
	sp.transformations = append(sp.transformations, Transform{
		Type: "map",
		Func: transform,
	})
	return sp
}

func (sp *StreamProcessor) Filter(predicate func([]byte) bool) *StreamProcessor {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	
	sp.transformations = append(sp.transformations, Transform{
		Type: "filter",
		Func: func(data []byte) []byte {
			if predicate(data) {
				return data
			}
			return nil
		},
	})
	return sp
}

func (sp *StreamProcessor) ProcessEvent(event map[string]interface{}) []byte {
	sp.mutex.RLock()
	defer sp.mutex.RUnlock()
	
	// Convert event to JSON
	data, err := json.Marshal(event)
	if err != nil {
		return nil
	}
	
	// Apply transformations
	for _, transform := range sp.transformations {
		data = transform.Func(data)
		if data == nil {
			return nil // Filtered out
		}
	}
	
	return data
}

// Aggregation Processor Implementation
type AggregationProcessor struct {
	windowSize    time.Duration
	aggregations  map[string]map[string]interface{}
	mutex         sync.RWMutex
}

type AggregationEvent struct {
	UserID    string
	Amount    float64
	Timestamp time.Time
}

func NewAggregationProcessor(windowSize time.Duration) *AggregationProcessor {
	return &AggregationProcessor{
		windowSize:   windowSize,
		aggregations: make(map[string]map[string]interface{}),
	}
}

func (ap *AggregationProcessor) ProcessEvent(event AggregationEvent) {
	ap.mutex.Lock()
	defer ap.mutex.Unlock()
	
	// Get or create aggregation for user
	if ap.aggregations[event.UserID] == nil {
		ap.aggregations[event.UserID] = map[string]interface{}{
			"count": 0,
			"sum":   0.0,
			"avg":   0.0,
		}
	}
	
	agg := ap.aggregations[event.UserID]
	
	// Update aggregation
	count := agg["count"].(int) + 1
	sum := agg["sum"].(float64) + event.Amount
	avg := sum / float64(count)
	
	agg["count"] = count
	agg["sum"] = sum
	agg["avg"] = avg
}

func (ap *AggregationProcessor) GetCurrentAggregations() map[string]map[string]interface{} {
	ap.mutex.RLock()
	defer ap.mutex.RUnlock()
	
	result := make(map[string]map[string]interface{})
	for userID, agg := range ap.aggregations {
		result[userID] = map[string]interface{}{
			"count": agg["count"],
			"sum":   agg["sum"],
			"avg":   agg["avg"],
		}
	}
	return result
}

// Event Sourcing Implementation
type EventSourcingProcessor struct {
	eventStore map[string][]Event
	mutex      sync.RWMutex
}

type Event struct {
	ID        string
	Type      string
	Data      map[string]interface{}
	Timestamp time.Time
	Version   int
}

func NewEventSourcingProcessor() *EventSourcingProcessor {
	return &EventSourcingProcessor{
		eventStore: make(map[string][]Event),
	}
}

func (esp *EventSourcingProcessor) AppendEvent(streamID string, event Event) {
	esp.mutex.Lock()
	defer esp.mutex.Unlock()
	
	if esp.eventStore[streamID] == nil {
		esp.eventStore[streamID] = make([]Event, 0)
	}
	
	event.Version = len(esp.eventStore[streamID]) + 1
	event.Timestamp = time.Now()
	esp.eventStore[streamID] = append(esp.eventStore[streamID], event)
}

func (esp *EventSourcingProcessor) GetStream(streamID string) []Event {
	esp.mutex.RLock()
	defer esp.mutex.RUnlock()
	
	return esp.eventStore[streamID]
}

func (esp *EventSourcingProcessor) GetCurrentState(streamID string) map[string]interface{} {
	esp.mutex.RLock()
	defer esp.mutex.RUnlock()
	
	state := make(map[string]interface{})
	events := esp.eventStore[streamID]
	
	for _, event := range events {
		// Apply event to state
		for key, value := range event.Data {
			state[key] = value
		}
	}
	
	return state
}

// CQRS Implementation
type CQRSProcessor struct {
	commandStore map[string]Command
	queryStore   map[string]Query
	userStore    map[string]map[string]interface{}
	mutex        sync.RWMutex
}

type Command struct {
	ID     string
	Type   string
	Data   map[string]interface{}
	Status string
}

type Query struct {
	ID     string
	Type   string
	Data   map[string]interface{}
	Result interface{}
}

func NewCQRSProcessor() *CQRSProcessor {
	return &CQRSProcessor{
		commandStore: make(map[string]Command),
		queryStore:   make(map[string]Query),
		userStore:    make(map[string]map[string]interface{}),
	}
}

func (cp *CQRSProcessor) HandleCommand(command Command) error {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()
	
	command.Status = "processing"
	cp.commandStore[command.ID] = command
	
	// Process command
	var result interface{}
	var err error
	
	switch command.Type {
	case "CreateUser":
		result, err = cp.createUser(command.Data)
	case "UpdateUser":
		result, err = cp.updateUser(command.Data)
	case "DeleteUser":
		result, err = cp.deleteUser(command.Data)
	default:
		err = fmt.Errorf("unknown command type: %s", command.Type)
	}
	
	if err != nil {
		command.Status = "failed"
		cp.commandStore[command.ID] = command
		return err
	}
	
	// Update query store
	query := Query{
		ID:     command.ID,
		Type:   command.Type,
		Data:   command.Data,
		Result: result,
	}
	
	cp.queryStore[command.ID] = query
	command.Status = "completed"
	cp.commandStore[command.ID] = command
	
	return nil
}

func (cp *CQRSProcessor) HandleQuery(query Query) (interface{}, error) {
	cp.mutex.RLock()
	defer cp.mutex.RUnlock()
	
	switch query.Type {
	case "GetUser":
		userID := query.Data["id"].(string)
		if user, exists := cp.userStore[userID]; exists {
			return user, nil
		}
		return nil, fmt.Errorf("user not found")
	case "ListUsers":
		return cp.userStore, nil
	default:
		return nil, fmt.Errorf("unknown query type: %s", query.Type)
	}
}

func (cp *CQRSProcessor) createUser(data map[string]interface{}) (interface{}, error) {
	userID := fmt.Sprintf("user%d", len(cp.userStore)+1)
	user := map[string]interface{}{
		"id":    userID,
		"name":  data["name"],
		"email": data["email"],
	}
	cp.userStore[userID] = user
	return user, nil
}

func (cp *CQRSProcessor) updateUser(data map[string]interface{}) (interface{}, error) {
	userID := data["id"].(string)
	if user, exists := cp.userStore[userID]; exists {
		for key, value := range data {
			if key != "id" {
				user[key] = value
			}
		}
		return user, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (cp *CQRSProcessor) deleteUser(data map[string]interface{}) (interface{}, error) {
	userID := data["id"].(string)
	if _, exists := cp.userStore[userID]; exists {
		delete(cp.userStore, userID)
		return map[string]interface{}{"deleted": userID}, nil
	}
	return nil, fmt.Errorf("user not found")
}

// Helper functions
func generateSampleEvents(count int) []map[string]interface{} {
	events := make([]map[string]interface{}, count)
	eventTypes := []string{"user_login", "user_logout", "user_action", "system_event"}
	
	for i := 0; i < count; i++ {
		events[i] = map[string]interface{}{
			"id":        fmt.Sprintf("evt_%d", i+1),
			"type":      eventTypes[rand.Intn(len(eventTypes))],
			"user_id":   fmt.Sprintf("user_%d", rand.Intn(3)+1),
			"timestamp": time.Now().Unix(),
			"data":      map[string]interface{}{"value": rand.Float64() * 100},
		}
	}
	
	return events
}
