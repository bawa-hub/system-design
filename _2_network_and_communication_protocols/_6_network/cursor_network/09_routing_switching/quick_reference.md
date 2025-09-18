# Routing & Switching Quick Reference Guide

## 🚀 Essential Concepts

### Switching (Layer 2)
- **MAC Address Learning**: Switches learn MAC addresses by observing traffic
- **Frame Forwarding**: Forward frames based on destination MAC address
- **VLAN Support**: Virtual LAN segmentation and management
- **STP**: Spanning Tree Protocol prevents loops
- **Port States**: Blocking, Listening, Learning, Forwarding, Disabled

### Routing (Layer 3)
- **Packet Forwarding**: Forward packets between different networks
- **Path Selection**: Choose best route to destination
- **Route Management**: Maintain routing tables
- **Protocol Support**: Implement routing protocols (RIP, OSPF, BGP)

## 🔧 Go Implementation Patterns

### Network Node Structure
```go
type NetworkNode struct {
    ID        string
    Type      string // "router" or "switch"
    Ports     map[string]*Port
    MACTable  map[string]string
    Routes    map[string]*Route
    Neighbors map[string]*NetworkNode
    VLANs     map[int]*VLAN
    mutex     sync.RWMutex
}
```

### Port Management
```go
type Port struct {
    ID     string
    Type   string // "ethernet", "serial", "fiber"
    Speed  int    // Mbps
    Duplex string // "half", "full"
    Status string // "up", "down", "admin_down"
    VLAN   int
}
```

### Route Management
```go
type Route struct {
    Destination string
    NextHop     string
    Interface   string
    Metric      int
    AdminDist   int
    Protocol    string
    Age         time.Time
}
```

## 🔄 Switching Implementation

### MAC Address Learning
```go
func (n *NetworkNode) LearnMAC(mac, portID string) {
    n.mutex.Lock()
    defer n.mutex.Unlock()
    
    n.MACTable[mac] = portID
}

func (n *NetworkNode) LookupMAC(mac string) (string, bool) {
    n.mutex.RLock()
    defer n.mutex.RUnlock()
    
    port, exists := n.MACTable[mac]
    return port, exists
}
```

### Frame Forwarding
```go
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
```

### VLAN Management
```go
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
```

## 🛣️ Routing Implementation

### Route Lookup
```go
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
```

### Route Addition
```go
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
```

## 🧮 Routing Algorithms

### Dijkstra's Algorithm
```go
func (g *Graph) DijkstraShortestPath(source, destination string) ([]string, int) {
    // Initialize distances
    distances := make(map[string]int)
    previous := make(map[string]string)
    visited := make(map[string]bool)
    
    // Set all distances to infinity
    for nodeID := range g.Nodes {
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
        item := heap.Pop(&pq).(*Item)
        current := item.Value
        
        if visited[current] {
            continue
        }
        
        visited[current] = true
        
        if current == destination {
            break
        }
        
        // Check all neighbors
        if node, exists := g.Nodes[current]; exists {
            for neighborID, neighbor := range node.Neighbors {
                if visited[neighborID] {
                    continue
                }
                
                edge := g.GetEdge(current, neighborID)
                if edge == nil || edge.Status != "up" {
                    continue
                }
                
                newDist := distances[current] + edge.Weight
                
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
        return path, -1
    }
    
    current := destination
    for current != "" {
        path = append([]string{current}, path...)
        current = previous[current]
    }
    
    return path, distances[destination]
}
```

### Bellman-Ford Algorithm
```go
func (g *Graph) BellmanFordShortestPath(source, destination string) ([]string, int) {
    // Initialize distances
    distances := make(map[string]int)
    previous := make(map[string]string)
    
    // Set all distances to infinity
    for nodeID := range g.Nodes {
        distances[nodeID] = math.MaxInt32
    }
    
    // Set source distance to 0
    distances[source] = 0
    
    // Relax edges |V| - 1 times
    for i := 0; i < len(g.Nodes)-1; i++ {
        for _, edge := range g.Edges {
            if edge.Status != "up" {
                continue
            }
            
            from := edge.From.ID
            to := edge.To.ID
            
            if distances[from] != math.MaxInt32 && distances[from]+edge.Weight < distances[to] {
                distances[to] = distances[from] + edge.Weight
                previous[to] = from
            }
        }
    }
    
    // Check for negative cycles
    for _, edge := range g.Edges {
        if edge.Status != "up" {
            continue
        }
        
        from := edge.From.ID
        to := edge.To.ID
        
        if distances[from] != math.MaxInt32 && distances[from]+edge.Weight < distances[to] {
            return []string{}, -1 // Negative cycle detected
        }
    }
    
    // Reconstruct path
    path := []string{}
    if distances[destination] == math.MaxInt32 {
        return path, -1
    }
    
    current := destination
    for current != "" {
        path = append([]string{current}, path...)
        current = previous[current]
    }
    
    return path, distances[destination]
}
```

### A* Algorithm
```go
func (g *Graph) AStarShortestPath(source, destination string) ([]string, int) {
    // Heuristic function (Euclidean distance)
    heuristic := func(nodeID string) int {
        if node, exists := g.Nodes[nodeID]; exists {
            dest, exists := g.Nodes[destination]
            if !exists {
                return 0
            }
            
            dx := node.Position.X - dest.Position.X
            dy := node.Position.Y - dest.Position.Y
            return int(math.Sqrt(dx*dx + dy*dy))
        }
        return 0
    }
    
    // Initialize distances
    distances := make(map[string]int)
    previous := make(map[string]string)
    visited := make(map[string]bool)
    
    // Set all distances to infinity
    for nodeID := range g.Nodes {
        distances[nodeID] = math.MaxInt32
    }
    
    // Set source distance to 0
    distances[source] = 0
    
    // Priority queue for unvisited nodes
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    
    // Add source to priority queue
    heap.Push(&pq, &Item{Value: source, Priority: heuristic(source)})
    
    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item)
        current := item.Value
        
        if visited[current] {
            continue
        }
        
        visited[current] = true
        
        if current == destination {
            break
        }
        
        // Check all neighbors
        if node, exists := g.Nodes[current]; exists {
            for neighborID, neighbor := range node.Neighbors {
                if visited[neighborID] {
                    continue
                }
                
                edge := g.GetEdge(current, neighborID)
                if edge == nil || edge.Status != "up" {
                    continue
                }
                
                newDist := distances[current] + edge.Weight
                
                if newDist < distances[neighborID] {
                    distances[neighborID] = newDist
                    previous[neighborID] = current
                    
                    // Calculate f-score (g + h)
                    fScore := newDist + heuristic(neighborID)
                    heap.Push(&pq, &Item{Value: neighborID, Priority: fScore})
                }
            }
        }
    }
    
    // Reconstruct path
    path := []string{}
    if distances[destination] == math.MaxInt32 {
        return path, -1
    }
    
    current := destination
    for current != "" {
        path = append([]string{current}, path...)
        current = previous[current]
    }
    
    return path, distances[destination]
}
```

## 🌐 Network Topology

### Graph Representation
```go
type Graph struct {
    Nodes map[string]*Node
    Edges map[string]*Edge
    mutex sync.RWMutex
}

type Node struct {
    ID        string
    Type      string
    Position  Point
    Neighbors map[string]*Node
    Routes    map[string]*Route
}

type Edge struct {
    From      *Node
    To        *Node
    Weight    int
    Bandwidth int
    Delay     int
    Status    string
}
```

### Topology Management
```go
func (g *Graph) AddNode(id, nodeType string, x, y float64) {
    g.mutex.Lock()
    defer g.mutex.Unlock()
    
    g.Nodes[id] = &Node{
        ID:        id,
        Type:      nodeType,
        Position:  Point{X: x, Y: y},
        Neighbors: make(map[string]*Node),
        Routes:    make(map[string]*Route),
    }
}

func (g *Graph) AddEdge(from, to string, weight, bandwidth, delay int) {
    g.mutex.Lock()
    defer g.mutex.Unlock()
    
    fromNode, exists1 := g.Nodes[from]
    toNode, exists2 := g.Nodes[to]
    
    if !exists1 || !exists2 {
        return
    }
    
    edgeID := fmt.Sprintf("%s-%s", from, to)
    edge := &Edge{
        From:      fromNode,
        To:        toNode,
        Weight:    weight,
        Bandwidth: bandwidth,
        Delay:     delay,
        Status:    "up",
    }
    
    g.Edges[edgeID] = edge
    
    // Add neighbors
    fromNode.Neighbors[to] = toNode
    toNode.Neighbors[from] = fromNode
}
```

## 🔄 Routing Protocols

### RIP Implementation
```go
type RIPRouter struct {
    Node        *NetworkNode
    Routes      map[string]*RIPRoute
    Neighbors   map[string]*RIPNeighbor
    UpdateTimer *time.Timer
    mutex       sync.RWMutex
}

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
```

### OSPF Implementation
```go
type OSPFRouter struct {
    Node        *NetworkNode
    RouterID    string
    AreaID      int
    LSDB        map[string]*LSA
    Neighbors   map[string]*OSPFNeighbor
    SPF         *SPFCalculator
    mutex       sync.RWMutex
}

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
```

## 📊 Network Analysis

### Connectivity Analysis
```go
func (na *NetworkAnalyzer) AnalyzeConnectivity() map[string]bool {
    connectivity := make(map[string]bool)
    
    for nodeID := range na.Graph.Nodes {
        reachable := true
        for otherNodeID := range na.Graph.Nodes {
            if nodeID == otherNodeID {
                continue
            }
            
            path, _ := na.Graph.DijkstraShortestPath(otherNodeID, nodeID)
            if len(path) == 0 {
                reachable = false
                break
            }
        }
        connectivity[nodeID] = reachable
    }
    
    return connectivity
}
```

### Performance Analysis
```go
func (na *NetworkAnalyzer) AnalyzePerformance() map[string]float64 {
    performance := make(map[string]float64)
    
    // Calculate average path length
    totalPaths := 0
    totalLength := 0
    
    for from := range na.Graph.Nodes {
        for to := range na.Graph.Nodes {
            if from == to {
                continue
            }
            
            path, _ := na.Graph.DijkstraShortestPath(from, to)
            if len(path) > 0 {
                totalPaths++
                totalLength += len(path) - 1
            }
        }
    }
    
    if totalPaths > 0 {
        performance["average_path_length"] = float64(totalLength) / float64(totalPaths)
    }
    
    return performance
}
```

## 🎯 Best Practices

1. **Use appropriate algorithms**: Dijkstra for single-source, Bellman-Ford for negative weights
2. **Implement proper concurrency**: Use mutexes for thread-safe operations
3. **Handle edge cases**: Check for cycles, disconnected components
4. **Optimize performance**: Use priority queues, efficient data structures
5. **Monitor network state**: Track link status, route changes
6. **Implement redundancy**: Multiple paths, failover mechanisms
7. **Use proper metrics**: Bandwidth, delay, reliability
8. **Test thoroughly**: Validate algorithms with various topologies

## 🚨 Common Issues

### Switching Issues
- **MAC table overflow**: Implement aging and limits
- **Broadcast storms**: Use STP and VLANs
- **VLAN misconfiguration**: Validate VLAN assignments
- **Port security**: Implement MAC address limits

### Routing Issues
- **Routing loops**: Use proper loop prevention
- **Convergence time**: Optimize algorithm performance
- **Route flapping**: Implement dampening
- **Memory usage**: Limit routing table size

### Algorithm Issues
- **Infinite loops**: Check for cycles and disconnected components
- **Performance**: Use appropriate data structures
- **Accuracy**: Validate algorithm implementations
- **Scalability**: Consider large network topologies

---

**Remember**: Routing and switching are the backbone of network communication. Master these concepts, and you'll understand how data moves through networks! 🚀
