# Lesson 12: Schema Registry and Data Evolution

## 🎯 Learning Objectives
By the end of this lesson, you will understand:
- Schema Registry concepts and architecture
- Schema evolution strategies and compatibility
- Avro, JSON Schema, and Protobuf support
- Schema versioning and migration
- Data serialization and deserialization
- Schema validation and enforcement
- Best practices for schema management

## 📚 Theory Section

### 1. **Schema Registry Overview**

#### **What is Schema Registry?**
Schema Registry is a centralized service for storing and retrieving schemas. It provides:

- **Schema Storage**: Centralized schema repository
- **Schema Evolution**: Backward and forward compatibility
- **Schema Validation**: Runtime schema validation
- **Version Management**: Schema versioning and migration
- **Client Integration**: Easy integration with Kafka clients

#### **Key Benefits:**
- **Data Consistency**: Ensures data format consistency
- **Schema Evolution**: Safe schema changes over time
- **Type Safety**: Strong typing for data structures
- **Documentation**: Self-documenting data formats
- **Validation**: Runtime data validation

### 2. **Schema Evolution Strategies**

#### **Compatibility Types:**
```go
type CompatibilityType string

const (
    BACKWARD    CompatibilityType = "BACKWARD"    // New schema can read data written by old schema
    FORWARD     CompatibilityType = "FORWARD"     // Old schema can read data written by new schema
    FULL        CompatibilityType = "FULL"        // Both backward and forward compatible
    NONE        CompatibilityType = "NONE"        // No compatibility checks
    BACKWARD_TRANSITIVE CompatibilityType = "BACKWARD_TRANSITIVE" // Backward compatible with all previous versions
    FORWARD_TRANSITIVE  CompatibilityType = "FORWARD_TRANSITIVE"  // Forward compatible with all future versions
    FULL_TRANSITIVE     CompatibilityType = "FULL_TRANSITIVE"     // Full compatibility with all versions
)
```

#### **Schema Evolution Rules:**
```go
type SchemaEvolutionRules struct {
    // Backward Compatibility Rules
    BackwardCompatible bool
    // - Can add new fields with default values
    // - Can remove fields
    // - Cannot change field types
    // - Cannot make required fields optional
    
    // Forward Compatibility Rules
    ForwardCompatible bool
    // - Can add new optional fields
    // - Cannot remove fields
    // - Cannot change field types
    // - Cannot make optional fields required
    
    // Full Compatibility Rules
    FullCompatible bool
    // - Must satisfy both backward and forward compatibility
    // - Most restrictive but safest
}
```

### 3. **Avro Schema Management**

#### **Avro Schema Definition:**
```go
type AvroSchema struct {
    Type      string                 `json:"type"`
    Name      string                 `json:"name"`
    Namespace string                 `json:"namespace"`
    Fields    []AvroField           `json:"fields"`
    Doc       string                 `json:"doc,omitempty"`
}

type AvroField struct {
    Name    string      `json:"name"`
    Type    interface{} `json:"type"`
    Default interface{} `json:"default,omitempty"`
    Doc     string      `json:"doc,omitempty"`
}

// Example Avro Schema
var UserEventSchema = AvroSchema{
    Type:      "record",
    Name:      "UserEvent",
    Namespace: "com.example.events",
    Fields: []AvroField{
        {
            Name: "user_id",
            Type: "string",
            Doc:  "Unique identifier for the user",
        },
        {
            Name: "event_type",
            Type: "string",
            Doc:  "Type of event (login, logout, action)",
        },
        {
            Name: "timestamp",
            Type: "long",
            Doc:  "Event timestamp in milliseconds",
        },
        {
            Name: "metadata",
            Type: []interface{}{
                "null",
                map[string]interface{}{
                    "type": "map",
                    "values": "string",
                },
            },
            Default: nil,
            Doc:     "Optional metadata for the event",
        },
    },
}
```

#### **Schema Registry Client:**
```go
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

func (src *SchemaRegistryClient) RegisterSchema(subject string, schema string) (*SchemaInfo, error) {
    url := fmt.Sprintf("%s/subjects/%s/versions", src.baseURL, subject)
    
    payload := map[string]string{
        "schema": schema,
    }
    
    jsonData, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }
    
    resp, err := src.httpClient.Post(url, "application/vnd.schemaregistry.v1+json", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var schemaInfo SchemaInfo
    if err := json.NewDecoder(resp.Body).Decode(&schemaInfo); err != nil {
        return nil, err
    }
    
    return &schemaInfo, nil
}

func (src *SchemaRegistryClient) GetSchema(subject string, version int) (*SchemaInfo, error) {
    // Check cache first
    src.mutex.RLock()
    cacheKey := fmt.Sprintf("%s:%d", subject, version)
    if cached, exists := src.cache[cacheKey]; exists {
        src.mutex.RUnlock()
        return cached, nil
    }
    src.mutex.RUnlock()
    
    // Fetch from registry
    url := fmt.Sprintf("%s/subjects/%s/versions/%d", src.baseURL, subject, version)
    resp, err := src.httpClient.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var schemaInfo SchemaInfo
    if err := json.NewDecoder(resp.Body).Decode(&schemaInfo); err != nil {
        return nil, err
    }
    
    // Cache the result
    src.mutex.Lock()
    src.cache[cacheKey] = &schemaInfo
    src.mutex.Unlock()
    
    return &schemaInfo, nil
}
```

### 4. **Data Serialization and Deserialization**

#### **Avro Serialization:**
```go
type AvroSerializer struct {
    schemaRegistry *SchemaRegistryClient
    schemaCache    map[string]*SchemaInfo
    mutex          sync.RWMutex
}

func (as *AvroSerializer) Serialize(subject string, data interface{}) ([]byte, error) {
    // Get schema
    schemaInfo, err := as.getSchema(subject)
    if err != nil {
        return nil, err
    }
    
    // Parse Avro schema
    schema, err := avro.ParseSchema(schemaInfo.Schema)
    if err != nil {
        return nil, err
    }
    
    // Serialize data
    writer := avro.NewSpecificDatumWriter(schema)
    buffer := new(bytes.Buffer)
    encoder := avro.NewBinaryEncoder(buffer)
    
    // Write schema ID first (Confluent format)
    binary.Write(buffer, binary.BigEndian, uint32(schemaInfo.ID))
    
    // Write serialized data
    if err := writer.Write(data, encoder); err != nil {
        return nil, err
    }
    
    return buffer.Bytes(), nil
}

func (as *AvroSerializer) Deserialize(subject string, data []byte) (interface{}, error) {
    // Read schema ID
    var schemaID uint32
    reader := bytes.NewReader(data)
    if err := binary.Read(reader, binary.BigEndian, &schemaID); err != nil {
        return nil, err
    }
    
    // Get schema by ID
    schemaInfo, err := as.getSchemaByID(int(schemaID))
    if err != nil {
        return nil, err
    }
    
    // Parse Avro schema
    schema, err := avro.ParseSchema(schemaInfo.Schema)
    if err != nil {
        return nil, err
    }
    
    // Deserialize data
    reader = avro.NewSpecificDatumReader(schema)
    decoder := avro.NewBinaryDecoder(reader)
    
    var result interface{}
    if err := reader.Read(&result, decoder); err != nil {
        return nil, err
    }
    
    return result, nil
}
```

#### **JSON Schema Support:**
```go
type JSONSchemaSerializer struct {
    schemaRegistry *SchemaRegistryClient
    validator      *jsonschema.Schema
}

func (jss *JSONSchemaSerializer) Validate(data interface{}) error {
    if jss.validator == nil {
        return nil
    }
    
    result := jss.validator.Validate(data)
    if result.Valid() {
        return nil
    }
    
    var errors []string
    for _, err := range result.Errors() {
        errors = append(errors, err.String())
    }
    
    return fmt.Errorf("validation failed: %s", strings.Join(errors, ", "))
}

func (jss *JSONSchemaSerializer) Serialize(subject string, data interface{}) ([]byte, error) {
    // Validate data against schema
    if err := jss.Validate(data); err != nil {
        return nil, err
    }
    
    // Serialize to JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }
    
    return jsonData, nil
}
```

### 5. **Schema Versioning and Migration**

#### **Schema Version Manager:**
```go
type SchemaVersionManager struct {
    schemaRegistry *SchemaRegistryClient
    migrations     map[string][]Migration
    mutex          sync.RWMutex
}

type Migration struct {
    FromVersion int
    ToVersion   int
    Transform   func(interface{}) (interface{}, error)
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
```

#### **Data Migration Example:**
```go
// Example migration from version 1 to version 2
func migrateUserEventV1ToV2(data interface{}) (interface{}, error) {
    // Cast to map for manipulation
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

// Example migration from version 2 to version 3
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
```

### 6. **Schema Validation and Enforcement**

#### **Schema Validator:**
```go
type SchemaValidator struct {
    schemaRegistry *SchemaRegistryClient
    validators     map[string]*jsonschema.Schema
    mutex          sync.RWMutex
}

func (sv *SchemaValidator) ValidateMessage(subject string, data interface{}) error {
    // Get schema
    schemaInfo, err := sv.schemaRegistry.GetLatestSchema(subject)
    if err != nil {
        return err
    }
    
    // Get or create validator
    validator, err := sv.getValidator(subject, schemaInfo.Schema)
    if err != nil {
        return err
    }
    
    // Validate data
    result := validator.Validate(data)
    if result.Valid() {
        return nil
    }
    
    var errors []string
    for _, err := range result.Errors() {
        errors = append(errors, err.String())
    }
    
    return fmt.Errorf("schema validation failed: %s", strings.Join(errors, ", "))
}

func (sv *SchemaValidator) getValidator(subject, schema string) (*jsonschema.Schema, error) {
    sv.mutex.RLock()
    if validator, exists := sv.validators[subject]; exists {
        sv.mutex.RUnlock()
        return validator, nil
    }
    sv.mutex.RUnlock()
    
    // Create new validator
    validator, err := jsonschema.NewSchema(schema)
    if err != nil {
        return nil, err
    }
    
    // Cache validator
    sv.mutex.Lock()
    sv.validators[subject] = validator
    sv.mutex.Unlock()
    
    return validator, nil
}
```

### 7. **Best Practices for Schema Management**

#### **Schema Design Guidelines:**
```go
type SchemaDesignGuidelines struct {
    // Naming Conventions
    UseDescriptiveNames    bool
    UseConsistentNaming    bool
    AvoidAbbreviations     bool
    
    // Field Design
    UseOptionalFields      bool
    ProvideDefaultValues   bool
    UseEnumsForConstants   bool
    
    // Evolution Strategy
    PlanForEvolution       bool
    UseCompatibleChanges   bool
    DocumentBreakingChanges bool
    
    // Performance
    MinimizeSchemaSize     bool
    UseEfficientTypes      bool
    CacheSchemas           bool
}
```

#### **Schema Governance:**
```go
type SchemaGovernance struct {
    // Access Control
    ReadPermissions  map[string][]string
    WritePermissions map[string][]string
    
    // Validation Rules
    RequiredFields   []string
    ForbiddenFields  []string
    FieldConstraints map[string]interface{}
    
    // Lifecycle Management
    RetentionPolicy  time.Duration
    ArchivePolicy    time.Duration
    CleanupPolicy    time.Duration
}
```

## 🧪 Hands-on Experiments

### Experiment 1: Schema Registry Setup

**Goal**: Set up Schema Registry and register schemas

**Steps**:
1. Start Schema Registry service
2. Register Avro schemas for different subjects
3. Test schema registration and retrieval
4. Verify schema versioning

**Expected Results**:
- Schema Registry running and accessible
- Schemas registered successfully
- Schema versioning working correctly

### Experiment 2: Schema Evolution

**Goal**: Implement schema evolution with compatibility

**Steps**:
1. Create initial schema version
2. Evolve schema with backward compatibility
3. Test data migration between versions
4. Verify compatibility rules

**Expected Results**:
- Schema evolution working correctly
- Data migration successful
- Compatibility rules enforced

### Experiment 3: Data Serialization

**Goal**: Implement Avro serialization and deserialization

**Steps**:
1. Create Avro schemas
2. Implement serialization logic
3. Implement deserialization logic
4. Test with different data types

**Expected Results**:
- Data serialization working
- Data deserialization working
- Schema validation working

### Experiment 4: Schema Validation

**Goal**: Implement runtime schema validation

**Steps**:
1. Create schema validators
2. Implement validation logic
3. Test with valid and invalid data
4. Handle validation errors

**Expected Results**:
- Schema validation working
- Error handling implemented
- Performance optimized

## 📊 Schema Management Metrics

### **Key Performance Indicators:**
- **Schema Registration Rate**: Schemas registered per minute
- **Schema Retrieval Latency**: Time to retrieve schemas
- **Validation Success Rate**: Percentage of valid messages
- **Migration Success Rate**: Percentage of successful migrations
- **Cache Hit Rate**: Schema cache hit percentage

### **Monitoring Schema Registry:**
```go
type SchemaRegistryMetrics struct {
    SchemasRegistered    int64
    SchemasRetrieved     int64
    ValidationFailures   int64
    MigrationFailures    int64
    CacheHits           int64
    CacheMisses         int64
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
```

## 🎯 Key Takeaways

1. **Schema Registry** provides centralized schema management
2. **Schema Evolution** enables safe schema changes over time
3. **Compatibility Rules** ensure data consistency across versions
4. **Data Migration** handles schema version transitions
5. **Schema Validation** ensures data quality at runtime
6. **Best Practices** are essential for schema management
7. **Monitoring** is crucial for schema registry health

## 📝 Lesson 12 Assessment Questions

1. **What is Schema Registry and why is it important?**
2. **What are the different types of schema compatibility?**
3. **How do you implement schema evolution safely?**
4. **What is the difference between Avro, JSON Schema, and Protobuf?**
5. **How do you handle data migration between schema versions?**
6. **What are the best practices for schema design?**
7. **How do you implement schema validation at runtime?**
8. **What metrics should you monitor for schema management?**

---

## 🔄 Next Lesson Preview: Kafka Security and Authentication

**What we'll learn**:
- Kafka security architecture
- Authentication mechanisms (SASL, SSL/TLS)
- Authorization and access control
- Encryption and data protection
- Security best practices
- Compliance and auditing

**Hands-on experiments**:
- Set up secure Kafka cluster
- Implement authentication and authorization
- Configure encryption and data protection
- Test security configurations
