# Topic 9: Routing & Switching - Network Path Management Mastery

## Routing & Switching Overview

Routing and switching are the fundamental mechanisms that enable data to move through networks efficiently. They operate at different layers of the OSI model and serve different purposes in network communication.

### Key Concepts
- **Routing**: Process of selecting paths for data packets across networks
- **Switching**: Process of forwarding data frames within a network segment
- **Forwarding**: The actual movement of data from input to output
- **Path Selection**: Choosing the best route based on various criteria
- **Network Topology**: Physical and logical arrangement of network elements

## Network Switching

### Switching Fundamentals
Switching operates at Layer 2 (Data Link Layer) and is responsible for forwarding frames within a local area network (LAN).

#### Switch Functions
- **Frame Forwarding**: Move frames from input port to output port
- **Address Learning**: Build MAC address table by observing traffic
- **Loop Prevention**: Prevent broadcast storms using STP
- **VLAN Support**: Virtual LAN segmentation and management
- **Quality of Service**: Traffic prioritization and bandwidth management

#### Switch Types
- **Unmanaged Switches**: Basic plug-and-play switches
- **Managed Switches**: Configurable switches with advanced features
- **Layer 2 Switches**: Operate at Data Link Layer
- **Layer 3 Switches**: Combine switching and routing capabilities
- **PoE Switches**: Power over Ethernet for connected devices

### MAC Address Learning
Switches learn MAC addresses by examining the source address of incoming frames and associating them with the port they arrived on.

#### Learning Process
1. **Frame Arrival**: Switch receives frame on port
2. **Source Learning**: Extract source MAC address
3. **Table Update**: Add MAC-port mapping to table
4. **Aging**: Remove old entries after timeout period
5. **Flooding**: Send unknown destination frames to all ports

#### MAC Address Table
- **MAC Address**: 48-bit hardware address
- **Port**: Physical port number
- **VLAN**: Virtual LAN identifier
- **Age**: Time since last seen
- **Type**: Static or dynamic entry

### VLAN (Virtual LAN)
VLANs logically segment a physical network into multiple broadcast domains.

#### VLAN Benefits
- **Security**: Isolate traffic between VLANs
- **Performance**: Reduce broadcast domain size
- **Management**: Logical grouping of devices
- **Flexibility**: Move devices without rewiring

#### VLAN Types
- **Port-based VLAN**: Assign ports to VLANs
- **MAC-based VLAN**: Assign devices by MAC address
- **Protocol-based VLAN**: Assign by protocol type
- **IP-based VLAN**: Assign by IP subnet

#### VLAN Trunking
- **802.1Q**: Standard for VLAN tagging
- **ISL**: Cisco's Inter-Switch Link protocol
- **Native VLAN**: Untagged traffic on trunk
- **VLAN Range**: 1-4094 (1 and 1002-1005 reserved)

### Spanning Tree Protocol (STP)
STP prevents loops in switched networks by creating a loop-free topology.

#### STP Concepts
- **Root Bridge**: Central reference point for STP
- **Root Port**: Port with lowest cost to root bridge
- **Designated Port**: Port with lowest cost to segment
- **Blocking Port**: Port blocked to prevent loops
- **BPDU**: Bridge Protocol Data Unit for STP communication

#### STP States
- **Blocking**: Port blocked, no forwarding
- **Listening**: Port listening for BPDUs
- **Learning**: Port learning MAC addresses
- **Forwarding**: Port forwarding traffic
- **Disabled**: Port administratively disabled

#### STP Variants
- **STP (802.1D)**: Original spanning tree
- **RSTP (802.1W)**: Rapid spanning tree
- **MSTP (802.1S)**: Multiple spanning tree
- **PVST+**: Per-VLAN spanning tree

## Network Routing

### Routing Fundamentals
Routing operates at Layer 3 (Network Layer) and is responsible for forwarding packets between different networks.

#### Router Functions
- **Packet Forwarding**: Move packets between networks
- **Path Selection**: Choose best route to destination
- **Route Management**: Maintain routing tables
- **Protocol Support**: Implement routing protocols
- **Traffic Control**: Manage bandwidth and QoS

#### Routing Types
- **Static Routing**: Manually configured routes
- **Dynamic Routing**: Automatically learned routes
- **Default Routing**: Catch-all route for unknown destinations
- **Host Routing**: Route to specific host
- **Network Routing**: Route to network subnet

### Routing Tables
Routing tables contain information about how to reach different destinations.

#### Table Entries
- **Destination**: Network or host address
- **Next Hop**: Next router in path
- **Interface**: Outgoing interface
- **Metric**: Cost to reach destination
- **Administrative Distance**: Trustworthiness of route source

#### Route Types
- **Directly Connected**: Local network interfaces
- **Static Routes**: Manually configured
- **Dynamic Routes**: Learned from routing protocols
- **Default Routes**: 0.0.0.0/0 destination

### Routing Algorithms

#### Distance Vector Algorithms
Distance vector algorithms use the Bellman-Ford algorithm to find shortest paths.

**RIP (Routing Information Protocol)**
- **Metric**: Hop count (maximum 15)
- **Update Interval**: 30 seconds
- **Convergence**: Slow (up to 3 minutes)
- **Variant**: RIPv1, RIPv2, RIPng

**IGRP (Interior Gateway Routing Protocol)**
- **Metric**: Composite (bandwidth, delay, reliability, load)
- **Update Interval**: 90 seconds
- **Convergence**: Moderate
- **Variant**: Cisco proprietary

**EIGRP (Enhanced IGRP)**
- **Metric**: Composite (bandwidth, delay)
- **Update**: Triggered updates
- **Convergence**: Fast
- **Variant**: Cisco proprietary

#### Link State Algorithms
Link state algorithms use Dijkstra's algorithm to find shortest paths.

**OSPF (Open Shortest Path First)**
- **Metric**: Cost (bandwidth-based)
- **Update**: Link state advertisements (LSA)
- **Convergence**: Fast
- **Variant**: OSPFv2 (IPv4), OSPFv3 (IPv6)

**IS-IS (Intermediate System to Intermediate System)**
- **Metric**: Cost (configurable)
- **Update**: Link state packets (LSP)
- **Convergence**: Fast
- **Variant**: ISO 10589 standard

#### Path Vector Algorithms
Path vector algorithms prevent routing loops by tracking path information.

**BGP (Border Gateway Protocol)**
- **Metric**: Path attributes
- **Update**: Incremental updates
- **Convergence**: Variable
- **Variant**: BGP-4 (IPv4), BGP-4+ (IPv6)

### Routing Protocols

#### Interior Gateway Protocols (IGP)
IGPs are used within autonomous systems (AS).

**RIP (Routing Information Protocol)**
- **Type**: Distance vector
- **Metric**: Hop count
- **Max Hops**: 15
- **Update**: Every 30 seconds
- **Use Case**: Small networks

**OSPF (Open Shortest Path First)**
- **Type**: Link state
- **Metric**: Cost
- **Update**: Triggered
- **Use Case**: Large networks

**EIGRP (Enhanced IGRP)**
- **Type**: Advanced distance vector
- **Metric**: Composite
- **Update**: Triggered
- **Use Case**: Cisco networks

#### Exterior Gateway Protocols (EGP)
EGPs are used between autonomous systems.

**BGP (Border Gateway Protocol)**
- **Type**: Path vector
- **Metric**: Path attributes
- **Update**: Incremental
- **Use Case**: Internet routing

### Network Topology

#### Physical Topology
Physical topology describes the physical arrangement of network devices.

**Bus Topology**
- **Structure**: Single cable connecting all devices
- **Advantages**: Simple, inexpensive
- **Disadvantages**: Single point of failure, limited scalability
- **Use Case**: Small networks

**Star Topology**
- **Structure**: Central hub/switch connecting all devices
- **Advantages**: Easy management, fault isolation
- **Disadvantages**: Central point of failure
- **Use Case**: Most common LAN topology

**Ring Topology**
- **Structure**: Devices connected in circular fashion
- **Advantages**: Equal access, no collisions
- **Disadvantages**: Single point of failure, difficult to modify
- **Use Case**: Token Ring networks

**Mesh Topology**
- **Structure**: Every device connected to every other device
- **Advantages**: High redundancy, fault tolerance
- **Disadvantages**: Expensive, complex
- **Use Case**: Critical networks

#### Logical Topology
Logical topology describes how data flows through the network.

**Flat Network**
- **Structure**: Single broadcast domain
- **Advantages**: Simple, easy to manage
- **Disadvantages**: Limited scalability, security issues
- **Use Case**: Small networks

**Hierarchical Network**
- **Structure**: Multiple layers (core, distribution, access)
- **Advantages**: Scalable, manageable
- **Disadvantages**: Complex, expensive
- **Use Case**: Large networks

**Collapsed Core**
- **Structure**: Core and distribution combined
- **Advantages**: Cost-effective, simpler
- **Disadvantages**: Limited scalability
- **Use Case**: Medium networks

### Quality of Service (QoS)

#### QoS Concepts
QoS ensures that critical traffic receives priority treatment in the network.

**Traffic Classification**
- **Voice**: Real-time voice traffic
- **Video**: Real-time video traffic
- **Data**: Best-effort data traffic
- **Management**: Network management traffic

**QoS Mechanisms**
- **Classification**: Identify traffic types
- **Marking**: Set priority levels
- **Policing**: Enforce traffic limits
- **Shaping**: Smooth traffic bursts
- **Queuing**: Manage traffic queues
- **Scheduling**: Prioritize traffic transmission

#### QoS Models
**Best Effort**
- **Treatment**: No guarantees
- **Use Case**: General internet traffic
- **Advantages**: Simple, no overhead
- **Disadvantages**: No service guarantees

**Integrated Services (IntServ)**
- **Treatment**: Per-flow guarantees
- **Use Case**: Real-time applications
- **Advantages**: Guaranteed service
- **Disadvantages**: Complex, not scalable

**Differentiated Services (DiffServ)**
- **Treatment**: Per-class guarantees
- **Use Case**: Enterprise networks
- **Advantages**: Scalable, practical
- **Disadvantages**: No per-flow guarantees

### Network Redundancy

#### Redundancy Concepts
Redundancy provides backup paths and failover capabilities.

**High Availability**
- **Uptime**: 99.9% or better
- **Downtime**: Less than 8.76 hours/year
- **Mechanisms**: Redundancy, failover, load balancing

**Fault Tolerance**
- **Single Points of Failure**: Eliminate SPOFs
- **Redundant Components**: Backup systems
- **Failover**: Automatic switching to backup
- **Load Balancing**: Distribute traffic across multiple paths

#### Redundancy Protocols
**HSRP (Hot Standby Router Protocol)**
- **Type**: First-hop redundancy
- **Variant**: Cisco proprietary
- **Use Case**: Default gateway redundancy

**VRRP (Virtual Router Redundancy Protocol)**
- **Type**: First-hop redundancy
- **Variant**: RFC 3768 standard
- **Use Case**: Multi-vendor environments

**GLBP (Gateway Load Balancing Protocol)**
- **Type**: Load balancing and redundancy
- **Variant**: Cisco proprietary
- **Use Case**: Load distribution

### Network Monitoring

#### Monitoring Concepts
Network monitoring ensures optimal performance and availability.

**SNMP (Simple Network Management Protocol)**
- **Version**: SNMPv1, SNMPv2c, SNMPv3
- **Components**: Manager, Agent, MIB
- **Use Case**: Network device monitoring

**RMON (Remote Monitoring)**
- **Type**: Remote network monitoring
- **Variant**: RMON1, RMON2
- **Use Case**: Traffic analysis

**NetFlow**
- **Type**: Traffic flow analysis
- **Variant**: NetFlow v5, v9, IPFIX
- **Use Case**: Traffic monitoring and analysis

#### Performance Metrics
**Bandwidth Utilization**
- **Measurement**: Percentage of available bandwidth used
- **Threshold**: 70-80% utilization
- **Action**: Upgrade or optimize

**Latency**
- **Measurement**: Round-trip time (RTT)
- **Threshold**: < 100ms for voice, < 150ms for video
- **Action**: Optimize routing, reduce hops

**Packet Loss**
- **Measurement**: Percentage of packets lost
- **Threshold**: < 0.1% for voice, < 0.5% for video
- **Action**: Check links, upgrade hardware

**Jitter**
- **Measurement**: Variation in packet arrival times
- **Threshold**: < 30ms for voice, < 50ms for video
- **Action**: Implement QoS, buffer management

### Go Implementation Concepts

#### Network Simulation
Go can be used to simulate network behavior and routing algorithms.

**Graph Representation**
- **Nodes**: Network devices (routers, switches)
- **Edges**: Network links
- **Weights**: Link costs or metrics
- **Algorithms**: Dijkstra, Bellman-Ford, Floyd-Warshall

**Routing Table Implementation**
- **Data Structure**: Map or slice for routes
- **Operations**: Add, remove, update, lookup
- **Concurrency**: Thread-safe operations
- **Persistence**: Save/load routing tables

**Protocol Simulation**
- **Message Passing**: Simulate protocol messages
- **State Machines**: Implement protocol states
- **Timers**: Handle protocol timers
- **Convergence**: Test algorithm convergence

#### Network Management
Go can be used to build network management tools.

**SNMP Client**
- **Operations**: Get, Set, Walk, Trap
- **Data Types**: Integer, String, OID
- **Error Handling**: SNMP error codes
- **Concurrency**: Multiple concurrent operations

**Network Discovery**
- **Ping**: Test connectivity
- **Traceroute**: Find network path
- **Port Scanning**: Discover services
- **Topology Mapping**: Build network map

**Performance Monitoring**
- **Bandwidth**: Measure link utilization
- **Latency**: Measure round-trip time
- **Packet Loss**: Measure dropped packets
- **Throughput**: Measure data transfer rate

## Interview Questions

### Basic Questions
1. What is the difference between routing and switching?
2. How does a switch learn MAC addresses?
3. What is the purpose of VLANs?

### Intermediate Questions
1. Explain the Spanning Tree Protocol and its states.
2. Compare distance vector and link state routing algorithms.
3. What are the advantages of hierarchical network design?

### Advanced Questions
1. Design a network topology for a large enterprise.
2. Implement a routing algorithm in Go.
3. How would you optimize network performance and reliability?

## Next Steps
After mastering routing and switching, proceed to:
- **Topic 10**: Network Performance & Optimization
- **Topic 11**: Distributed Systems Networking
- **Topic 12**: Cloud Networking & Containerization

Master routing and switching, and you'll understand how networks move data efficiently! 🚀
