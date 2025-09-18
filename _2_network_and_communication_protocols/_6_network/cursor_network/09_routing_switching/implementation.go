package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"sort"
	"sync"
	"time"
)

// NetworkNode represents a network device (router or switch)
type NetworkNode struct {
	ID       string
	Type     string // "router" or "switch"
	Ports    map[string]*Port
	MACTable map[string]string // MAC address -> port mapping
	Routes   map[string]*Route // destination -> route
	Neighbors map[string]*NetworkNode
	VLANs    map[int]*VLAN
	mutex    sync.RWMutex
}

// Port represents a network port
type Port struct {
	ID       string
	Type     string // "ethernet", "serial", "fiber"
	Speed    int    // Mbps
	Duplex   string // "half", "full"
	Status   string // "up", "down", "admin_down"
	VLAN     int
	Neighbor *NetworkNode
}

// Route represents a routing table entry
type Route struct {
	Destination string
	NextHop     string
	Interface   string
	Metric      int
	AdminDist   int
	Protocol    string
	Age         time.Time
}

// VLAN represents a Virtual LAN
type VLAN struct {
	ID          int
	Name        string
	Ports       map[string]bool
	Broadcast   bool
	Multicast   bool
	UnknownUnicast bool
}

// NetworkLink represents a connection between nodes
type NetworkLink struct {
	From     *NetworkNode
	To       *NetworkNode
	FromPort string
	ToPort   string
	Cost     int
	Bandwidth int // Mbps
	Status   string // "up", "down"
}

// NetworkTopology represents the entire network
type NetworkTopology struct {
	Nodes map[string]*NetworkNode
	Links map[string]*NetworkLink
	mutex sync.RWMutex
}

// NewNetworkNode creates a new network node
func NewNetworkNode(id, nodeType string) *NetworkNode {
	return &NetworkNode{
		ID:        id,
		Type:      nodeType,
		Ports:     make(map[string]*Port),
		MACTable:  make(map[string]string),
		Routes:    make(map[string]*Route),
		Neighbors: make(map[string]*NetworkNode),
		VLANs:     make(map[int]*VLAN),
	}
}

// AddPort adds a port to the node
func (n *NetworkNode) AddPort(portID, portType string, speed int) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	
	n.Ports[portID] = &Port{
		ID:     portID,
		Type:   portType,
		Speed:  speed,
		Duplex: "full",
		Status: "up",
		VLAN:   1, // Default VLAN
	}
}

// AddRoute adds a route to the routing table
func (n *NetworkNode) AddRoute(dest, nextHop, intf string, metric, adminDist int, protocol string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	
	n.Routes[dest] = &Route{
		Destination: dest,
		NextHop:     nextHop,
		Interface:   intf,
		Metric:      metric,
		AdminDist:   adminDist,
		Protocol:    protocol,
		Age:         time.Now(),
	}
}

// LookupRoute finds the best route to a destination
func (n *NetworkNode) LookupRoute(destination string) *Route {
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	
	// First try exact match
	if route, exists := n.Routes[destination]; exists {
		return route
	}
	
	// Then try longest prefix match
	var bestRoute *Route
	var bestPrefixLen int
	
	for dest, route := range n.Routes {
		if n.isSubnetMatch(destination, dest) {
			prefixLen := n.getPrefixLength(dest)
			if prefixLen > bestPrefixLen {
				bestPrefixLen = prefixLen
				bestRoute = route
			}
		}
	}
	
	return bestRoute
}

// isSubnetMatch checks if an IP matches a subnet
func (n *NetworkNode) isSubnetMatch(ip, subnet string) bool {
	// Simplified implementation - in reality, you'd parse IPs and subnets
	return len(ip) >= len(subnet) && ip[:len(subnet)] == subnet
}

// getPrefixLength returns the prefix length of a subnet
func (n *NetworkNode) getPrefixLength(subnet string) int {
	// Simplified implementation - in reality, you'd parse CIDR notation
	return len(subnet)
}

// LearnMAC learns a MAC address on a port
func (n *NetworkNode) LearnMAC(mac, portID string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	
	n.MACTable[mac] = portID
}

// LookupMAC finds the port for a MAC address
func (n *NetworkNode) LookupMAC(mac string) (string, bool) {
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	
	port, exists := n.MACTable[mac]
	return port, exists
}

// ForwardFrame forwards a frame based on MAC address
func (n *NetworkNode) ForwardFrame(srcMAC, dstMAC string, frame []byte) {
	if n.Type != "switch" {
		return
	}
	
	// Learn source MAC
	n.LearnMAC(srcMAC, "incoming")
	
	// Look up destination MAC
	if port, exists := n.LookupMAC(dstMAC); exists {
		// Forward to specific port
		fmt.Printf("Switch %s: Forwarding frame from %s to %s via port %s\n", 
			n.ID, srcMAC, dstMAC, port)
	} else {
		// Flood to all ports except incoming
		fmt.Printf("Switch %s: Flooding frame from %s to %s (unknown destination)\n", 
			n.ID, srcMAC, dstMAC)
	}
}

// CreateVLAN creates a new VLAN
func (n *NetworkNode) CreateVLAN(id int, name string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	
	n.VLANs[id] = &VLAN{
		ID:          id,
		Name:        name,
		Ports:       make(map[string]bool),
		Broadcast:   true,
		Multicast:   true,
		UnknownUnicast: true,
	}
}

// AddPortToVLAN adds a port to a VLAN
func (n *NetworkNode) AddPortToVLAN(portID string, vlanID int) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	
	if vlan, exists := n.VLANs[vlanID]; exists {
		vlan.Ports[portID] = true
		if port, exists := n.Ports[portID]; exists {
			port.VLAN = vlanID
		}
	}
}

// NewNetworkTopology creates a new network topology
func NewNetworkTopology() *NetworkTopology {
	return &NetworkTopology{
		Nodes: make(map[string]*NetworkNode),
		Links: make(map[string]*NetworkLink),
	}
}

// AddNode adds a node to the topology
func (nt *NetworkTopology) AddNode(node *NetworkNode) {
	nt.mutex.Lock()
	defer nt.mutex.Unlock()
	
	nt.Nodes[node.ID] = node
}

// AddLink adds a link between two nodes
func (nt *NetworkTopology) AddLink(from, to *NetworkNode, fromPort, toPort string, cost, bandwidth int) {
	nt.mutex.Lock()
	defer nt.mutex.Unlock()
	
	linkID := fmt.Sprintf("%s-%s", from.ID, to.ID)
	link := &NetworkLink{
		From:       from,
		To:         to,
		FromPort:   fromPort,
		ToPort:     toPort,
		Cost:       cost,
		Bandwidth:  bandwidth,
		Status:     "up",
	}
	
	nt.Links[linkID] = link
	
	// Add neighbors
	from.Neighbors[to.ID] = to
	to.Neighbors[from.ID] = from
}

// DijkstraShortestPath finds the shortest path between two nodes
func (nt *NetworkTopology) DijkstraShortestPath(source, destination string) ([]string, int) {
	nt.mutex.RLock()
	defer nt.mutex.RUnlock()
	
	// Initialize distances
	distances := make(map[string]int)
	previous := make(map[string]string)
	visited := make(map[string]bool)
	
	// Set all distances to infinity
	for nodeID := range nt.Nodes {
		distances[nodeID] = math.MaxInt32
	}
	
	// Set source distance to 0
	distances[source] = 0
	
	// Priority queue for unvisited nodes
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	
	// Add source to priority queue
	heap.Push(&pq, &Item{Value: source, Priority: 0})
	
	for pq.Len() > 0 {
		// Get node with minimum distance
		item := heap.Pop(&pq).(*Item)
		current := item.Value
		
		if visited[current] {
			continue
		}
		
		visited[current] = true
		
		// If we reached the destination, we're done
		if current == destination {
			break
		}
		
		// Check all neighbors
		if node, exists := nt.Nodes[current]; exists {
			for neighborID, neighbor := range node.Neighbors {
				if visited[neighborID] {
					continue
				}
				
				// Find link cost
				linkCost := nt.getLinkCost(current, neighborID)
				if linkCost == -1 {
					continue
				}
				
				// Calculate new distance
				newDist := distances[current] + linkCost
				
				// Update if we found a shorter path
				if newDist < distances[neighborID] {
					distances[neighborID] = newDist
					previous[neighborID] = current
					heap.Push(&pq, &Item{Value: neighborID, Priority: newDist})
				}
			}
		}
	}
	
	// Reconstruct path
	path := []string{}
	if distances[destination] == math.MaxInt32 {
		return path, -1 // No path found
	}
	
	current := destination
	for current != "" {
		path = append([]string{current}, path...)
		current = previous[current]
	}
	
	return path, distances[destination]
}

// getLinkCost returns the cost of a link between two nodes
func (nt *NetworkTopology) getLinkCost(from, to string) int {
	linkID := fmt.Sprintf("%s-%s", from, to)
	if link, exists := nt.Links[linkID]; exists && link.Status == "up" {
		return link.Cost
	}
	
	// Check reverse link
	linkID = fmt.Sprintf("%s-%s", to, from)
	if link, exists := nt.Links[linkID]; exists && link.Status == "up" {
		return link.Cost
	}
	
	return -1 // No link found
}

// PriorityQueue implementation for Dijkstra's algorithm
type Item struct {
	Value    string
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// RIP (Routing Information Protocol) implementation
type RIPRouter struct {
	Node        *NetworkNode
	Routes      map[string]*RIPRoute
	Neighbors   map[string]*RIPNeighbor
	UpdateTimer *time.Timer
	mutex       sync.RWMutex
}

type RIPRoute struct {
	Destination string
	NextHop     string
	Metric      int
	Age         time.Time
}

type RIPNeighbor struct {
	Address string
	LastSeen time.Time
	Routes   map[string]*RIPRoute
}

// NewRIPRouter creates a new RIP router
func NewRIPRouter(node *NetworkNode) *RIPRouter {
	router := &RIPRouter{
		Node:      node,
		Routes:    make(map[string]*RIPRoute),
		Neighbors: make(map[string]*RIPNeighbor),
	}
	
	// Start periodic updates
	router.startPeriodicUpdates()
	
	return router
}

// startPeriodicUpdates starts the periodic route updates
func (r *RIPRouter) startPeriodicUpdates() {
	r.UpdateTimer = time.NewTimer(30 * time.Second)
	go func() {
		for {
			select {
			case <-r.UpdateTimer.C:
				r.sendPeriodicUpdate()
				r.UpdateTimer.Reset(30 * time.Second)
			}
		}
	}()
}

// sendPeriodicUpdate sends periodic route updates to neighbors
func (r *RIPRouter) sendPeriodicUpdate() {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	fmt.Printf("RIP Router %s: Sending periodic update\n", r.Node.ID)
	
	for _, neighbor := range r.Neighbors {
		r.sendUpdateToNeighbor(neighbor)
	}
}

// sendUpdateToNeighbor sends route update to a specific neighbor
func (r *RIPRouter) sendUpdateToNeighbor(neighbor *RIPNeighbor) {
	fmt.Printf("RIP Router %s: Sending update to neighbor %s\n", r.Node.ID, neighbor.Address)
	
	// In a real implementation, this would send actual RIP packets
	// For simulation, we'll just print the routes
	for dest, route := range r.Routes {
		fmt.Printf("  Route: %s -> %s (metric: %d)\n", dest, route.NextHop, route.Metric)
	}
}

// AddRoute adds a route to the RIP routing table
func (r *RIPRouter) AddRoute(destination, nextHop string, metric int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	r.Routes[destination] = &RIPRoute{
		Destination: destination,
		NextHop:     nextHop,
		Metric:      metric,
		Age:         time.Now(),
	}
	
	// Update the node's routing table
	r.Node.AddRoute(destination, nextHop, "eth0", metric, 120, "RIP")
}

// OSPF (Open Shortest Path First) implementation
type OSPFRouter struct {
	Node        *NetworkNode
	RouterID    string
	AreaID      int
	LSDB        map[string]*LSA // Link State Database
	Neighbors   map[string]*OSPFNeighbor
	SPF         *SPFCalculator
	mutex       sync.RWMutex
}

type LSA struct {
	Type       int
	RouterID   string
	AreaID     int
	Sequence   int
	Age        int
	Data       interface{}
}

type OSPFNeighbor struct {
	RouterID    string
	State       string
	Priority    int
	LastSeen    time.Time
	LSAs        map[string]*LSA
}

type SPFCalculator struct {
	RouterID string
	LSDB     map[string]*LSA
}

// NewOSPFRouter creates a new OSPF router
func NewOSPFRouter(node *NetworkNode, routerID string, areaID int) *OSPFRouter {
	router := &OSPFRouter{
		Node:      node,
		RouterID:  routerID,
		AreaID:    areaID,
		LSDB:      make(map[string]*LSA),
		Neighbors: make(map[string]*OSPFNeighbor),
		SPF:       &SPFCalculator{RouterID: routerID, LSDB: make(map[string]*LSA)},
	}
	
	// Start OSPF processes
	go router.startOSPFProcesses()
	
	return router
}

// startOSPFProcesses starts OSPF background processes
func (o *OSPFRouter) startOSPFProcesses() {
	// Hello timer
	helloTimer := time.NewTicker(10 * time.Second)
	defer helloTimer.Stop()
	
	// LSA refresh timer
	lsaTimer := time.NewTicker(30 * time.Minute)
	defer lsaTimer.Stop()
	
	for {
		select {
		case <-helloTimer.C:
			o.sendHelloPackets()
		case <-lsaTimer.C:
			o.refreshLSAs()
		}
	}
}

// sendHelloPackets sends OSPF hello packets
func (o *OSPFRouter) sendHelloPackets() {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	
	fmt.Printf("OSPF Router %s: Sending hello packets\n", o.RouterID)
	
	for _, neighbor := range o.Neighbors {
		fmt.Printf("  Hello to neighbor %s\n", neighbor.RouterID)
	}
}

// refreshLSAs refreshes OSPF LSAs
func (o *OSPFRouter) refreshLSAs() {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	
	fmt.Printf("OSPF Router %s: Refreshing LSAs\n", o.RouterID)
	
	// In a real implementation, this would refresh LSA timers
	// and send LSA updates if needed
}

// AddNeighbor adds an OSPF neighbor
func (o *OSPFRouter) AddNeighbor(routerID string, priority int) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	
	o.Neighbors[routerID] = &OSPFNeighbor{
		RouterID: routerID,
		State:    "Down",
		Priority: priority,
		LastSeen: time.Now(),
		LSAs:     make(map[string]*LSA),
	}
}

// CalculateSPF calculates the shortest path first tree
func (o *OSPFRouter) CalculateSPF() {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	
	fmt.Printf("OSPF Router %s: Calculating SPF tree\n", o.RouterID)
	
	// In a real implementation, this would run Dijkstra's algorithm
	// on the link state database to calculate shortest paths
}

// NetworkSimulator simulates network behavior
type NetworkSimulator struct {
	Topology *NetworkTopology
	Routers  map[string]*RIPRouter
	OSPFRouters map[string]*OSPFRouter
	mutex    sync.RWMutex
}

// NewNetworkSimulator creates a new network simulator
func NewNetworkSimulator() *NetworkSimulator {
	return &NetworkSimulator{
		Topology:    NewNetworkTopology(),
		Routers:     make(map[string]*RIPRouter),
		OSPFRouters: make(map[string]*OSPFRouter),
	}
}

// AddRIPRouter adds a RIP router to the simulation
func (ns *NetworkSimulator) AddRIPRouter(node *NetworkNode) {
	ns.mutex.Lock()
	defer ns.mutex.Unlock()
	
	ns.Routers[node.ID] = NewRIPRouter(node)
}

// AddOSPFRouter adds an OSPF router to the simulation
func (ns *NetworkSimulator) AddOSPFRouter(node *NetworkNode, routerID string, areaID int) {
	ns.mutex.Lock()
	defer ns.mutex.Unlock()
	
	ns.OSPFRouters[node.ID] = NewOSPFRouter(node, routerID, areaID)
}

// SimulateTraffic simulates network traffic
func (ns *NetworkSimulator) SimulateTraffic(src, dst string, data []byte) {
	ns.mutex.RLock()
	defer ns.mutex.RUnlock()
	
	fmt.Printf("Simulating traffic from %s to %s\n", src, dst)
	
	// Find path using Dijkstra's algorithm
	path, cost := ns.Topology.DijkstraShortestPath(src, dst)
	
	if len(path) == 0 {
		fmt.Printf("No path found from %s to %s\n", src, dst)
		return
	}
	
	fmt.Printf("Path: %v (cost: %d)\n", path, cost)
	
	// Simulate packet forwarding along the path
	for i := 0; i < len(path)-1; i++ {
		current := path[i]
		next := path[i+1]
		
		fmt.Printf("  %s -> %s\n", current, next)
		
		// Simulate processing delay
		time.Sleep(10 * time.Millisecond)
	}
}

// Demonstrate network switching
func demonstrateSwitching() {
	fmt.Println("=== Network Switching Demo ===\n")
	
	// Create switches
	sw1 := NewNetworkNode("SW1", "switch")
	sw2 := NewNetworkNode("SW2", "switch")
	
	// Add ports
	sw1.AddPort("eth0", "ethernet", 1000)
	sw1.AddPort("eth1", "ethernet", 1000)
	sw1.AddPort("eth2", "ethernet", 1000)
	
	sw2.AddPort("eth0", "ethernet", 1000)
	sw2.AddPort("eth1", "ethernet", 1000)
	
	// Create VLANs
	sw1.CreateVLAN(10, "Sales")
	sw1.CreateVLAN(20, "Engineering")
	
	// Add ports to VLANs
	sw1.AddPortToVLAN("eth0", 10)
	sw1.AddPortToVLAN("eth1", 20)
	sw1.AddPortToVLAN("eth2", 10)
	
	// Simulate MAC learning
	sw1.LearnMAC("00:11:22:33:44:55", "eth0")
	sw1.LearnMAC("00:11:22:33:44:66", "eth1")
	sw1.LearnMAC("00:11:22:33:44:77", "eth2")
	
	// Simulate frame forwarding
	sw1.ForwardFrame("00:11:22:33:44:55", "00:11:22:33:44:77", []byte("Hello"))
	sw1.ForwardFrame("00:11:22:33:44:55", "00:11:22:33:44:99", []byte("Unknown"))
	
	// Display MAC table
	fmt.Printf("Switch %s MAC Table:\n", sw1.ID)
	for mac, port := range sw1.MACTable {
		fmt.Printf("  %s -> %s\n", mac, port)
	}
}

// Demonstrate network routing
func demonstrateRouting() {
	fmt.Println("=== Network Routing Demo ===\n")
	
	// Create routers
	r1 := NewNetworkNode("R1", "router")
	r2 := NewNetworkNode("R2", "router")
	r3 := NewNetworkNode("R3", "router")
	r4 := NewNetworkNode("R4", "router")
	
	// Add routes
	r1.AddRoute("192.168.1.0/24", "192.168.1.1", "eth0", 0, 0, "connected")
	r1.AddRoute("10.0.0.0/8", "192.168.1.2", "eth0", 1, 120, "RIP")
	r1.AddRoute("172.16.0.0/16", "192.168.1.3", "eth0", 2, 120, "RIP")
	
	r2.AddRoute("10.0.0.0/8", "10.0.0.1", "eth0", 0, 0, "connected")
	r2.AddRoute("192.168.1.0/24", "192.168.1.1", "eth1", 1, 120, "RIP")
	
	r3.AddRoute("172.16.0.0/16", "172.16.0.1", "eth0", 0, 0, "connected")
	r3.AddRoute("192.168.1.0/24", "192.168.1.1", "eth1", 2, 120, "RIP")
	
	// Test route lookups
	destinations := []string{"192.168.1.10", "10.0.0.5", "172.16.0.10", "8.8.8.8"}
	
	for _, dest := range destinations {
		route := r1.LookupRoute(dest)
		if route != nil {
			fmt.Printf("Route to %s: Next hop %s via %s (metric: %d)\n", 
				dest, route.NextHop, route.Interface, route.Metric)
		} else {
			fmt.Printf("No route found to %s\n", dest)
		}
	}
}

// Demonstrate network topology
func demonstrateTopology() {
	fmt.Println("=== Network Topology Demo ===\n")
	
	// Create network topology
	topology := NewNetworkTopology()
	
	// Create nodes
	nodes := []*NetworkNode{
		NewNetworkNode("R1", "router"),
		NewNetworkNode("R2", "router"),
		NewNetworkNode("R3", "router"),
		NewNetworkNode("R4", "router"),
		NewNetworkNode("SW1", "switch"),
		NewNetworkNode("SW2", "switch"),
	}
	
	// Add nodes to topology
	for _, node := range nodes {
		topology.AddNode(node)
	}
	
	// Add links
	topology.AddLink(nodes[0], nodes[1], "eth0", "eth0", 1, 1000)
	topology.AddLink(nodes[1], nodes[2], "eth1", "eth0", 2, 1000)
	topology.AddLink(nodes[2], nodes[3], "eth1", "eth0", 1, 1000)
	topology.AddLink(nodes[0], nodes[4], "eth1", "eth0", 1, 1000)
	topology.AddLink(nodes[4], nodes[5], "eth1", "eth0", 1, 1000)
	
	// Test shortest path calculation
	path, cost := topology.DijkstraShortestPath("R1", "R4")
	fmt.Printf("Shortest path from R1 to R4: %v (cost: %d)\n", path, cost)
	
	path, cost = topology.DijkstraShortestPath("R1", "SW2")
	fmt.Printf("Shortest path from R1 to SW2: %v (cost: %d)\n", path, cost)
	
	// Test unreachable destination
	path, cost = topology.DijkstraShortestPath("R1", "R5")
	if len(path) == 0 {
		fmt.Printf("No path found from R1 to R5\n")
	}
}

// Demonstrate RIP protocol
func demonstrateRIP() {
	fmt.Println("=== RIP Protocol Demo ===\n")
	
	// Create routers
	r1 := NewNetworkNode("R1", "router")
	r2 := NewNetworkNode("R2", "router")
	r3 := NewNetworkNode("R3", "router")
	
	// Create RIP routers
	rip1 := NewRIPRouter(r1)
	rip2 := NewRIPRouter(r2)
	rip3 := NewRIPRouter(r3)
	
	// Add routes
	rip1.AddRoute("192.168.1.0/24", "192.168.1.1", 0) // Directly connected
	rip1.AddRoute("10.0.0.0/8", "192.168.1.2", 1)    // Via R2
	
	rip2.AddRoute("10.0.0.0/8", "10.0.0.1", 0)       // Directly connected
	rip2.AddRoute("192.168.1.0/24", "192.168.1.1", 1) // Via R1
	rip2.AddRoute("172.16.0.0/16", "10.0.0.3", 1)     // Via R3
	
	rip3.AddRoute("172.16.0.0/16", "172.16.0.1", 0)   // Directly connected
	rip3.AddRoute("10.0.0.0/8", "10.0.0.2", 1)        // Via R2
	
	// Display routing tables
	fmt.Printf("R1 Routing Table:\n")
	for dest, route := range rip1.Routes {
		fmt.Printf("  %s -> %s (metric: %d)\n", dest, route.NextHop, route.Metric)
	}
	
	fmt.Printf("\nR2 Routing Table:\n")
	for dest, route := range rip2.Routes {
		fmt.Printf("  %s -> %s (metric: %d)\n", dest, route.NextHop, route.Metric)
	}
	
	fmt.Printf("\nR3 Routing Table:\n")
	for dest, route := range rip3.Routes {
		fmt.Printf("  %s -> %s (metric: %d)\n", dest, route.NextHop, route.Metric)
	}
}

// Demonstrate OSPF protocol
func demonstrateOSPF() {
	fmt.Println("=== OSPF Protocol Demo ===\n")
	
	// Create routers
	r1 := NewNetworkNode("R1", "router")
	r2 := NewNetworkNode("R2", "router")
	r3 := NewNetworkNode("R3", "router")
	
	// Create OSPF routers
	ospf1 := NewOSPFRouter(r1, "1.1.1.1", 0)
	ospf2 := NewOSPFRouter(r2, "2.2.2.2", 0)
	ospf3 := NewOSPFRouter(r3, "3.3.3.3", 0)
	
	// Add neighbors
	ospf1.AddNeighbor("2.2.2.2", 1)
	ospf1.AddNeighbor("3.3.3.3", 1)
	
	ospf2.AddNeighbor("1.1.1.1", 1)
	ospf2.AddNeighbor("3.3.3.3", 1)
	
	ospf3.AddNeighbor("1.1.1.1", 1)
	ospf3.AddNeighbor("2.2.2.2", 1)
	
	// Display OSPF information
	fmt.Printf("OSPF Router %s:\n", ospf1.RouterID)
	fmt.Printf("  Area ID: %d\n", ospf1.AreaID)
	fmt.Printf("  Neighbors: %d\n", len(ospf1.Neighbors))
	
	for neighborID, neighbor := range ospf1.Neighbors {
		fmt.Printf("    %s (state: %s, priority: %d)\n", 
			neighborID, neighbor.State, neighbor.Priority)
	}
}

// Demonstrate network simulation
func demonstrateSimulation() {
	fmt.Println("=== Network Simulation Demo ===\n")
	
	// Create network simulator
	simulator := NewNetworkSimulator()
	
	// Create topology
	topology := simulator.Topology
	
	// Create nodes
	r1 := NewNetworkNode("R1", "router")
	r2 := NewNetworkNode("R2", "router")
	r3 := NewNetworkNode("R3", "router")
	r4 := NewNetworkNode("R4", "router")
	
	// Add nodes
	topology.AddNode(r1)
	topology.AddNode(r2)
	topology.AddNode(r3)
	topology.AddNode(r4)
	
	// Add links
	topology.AddLink(r1, r2, "eth0", "eth0", 1, 1000)
	topology.AddLink(r2, r3, "eth1", "eth0", 2, 1000)
	topology.AddLink(r3, r4, "eth1", "eth0", 1, 1000)
	topology.AddLink(r1, r4, "eth1", "eth1", 3, 1000)
	
	// Add RIP routers
	simulator.AddRIPRouter(r1)
	simulator.AddRIPRouter(r2)
	simulator.AddRIPRouter(r3)
	simulator.AddRIPRouter(r4)
	
	// Simulate traffic
	simulator.SimulateTraffic("R1", "R4", []byte("Hello from R1 to R4"))
	simulator.SimulateTraffic("R2", "R4", []byte("Hello from R2 to R4"))
	simulator.SimulateTraffic("R1", "R3", []byte("Hello from R1 to R3"))
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 9: Routing & Switching")
	fmt.Println("================================================================\n")
	
	// Run all demonstrations
	demonstrateSwitching()
	fmt.Println()
	demonstrateRouting()
	fmt.Println()
	demonstrateTopology()
	fmt.Println()
	demonstrateRIP()
	fmt.Println()
	demonstrateOSPF()
	fmt.Println()
	demonstrateSimulation()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Switching operates at Layer 2 and forwards frames within LANs")
	fmt.Println("2. Routing operates at Layer 3 and forwards packets between networks")
	fmt.Println("3. VLANs provide logical network segmentation")
	fmt.Println("4. STP prevents loops in switched networks")
	fmt.Println("5. Routing algorithms find optimal paths through networks")
	fmt.Println("6. RIP and OSPF are common routing protocols")
	fmt.Println("7. Network topology affects performance and reliability")
	fmt.Println("8. Go can simulate complex network behavior")
	
	fmt.Println("\n📚 Next Topic: Network Performance & Optimization")
}
