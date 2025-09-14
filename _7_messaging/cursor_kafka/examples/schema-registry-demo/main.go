package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("📋 Schema Registry and Data Evolution Demo")
	fmt.Println(strings.Repeat("=", 55))
	fmt.Println("Managing schemas and data evolution with Kafka...")
	fmt.Println()

	// Run different schema registry demonstrations
	demonstrateSchemaRegistration()
	demonstrateSchemaEvolution()
	demonstrateDataSerialization()
	demonstrateSchemaValidation()
}

func demonstrateSchemaRegistration() {
	fmt.Println("📝 Schema Registration Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create schema registry client
	client := NewSchemaRegistryClient("http://localhost:8081")
	
	// Define schemas
	schemas := map[string]string{
		"user-events": `{
			"type": "record",
			"name": "UserEvent",
			"namespace": "com.example.events",
			"fields": [
				{"name": "user_id", "type": "string"},
				{"name": "event_type", "type": "string"},
				{"name": "timestamp", "type": "long"},
				{"name": "metadata", "type": ["null", {"type": "map", "values": "string"}], "default": null}
			]
		}`,
		"order-events": `{
			"type": "record",
			"name": "OrderEvent",
			"namespace": "com.example.events",
			"fields": [
				{"name": "order_id", "type": "string"},
				{"name": "user_id", "type": "string"},
				{"name": "amount", "type": "double"},
				{"name": "status", "type": "string"},
				{"name": "timestamp", "type": "long"}
			]
		}`,
	}
	
	// Register schemas
	fmt.Println("🔧 Registering schemas...")
	for subject, schema := range schemas {
		fmt.Printf("   Registering schema for subject: %s\n", subject)
		
		schemaInfo, err := client.RegisterSchema(subject, schema)
		if err != nil {
			fmt.Printf("   ❌ Error: %v\n", err)
		} else {
			fmt.Printf("   ✅ Schema registered - ID: %d, Version: %d\n", 
				schemaInfo.ID, schemaInfo.Version)
		}
	}
	fmt.Println()
}

func demonstrateSchemaEvolution() {
	fmt.Println("🔄 Schema Evolution Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create schema version manager
	manager := NewSchemaVersionManager()
	
	// Register migrations
	manager.RegisterMigration("user-events", Migration{
		FromVersion: 1,
		ToVersion:   2,
		Transform:   migrateUserEventV1ToV2,
	})
	
	manager.RegisterMigration("user-events", Migration{
		FromVersion: 2,
		ToVersion:   3,
		Transform:   migrateUserEventV2ToV3,
	})
	
	// Test schema evolution
	fmt.Println("🔧 Testing schema evolution...")
	
	// Version 1 data
	v1Data := map[string]interface{}{
		"user_id":    "user123",
		"event_type": "login",
		"timestamp":  time.Now().UnixMilli(),
		"metadata":   map[string]string{"ip": "192.168.1.1"},
	}
	
	fmt.Printf("   Original data (V1): %+v\n", v1Data)
	
	// Migrate to version 2
	v2Data, err := manager.MigrateData("user-events", v1Data, 1, 2)
	if err != nil {
		fmt.Printf("   ❌ Migration error: %v\n", err)
	} else {
		fmt.Printf("   Migrated data (V2): %+v\n", v2Data)
	}
	
	// Migrate to version 3
	v3Data, err := manager.MigrateData("user-events", v2Data, 2, 3)
	if err != nil {
		fmt.Printf("   ❌ Migration error: %v\n", err)
	} else {
		fmt.Printf("   Migrated data (V3): %+v\n", v3Data)
	}
	fmt.Println()
}

func demonstrateDataSerialization() {
	fmt.Println("💾 Data Serialization Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create serializer
	serializer := NewDataSerializer()
	
	// Test data
	testData := map[string]interface{}{
		"user_id":    "user456",
		"event_type": "purchase",
		"timestamp":  time.Now().UnixMilli(),
		"metadata": map[string]string{
			"product_id": "prod123",
			"category":   "electronics",
		},
	}
	
	fmt.Printf("🔧 Serializing data: %+v\n", testData)
	
	// Serialize data
	serialized, err := serializer.Serialize("user-events", testData)
	if err != nil {
		fmt.Printf("   ❌ Serialization error: %v\n", err)
	} else {
		fmt.Printf("   ✅ Serialized data length: %d bytes\n", len(serialized))
	}
	
	// Deserialize data
	deserialized, err := serializer.Deserialize("user-events", serialized)
	if err != nil {
		fmt.Printf("   ❌ Deserialization error: %v\n", err)
	} else {
		fmt.Printf("   ✅ Deserialized data: %+v\n", deserialized)
	}
	fmt.Println()
}

func demonstrateSchemaValidation() {
	fmt.Println("✅ Schema Validation Demonstration")
	fmt.Println(strings.Repeat("-", 45))

	// Create validator
	validator := NewSchemaValidator()
	
	// Test cases
	testCases := []struct {
		name string
		data map[string]interface{}
		valid bool
	}{
		{
			name: "Valid user event",
			data: map[string]interface{}{
				"user_id":    "user789",
				"event_type": "logout",
				"timestamp":  time.Now().UnixMilli(),
			},
			valid: true,
		},
		{
			name: "Invalid user event (missing required field)",
			data: map[string]interface{}{
				"event_type": "logout",
				"timestamp":  time.Now().UnixMilli(),
			},
			valid: false,
		},
		{
			name: "Invalid user event (wrong type)",
			data: map[string]interface{}{
				"user_id":    123, // Should be string
				"event_type": "logout",
				"timestamp":  time.Now().UnixMilli(),
			},
			valid: false,
		},
	}
	
	fmt.Println("🔧 Testing schema validation...")
	for _, tc := range testCases {
		fmt.Printf("   Testing: %s\n", tc.name)
		
		err := validator.ValidateMessage("user-events", tc.data)
		if err != nil {
			if tc.valid {
				fmt.Printf("   ❌ Unexpected validation error: %v\n", err)
			} else {
				fmt.Printf("   ✅ Expected validation error: %v\n", err)
			}
		} else {
			if tc.valid {
				fmt.Printf("   ✅ Validation passed\n")
			} else {
				fmt.Printf("   ❌ Validation should have failed\n")
			}
		}
	}
	fmt.Println()
}

// Schema Registry Client Implementation
type SchemaRegistryClient struct {
	baseURL    string
	httpClient *http.Client
	cache      map[string]*SchemaInfo
	mutex      sync.RWMutex
}

type SchemaInfo struct {
	Subject   string `json:"subject"`
	Version   int    `json:"version"`
	ID        int    `json:"id"`
	Schema    string `json:"schema"`
	Type      string `json:"type"`
}

func NewSchemaRegistryClient(baseURL string) *SchemaRegistryClient {
	return &SchemaRegistryClient{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 10 * time.Second},
		cache:      make(map[string]*SchemaInfo),
	}
}

func (src *SchemaRegistryClient) RegisterSchema(subject string, schema string) (*SchemaInfo, error) {
	// Simulate schema registration (in real implementation, this would make HTTP call)
	schemaInfo := &SchemaInfo{
		Subject: subject,
		Version: 1,
		ID:      rand.Intn(1000) + 1,
		Schema:  schema,
		Type:    "AVRO",
	}
	
	// Cache the schema
	src.mutex.Lock()
	src.cache[subject] = schemaInfo
	src.mutex.Unlock()
	
	return schemaInfo, nil
}

func (src *SchemaRegistryClient) GetSchema(subject string, version int) (*SchemaInfo, error) {
	// Check cache first
	src.mutex.RLock()
	if cached, exists := src.cache[subject]; exists {
		src.mutex.RUnlock()
		return cached, nil
	}
	src.mutex.RUnlock()
	
	// Simulate schema retrieval
	return nil, fmt.Errorf("schema not found")
}

func (src *SchemaRegistryClient) GetLatestSchema(subject string) (*SchemaInfo, error) {
	return src.GetSchema(subject, -1)
}

// Schema Version Manager Implementation
type SchemaVersionManager struct {
	migrations map[string][]Migration
	mutex      sync.RWMutex
}

type Migration struct {
	FromVersion int
	ToVersion   int
	Transform   func(interface{}) (interface{}, error)
}

func NewSchemaVersionManager() *SchemaVersionManager {
	return &SchemaVersionManager{
		migrations: make(map[string][]Migration),
	}
}

func (svm *SchemaVersionManager) RegisterMigration(subject string, migration Migration) {
	svm.mutex.Lock()
	defer svm.mutex.Unlock()
	
	if svm.migrations[subject] == nil {
		svm.migrations[subject] = make([]Migration, 0)
	}
	
	svm.migrations[subject] = append(svm.migrations[subject], migration)
}

func (svm *SchemaVersionManager) MigrateData(subject string, data interface{}, fromVersion, toVersion int) (interface{}, error) {
	svm.mutex.RLock()
	migrations := svm.migrations[subject]
	svm.mutex.RUnlock()
	
	currentData := data
	currentVersion := fromVersion
	
	for _, migration := range migrations {
		if migration.FromVersion == currentVersion && migration.ToVersion <= toVersion {
			migratedData, err := migration.Transform(currentData)
			if err != nil {
				return nil, fmt.Errorf("migration failed from version %d to %d: %w", 
					migration.FromVersion, migration.ToVersion, err)
			}
			
			currentData = migratedData
			currentVersion = migration.ToVersion
		}
	}
	
	return currentData, nil
}

// Data Serializer Implementation
type DataSerializer struct {
	schemaRegistry *SchemaRegistryClient
}

func NewDataSerializer() *DataSerializer {
	return &DataSerializer{
		schemaRegistry: NewSchemaRegistryClient("http://localhost:8081"),
	}
}

func (ds *DataSerializer) Serialize(subject string, data interface{}) ([]byte, error) {
	// Get schema
	schemaInfo, err := ds.schemaRegistry.GetLatestSchema(subject)
	if err != nil {
		return nil, err
	}
	
	// Serialize data to JSON (simplified)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	
	// Create serialized message with schema ID
	var buffer bytes.Buffer
	
	// Write schema ID (4 bytes)
	binary.Write(&buffer, binary.BigEndian, uint32(schemaInfo.ID))
	
	// Write JSON data
	buffer.Write(jsonData)
	
	return buffer.Bytes(), nil
}

func (ds *DataSerializer) Deserialize(subject string, data []byte) (interface{}, error) {
	// Read schema ID
	var schemaID uint32
	reader := bytes.NewReader(data)
	if err := binary.Read(reader, binary.BigEndian, &schemaID); err != nil {
		return nil, err
	}
	
	// Get schema by ID (simplified)
	_, err := ds.schemaRegistry.GetLatestSchema(subject)
	if err != nil {
		return nil, err
	}
	
	// Read JSON data
	jsonData := make([]byte, reader.Len())
	reader.Read(jsonData)
	
	// Deserialize JSON
	var result interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}
	
	return result, nil
}

// Schema Validator Implementation
type SchemaValidator struct {
	schemaRegistry *SchemaRegistryClient
}

func NewSchemaValidator() *SchemaValidator {
	return &SchemaValidator{
		schemaRegistry: NewSchemaRegistryClient("http://localhost:8081"),
	}
}

func (sv *SchemaValidator) ValidateMessage(subject string, data interface{}) error {
	// Get schema
	schemaInfo, err := sv.schemaRegistry.GetLatestSchema(subject)
	if err != nil {
		return err
	}
	
	// Parse schema (simplified validation)
	var schema map[string]interface{}
	if err := json.Unmarshal([]byte(schemaInfo.Schema), &schema); err != nil {
		return err
	}
	
	// Validate required fields
	fields, ok := schema["fields"].([]interface{})
	if !ok {
		return fmt.Errorf("invalid schema format")
	}
	
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("data must be a map")
	}
	
	// Check required fields
	for _, field := range fields {
		fieldMap, ok := field.(map[string]interface{})
		if !ok {
			continue
		}
		
		fieldName, ok := fieldMap["name"].(string)
		if !ok {
			continue
		}
		
		// Check if field exists in data
		if _, exists := dataMap[fieldName]; !exists {
			return fmt.Errorf("missing required field: %s", fieldName)
		}
	}
	
	return nil
}

// Migration functions
func migrateUserEventV1ToV2(data interface{}) (interface{}, error) {
	event, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data type for migration")
	}
	
	// Add new field with default value
	event["session_id"] = "unknown"
	
	// Rename field
	if oldValue, exists := event["user_id"]; exists {
		event["user_identifier"] = oldValue
		delete(event, "user_id")
	}
	
	return event, nil
}

func migrateUserEventV2ToV3(data interface{}) (interface{}, error) {
	event, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data type for migration")
	}
	
	// Add new optional field
	event["device_info"] = map[string]interface{}{
		"platform": "unknown",
		"version":  "unknown",
	}
	
	return event, nil
}

// Schema Registry Metrics
type SchemaRegistryMetrics struct {
	SchemasRegistered  int64
	SchemasRetrieved   int64
	ValidationFailures int64
	MigrationFailures  int64
	CacheHits          int64
	CacheMisses        int64
}

func (srm *SchemaRegistryMetrics) RecordSchemaRegistration() {
	atomic.AddInt64(&srm.SchemasRegistered, 1)
}

func (srm *SchemaRegistryMetrics) RecordSchemaRetrieval() {
	atomic.AddInt64(&srm.SchemasRetrieved, 1)
}

func (srm *SchemaRegistryMetrics) RecordValidationFailure() {
	atomic.AddInt64(&srm.ValidationFailures, 1)
}

func (srm *SchemaRegistryMetrics) RecordMigrationFailure() {
	atomic.AddInt64(&srm.MigrationFailures, 1)
}

func (srm *SchemaRegistryMetrics) RecordCacheHit() {
	atomic.AddInt64(&srm.CacheHits, 1)
}

func (srm *SchemaRegistryMetrics) RecordCacheMiss() {
	atomic.AddInt64(&srm.CacheMisses, 1)
}
