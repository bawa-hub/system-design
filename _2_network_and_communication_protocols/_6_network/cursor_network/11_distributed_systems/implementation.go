package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// Node represents a node in the distributed system
type Node struct {
	ID          string
	Address     string
	Port        int
	State       string // "leader", "follower", "candidate"
	Term        int64
	VotedFor    string
	Log         []LogEntry
	CommitIndex int64
	LastApplied int64
	NextIndex   map[string]int64
	MatchIndex  map[string]int64
	Peers       map[string]*Node
	IsHealthy   bool
	LastHeartbeat time.Time
	mutex       sync.RWMutex
}

// LogEntry represents a log entry
type LogEntry struct {
	Term    int64
	Index   int64
	Command string
	Data    interface{}
}

// RaftNode represents a Raft consensus node
type RaftNode struct {
	Node
	ElectionTimeout  time.Duration
	HeartbeatTimeout time.Duration
	VoteCount        int32
	LeaderID         string
	StopChan         chan struct{}
	RequestVoteChan  chan RequestVoteRequest
	AppendEntriesChan chan AppendEntriesRequest
	ClientRequestChan chan ClientRequest
}

// RequestVoteRequest represents a request vote message
type RequestVoteRequest struct {
	Term         int64
	CandidateID  string
	LastLogIndex int64
	LastLogTerm  int64
}

// RequestVoteResponse represents a request vote response
type RequestVoteResponse struct {
	Term        int64
	VoteGranted bool
}

// AppendEntriesRequest represents an append entries message
type AppendEntriesRequest struct {
	Term         int64
	LeaderID     string
	PrevLogIndex int64
	PrevLogTerm  int64
	Entries      []LogEntry
	LeaderCommit int64
}

// AppendEntriesResponse represents an append entries response
type AppendEntriesResponse struct {
	Term    int64
	Success bool
}

// ClientRequest represents a client request
type ClientRequest struct {
	Command string
	Data    interface{}
	Reply   chan ClientResponse
}

// ClientResponse represents a client response
type ClientResponse struct {
	Success bool
	Data    interface{}
	Error   string
}

// NewRaftNode creates a new Raft node
func NewRaftNode(id, address string, port int) *RaftNode {
	return &RaftNode{
		Node: Node{
			ID:          id,
			Address:     address,
			Port:        port,
			State:       "follower",
			Term:        0,
			VotedFor:    "",
			Log:         make([]LogEntry, 0),
			CommitIndex: -1,
			LastApplied: -1,
			NextIndex:   make(map[string]int64),
			MatchIndex:  make(map[string]int64),
			Peers:       make(map[string]*Node),
			IsHealthy:   true,
		},
		ElectionTimeout:   time.Duration(150+rand.Intn(150)) * time.Millisecond,
		HeartbeatTimeout: 50 * time.Millisecond,
		StopChan:         make(chan struct{}),
		RequestVoteChan:  make(chan RequestVoteRequest, 100),
		AppendEntriesChan: make(chan AppendEntriesRequest, 100),
		ClientRequestChan: make(chan ClientRequest, 100),
	}
}

// Start starts the Raft node
func (rn *RaftNode) Start() {
	go rn.run()
	go rn.startHTTPServer()
}

// Stop stops the Raft node
func (rn *RaftNode) Stop() {
	close(rn.StopChan)
}

// run runs the main Raft loop
func (rn *RaftNode) run() {
	for {
		select {
		case <-rn.StopChan:
			return
		default:
			rn.mutex.RLock()
			state := rn.State
			rn.mutex.RUnlock()
			
			switch state {
			case "follower":
				rn.runFollower()
			case "candidate":
				rn.runCandidate()
			case "leader":
				rn.runLeader()
			}
		}
	}
}

// runFollower runs the follower state
func (rn *RaftNode) runFollower() {
	timeout := time.NewTimer(rn.ElectionTimeout)
	defer timeout.Stop()
	
	for {
		select {
		case <-rn.StopChan:
			return
		case <-timeout.C:
			// Election timeout, become candidate
			rn.becomeCandidate()
			return
		case req := <-rn.RequestVoteChan:
			rn.handleRequestVote(req)
		case req := <-rn.AppendEntriesChan:
			rn.handleAppendEntries(req)
		case req := <-rn.ClientRequestChan:
			// Redirect to leader
			rn.redirectToLeader(req)
		}
	}
}

// runCandidate runs the candidate state
func (rn *RaftNode) runCandidate() {
	rn.mutex.Lock()
	rn.State = "candidate"
	rn.Term++
	rn.VotedFor = rn.ID
	rn.VoteCount = 1
	rn.mutex.Unlock()
	
	// Request votes from all peers
	rn.requestVotes()
	
	timeout := time.NewTimer(rn.ElectionTimeout)
	defer timeout.Stop()
	
	for {
		select {
		case <-rn.StopChan:
			return
		case <-timeout.C:
			// Election timeout, start new election
			rn.becomeCandidate()
			return
		case req := <-rn.RequestVoteChan:
			rn.handleRequestVote(req)
		case req := <-rn.AppendEntriesChan:
			rn.handleAppendEntries(req)
		case req := <-rn.ClientRequestChan:
			// Redirect to leader
			rn.redirectToLeader(req)
		}
	}
}

// runLeader runs the leader state
func (rn *RaftNode) runLeader() {
	rn.mutex.Lock()
	rn.State = "leader"
	rn.LeaderID = rn.ID
	
	// Initialize nextIndex and matchIndex
	for peerID := range rn.Peers {
		rn.NextIndex[peerID] = int64(len(rn.Log))
		rn.MatchIndex[peerID] = -1
	}
	rn.mutex.Unlock()
	
	// Send initial heartbeat
	rn.sendHeartbeat()
	
	ticker := time.NewTicker(rn.HeartbeatTimeout)
	defer ticker.Stop()
	
	for {
		select {
		case <-rn.StopChan:
			return
		case <-ticker.C:
			rn.sendHeartbeat()
		case req := <-rn.RequestVoteChan:
			rn.handleRequestVote(req)
		case req := <-rn.AppendEntriesChan:
			rn.handleAppendEntries(req)
		case req := <-rn.ClientRequestChan:
			rn.handleClientRequest(req)
		}
	}
}

// becomeCandidate transitions to candidate state
func (rn *RaftNode) becomeCandidate() {
	rn.mutex.Lock()
	rn.State = "candidate"
	rn.Term++
	rn.VotedFor = rn.ID
	rn.VoteCount = 1
	rn.mutex.Unlock()
	
	// Request votes from all peers
	rn.requestVotes()
}

// requestVotes requests votes from all peers
func (rn *RaftNode) requestVotes() {
	rn.mutex.RLock()
	term := rn.Term
	lastLogIndex := int64(len(rn.Log)) - 1
	lastLogTerm := int64(0)
	if lastLogIndex >= 0 {
		lastLogTerm = rn.Log[lastLogIndex].Term
	}
	rn.mutex.RUnlock()
	
	req := RequestVoteRequest{
		Term:         term,
		CandidateID:  rn.ID,
		LastLogIndex: lastLogIndex,
		LastLogTerm:  lastLogTerm,
	}
	
	// Send to all peers
	for peerID, peer := range rn.Peers {
		go rn.sendRequestVote(peerID, peer, req)
	}
}

// sendRequestVote sends a request vote to a peer
func (rn *RaftNode) sendRequestVote(peerID string, peer *Node, req RequestVoteRequest) {
	// In a real implementation, this would send over the network
	// For simulation, we'll just process locally
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	
	// Simulate peer response
	if rand.Float64() < 0.8 { // 80% chance of vote
		rn.mutex.Lock()
		rn.VoteCount++
		if rn.VoteCount > int32(len(rn.Peers)/2) {
			rn.becomeLeader()
		}
		rn.mutex.Unlock()
	}
}

// becomeLeader transitions to leader state
func (rn *RaftNode) becomeLeader() {
	rn.mutex.Lock()
	rn.State = "leader"
	rn.LeaderID = rn.ID
	rn.mutex.Unlock()
}

// sendHeartbeat sends heartbeat to all peers
func (rn *RaftNode) sendHeartbeat() {
	rn.mutex.RLock()
	term := rn.Term
	leaderID := rn.ID
	leaderCommit := rn.CommitIndex
	rn.mutex.RUnlock()
	
	req := AppendEntriesRequest{
		Term:         term,
		LeaderID:     leaderID,
		PrevLogIndex: -1,
		PrevLogTerm:  -1,
		Entries:      []LogEntry{},
		LeaderCommit: leaderCommit,
	}
	
	// Send to all peers
	for peerID, peer := range rn.Peers {
		go rn.sendAppendEntries(peerID, peer, req)
	}
}

// sendAppendEntries sends append entries to a peer
func (rn *RaftNode) sendAppendEntries(peerID string, peer *Node, req AppendEntriesRequest) {
	// In a real implementation, this would send over the network
	// For simulation, we'll just process locally
	time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	
	// Simulate peer response
	rn.mutex.Lock()
	rn.Peers[peerID].LastHeartbeat = time.Now()
	rn.mutex.Unlock()
}

// handleRequestVote handles a request vote message
func (rn *RaftNode) handleRequestVote(req RequestVoteRequest) {
	rn.mutex.Lock()
	defer rn.mutex.Unlock()
	
	if req.Term > rn.Term {
		rn.Term = req.Term
		rn.State = "follower"
		rn.VotedFor = ""
	}
	
	if req.Term == rn.Term && (rn.VotedFor == "" || rn.VotedFor == req.CandidateID) {
		rn.VotedFor = req.CandidateID
		// Send vote response
	}
}

// handleAppendEntries handles an append entries message
func (rn *RaftNode) handleAppendEntries(req AppendEntriesRequest) {
	rn.mutex.Lock()
	defer rn.mutex.Unlock()
	
	if req.Term >= rn.Term {
		rn.Term = req.Term
		rn.State = "follower"
		rn.LeaderID = req.LeaderID
		rn.VotedFor = ""
		
		// Update commit index
		if req.LeaderCommit > rn.CommitIndex {
			rn.CommitIndex = req.LeaderCommit
		}
		
		// Send success response
	}
}

// handleClientRequest handles a client request
func (rn *RaftNode) handleClientRequest(req ClientRequest) {
	rn.mutex.Lock()
	defer rn.mutex.Unlock()
	
	// Add to log
	logEntry := LogEntry{
		Term:    rn.Term,
		Index:   int64(len(rn.Log)),
		Command: req.Command,
		Data:    req.Data,
	}
	rn.Log = append(rn.Log, logEntry)
	
	// Send response
	req.Reply <- ClientResponse{
		Success: true,
		Data:    "Request processed",
	}
}

// redirectToLeader redirects client request to leader
func (rn *RaftNode) redirectToLeader(req ClientRequest) {
	rn.mutex.RLock()
	leaderID := rn.LeaderID
	rn.mutex.RUnlock()
	
	if leaderID != "" {
		// Redirect to leader
		req.Reply <- ClientResponse{
			Success: false,
			Error:   fmt.Sprintf("Redirect to leader: %s", leaderID),
		}
	} else {
		req.Reply <- ClientResponse{
			Success: false,
			Error:   "No leader available",
		}
	}
}

// startHTTPServer starts the HTTP server for the node
func (rn *RaftNode) startHTTPServer() {
	http.HandleFunc("/status", rn.handleStatus)
	http.HandleFunc("/request", rn.handleRequest)
	http.HandleFunc("/vote", rn.handleVote)
	http.HandleFunc("/append", rn.handleAppend)
	
	addr := fmt.Sprintf(":%d", rn.Port)
	log.Printf("Raft node %s starting HTTP server on %s", rn.ID, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// handleStatus handles status requests
func (rn *RaftNode) handleStatus(w http.ResponseWriter, r *http.Request) {
	rn.mutex.RLock()
	defer rn.mutex.RUnlock()
	
	status := map[string]interface{}{
		"id":            rn.ID,
		"state":         rn.State,
		"term":          rn.Term,
		"voted_for":     rn.VotedFor,
		"log_length":    len(rn.Log),
		"commit_index":  rn.CommitIndex,
		"last_applied":  rn.LastApplied,
		"is_healthy":    rn.IsHealthy,
		"last_heartbeat": rn.LastHeartbeat,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

// handleRequest handles client requests
func (rn *RaftNode) handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req ClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	req.Reply = make(chan ClientResponse, 1)
	rn.ClientRequestChan <- req
	
	select {
	case resp := <-req.Reply:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	case <-time.After(5 * time.Second):
		http.Error(w, "Request timeout", http.StatusRequestTimeout)
	}
}

// handleVote handles vote requests
func (rn *RaftNode) handleVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req RequestVoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	rn.RequestVoteChan <- req
	
	// Send response
	resp := RequestVoteResponse{
		Term:        rn.Term,
		VoteGranted: true, // Simplified
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// handleAppend handles append entries requests
func (rn *RaftNode) handleAppend(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req AppendEntriesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	rn.AppendEntriesChan <- req
	
	// Send response
	resp := AppendEntriesResponse{
		Term:    rn.Term,
		Success: true, // Simplified
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// ServiceRegistry represents a service registry
type ServiceRegistry struct {
	Services map[string]*Service
	mutex    sync.RWMutex
}

// Service represents a registered service
type Service struct {
	ID          string
	Name        string
	Address     string
	Port        int
	HealthCheck string
	Tags        []string
	LastSeen    time.Time
	IsHealthy   bool
}

// NewServiceRegistry creates a new service registry
func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		Services: make(map[string]*Service),
	}
}

// Register registers a service
func (sr *ServiceRegistry) Register(service *Service) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	
	service.LastSeen = time.Now()
	service.IsHealthy = true
	sr.Services[service.ID] = service
}

// Deregister deregisters a service
func (sr *ServiceRegistry) Deregister(serviceID string) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	
	delete(sr.Services, serviceID)
}

// GetService gets a service by ID
func (sr *ServiceRegistry) GetService(serviceID string) *Service {
	sr.mutex.RLock()
	defer sr.mutex.RUnlock()
	
	return sr.Services[serviceID]
}

// GetServicesByName gets services by name
func (sr *ServiceRegistry) GetServicesByName(name string) []*Service {
	sr.mutex.RLock()
	defer sr.mutex.RUnlock()
	
	var services []*Service
	for _, service := range sr.Services {
		if service.Name == name && service.IsHealthy {
			services = append(services, service)
		}
	}
	return services
}

// GetAllServices gets all services
func (sr *ServiceRegistry) GetAllServices() []*Service {
	sr.mutex.RLock()
	defer sr.mutex.RUnlock()
	
	var services []*Service
	for _, service := range sr.Services {
		services = append(services, service)
	}
	return services
}

// HealthCheck performs health check on a service
func (sr *ServiceRegistry) HealthCheck(serviceID string) bool {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	
	service, exists := sr.Services[serviceID]
	if !exists {
		return false
	}
	
	// Simulate health check
	service.IsHealthy = rand.Float64() > 0.1 // 90% chance of being healthy
	service.LastSeen = time.Now()
	
	return service.IsHealthy
}

// LoadBalancer represents a load balancer
type LoadBalancer struct {
	Services    map[string][]*Service
	Algorithm   string
	Current     map[string]int32
	mutex       sync.RWMutex
}

// NewLoadBalancer creates a new load balancer
func NewLoadBalancer(algorithm string) *LoadBalancer {
	return &LoadBalancer{
		Services:  make(map[string][]*Service),
		Algorithm: algorithm,
		Current:   make(map[string]int32),
	}
}

// AddService adds a service to the load balancer
func (lb *LoadBalancer) AddService(serviceName string, service *Service) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	
	lb.Services[serviceName] = append(lb.Services[serviceName], service)
}

// GetService gets a service using the load balancing algorithm
func (lb *LoadBalancer) GetService(serviceName string) *Service {
	lb.mutex.RLock()
	defer lb.mutex.RUnlock()
	
	services, exists := lb.Services[serviceName]
	if !exists || len(services) == 0 {
		return nil
	}
	
	// Filter healthy services
	var healthyServices []*Service
	for _, service := range services {
		if service.IsHealthy {
			healthyServices = append(healthyServices, service)
		}
	}
	
	if len(healthyServices) == 0 {
		return nil
	}
	
	switch lb.Algorithm {
	case "round_robin":
		return lb.getRoundRobinService(serviceName, healthyServices)
	case "random":
		return lb.getRandomService(healthyServices)
	case "least_connections":
		return lb.getLeastConnectionsService(healthyServices)
	default:
		return lb.getRoundRobinService(serviceName, healthyServices)
	}
}

// getRoundRobinService gets the next service in round-robin fashion
func (lb *LoadBalancer) getRoundRobinService(serviceName string, services []*Service) *Service {
	if len(services) == 0 {
		return nil
	}
	
	index := int(atomic.AddInt32(&lb.Current[serviceName], 1)) % len(services)
	return services[index]
}

// getRandomService gets a random service
func (lb *LoadBalancer) getRandomService(services []*Service) *Service {
	if len(services) == 0 {
		return nil
	}
	
	index := rand.Intn(len(services))
	return services[index]
}

// getLeastConnectionsService gets the service with the least connections
func (lb *LoadBalancer) getLeastConnectionsService(services []*Service) *Service {
	if len(services) == 0 {
		return nil
	}
	
	// Simplified - in reality, you would track actual connections
	return services[0]
}

// MessageQueue represents a message queue
type MessageQueue struct {
	Messages chan Message
	Subscribers map[string][]chan Message
	mutex    sync.RWMutex
}

// Message represents a message in the queue
type Message struct {
	ID        string
	Topic     string
	Data      interface{}
	Timestamp time.Time
	TTL       time.Duration
}

// NewMessageQueue creates a new message queue
func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		Messages:    make(chan Message, 1000),
		Subscribers: make(map[string][]chan Message),
	}
}

// Publish publishes a message to a topic
func (mq *MessageQueue) Publish(topic string, data interface{}) error {
	message := Message{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		Topic:     topic,
		Data:      data,
		Timestamp: time.Now(),
		TTL:       24 * time.Hour,
	}
	
	select {
	case mq.Messages <- message:
		return nil
	default:
		return fmt.Errorf("queue full")
	}
}

// Subscribe subscribes to a topic
func (mq *MessageQueue) Subscribe(topic string) <-chan Message {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	
	ch := make(chan Message, 100)
	mq.Subscribers[topic] = append(mq.Subscribers[topic], ch)
	return ch
}

// Start starts the message queue
func (mq *MessageQueue) Start() {
	go mq.processMessages()
}

// processMessages processes messages and distributes to subscribers
func (mq *MessageQueue) processMessages() {
	for message := range mq.Messages {
		mq.mutex.RLock()
		subscribers := mq.Subscribers[message.Topic]
		mq.mutex.RUnlock()
		
		for _, subscriber := range subscribers {
			select {
			case subscriber <- message:
			default:
				// Subscriber channel full, skip
			}
		}
	}
}

// DistributedCache represents a distributed cache
type DistributedCache struct {
	Data  map[string]CacheItem
	mutex sync.RWMutex
}

// CacheItem represents a cache item
type CacheItem struct {
	Value     interface{}
	ExpiresAt time.Time
	Version   int64
}

// NewDistributedCache creates a new distributed cache
func NewDistributedCache() *DistributedCache {
	return &DistributedCache{
		Data: make(map[string]CacheItem),
	}
}

// Set sets a value in the cache
func (dc *DistributedCache) Set(key string, value interface{}, ttl time.Duration) {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	
	dc.Data[key] = CacheItem{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
		Version:   time.Now().UnixNano(),
	}
}

// Get gets a value from the cache
func (dc *DistributedCache) Get(key string) (interface{}, bool) {
	dc.mutex.RLock()
	defer dc.mutex.RUnlock()
	
	item, exists := dc.Data[key]
	if !exists {
		return nil, false
	}
	
	if time.Now().After(item.ExpiresAt) {
		// Item expired
		dc.mutex.RUnlock()
		dc.mutex.Lock()
		delete(dc.Data, key)
		dc.mutex.Unlock()
		dc.mutex.RLock()
		return nil, false
	}
	
	return item.Value, true
}

// Delete deletes a value from the cache
func (dc *DistributedCache) Delete(key string) {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	
	delete(dc.Data, key)
}

// Clear clears all values from the cache
func (dc *DistributedCache) Clear() {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()
	
	dc.Data = make(map[string]CacheItem)
}

// Demonstrate Raft consensus
func demonstrateRaftConsensus() {
	fmt.Println("=== Raft Consensus Demo ===\n")
	
	// Create Raft nodes
	nodes := make([]*RaftNode, 3)
	for i := 0; i < 3; i++ {
		nodes[i] = NewRaftNode(fmt.Sprintf("node-%d", i), "localhost", 8080+i)
	}
	
	// Add peers to each node
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j {
				nodes[i].Peers[fmt.Sprintf("node-%d", j)] = &nodes[j].Node
			}
		}
	}
	
	// Start nodes
	for _, node := range nodes {
		go node.Start()
	}
	
	// Wait for election
	time.Sleep(2 * time.Second)
	
	// Check node states
	for _, node := range nodes {
		node.mutex.RLock()
		fmt.Printf("Node %s: State=%s, Term=%d, LogLength=%d\n", 
			node.ID, node.State, node.Term, len(node.Log))
		node.mutex.RUnlock()
	}
	
	// Stop nodes
	for _, node := range nodes {
		node.Stop()
	}
}

// Demonstrate service discovery
func demonstrateServiceDiscovery() {
	fmt.Println("=== Service Discovery Demo ===\n")
	
	// Create service registry
	registry := NewServiceRegistry()
	
	// Register services
	services := []*Service{
		{ID: "service-1", Name: "user-service", Address: "localhost", Port: 8001, HealthCheck: "/health"},
		{ID: "service-2", Name: "user-service", Address: "localhost", Port: 8002, HealthCheck: "/health"},
		{ID: "service-3", Name: "order-service", Address: "localhost", Port: 8003, HealthCheck: "/health"},
		{ID: "service-4", Name: "payment-service", Address: "localhost", Port: 8004, HealthCheck: "/health"},
	}
	
	for _, service := range services {
		registry.Register(service)
	}
	
	// Display all services
	fmt.Println("All Services:")
	for _, service := range registry.GetAllServices() {
		fmt.Printf("  %s: %s:%d (healthy: %t)\n", 
			service.ID, service.Address, service.Port, service.IsHealthy)
	}
	
	// Get services by name
	fmt.Println("\nUser Services:")
	userServices := registry.GetServicesByName("user-service")
	for _, service := range userServices {
		fmt.Printf("  %s: %s:%d\n", service.ID, service.Address, service.Port)
	}
	
	// Perform health checks
	fmt.Println("\nHealth Checks:")
	for _, service := range registry.GetAllServices() {
		healthy := registry.HealthCheck(service.ID)
		fmt.Printf("  %s: %t\n", service.ID, healthy)
	}
}

// Demonstrate load balancing
func demonstrateLoadBalancing() {
	fmt.Println("=== Load Balancing Demo ===\n")
	
	// Create load balancer
	lb := NewLoadBalancer("round_robin")
	
	// Add services
	services := []*Service{
		{ID: "service-1", Name: "api-service", Address: "localhost", Port: 8001, IsHealthy: true},
		{ID: "service-2", Name: "api-service", Address: "localhost", Port: 8002, IsHealthy: true},
		{ID: "service-3", Name: "api-service", Address: "localhost", Port: 8003, IsHealthy: true},
	}
	
	for _, service := range services {
		lb.AddService("api-service", service)
	}
	
	// Test load balancing
	fmt.Println("Load Balancing Test:")
	for i := 0; i < 10; i++ {
		service := lb.GetService("api-service")
		if service != nil {
			fmt.Printf("  Request %d: %s:%d\n", i+1, service.Address, service.Port)
		} else {
			fmt.Printf("  Request %d: No service available\n", i+1)
		}
	}
}

// Demonstrate message queuing
func demonstrateMessageQueuing() {
	fmt.Println("=== Message Queuing Demo ===\n")
	
	// Create message queue
	mq := NewMessageQueue()
	mq.Start()
	
	// Subscribe to topics
	userTopic := mq.Subscribe("user-events")
	orderTopic := mq.Subscribe("order-events")
	
	// Publish messages
	go func() {
		for i := 0; i < 5; i++ {
			mq.Publish("user-events", map[string]interface{}{
				"user_id": i,
				"action":  "login",
				"time":    time.Now(),
			})
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	go func() {
		for i := 0; i < 3; i++ {
			mq.Publish("order-events", map[string]interface{}{
				"order_id": i,
				"status":   "created",
				"time":     time.Now(),
			})
			time.Sleep(150 * time.Millisecond)
		}
	}()
	
	// Consume messages
	go func() {
		for message := range userTopic {
			fmt.Printf("User Event: %+v\n", message.Data)
		}
	}()
	
	go func() {
		for message := range orderTopic {
			fmt.Printf("Order Event: %+v\n", message.Data)
		}
	}()
	
	// Wait for messages
	time.Sleep(2 * time.Second)
}

// Demonstrate distributed caching
func demonstrateDistributedCaching() {
	fmt.Println("=== Distributed Caching Demo ===\n")
	
	// Create distributed cache
	cache := NewDistributedCache()
	
	// Set values
	cache.Set("user:1", map[string]interface{}{"name": "John", "email": "john@example.com"}, 5*time.Minute)
	cache.Set("user:2", map[string]interface{}{"name": "Jane", "email": "jane@example.com"}, 5*time.Minute)
	cache.Set("config:app", map[string]interface{}{"version": "1.0", "debug": true}, 10*time.Minute)
	
	// Get values
	fmt.Println("Cache Operations:")
	if value, found := cache.Get("user:1"); found {
		fmt.Printf("  user:1 = %+v\n", value)
	}
	
	if value, found := cache.Get("user:2"); found {
		fmt.Printf("  user:2 = %+v\n", value)
	}
	
	if value, found := cache.Get("config:app"); found {
		fmt.Printf("  config:app = %+v\n", value)
	}
	
	// Test expiration
	cache.Set("temp:key", "temporary value", 100*time.Millisecond)
	time.Sleep(200 * time.Millisecond)
	
	if value, found := cache.Get("temp:key"); found {
		fmt.Printf("  temp:key = %+v (should be expired)\n", value)
	} else {
		fmt.Println("  temp:key = expired (as expected)")
	}
	
	// Clear cache
	cache.Clear()
	fmt.Println("  Cache cleared")
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 11: Distributed Systems Networking")
	fmt.Println("=======================================================================\n")
	
	// Run all demonstrations
	demonstrateRaftConsensus()
	fmt.Println()
	demonstrateServiceDiscovery()
	fmt.Println()
	demonstrateLoadBalancing()
	fmt.Println()
	demonstrateMessageQueuing()
	fmt.Println()
	demonstrateDistributedCaching()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Distributed systems require consensus algorithms for coordination")
	fmt.Println("2. Service discovery enables dynamic service location")
	fmt.Println("3. Load balancing distributes traffic across multiple services")
	fmt.Println("4. Message queuing enables asynchronous communication")
	fmt.Println("5. Distributed caching improves performance and reduces load")
	fmt.Println("6. Raft provides understandable consensus for distributed systems")
	fmt.Println("7. Go provides excellent tools for building distributed systems")
	fmt.Println("8. Understanding distributed patterns is crucial for scalability")
	
	fmt.Println("\n📚 Next Topic: Cloud Networking & Containerization")
}
