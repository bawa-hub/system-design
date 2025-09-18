package main

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

// IPAddress represents an IPv4 address with utilities
type IPAddress struct {
	Address net.IP
	Mask    net.IPMask
}

// NewIPAddress creates a new IPAddress from string and CIDR notation
func NewIPAddress(address string) (*IPAddress, error) {
	ip, network, err := net.ParseCIDR(address)
	if err != nil {
		return nil, err
	}
	
	return &IPAddress{
		Address: ip,
		Mask:    network.Mask,
	}, nil
}

// NewIPAddressWithMask creates a new IPAddress with custom mask
func NewIPAddressWithMask(address string, mask net.IPMask) (*IPAddress, error) {
	ip := net.ParseIP(address)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address: %s", address)
	}
	
	return &IPAddress{
		Address: ip,
		Mask:    mask,
	}, nil
}

// GetNetworkAddress returns the network address
func (ip *IPAddress) GetNetworkAddress() net.IP {
	return ip.Address.Mask(ip.Mask)
}

// GetBroadcastAddress returns the broadcast address
func (ip *IPAddress) GetBroadcastAddress() net.IP {
	network := ip.GetNetworkAddress()
	ones, bits := ip.Mask.Size()
	
	// Calculate the number of host bits
	hostBits := bits - ones
	
	// Create broadcast address by setting all host bits to 1
	broadcast := make(net.IP, len(network))
	copy(broadcast, network)
	
	for i := 0; i < hostBits; i++ {
		byteIndex := (bits - 1 - i) / 8
		bitIndex := 7 - ((bits - 1 - i) % 8)
		broadcast[byteIndex] |= 1 << bitIndex
	}
	
	return broadcast
}

// GetFirstHostAddress returns the first usable host address
func (ip *IPAddress) GetFirstHostAddress() net.IP {
	network := ip.GetNetworkAddress()
	firstHost := make(net.IP, len(network))
	copy(firstHost, network)
	
	// Add 1 to the last octet
	firstHost[len(firstHost)-1]++
	return firstHost
}

// GetLastHostAddress returns the last usable host address
func (ip *IPAddress) GetLastHostAddress() net.IP {
	broadcast := ip.GetBroadcastAddress()
	lastHost := make(net.IP, len(broadcast))
	copy(lastHost, broadcast)
	
	// Subtract 1 from the last octet
	lastHost[len(lastHost)-1]--
	return lastHost
}

// GetHostCount returns the number of usable host addresses
func (ip *IPAddress) GetHostCount() int {
	ones, bits := ip.Mask.Size()
	hostBits := bits - ones
	return int(math.Pow(2, float64(hostBits))) - 2
}

// GetSubnetCount returns the number of subnets
func (ip *IPAddress) GetSubnetCount() int {
	ones, _ := ip.Mask.Size()
	return int(math.Pow(2, float64(32-ones)))
}

// IsInSameSubnet checks if another IP is in the same subnet
func (ip *IPAddress) IsInSameSubnet(otherIP net.IP) bool {
	network1 := ip.GetNetworkAddress()
	network2 := otherIP.Mask(ip.Mask)
	return network1.Equal(network2)
}

// String returns the string representation
func (ip *IPAddress) String() string {
	ones, _ := ip.Mask.Size()
	return fmt.Sprintf("%s/%d", ip.Address.String(), ones)
}

// SubnetCalculator provides subnetting utilities
type SubnetCalculator struct{}

// CalculateSubnets calculates subnets for a given network
func (sc *SubnetCalculator) CalculateSubnets(network string, requiredSubnets int) ([]*IPAddress, error) {
	baseIP, baseNetwork, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}
	
	// Calculate required bits for subnets
	subnetBits := int(math.Ceil(math.Log2(float64(requiredSubnets))))
	
	// Get current prefix length
	ones, _ := baseNetwork.Mask.Size()
	
	// Calculate new prefix length
	newPrefixLength := ones + subnetBits
	
	if newPrefixLength > 32 {
		return nil, fmt.Errorf("not enough address space for %d subnets", requiredSubnets)
	}
	
	// Create new mask
	newMask := net.CIDRMask(newPrefixLength, 32)
	
	var subnets []*IPAddress
	
	// Calculate subnet size
	subnetSize := int(math.Pow(2, float64(32-newPrefixLength)))
	
	// Generate subnets
	for i := 0; i < requiredSubnets; i++ {
		subnetAddr := make(net.IP, len(baseIP))
		copy(subnetAddr, baseIP)
		
		// Calculate subnet address
		subnetOffset := i * subnetSize
		for j := 0; j < 4; j++ {
			subnetAddr[3-j] += byte((subnetOffset >> (j * 8)) & 0xFF)
		}
		
		subnet := &IPAddress{
			Address: subnetAddr,
			Mask:    newMask,
		}
		
		subnets = append(subnets, subnet)
	}
	
	return subnets, nil
}

// VLSMSubnet represents a VLSM subnet requirement
type VLSMSubnet struct {
	Name        string
	HostsNeeded int
	Subnet      *IPAddress
}

// CalculateVLSM calculates VLSM subnets for different host requirements
func (sc *SubnetCalculator) CalculateVLSM(network string, requirements []VLSMSubnet) ([]VLSMSubnet, error) {
	// Sort requirements by host count (descending)
	for i := 0; i < len(requirements)-1; i++ {
		for j := i + 1; j < len(requirements); j++ {
			if requirements[i].HostsNeeded < requirements[j].HostsNeeded {
				requirements[i], requirements[j] = requirements[j], requirements[i]
			}
		}
	}
	
	baseIP, baseNetwork, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}
	
	currentIP := make(net.IP, len(baseIP))
	copy(currentIP, baseIP)
	
	var result []VLSMSubnet
	
	for _, req := range requirements {
		// Calculate required host bits
		hostBits := int(math.Ceil(math.Log2(float64(req.HostsNeeded + 2)))) // +2 for network and broadcast
		prefixLength := 32 - hostBits
		
		// Create mask
		mask := net.CIDRMask(prefixLength, 32)
		
		// Create subnet
		subnet := &IPAddress{
			Address: currentIP,
			Mask:    mask,
		}
		
		// Update result
		req.Subnet = subnet
		result = append(result, req)
		
		// Calculate next subnet address
		subnetSize := int(math.Pow(2, float64(hostBits)))
		for j := 0; j < 4; j++ {
			currentIP[3-j] += byte((subnetSize >> (j * 8)) & 0xFF)
		}
	}
	
	return result, nil
}

// IPAddressValidator provides IP address validation utilities
type IPAddressValidator struct{}

// IsValidIP checks if an IP address is valid
func (v *IPAddressValidator) IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsPrivateIP checks if an IP address is private
func (v *IPAddressValidator) IsPrivateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	
	// Check private ranges
	privateRanges := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}
	
	for _, rangeStr := range privateRanges {
		_, network, _ := net.ParseCIDR(rangeStr)
		if network.Contains(parsedIP) {
			return true
		}
	}
	
	return false
}

// IsPublicIP checks if an IP address is public
func (v *IPAddressValidator) IsPublicIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	
	// Check if it's a valid IP and not private
	return v.IsValidIP(ip) && !v.IsPrivateIP(ip) && !parsedIP.IsLoopback() && !parsedIP.IsMulticast()
}

// GetIPClass returns the class of an IP address
func (v *IPAddressValidator) GetIPClass(ip string) string {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return "Invalid"
	}
	
	firstOctet := int(parsedIP.To4()[0])
	
	switch {
	case firstOctet >= 1 && firstOctet <= 126:
		return "Class A"
	case firstOctet >= 128 && firstOctet <= 191:
		return "Class B"
	case firstOctet >= 192 && firstOctet <= 223:
		return "Class C"
	case firstOctet >= 224 && firstOctet <= 239:
		return "Class D (Multicast)"
	case firstOctet >= 240 && firstOctet <= 255:
		return "Class E (Reserved)"
	default:
		return "Unknown"
	}
}

// Demonstrate basic subnetting
func demonstrateBasicSubnetting() {
	fmt.Println("=== Basic Subnetting Demo ===\n")
	
	// Example: 192.168.1.0/24 needs 4 subnets
	network := "192.168.1.0/24"
	requiredSubnets := 4
	
	calculator := &SubnetCalculator{}
	subnets, err := calculator.CalculateSubnets(network, requiredSubnets)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Network: %s\n", network)
	fmt.Printf("Required subnets: %d\n\n", requiredSubnets)
	
	for i, subnet := range subnets {
		fmt.Printf("Subnet %d:\n", i+1)
		fmt.Printf("  Network: %s\n", subnet.GetNetworkAddress())
		fmt.Printf("  Broadcast: %s\n", subnet.GetBroadcastAddress())
		fmt.Printf("  First Host: %s\n", subnet.GetFirstHostAddress())
		fmt.Printf("  Last Host: %s\n", subnet.GetLastHostAddress())
		fmt.Printf("  Usable Hosts: %d\n", subnet.GetHostCount())
		fmt.Printf("  CIDR: %s\n\n", subnet.String())
	}
}

// Demonstrate VLSM
func demonstrateVLSM() {
	fmt.Println("=== VLSM Demo ===\n")
	
	network := "192.168.1.0/24"
	requirements := []VLSMSubnet{
		{Name: "Sales", HostsNeeded: 50},
		{Name: "Marketing", HostsNeeded: 25},
		{Name: "IT", HostsNeeded: 10},
		{Name: "HR", HostsNeeded: 5},
	}
	
	calculator := &SubnetCalculator{}
	subnets, err := calculator.CalculateVLSM(network, requirements)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Network: %s\n", network)
	fmt.Println("VLSM Subnets:\n")
	
	for _, subnet := range subnets {
		fmt.Printf("%s:\n", subnet.Name)
		fmt.Printf("  Hosts Needed: %d\n", subnet.HostsNeeded)
		fmt.Printf("  Network: %s\n", subnet.Subnet.GetNetworkAddress())
		fmt.Printf("  Broadcast: %s\n", subnet.Subnet.GetBroadcastAddress())
		fmt.Printf("  First Host: %s\n", subnet.Subnet.GetFirstHostAddress())
		fmt.Printf("  Last Host: %s\n", subnet.Subnet.GetLastHostAddress())
		fmt.Printf("  Usable Hosts: %d\n", subnet.Subnet.GetHostCount())
		fmt.Printf("  CIDR: %s\n\n", subnet.Subnet.String())
	}
}

// Demonstrate IP address validation
func demonstrateIPValidation() {
	fmt.Println("=== IP Address Validation Demo ===\n")
	
	validator := &IPAddressValidator{}
	testIPs := []string{
		"192.168.1.1",
		"10.0.0.1",
		"172.16.0.1",
		"8.8.8.8",
		"127.0.0.1",
		"256.1.1.1",
		"invalid",
		"2001:db8::1",
	}
	
	for _, ip := range testIPs {
		fmt.Printf("IP: %s\n", ip)
		fmt.Printf("  Valid: %t\n", validator.IsValidIP(ip))
		fmt.Printf("  Private: %t\n", validator.IsPrivateIP(ip))
		fmt.Printf("  Public: %t\n", validator.IsPublicIP(ip))
		fmt.Printf("  Class: %s\n\n", validator.GetIPClass(ip))
	}
}

// Demonstrate subnet membership
func demonstrateSubnetMembership() {
	fmt.Println("=== Subnet Membership Demo ===\n")
	
	// Create a subnet
	subnet, err := NewIPAddress("192.168.1.0/24")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	testIPs := []string{
		"192.168.1.1",
		"192.168.1.100",
		"192.168.1.254",
		"192.168.2.1",
		"10.0.0.1",
	}
	
	fmt.Printf("Subnet: %s\n", subnet.String())
	fmt.Printf("Network: %s\n", subnet.GetNetworkAddress())
	fmt.Printf("Broadcast: %s\n\n", subnet.GetBroadcastAddress())
	
	for _, ipStr := range testIPs {
		ip := net.ParseIP(ipStr)
		if ip != nil {
			inSubnet := subnet.IsInSameSubnet(ip)
			fmt.Printf("IP %s is %s the subnet\n", ipStr, map[bool]string{true: "in", false: "not in"}[inSubnet])
		}
	}
}

// Calculate subnet requirements
func calculateSubnetRequirements() {
	fmt.Println("=== Subnet Requirements Calculator ===\n")
	
	// Example: Company needs to subnet 10.0.0.0/8
	baseNetwork := "10.0.0.0/8"
	
	departments := []struct {
		Name        string
		HostsNeeded int
	}{
		{"Headquarters", 1000},
		{"Branch Office 1", 500},
		{"Branch Office 2", 250},
		{"Sales Team", 100},
		{"IT Department", 50},
		{"HR Department", 25},
		{"Guest Network", 10},
	}
	
	calculator := &SubnetCalculator{}
	
	// Convert to VLSM requirements
	var requirements []VLSMSubnet
	for _, dept := range departments {
		requirements = append(requirements, VLSMSubnet{
			Name:        dept.Name,
			HostsNeeded: dept.HostsNeeded,
		})
	}
	
	subnets, err := calculator.CalculateVLSM(baseNetwork, requirements)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Base Network: %s\n", baseNetwork)
	fmt.Println("Department Subnets:\n")
	
	for _, subnet := range subnets {
		fmt.Printf("%s:\n", subnet.Name)
		fmt.Printf("  Required Hosts: %d\n", subnet.HostsNeeded)
		fmt.Printf("  Allocated Hosts: %d\n", subnet.Subnet.GetHostCount())
		fmt.Printf("  Network: %s\n", subnet.Subnet.GetNetworkAddress())
		fmt.Printf("  CIDR: %s\n", subnet.Subnet.String())
		fmt.Printf("  Efficiency: %.1f%%\n", float64(subnet.HostsNeeded)/float64(subnet.Subnet.GetHostCount())*100)
		fmt.Println()
	}
}

func main() {
	fmt.Println("🚀 Computer Networking Mastery - Topic 2: IP Addressing & Subnetting")
	fmt.Println("==================================================================\n")
	
	// Run all demonstrations
	demonstrateBasicSubnetting()
	demonstrateVLSM()
	demonstrateIPValidation()
	demonstrateSubnetMembership()
	calculateSubnetRequirements()
	
	fmt.Println("🎯 Key Takeaways:")
	fmt.Println("1. Subnetting allows efficient IP address utilization")
	fmt.Println("2. VLSM provides maximum flexibility for different subnet sizes")
	fmt.Println("3. Understanding IP classes helps with network design")
	fmt.Println("4. Private IP ranges are used for internal networks")
	fmt.Println("5. Proper subnetting is crucial for network scalability")
	
	fmt.Println("\n📚 Next Topic: TCP Protocol Deep Dive")
}
