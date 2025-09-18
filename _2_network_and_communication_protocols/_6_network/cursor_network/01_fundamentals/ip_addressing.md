# IP Addressing & Subnetting Mastery

## IPv4 Addressing

### Understanding IP Addresses
An IPv4 address is a 32-bit number divided into four 8-bit octets, represented in dotted decimal notation (e.g., 192.168.1.1).

### Address Classes (Historical)
- **Class A**: 1.0.0.0 to 126.255.255.255 (8-bit network, 24-bit host)
- **Class B**: 128.0.0.0 to 191.255.255.255 (16-bit network, 16-bit host)
- **Class C**: 192.0.0.0 to 223.255.255.255 (24-bit network, 8-bit host)
- **Class D**: 224.0.0.0 to 239.255.255.255 (Multicast)
- **Class E**: 240.0.0.0 to 255.255.255.255 (Reserved)

### Private Address Ranges (RFC 1918)
- **10.0.0.0/8**: 10.0.0.0 to 10.255.255.255
- **172.16.0.0/12**: 172.16.0.0 to 172.31.255.255
- **192.168.0.0/16**: 192.168.0.0 to 192.168.255.255

### Special Addresses
- **0.0.0.0/0**: Default route
- **127.0.0.0/8**: Loopback addresses
- **169.254.0.0/16**: Link-local addresses (APIPA)
- **255.255.255.255**: Limited broadcast

## CIDR (Classless Inter-Domain Routing)

### CIDR Notation
Format: `IP_ADDRESS/PREFIX_LENGTH`
- Example: `192.168.1.0/24`
- Prefix length indicates number of network bits
- Remaining bits are host bits

### Subnet Masks
- **/8**: 255.0.0.0
- **/16**: 255.255.0.0
- **/24**: 255.255.255.0
- **/32**: 255.255.255.255 (host route)

## Subnetting Fundamentals

### Why Subnet?
1. **Efficient IP usage**: Avoid wasting addresses
2. **Security**: Isolate network segments
3. **Performance**: Reduce broadcast domains
4. **Management**: Logical organization

### Subnetting Process
1. Determine required subnets
2. Calculate subnet mask
3. Identify network addresses
4. Identify broadcast addresses
5. Identify usable host ranges

### Subnetting Formulas
- **Number of subnets**: 2^(subnet_bits)
- **Number of hosts per subnet**: 2^(host_bits) - 2
- **Subnet mask**: 2^32 - 2^(32-prefix_length)

## VLSM (Variable Length Subnet Masking)

VLSM allows different subnets to have different subnet masks within the same network.

### VLSM Benefits
- Maximum IP address utilization
- Hierarchical addressing
- Efficient routing

## IPv6 Addressing

### IPv6 Basics
- 128-bit addresses (vs 32-bit IPv4)
- Hexadecimal notation with colons
- Example: `2001:0db8:85a3:0000:0000:8a2e:0370:7334`
- Shortened: `2001:db8:85a3::8a2e:370:7334`

### IPv6 Address Types
- **Unicast**: Single interface
- **Multicast**: Multiple interfaces
- **Anycast**: Nearest interface

### IPv6 Special Addresses
- **::/128**: Unspecified address
- **::1/128**: Loopback address
- **2000::/3**: Global unicast
- **fc00::/7**: Unique local
- **fe80::/10**: Link-local

## NAT (Network Address Translation)

### NAT Types
1. **Static NAT**: One-to-one mapping
2. **Dynamic NAT**: Pool-based mapping
3. **PAT (Port Address Translation)**: Many-to-one mapping

### NAT Benefits
- IP address conservation
- Security through obscurity
- Simplified addressing

## Practical Examples

### Example 1: Basic Subnetting
**Given**: 192.168.1.0/24, need 4 subnets

**Solution**:
- Need 2 bits for subnets (2^2 = 4)
- New prefix: /26
- Subnet mask: 255.255.255.192
- Subnets:
  - 192.168.1.0/26 (0-63)
  - 192.168.1.64/26 (64-127)
  - 192.168.1.128/26 (128-191)
  - 192.168.1.192/26 (192-255)

### Example 2: VLSM
**Given**: 192.168.1.0/24, need:
- 2 subnets with 50 hosts each
- 4 subnets with 10 hosts each

**Solution**:
1. For 50 hosts: need 6 host bits (2^6-2 = 62 hosts)
   - Subnets: /26 (255.255.255.192)
   - 192.168.1.0/26 and 192.168.1.64/26

2. For 10 hosts: need 4 host bits (2^4-2 = 14 hosts)
   - Subnets: /28 (255.255.255.240)
   - 192.168.1.128/28, 192.168.1.144/28, 192.168.1.160/28, 192.168.1.176/28

## Interview Questions

### Basic Questions
1. What's the difference between a network address and a broadcast address?
2. How do you calculate the number of usable hosts in a subnet?
3. What is the purpose of a subnet mask?

### Intermediate Questions
1. Explain the difference between classful and classless addressing.
2. How does VLSM improve IP address utilization?
3. What are the advantages and disadvantages of NAT?

### Advanced Questions
1. Design a network addressing scheme for a company with multiple departments.
2. Explain how IPv6 addresses are allocated and managed.
3. How would you troubleshoot IP addressing issues in a large network?

## Next Steps
Now that you understand IP addressing, we'll move to:
1. TCP and UDP protocols
2. Routing and switching
3. Network programming in Go
4. Advanced networking concepts

Master these fundamentals, and you'll be unstoppable in networking interviews! 🚀
