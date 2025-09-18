package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

// Graph represents a network graph
type Graph struct {
	Nodes map[string]*Node
	Edges map[string]*Edge
	mutex sync.RWMutex
}

// Node represents a network node
type Node struct {
	ID       string
	Type     string // "router", "switch", "host"
	Position Point
	Neighbors map[string]*Node
	Routes   map[string]*Route
}

// Edge represents a network link
type Edge struct {
	From     *Node
	To       *Node
	Weight   int
	Bandwidth int
	Delay    int
	Status   string // "up", "down"
}

// Point represents a 2D coordinate
type Point struct {
	X, Y float64
}

// Route represents a routing table entry
type Route struct {
	Destination string
	NextHop     string
	Cost        int
	Path        []string
	Protocol    string
	Age         time.Time
}

// NewGraph creates a new network graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		Edges: make(map[string]*Edge),
	}
}

// AddNode adds a node to the graph
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

// AddEdge adds an edge between two nodes
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

// GetEdge returns an edge between two nodes
func (g *Graph) GetEdge(from, to string) *Edge {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	edgeID := fmt.Sprintf("%s-%s", from, to)
	if edge, exists := g.Edges[edgeID]; exists {
		return edge
	}
	
	// Check reverse edge
	edgeID = fmt.Sprintf("%s-%s", to, from)
	if edge, exists := g.Edges[edgeID]; exists {
		return edge
	}
	
	return nil
}

// DijkstraShortestPath finds the shortest path using Dijkstra's algorithm
func (g *Graph) DijkstraShortestPath(source, destination string) ([]string, int) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
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
		if node, exists := g.Nodes[current]; exists {
			for neighborID, neighbor := range node.Neighbors {
				if visited[neighborID] {
					continue
				}
				
				// Find edge weight
				edge := g.GetEdge(current, neighborID)
				if edge == nil || edge.Status != "up" {
					continue
				}
				
				// Calculate new distance
				newDist := distances[current] + edge.Weight
				
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

// BellmanFordShortestPath finds the shortest path using Bellman-Ford algorithm
func (g *Graph) BellmanFordShortestPath(source, destination string) ([]string, int) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
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
		return path, -1 // No path found
	}
	
	current := destination
	for current != "" {
		path = append([]string{current}, path...)
		current = previous[current]
	}
	
	return path, distances[destination]
}

// FloydWarshallAllPairs finds shortest paths between all pairs of nodes
func (g *Graph) FloydWarshallAllPairs() map[string]map[string]int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	// Initialize distance matrix
	distances := make(map[string]map[string]int)
	
	// Initialize with infinity
	for from := range g.Nodes {
		distances[from] = make(map[string]int)
		for to := range g.Nodes {
			if from == to {
				distances[from][to] = 0
			} else {
				distances[from][to] = math.MaxInt32
			}
		}
	}
	
	// Initialize with edge weights
	for _, edge := range g.Edges {
		if edge.Status == "up" {
			distances[edge.From.ID][edge.To.ID] = edge.Weight
		}
	}
	
	// Floyd-Warshall algorithm
	for k := range g.Nodes {
		for i := range g.Nodes {
			for j := range g.Nodes {
				if distances[i][k] != math.MaxInt32 && distances[k][j] != math.MaxInt32 {
					if distances[i][k]+distances[k][j] < distances[i][j] {
						distances[i][j] = distances[i][k] + distances[k][j]
					}
				}
			}
		}
	}
	
	return distances
}

// AStarShortestPath finds the shortest path using A* algorithm
func (g *Graph) AStarShortestPath(source, destination string) ([]string, int) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
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
		// Get node with minimum f-score
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
		if node, exists := g.Nodes[current]; exists {
			for neighborID, neighbor := range node.Neighbors {
				if visited[neighborID] {
					continue
				}
				
				// Find edge weight
				edge := g.GetEdge(current, neighborID)
				if edge == nil || edge.Status != "up" {
					continue
				}
				
				// Calculate new distance
				newDist := distances[current] + edge.Weight
				
				// Update if we found a shorter path
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
		return path, -1 // No path found
	}
	
	current := destination
	for current != "" {
		path = append([]string{current}, path...)
		current = previous[current]
	}
	
	return path, distances[destination]
}

// KruskalMST finds the minimum spanning tree using Kruskal's algorithm
func (g *Graph) KruskalMST() []*Edge {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	// Create list of all edges
	edges := make([]*Edge, 0, len(g.Edges))
	for _, edge := range g.Edges {
		if edge.Status == "up" {
			edges = append(edges, edge)
		}
	}
	
	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})
	
	// Union-Find data structure
	parent := make(map[string]string)
	rank := make(map[string]int)
	
	// Initialize each node as its own parent
	for nodeID := range g.Nodes {
		parent[nodeID] = nodeID
		rank[nodeID] = 0
	}
	
	// Find function
	find := func(x string) string {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	
	// Union function
	union := func(x, y string) {
		px := find(x)
		py := find(y)
		
		if px == py {
			return
		}
		
		if rank[px] < rank[py] {
			parent[px] = py
		} else if rank[px] > rank[py] {
			parent[py] = px
		} else {
			parent[py] = px
			rank[px]++
		}
	}
	
	// Kruskal's algorithm
	mst := make([]*Edge, 0)
	for _, edge := range edges {
		from := edge.From.ID
		to := edge.To.ID
		
		if find(from) != find(to) {
			mst = append(mst, edge)
			union(from, to)
		}
	}
	
	return mst
}

// PrimMST finds the minimum spanning tree using Prim's algorithm
func (g *Graph) PrimMST(start string) []*Edge {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	// Initialize
	visited := make(map[string]bool)
	mst := make([]*Edge, 0)
	
	// Priority queue for edges
	pq := make(EdgePriorityQueue, 0)
	heap.Init(&pq)
	
	// Start with the given node
	visited[start] = true
	
	// Add all edges from start node to priority queue
	if startNode, exists := g.Nodes[start]; exists {
		for neighborID := range startNode.Neighbors {
			edge := g.GetEdge(start, neighborID)
			if edge != nil && edge.Status == "up" {
				heap.Push(&pq, &EdgeItem{Edge: edge, Priority: edge.Weight})
			}
		}
	}
	
	// Prim's algorithm
	for pq.Len() > 0 && len(mst) < len(g.Nodes)-1 {
		item := heap.Pop(&pq).(*EdgeItem)
		edge := item.Edge
		
		from := edge.From.ID
		to := edge.To.ID
		
		// Skip if both nodes are already visited
		if visited[from] && visited[to] {
			continue
		}
		
		// Add edge to MST
		mst = append(mst, edge)
		
		// Mark the unvisited node as visited
		if !visited[from] {
			visited[from] = true
			// Add edges from the newly visited node
			if fromNode, exists := g.Nodes[from]; exists {
				for neighborID := range fromNode.Neighbors {
					neighborEdge := g.GetEdge(from, neighborID)
					if neighborEdge != nil && neighborEdge.Status == "up" && !visited[neighborID] {
						heap.Push(&pq, &EdgeItem{Edge: neighborEdge, Priority: neighborEdge.Weight})
					}
				}
			}
		} else {
			visited[to] = true
			// Add edges from the newly visited node
			if toNode, exists := g.Nodes[to]; exists {
				for neighborID := range toNode.Neighbors {
					neighborEdge := g.GetEdge(to, neighborID)
					if neighborEdge != nil && neighborEdge.Status == "up" && !visited[neighborID] {
						heap.Push(&pq, &EdgeItem{Edge: neighborEdge, Priority: neighborEdge.Weight})
					}
				}
			}
		}
	}
	
	return mst
}

// EdgePriorityQueue for Prim's algorithm
type EdgeItem struct {
	Edge     *Edge
	Priority int
	Index    int
}

type EdgePriorityQueue []*EdgeItem

func (pq EdgePriorityQueue) Len() int { return len(pq) }

func (pq EdgePriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq EdgePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *EdgePriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*EdgeItem)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *EdgePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// NetworkAnalyzer provides network analysis capabilities
type NetworkAnalyzer struct {
	Graph *Graph
}

// NewNetworkAnalyzer creates a new network analyzer
func NewNetworkAnalyzer(graph *Graph) *NetworkAnalyzer {
	return &NetworkAnalyzer{Graph: graph}
}

// AnalyzeConnectivity analyzes network connectivity
func (na *NetworkAnalyzer) AnalyzeConnectivity() map[string]bool {
	na.Graph.mutex.RLock()
	defer na.Graph.mutex.RUnlock()
	
	connectivity := make(map[string]bool)
	
	for nodeID := range na.Graph.Nodes {
		// Check if node is reachable from all other nodes
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

// AnalyzeRedundancy analyzes network redundancy
func (na *NetworkAnalyzer) AnalyzeRedundancy() map[string]int {
	na.Graph.mutex.RLock()
	defer na.Graph.mutex.RUnlock()
	
	redundancy := make(map[string]int)
	
	for nodeID := range na.Graph.Nodes {
		// Count alternative paths to each other node
		altPaths := 0
		for otherNodeID := range na.Graph.Nodes {
			if nodeID == otherNodeID {
				continue
			}
			
			// Find all paths using different algorithms
			path1, _ := na.Graph.DijkstraShortestPath(nodeID, otherNodeID)
			path2, _ := na.Graph.BellmanFordShortestPath(nodeID, otherNodeID)
			
			if len(path1) > 0 && len(path2) > 0 && !na.pathsEqual(path1, path2) {
				altPaths++
			}
		}
		redundancy[nodeID] = altPaths
	}
	
	return redundancy
}

// pathsEqual checks if two paths are equal
func (na *NetworkAnalyzer) pathsEqual(path1, path2 []string) bool {
	if len(path1) != len(path2) {
		return false
	}
	
	for i := range path1 {
		if path1[i] != path2[i] {
			return false
		}
	}
	
	return true
}

// AnalyzePerformance analyzes network performance
func (na *NetworkAnalyzer) AnalyzePerformance() map[string]float64 {
	na.Graph.mutex.RLock()
	defer na.Graph.mutex.RUnlock()
	
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
	
	// Calculate network diameter
	maxPathLength := 0
	for from := range na.Graph.Nodes {
		for to := range na.Graph.Nodes {
			if from == to {
				continue
			}
			
			path, _ := na.Graph.DijkstraShortestPath(from, to)
			if len(path) > 0 && len(path)-1 > maxPathLength {
				maxPathLength = len(path) - 1
			}
		}
	}
	performance["diameter"] = float64(maxPathLength)
	
	return performance
}

// Demonstrate routing algorithms
func demonstrateRoutingAlgorithms() {
	fmt.Println("=== Routing Algorithms Demo ===\n")
	
	// Create network graph
	graph := NewGraph()
	
	// Add nodes
	graph.AddNode("A", "router", 0, 0)
	graph.AddNode("B", "router", 1, 0)
	graph.AddNode("C", "router", 2, 0)
	graph.AddNode("D", "router", 0, 1)
	graph.AddNode("E", "router", 1, 1)
	graph.AddNode("F", "router", 2, 1)
	
	// Add edges
	graph.AddEdge("A", "B", 4, 1000, 10)
	graph.AddEdge("A", "D", 2, 1000, 5)
	graph.AddEdge("B", "C", 3, 1000, 8)
	graph.AddEdge("B", "E", 5, 1000, 12)
	graph.AddEdge("C", "F", 2, 1000, 6)
	graph.AddEdge("D", "E", 1, 1000, 3)
	graph.AddEdge("E", "F", 3, 1000, 9)
	
	// Test Dijkstra's algorithm
	fmt.Println("Dijkstra's Algorithm:")
	path, cost := graph.DijkstraShortestPath("A", "F")
	fmt.Printf("  Path from A to F: %v (cost: %d)\n", path, cost)
	
	// Test Bellman-Ford algorithm
	fmt.Println("\nBellman-Ford Algorithm:")
	path, cost = graph.BellmanFordShortestPath("A", "F")
	fmt.Printf("  Path from A to F: %v (cost: %d)\n", path, cost)
	
	// Test A* algorithm
	fmt.Println("\nA* Algorithm:")
	path, cost = graph.AStarShortestPath("A", "F")
	fmt.Printf("  Path from A to F: %v (cost: %d)\n", path, cost)
	
	// Test Floyd-Warshall algorithm
	fmt.Println("\nFloyd-Warshall Algorithm:")
	distances := graph.FloydWarshallAllPairs()
	fmt.Printf("  Distance matrix:\n")
	for from := range distances {
		fmt.Printf("    %s: ", from)
		for to := range distances[from] {
			if distances[from][to] == math.MaxInt32 {
				fmt.Printf("∞ ")
			} else {
				fmt.Printf("%d ", distances[from][to])
			}
		}
		fmt.Println()
	}
}

// Demonstrate MST algorithms
func demonstrateMSTAlgorithms() {
	fmt.Println("=== MST Algorithms Demo ===\n")
	
	// Create network graph
	graph := NewGraph()
	
	// Add nodes
	graph.AddNode("A", "router", 0, 0)
	graph.AddNode("B", "router", 1, 0)
	graph.AddNode("C", "router", 2, 0)
	graph.AddNode("D", "router", 0, 1)
	graph.AddNode("E", "router", 1, 1)
	graph.AddNode("F", "router", 2, 1)
	
	// Add edges
	graph.AddEdge("A", "B", 4, 1000, 10)
	graph.AddEdge("A", "D", 2, 1000, 5)
	graph.AddEdge("B", "C", 3, 1000, 8)
	graph.AddEdge("B", "E", 5, 1000, 12)
	graph.AddEdge("C", "F", 2, 1000, 6)
	graph.AddEdge("D", "E", 1, 1000, 3)
	graph.AddEdge("E", "F", 3, 1000, 9)
	
	// Test Kruskal's algorithm
	fmt.Println("Kruskal's MST:")
	mst := graph.KruskalMST()
	totalWeight := 0
	for _, edge := range mst {
		fmt.Printf("  %s - %s (weight: %d)\n", edge.From.ID, edge.To.ID, edge.Weight)
		totalWeight += edge.Weight
	}
	fmt.Printf("  Total weight: %d\n", totalWeight)
	
	// Test Prim's algorithm
	fmt.Println("\nPrim's MST:")
	mst = graph.PrimMST("A")
	totalWeight = 0
	for _, edge := range mst {
		fmt.Printf("  %s - %s (weight: %d)\n", edge.From.ID, edge.To.ID, edge.Weight)
		totalWeight += edge.Weight
	}
	fmt.Printf("  Total weight: %d\n", totalWeight)
}

// Demonstrate network analysis
func demonstrateNetworkAnalysis() {
	fmt.Println("=== Network Analysis Demo ===\n")
	
	// Create network graph
	graph := NewGraph()
	
	// Add nodes
	graph.AddNode("A", "router", 0, 0)
	graph.AddNode("B", "router", 1, 0)
	graph.AddNode("C", "router", 2, 0)
	graph.AddNode("D", "router", 0, 1)
	graph.AddNode("E", "router", 1, 1)
	graph.AddNode("F", "router", 2, 1)
	
	// Add edges
	graph.AddEdge("A", "B", 4, 1000, 10)
	graph.AddEdge("A", "D", 2, 1000, 5)
	graph.AddEdge("B", "C", 3, 1000, 8)
	graph.AddEdge("B", "E", 5, 1000, 12)
	graph.AddEdge("C", "F", 2, 1000, 6)
	graph.AddEdge("D", "E", 1, 1000, 3)
	graph.AddEdge("E", "F", 3, 1000, 9)
	
	// Create network analyzer
	analyzer := NewNetworkAnalyzer(graph)
	
	// Analyze connectivity
	fmt.Println("Connectivity Analysis:")
	connectivity := analyzer.AnalyzeConnectivity()
	for nodeID, reachable := range connectivity {
		fmt.Printf("  %s: %t\n", nodeID, reachable)
	}
	
	// Analyze redundancy
	fmt.Println("\nRedundancy Analysis:")
	redundancy := analyzer.AnalyzeRedundancy()
	for nodeID, altPaths := range redundancy {
		fmt.Printf("  %s: %d alternative paths\n", nodeID, altPaths)
	}
	
	// Analyze performance
	fmt.Println("\nPerformance Analysis:")
	performance := analyzer.AnalyzePerformance()
	for metric, value := range performance {
		fmt.Printf("  %s: %.2f\n", metric, value)
	}
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 9: Routing Algorithms")
	fmt.Println("============================================================\n")
	
	// Run all demonstrations
	demonstrateRoutingAlgorithms()
	fmt.Println()
	demonstrateMSTAlgorithms()
	fmt.Println()
	demonstrateNetworkAnalysis()
	
	fmt.Println("\n🎯 Key Takeaways:")
	fmt.Println("1. Dijkstra's algorithm finds shortest paths from a single source")
	fmt.Println("2. Bellman-Ford algorithm handles negative weights and detects cycles")
	fmt.Println("3. A* algorithm uses heuristics for faster pathfinding")
	fmt.Println("4. Floyd-Warshall algorithm finds shortest paths between all pairs")
	fmt.Println("5. Kruskal's and Prim's algorithms find minimum spanning trees")
	fmt.Println("6. Network analysis helps optimize performance and reliability")
	fmt.Println("7. Go provides excellent support for graph algorithms")
	fmt.Println("8. Understanding these algorithms is crucial for network design")
	
	fmt.Println("\n📚 Next Topic: Network Performance & Optimization")
}
