# Lesson 14: Multi-cluster and Cross-Datacenter

## 🎯 Learning Objectives
By the end of this lesson, you will understand:
- Multi-cluster Kafka architectures and patterns
- Cross-datacenter replication strategies
- Disaster recovery and high availability
- Cluster federation and data synchronization
- Network partitioning and split-brain scenarios
- Global data distribution patterns
- Enterprise deployment best practices

## 📚 Theory Section

### 1. **Multi-cluster Architecture Overview**

#### **Why Multi-cluster?**
- **Geographic Distribution**: Serve users in different regions
- **Disaster Recovery**: Backup clusters in different locations
- **Compliance**: Data residency requirements
- **Scalability**: Distribute load across multiple clusters
- **Isolation**: Separate environments (dev, staging, prod)

#### **Multi-cluster Patterns:**
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Primary DC    │    │  Secondary DC   │    │   Tertiary DC   │
│                 │    │                 │    │                 │
│  ┌───────────┐  │    │  ┌───────────┐  │    │  ┌───────────┐  │
│  │  Cluster  │  │    │  │  Cluster  │  │    │  │  Cluster  │  │
│  │     A     │  │    │  │     B     │  │    │  │     C     │  │
│  └───────────┘  │    │  └───────────┘  │    │  └───────────┘  │
│                 │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                    ┌─────────────────┐
                    │  Replication    │
                    │    Manager      │
                    └─────────────────┘
```

### 2. **Cross-Datacenter Replication**

#### **MirrorMaker 2.0:**
```go
type MirrorMaker2Config struct {
    Clusters map[string]ClusterConfig `json:"clusters"`
    Replication map[string]ReplicationConfig `json:"replication"`
    Connectors  map[string]ConnectorConfig `json:"connectors"`
}

type ClusterConfig struct {
    BootstrapServers string `json:"bootstrap.servers"`
    SecurityProtocol string `json:"security.protocol"`
    SASLMechanism   string `json:"sasl.mechanism"`
    SASLUsername    string `json:"sasl.username"`
    SASLPassword    string `json:"sasl.password"`
}

type ReplicationConfig struct {
    SourceCluster    string   `json:"source.cluster"`
    TargetCluster    string   `json:"target.cluster"`
    Topics           []string `json:"topics"`
    TopicsRegex      string   `json:"topics.regex"`
    ConsumerGroups   []string `json:"consumer.groups"`
    OffsetSync       bool     `json:"offset.sync"`
    CheckpointSync   bool     `json:"checkpoint.sync"`
}
```

#### **Replication Strategies:**
```go
type ReplicationStrategy string

const (
    // Active-Passive: One active cluster, others standby
    ACTIVE_PASSIVE ReplicationStrategy = "ACTIVE_PASSIVE"
    
    // Active-Active: Multiple active clusters
    ACTIVE_ACTIVE ReplicationStrategy = "ACTIVE_ACTIVE"
    
    // Hub-Spoke: Central hub with multiple spokes
    HUB_SPOKE ReplicationStrategy = "HUB_SPOKE"
    
    // Mesh: All clusters replicate to each other
    MESH ReplicationStrategy = "MESH"
)

type ReplicationManager struct {
    strategy    ReplicationStrategy
    clusters    map[string]*Cluster
    replicators map[string]*Replicator
    mutex       sync.RWMutex
}

func (rm *ReplicationManager) StartReplication() error {
    switch rm.strategy {
    case ACTIVE_PASSIVE:
        return rm.startActivePassiveReplication()
    case ACTIVE_ACTIVE:
        return rm.startActiveActiveReplication()
    case HUB_SPOKE:
        return rm.startHubSpokeReplication()
    case MESH:
        return rm.startMeshReplication()
    default:
        return fmt.Errorf("unknown replication strategy: %s", rm.strategy)
    }
}
```

### 3. **Disaster Recovery Strategies**

#### **RTO and RPO Definitions:**
```go
type DisasterRecoveryConfig struct {
    // Recovery Time Objective (RTO) - Maximum acceptable downtime
    RTO time.Duration
    
    // Recovery Point Objective (RPO) - Maximum acceptable data loss
    RPO time.Duration
    
    // Backup frequency
    BackupFrequency time.Duration
    
    // Replication lag tolerance
    MaxReplicationLag time.Duration
    
    // Failover strategy
    FailoverStrategy FailoverStrategy
}

type FailoverStrategy string

const (
    AUTOMATIC_FAILOVER FailoverStrategy = "AUTOMATIC"
    MANUAL_FAILOVER   FailoverStrategy = "MANUAL"
    GRACEFUL_FAILOVER FailoverStrategy = "GRACEFUL"
)
```

#### **Disaster Recovery Manager:**
```go
type DisasterRecoveryManager struct {
    primaryCluster    *Cluster
    secondaryClusters []*Cluster
    config           DisasterRecoveryConfig
    healthChecker    *HealthChecker
    mutex            sync.RWMutex
}

func (drm *DisasterRecoveryManager) MonitorHealth() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        health := drm.healthChecker.CheckClusterHealth(drm.primaryCluster)
        
        if !health.IsHealthy {
            drm.handleClusterFailure(health)
        }
    }
}

func (drm *DisasterRecoveryManager) handleClusterFailure(health ClusterHealth) {
    drm.mutex.Lock()
    defer drm.mutex.Unlock()
    
    switch drm.config.FailoverStrategy {
    case AUTOMATIC_FAILOVER:
        drm.performAutomaticFailover()
    case MANUAL_FAILOVER:
        drm.triggerManualFailover()
    case GRACEFUL_FAILOVER:
        drm.performGracefulFailover()
    }
}

func (drm *DisasterRecoveryManager) performAutomaticFailover() error {
    // Select best secondary cluster
    bestCluster := drm.selectBestSecondaryCluster()
    if bestCluster == nil {
        return fmt.Errorf("no healthy secondary cluster available")
    }
    
    // Update DNS/routing to point to secondary cluster
    if err := drm.updateRouting(bestCluster); err != nil {
        return err
    }
    
    // Start replication from secondary to primary
    if err := drm.startReverseReplication(bestCluster); err != nil {
        return err
    }
    
    return nil
}
```

### 4. **Cluster Federation**

#### **Federation Manager:**
```go
type FederationManager struct {
    clusters    map[string]*Cluster
    connectors  map[string]*Connector
    topics      map[string]*FederatedTopic
    mutex       sync.RWMutex
}

type FederatedTopic struct {
    Name        string
    Clusters    []string
    PrimaryCluster string
    ReplicationFactor int
    Partitions   int
    Config      map[string]string
}

func (fm *FederationManager) CreateFederatedTopic(topic FederatedTopic) error {
    fm.mutex.Lock()
    defer fm.mutex.Unlock()
    
    // Create topic in primary cluster
    if err := fm.createTopicInCluster(topic.PrimaryCluster, topic); err != nil {
        return err
    }
    
    // Create topic in secondary clusters
    for _, clusterID := range topic.Clusters {
        if clusterID == topic.PrimaryCluster {
            continue
        }
        
        if err := fm.createTopicInCluster(clusterID, topic); err != nil {
            return err
        }
    }
    
    // Set up replication
    if err := fm.setupReplication(topic); err != nil {
        return err
    }
    
    fm.topics[topic.Name] = &topic
    return nil
}

func (fm *FederationManager) setupReplication(topic FederatedTopic) error {
    // Create MirrorMaker connectors for each cluster
    for _, clusterID := range topic.Clusters {
        if clusterID == topic.PrimaryCluster {
            continue
        }
        
        connector := &Connector{
            Name: fmt.Sprintf("replicate-%s-to-%s", topic.PrimaryCluster, clusterID),
            Type: "MirrorSourceConnector",
            Config: map[string]string{
                "source.cluster.bootstrap.servers": fm.clusters[topic.PrimaryCluster].BootstrapServers,
                "target.cluster.bootstrap.servers": fm.clusters[clusterID].BootstrapServers,
                "topics": topic.Name,
                "consumer.group.id": fmt.Sprintf("mirror-%s-%s", topic.PrimaryCluster, clusterID),
            },
        }
        
        if err := fm.createConnector(connector); err != nil {
            return err
        }
    }
    
    return nil
}
```

### 5. **Data Synchronization**

#### **Synchronization Manager:**
```go
type SynchronizationManager struct {
    clusters    map[string]*Cluster
    syncConfig  SyncConfig
    mutex       sync.RWMutex
}

type SyncConfig struct {
    SyncInterval    time.Duration
    BatchSize       int
    CompressionType string
    RetryAttempts   int
    RetryDelay      time.Duration
}

func (sm *SynchronizationManager) SyncTopic(sourceCluster, targetCluster, topic string) error {
    sourceClient := sm.clusters[sourceCluster].Client
    targetClient := sm.clusters[targetCluster].Client
    
    // Get topic metadata
    metadata, err := sourceClient.Metadata(topic, nil, time.Second*10)
    if err != nil {
        return err
    }
    
    // Create consumer for source cluster
    consumer, err := sourceClient.Consumer(&sarama.ConsumerConfig{
        Group: "sync-consumer",
        Topics: []string{topic},
        Offset: sarama.OffsetOldest,
    })
    if err != nil {
        return err
    }
    defer consumer.Close()
    
    // Create producer for target cluster
    producer, err := targetClient.Producer(&sarama.ProducerConfig{
        Topic: topic,
        Compression: sarama.CompressionSnappy,
    })
    if err != nil {
        return err
    }
    defer producer.Close()
    
    // Sync messages
    return sm.syncMessages(consumer, producer, topic)
}

func (sm *SynchronizationManager) syncMessages(consumer sarama.Consumer, producer sarama.Producer, topic string) error {
    ticker := time.NewTicker(sm.syncConfig.SyncInterval)
    defer ticker.Stop()
    
    for range ticker.C {
        // Consume messages from source
        messages, err := consumer.Consume(sm.syncConfig.BatchSize)
        if err != nil {
            return err
        }
        
        if len(messages) == 0 {
            continue
        }
        
        // Produce messages to target
        for _, message := range messages {
            if err := producer.SendMessage(message); err != nil {
                return err
            }
        }
    }
    
    return nil
}
```

### 6. **Network Partitioning and Split-Brain**

#### **Split-Brain Detection:**
```go
type SplitBrainDetector struct {
    clusters    map[string]*Cluster
    quorumSize  int
    mutex       sync.RWMutex
}

func (sbd *SplitBrainDetector) DetectSplitBrain() (bool, []string) {
    sbd.mutex.RLock()
    defer sbd.mutex.RUnlock()
    
    var healthyClusters []string
    var unhealthyClusters []string
    
    for clusterID, cluster := range sbd.clusters {
        if sbd.isClusterHealthy(cluster) {
            healthyClusters = append(healthyClusters, clusterID)
        } else {
            unhealthyClusters = append(unhealthyClusters, clusterID)
        }
    }
    
    // Check if we have quorum
    if len(healthyClusters) < sbd.quorumSize {
        return true, unhealthyClusters
    }
    
    return false, unhealthyClusters
}

func (sbd *SplitBrainDetector) isClusterHealthy(cluster *Cluster) bool {
    // Check if cluster is reachable
    if !cluster.IsReachable() {
        return false
    }
    
    // Check if cluster has quorum
    if !cluster.HasQuorum() {
        return false
    }
    
    // Check replication lag
    if cluster.GetReplicationLag() > 30*time.Second {
        return false
    }
    
    return true
}
```

#### **Split-Brain Resolution:**
```go
type SplitBrainResolver struct {
    detector    *SplitBrainDetector
    config      SplitBrainConfig
    mutex       sync.RWMutex
}

type SplitBrainConfig struct {
    ResolutionStrategy ResolutionStrategy
    PrimaryCluster    string
    MaxPartitionTime  time.Duration
}

type ResolutionStrategy string

const (
    MANUAL_RESOLUTION    ResolutionStrategy = "MANUAL"
    AUTOMATIC_RESOLUTION ResolutionStrategy = "AUTOMATIC"
    QUORUM_RESOLUTION    ResolutionStrategy = "QUORUM"
)

func (sbr *SplitBrainResolver) ResolveSplitBrain() error {
    sbr.mutex.Lock()
    defer sbr.mutex.Unlock()
    
    isSplit, unhealthyClusters := sbr.detector.DetectSplitBrain()
    if !isSplit {
        return nil
    }
    
    switch sbr.config.ResolutionStrategy {
    case MANUAL_RESOLUTION:
        return sbr.triggerManualResolution(unhealthyClusters)
    case AUTOMATIC_RESOLUTION:
        return sbr.performAutomaticResolution(unhealthyClusters)
    case QUORUM_RESOLUTION:
        return sbr.performQuorumResolution(unhealthyClusters)
    default:
        return fmt.Errorf("unknown resolution strategy: %s", sbr.config.ResolutionStrategy)
    }
}
```

### 7. **Global Data Distribution**

#### **Global Distribution Manager:**
```go
type GlobalDistributionManager struct {
    regions    map[string]*Region
    replicators map[string]*Replicator
    mutex      sync.RWMutex
}

type Region struct {
    ID          string
    Name        string
    Cluster     *Cluster
    Latency     time.Duration
    Throughput  int64
    Health      ClusterHealth
}

func (gdm *GlobalDistributionManager) DistributeTopic(topic string, regions []string) error {
    gdm.mutex.Lock()
    defer gdm.mutex.Unlock()
    
    // Create topic in all regions
    for _, regionID := range regions {
        region := gdm.regions[regionID]
        if err := gdm.createTopicInRegion(region, topic); err != nil {
            return err
        }
    }
    
    // Set up cross-region replication
    if err := gdm.setupCrossRegionReplication(topic, regions); err != nil {
        return err
    }
    
    return nil
}

func (gdm *GlobalDistributionManager) setupCrossRegionReplication(topic string, regions []string) error {
    // Create replication topology
    for i, sourceRegion := range regions {
        for j, targetRegion := range regions {
            if i == j {
                continue
            }
            
            replicator := &Replicator{
                SourceRegion: sourceRegion,
                TargetRegion: targetRegion,
                Topic:        topic,
                Config:       gdm.getReplicationConfig(sourceRegion, targetRegion),
            }
            
            if err := gdm.startReplicator(replicator); err != nil {
                return err
            }
        }
    }
    
    return nil
}
```

## 🧪 Hands-on Experiments

### Experiment 1: Multi-cluster Setup

**Goal**: Set up multiple Kafka clusters

**Steps**:
1. Configure multiple Kafka clusters
2. Set up cross-cluster connectivity
3. Test cluster communication
4. Verify cluster health

**Expected Results**:
- Multiple clusters running
- Cross-cluster communication working
- Health monitoring active

### Experiment 2: Cross-Datacenter Replication

**Goal**: Implement cross-datacenter replication

**Steps**:
1. Configure MirrorMaker 2.0
2. Set up topic replication
3. Test data replication
4. Monitor replication lag

**Expected Results**:
- Data replicating between clusters
- Replication lag within acceptable limits
- Monitoring and alerting working

### Experiment 3: Disaster Recovery

**Goal**: Test disaster recovery procedures

**Steps**:
1. Simulate cluster failure
2. Test failover procedures
3. Verify data consistency
4. Test recovery procedures

**Expected Results**:
- Automatic failover working
- Data consistency maintained
- Recovery procedures successful

### Experiment 4: Split-Brain Resolution

**Goal**: Handle network partitioning scenarios

**Steps**:
1. Simulate network partition
2. Test split-brain detection
3. Implement resolution strategy
4. Verify cluster recovery

**Expected Results**:
- Split-brain detected correctly
- Resolution strategy working
- Cluster recovery successful

## 📊 Multi-cluster Metrics

### **Key Performance Indicators:**
- **Replication Lag**: Time delay between clusters
- **Throughput**: Messages per second across clusters
- **Availability**: Cluster uptime percentage
- **Consistency**: Data consistency across clusters
- **Recovery Time**: Time to recover from failures

### **Monitoring Dashboard:**
```go
type MultiClusterDashboard struct {
    clusters    map[string]*ClusterMetrics
    replication map[string]*ReplicationMetrics
    health      map[string]*HealthMetrics
}

type ClusterMetrics struct {
    MessagesPerSecond int64
    BytesPerSecond    int64
    PartitionCount    int64
    ConsumerLag       int64
    Uptime           time.Duration
}

type ReplicationMetrics struct {
    SourceCluster    string
    TargetCluster    string
    ReplicationLag   time.Duration
    Throughput       int64
    ErrorRate        float64
    LastSyncTime     time.Time
}
```

## 🎯 Key Takeaways

1. **Multi-cluster architecture** enables geographic distribution and disaster recovery
2. **Cross-datacenter replication** ensures data availability across regions
3. **Disaster recovery** requires careful planning and testing
4. **Cluster federation** enables unified management of multiple clusters
5. **Split-brain scenarios** must be handled gracefully
6. **Global data distribution** requires careful consideration of latency and consistency
7. **Monitoring and alerting** are crucial for multi-cluster deployments

## 📝 Lesson 14 Assessment Questions

1. **What are the benefits of multi-cluster Kafka deployments?**
2. **How do you implement cross-datacenter replication?**
3. **What are the different disaster recovery strategies?**
4. **How do you handle split-brain scenarios?**
5. **What is cluster federation and how does it work?**
6. **How do you monitor multi-cluster deployments?**
7. **What are the challenges of global data distribution?**
8. **How do you ensure data consistency across clusters?**

---

## 🎉 Phase 3 Complete - Congratulations!

**You have successfully completed Phase 3: Advanced Level - Stream Processing and Enterprise Features!**

### **What You've Mastered:**
- ✅ Kafka Streams and real-time stream processing
- ✅ Schema Registry and data evolution
- ✅ Kafka security and authentication
- ✅ Multi-cluster and cross-datacenter deployments

### **You are now a Kafka Expert!** 🚀

**Ready for the next phase of your learning journey or ready to build production systems!**
