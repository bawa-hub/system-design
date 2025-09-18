#include <stdio.h>
#include <stdint.h>
#include <string.h>

// CPUID instruction wrapper
static inline void cpuid(uint32_t eax, uint32_t ecx, 
                        uint32_t *eax_out, uint32_t *ebx_out,
                        uint32_t *ecx_out, uint32_t *edx_out) {
    asm volatile("cpuid"
                 : "=a"(*eax_out), "=b"(*ebx_out), 
                   "=c"(*ecx_out), "=d"(*edx_out)
                 : "a"(eax), "c"(ecx));
}

// Get CPU vendor string
void get_cpu_vendor(char *vendor_string) {
    uint32_t eax, ebx, ecx, edx;
    cpuid(0, 0, &eax, &ebx, &ecx, &edx);
    
    *((uint32_t*)vendor_string) = ebx;
    *((uint32_t*)(vendor_string + 4)) = edx;
    *((uint32_t*)(vendor_string + 8)) = ecx;
    vendor_string[12] = '\0';
}

// Get CPU brand string
void get_cpu_brand(char *brand_string) {
    uint32_t eax, ebx, ecx, edx;
    char brand[48] = {0};
    
    // Get brand string using CPUID function 0x80000002-0x80000004
    for (int i = 0; i < 3; i++) {
        cpuid(0x80000002 + i, 0, &eax, &ebx, &ecx, &edx);
        *((uint32_t*)(brand + i * 16)) = eax;
        *((uint32_t*)(brand + i * 16 + 4)) = ebx;
        *((uint32_t*)(brand + i * 16 + 8)) = ecx;
        *((uint32_t*)(brand + i * 16 + 12)) = edx;
    }
    
    strcpy(brand_string, brand);
}

// Detect CPU features
void detect_cpu_features() {
    uint32_t eax, ebx, ecx, edx;
    
    printf("=== CPU Features ===\n");
    
    // Get basic features (CPUID function 1)
    cpuid(1, 0, &eax, &ebx, &ecx, &edx);
    
    printf("Basic Features (EDX):\n");
    if (edx & (1 << 0))  printf("  - FPU (Floating Point Unit)\n");
    if (edx & (1 << 3))  printf("  - PSE (Page Size Extension)\n");
    if (edx & (1 << 4))  printf("  - TSC (Time Stamp Counter)\n");
    if (edx & (1 << 5))  printf("  - MSR (Model Specific Registers)\n");
    if (edx & (1 << 8))  printf("  - CMPXCHG8B (Compare and Exchange 8 bytes)\n");
    if (edx & (1 << 15)) printf("  - CMOV (Conditional Move)\n");
    if (edx & (1 << 23)) printf("  - MMX (MultiMedia eXtensions)\n");
    if (edx & (1 << 24)) printf("  - FXSR (Fast Save/Restore)\n");
    if (edx & (1 << 25)) printf("  - SSE (Streaming SIMD Extensions)\n");
    if (edx & (1 << 26)) printf("  - SSE2 (Streaming SIMD Extensions 2)\n");
    
    printf("\nExtended Features (ECX):\n");
    if (ecx & (1 << 0))  printf("  - SSE3 (Streaming SIMD Extensions 3)\n");
    if (ecx & (1 << 9))  printf("  - SSSE3 (Supplemental SSE3)\n");
    if (ecx & (1 << 19)) printf("  - SSE4.1 (Streaming SIMD Extensions 4.1)\n");
    if (ecx & (1 << 20)) printf("  - SSE4.2 (Streaming SIMD Extensions 4.2)\n");
    if (ecx & (1 << 23)) printf("  - POPCNT (Population Count)\n");
    if (ecx & (1 << 28)) printf("  - AVX (Advanced Vector Extensions)\n");
    if (ecx & (1 << 29)) printf("  - F16C (Half-precision FMA)\n");
    if (ecx & (1 << 30)) printf("  - RDRAND (Random Number Generator)\n");
    
    // Get extended features (CPUID function 7)
    cpuid(7, 0, &eax, &ebx, &ecx, &edx);
    
    printf("\nAdvanced Features (EBX):\n");
    if (ebx & (1 << 0))  printf("  - FSGSBASE (FS/GS Base)\n");
    if (ebx & (1 << 3))  printf("  - BMI1 (Bit Manipulation Instructions 1)\n");
    if (ebx & (1 << 5))  printf("  - AVX2 (Advanced Vector Extensions 2)\n");
    if (ebx & (1 << 8))  printf("  - BMI2 (Bit Manipulation Instructions 2)\n");
    if (ebx & (1 << 16)) printf("  - AVX512F (AVX-512 Foundation)\n");
    if (ebx & (1 << 17)) printf("  - AVX512DQ (AVX-512 Doubleword and Quadword)\n");
    if (ebx & (1 << 18)) printf("  - RDSEED (Random Seed)\n");
    if (ebx & (1 << 19)) printf("  - ADX (Multi-Precision Add-Carry)\n");
    if (ebx & (1 << 20)) printf("  - SMAP (Supervisor Mode Access Prevention)\n");
    if (ebx & (1 << 21)) printf("  - AVX512IFMA (AVX-512 Integer Fused Multiply-Add)\n");
    if (ebx & (1 << 22)) printf("  - PCOMMIT (Persistent Commit)\n");
    if (ebx & (1 << 23)) printf("  - CLFLUSHOPT (Cache Line Flush Optimized)\n");
    if (ebx & (1 << 24)) printf("  - CLWB (Cache Line Write Back)\n");
    if (ebx & (1 << 25)) printf("  - INTEL_PT (Intel Processor Trace)\n");
    if (ebx & (1 << 26)) printf("  - AVX512PF (AVX-512 Prefetch)\n");
    if (ebx & (1 << 27)) printf("  - AVX512ER (AVX-512 Exponential and Reciprocal)\n");
    if (ebx & (1 << 28)) printf("  - AVX512CD (AVX-512 Conflict Detection)\n");
    if (ebx & (1 << 29)) printf("  - SHA (SHA Extensions)\n");
    if (ebx & (1 << 30)) printf("  - AVX512BW (AVX-512 Byte and Word)\n");
    if (ebx & (1 << 31)) printf("  - AVX512VL (AVX-512 Vector Length)\n");
}

// Get cache information
void get_cache_info() {
    uint32_t eax, ebx, ecx, edx;
    
    printf("\n=== Cache Information ===\n");
    
    // Get cache descriptors
    cpuid(2, 0, &eax, &ebx, &ecx, &edx);
    
    printf("Cache Descriptors (EAX): 0x%08x\n", eax);
    printf("Cache Descriptors (EBX): 0x%08x\n", ebx);
    printf("Cache Descriptors (ECX): 0x%08x\n", ecx);
    printf("Cache Descriptors (EDX): 0x%08x\n", edx);
    
    // Get deterministic cache parameters
    cpuid(4, 0, &eax, &ebx, &ecx, &edx);
    
    printf("\nL1 Data Cache:\n");
    printf("  Line Size: %d bytes\n", ((ebx & 0xFFF) + 1));
    printf("  Ways: %d\n", ((ebx >> 22) & 0x3FF) + 1);
    printf("  Sets: %d\n", ((ecx & 0xFFF) + 1));
    printf("  Size: %d KB\n", ((((ebx >> 22) & 0x3FF) + 1) * 
                               ((ecx & 0xFFF) + 1) * 
                               ((ebx & 0xFFF) + 1)) / 1024);
}

// Get CPU topology information
void get_cpu_topology() {
    uint32_t eax, ebx, ecx, edx;
    
    printf("\n=== CPU Topology ===\n");
    
    // Get number of logical processors
    cpuid(1, 0, &eax, &ebx, &ecx, &edx);
    int logical_processors = ((ebx >> 16) & 0xFF);
    printf("Logical Processors: %d\n", logical_processors);
    
    // Get APIC ID
    int apic_id = (ebx >> 24) & 0xFF;
    printf("APIC ID: %d\n", apic_id);
    
    // Get processor signature
    printf("Processor Signature: 0x%08x\n", eax);
    int family = ((eax >> 8) & 0xF) + ((eax >> 20) & 0xFF);
    int model = ((eax >> 4) & 0xF) + (((eax >> 16) & 0xF) << 4);
    int stepping = eax & 0xF;
    
    printf("Family: %d, Model: %d, Stepping: %d\n", family, model, stepping);
}

// Get TSC frequency
void get_tsc_frequency() {
    uint32_t eax, ebx, ecx, edx;
    
    printf("\n=== TSC Information ===\n");
    
    // Check if TSC is available
    cpuid(1, 0, &eax, &ebx, &ecx, &edx);
    if (!(edx & (1 << 4))) {
        printf("TSC not available\n");
        return;
    }
    
    printf("TSC is available\n");
    
    // Get TSC frequency (this is a simplified approach)
    // In practice, you'd need to calibrate the TSC against a known timer
    printf("TSC frequency calibration would require timing measurements\n");
}

int main() {
    char vendor[13];
    char brand[49];
    
    printf("=== CPU Information Detector ===\n\n");
    
    // Get CPU vendor
    get_cpu_vendor(vendor);
    printf("CPU Vendor: %s\n", vendor);
    
    // Get CPU brand
    get_cpu_brand(brand);
    printf("CPU Brand: %s\n", brand);
    
    // Detect features
    detect_cpu_features();
    
    // Get cache info
    get_cache_info();
    
    // Get topology
    get_cpu_topology();
    
    // Get TSC info
    get_tsc_frequency();
    
    return 0;
}
