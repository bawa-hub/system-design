# Computer Architecture Fundamentals

## Table of Contents
1. [Introduction](#introduction)
2. [CPU Architecture](#cpu-architecture)
3. [Memory Hierarchy](#memory-hierarchy)
4. [I/O Subsystems](#io-subsystems)
5. [Assembly Language Programming](#assembly-language-programming)
6. [Hardware-Software Interface](#hardware-software-interface)
7. [Practical Implementation](#practical-implementation)

## Introduction

Understanding computer architecture is the foundation of systems programming. Every operating system is built upon the principles of how hardware components interact and communicate. As a systems engineer, you need to understand not just what the hardware does, but how it does it and how software can control and optimize it.

### Why This Matters
- **Performance**: Understanding hardware helps you write faster, more efficient code
- **Debugging**: When things go wrong, you need to understand the hardware layer
- **Innovation**: Many breakthroughs come from understanding hardware limitations and capabilities
- **System Design**: You can't design good software without understanding the underlying hardware

## CPU Architecture

### Basic CPU Components

#### 1. Control Unit (CU)
The control unit is the brain of the CPU that:
- Fetches instructions from memory
- Decodes instructions to understand what operation to perform
- Coordinates the execution of instructions
- Manages the flow of data between CPU components

#### 2. Arithmetic Logic Unit (ALU)
The ALU performs:
- Arithmetic operations (addition, subtraction, multiplication, division)
- Logical operations (AND, OR, NOT, XOR)
- Comparison operations (greater than, less than, equal to)
- Bit manipulation operations (shifts, rotates)

#### 3. Registers
Registers are small, fast storage locations within the CPU:

**General Purpose Registers:**
- Store data and intermediate results
- Examples: EAX, EBX, ECX, EDX in x86

**Special Purpose Registers:**
- **Program Counter (PC)**: Points to the next instruction to execute
- **Stack Pointer (SP)**: Points to the top of the stack
- **Status Register (Flags)**: Contains condition codes (zero, carry, overflow, etc.)
- **Instruction Register (IR)**: Holds the current instruction being executed

#### 4. Cache Memory
- **L1 Cache**: Fastest, smallest (typically 32KB-64KB)
- **L2 Cache**: Medium speed, medium size (typically 256KB-2MB)
- **L3 Cache**: Slower, larger (typically 8MB-32MB)

### Instruction Execution Cycle

The CPU follows a basic cycle for each instruction:

1. **Fetch**: Get the instruction from memory
2. **Decode**: Determine what the instruction does
3. **Execute**: Perform the operation
4. **Write Back**: Store the result

```
PC → Memory → IR → Decoder → ALU/CU → Registers/Memory
 ↑                                    ↓
 └─────────── Update PC ←─────────────┘
```

### Pipelining

Modern CPUs use pipelining to execute multiple instructions simultaneously:

```
Clock Cycle:  1    2    3    4    5    6    7
Instruction 1: F    D    E    W
Instruction 2:      F    D    E    W
Instruction 3:           F    D    E    W
Instruction 4:                F    D    E    W
```

**Benefits:**
- Higher throughput
- Better resource utilization
- Faster overall execution

**Challenges:**
- Pipeline hazards (data, control, structural)
- Branch prediction complexity
- Increased hardware complexity

## Memory Hierarchy

### Memory Types and Characteristics

#### 1. Registers
- **Speed**: Fastest (1 CPU cycle)
- **Size**: Smallest (typically 16-64 registers)
- **Cost**: Most expensive per byte
- **Volatility**: Volatile (data lost on power off)

#### 2. Cache Memory
- **L1**: 1-2 CPU cycles, 32KB-64KB
- **L2**: 10-20 CPU cycles, 256KB-2MB
- **L3**: 40-75 CPU cycles, 8MB-32MB

#### 3. Main Memory (RAM)
- **Speed**: 100-300 CPU cycles
- **Size**: 4GB-128GB typical
- **Cost**: Moderate
- **Volatility**: Volatile

#### 4. Secondary Storage
- **Speed**: 10,000+ CPU cycles
- **Size**: 500GB-10TB typical
- **Cost**: Cheapest per byte
- **Volatility**: Non-volatile

### Memory Access Patterns

#### Spatial Locality
Programs tend to access memory locations that are close to recently accessed locations.

**Example:**
```c
int array[1000];
for (int i = 0; i < 1000; i++) {
    array[i] = i;  // Sequential access pattern
}
```

#### Temporal Locality
Recently accessed memory locations are likely to be accessed again soon.

**Example:**
```c
int sum = 0;
for (int i = 0; i < 1000; i++) {
    sum += array[i];  // 'sum' accessed repeatedly
}
```

### Cache Organization

#### Direct-Mapped Cache
Each memory block maps to exactly one cache line.

```
Memory Address: [Tag][Index][Offset]
Cache Line:     [Tag][Data]
```

#### Set-Associative Cache
Each memory block can map to one of several cache lines in a set.

```
Memory Address: [Tag][Set][Offset]
Cache Set:      [Line1][Line2][Line3][Line4]
```

#### Fully Associative Cache
Any memory block can be placed in any cache line.

### Virtual Memory

Virtual memory provides:
- **Abstraction**: Programs see a large, uniform address space
- **Protection**: Memory access control and isolation
- **Efficiency**: Only active pages need to be in physical memory

#### Address Translation
```
Virtual Address → Page Table → Physical Address
[VPN][Offset] → [PTE] → [PPN][Offset]
```

## I/O Subsystems

### I/O Methods

#### 1. Programmed I/O (PIO)
CPU directly controls I/O operations:
- CPU polls device status
- CPU transfers data byte by byte
- Simple but inefficient

#### 2. Interrupt-Driven I/O
Device interrupts CPU when ready:
- CPU can do other work while waiting
- More efficient than polling
- Requires interrupt handling

#### 3. Direct Memory Access (DMA)
DMA controller handles data transfer:
- CPU sets up transfer parameters
- DMA controller moves data directly
- CPU only interrupted when transfer complete

### I/O Buses

#### System Bus
- Connects CPU to main memory
- High speed, low latency
- Examples: Front Side Bus (FSB), QuickPath Interconnect (QPI)

#### I/O Bus
- Connects CPU to I/O devices
- Lower speed, higher latency
- Examples: PCI, PCIe, USB, SATA

### Device Controllers
Hardware components that interface between CPU and devices:
- **Disk Controller**: Manages hard drive operations
- **Network Controller**: Handles network communication
- **Graphics Controller**: Manages display output

## Assembly Language Programming

### Why Learn Assembly?
- **Understanding**: See exactly what your high-level code does
- **Optimization**: Write critical sections in assembly
- **Debugging**: Understand low-level program behavior
- **Systems Programming**: Required for OS development

### x86-64 Assembly Basics

#### Registers
```assembly
; 64-bit general purpose registers
RAX, RBX, RCX, RDX, RSI, RDI, RBP, RSP
R8, R9, R10, R11, R12, R13, R14, R15

; 32-bit versions (lower 32 bits)
EAX, EBX, ECX, EDX, ESI, EDI, EBP, ESP

; 16-bit versions (lower 16 bits)
AX, BX, CX, DX, SI, DI, BP, SP

; 8-bit versions (lower 8 bits)
AL, BL, CL, DL, AH, BH, CH, DH
```

#### Basic Instructions

**Data Movement:**
```assembly
mov rax, 42        ; Load immediate value
mov rax, rbx       ; Copy register to register
mov rax, [rbx]     ; Load from memory
mov [rbx], rax     ; Store to memory
```

**Arithmetic:**
```assembly
add rax, rbx       ; rax = rax + rbx
sub rax, rbx       ; rax = rax - rbx
mul rbx            ; rax = rax * rbx (64-bit result in rdx:rax)
div rbx            ; rax = rdx:rax / rbx, remainder in rdx
```

**Control Flow:**
```assembly
cmp rax, rbx       ; Compare rax and rbx
je label           ; Jump if equal
jne label          ; Jump if not equal
jmp label          ; Unconditional jump
call function      ; Call function
ret                ; Return from function
```

### System Calls

System calls are the interface between user programs and the operating system kernel.

#### x86-64 System Call Convention
- **System call number**: RAX
- **Arguments**: RDI, RSI, RDX, R10, R8, R9
- **Return value**: RAX
- **Instruction**: `syscall`

#### Example: Write System Call
```assembly
section .data
    message db 'Hello, World!', 10
    length equ $ - message

section .text
    global _start

_start:
    mov rax, 1          ; sys_write
    mov rdi, 1          ; stdout
    mov rsi, message    ; buffer
    mov rdx, length     ; count
    syscall

    mov rax, 60         ; sys_exit
    mov rdi, 0          ; exit status
    syscall
```

## Hardware-Software Interface

### Interrupts

Interrupts allow hardware to signal the CPU when attention is needed.

#### Types of Interrupts
1. **Hardware Interrupts**: From external devices
2. **Software Interrupts**: From programs (system calls)
3. **Exceptions**: From CPU itself (page faults, division by zero)

#### Interrupt Handling Process
1. CPU saves current state (registers, flags)
2. CPU jumps to interrupt handler
3. Handler processes the interrupt
4. CPU restores saved state
5. CPU continues execution

### Memory-Mapped I/O

Instead of special I/O instructions, devices are accessed through memory addresses.

```c
// Example: Accessing a UART (serial port)
#define UART_BASE 0x3F201000
#define UART_DATA_REG (UART_BASE + 0x00)
#define UART_STATUS_REG (UART_BASE + 0x18)

// Write a character
void uart_putc(char c) {
    // Wait until transmitter is ready
    while (*(volatile uint32_t*)UART_STATUS_REG & 0x20);
    
    // Write the character
    *(volatile uint32_t*)UART_DATA_REG = c;
}
```

### Port-Mapped I/O

Uses special I/O instructions to communicate with devices.

```assembly
; x86 I/O instructions
in al, 0x60        ; Read from port 0x60
out 0x60, al       ; Write to port 0x60
```

## Practical Implementation

### Project 1: CPU Information Detector

Let's build a program that detects CPU features and capabilities:

```c
// cpu_info.c
#include <stdio.h>
#include <stdint.h>

// CPUID instruction wrapper
static inline void cpuid(uint32_t eax, uint32_t ecx, 
                        uint32_t *eax_out, uint32_t *ebx_out,
                        uint32_t *ecx_out, uint32_t *edx_out) {
    asm volatile("cpuid"
                 : "=a"(*eax_out), "=b"(*ebx_out), 
                   "=c"(*ecx_out), "=d"(*edx_out)
                 : "a"(eax), "c"(ecx));
}

void detect_cpu_features() {
    uint32_t eax, ebx, ecx, edx;
    
    // Get CPU vendor string
    cpuid(0, 0, &eax, &ebx, &ecx, &edx);
    
    char vendor[13] = {0};
    *((uint32_t*)vendor) = ebx;
    *((uint32_t*)(vendor + 4)) = edx;
    *((uint32_t*)(vendor + 8)) = ecx;
    
    printf("CPU Vendor: %s\n", vendor);
    
    // Get CPU features
    cpuid(1, 0, &eax, &ebx, &ecx, &edx);
    
    printf("CPU Features:\n");
    if (edx & (1 << 23)) printf("  - MMX\n");
    if (edx & (1 << 25)) printf("  - SSE\n");
    if (edx & (1 << 26)) printf("  - SSE2\n");
    if (ecx & (1 << 0))  printf("  - SSE3\n");
    if (ecx & (1 << 19)) printf("  - SSE4.1\n");
    if (ecx & (1 << 20)) printf("  - SSE4.2\n");
    if (ecx & (1 << 28)) printf("  - AVX\n");
}

int main() {
    detect_cpu_features();
    return 0;
}
```

### Project 2: Memory Access Pattern Analyzer

```c
// memory_patterns.c
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <sys/time.h>

#define ARRAY_SIZE 1024 * 1024  // 1MB array
#define ITERATIONS 1000

double get_time() {
    struct timeval tv;
    gettimeofday(&tv, NULL);
    return tv.tv_sec + tv.tv_usec / 1000000.0;
}

void sequential_access(int *array, int size) {
    for (int i = 0; i < size; i++) {
        array[i] = i;
    }
}

void random_access(int *array, int size) {
    for (int i = 0; i < size; i++) {
        int index = rand() % size;
        array[index] = i;
    }
}

void measure_access_patterns() {
    int *array = malloc(ARRAY_SIZE * sizeof(int));
    if (!array) {
        printf("Memory allocation failed\n");
        return;
    }
    
    // Sequential access
    double start = get_time();
    for (int i = 0; i < ITERATIONS; i++) {
        sequential_access(array, ARRAY_SIZE);
    }
    double sequential_time = get_time() - start;
    
    // Random access
    start = get_time();
    for (int i = 0; i < ITERATIONS; i++) {
        random_access(array, ARRAY_SIZE);
    }
    double random_time = get_time() - start;
    
    printf("Sequential access time: %.6f seconds\n", sequential_time);
    printf("Random access time: %.6f seconds\n", random_time);
    printf("Speedup: %.2fx\n", random_time / sequential_time);
    
    free(array);
}

int main() {
    srand(time(NULL));
    measure_access_patterns();
    return 0;
}
```

### Project 3: Assembly Language Calculator

```assembly
; calculator.s
section .data
    prompt db 'Enter two numbers: ', 0
    result db 'Result: ', 0
    newline db 10, 0

section .bss
    num1 resb 4
    num2 resb 4
    result_val resb 4

section .text
    global _start

_start:
    ; Print prompt
    mov rax, 1          ; sys_write
    mov rdi, 1          ; stdout
    mov rsi, prompt
    mov rdx, 19
    syscall

    ; Read first number
    mov rax, 0          ; sys_read
    mov rdi, 0          ; stdin
    mov rsi, num1
    mov rdx, 4
    syscall

    ; Read second number
    mov rax, 0          ; sys_read
    mov rdi, 0          ; stdin
    mov rsi, num2
    mov rdx, 4
    syscall

    ; Convert ASCII to integer and add
    mov eax, [num1]
    sub eax, '0'        ; Convert ASCII to integer
    mov ebx, [num2]
    sub ebx, '0'        ; Convert ASCII to integer
    add eax, ebx        ; Add the numbers
    add eax, '0'        ; Convert back to ASCII
    mov [result_val], eax

    ; Print result
    mov rax, 1          ; sys_write
    mov rdi, 1          ; stdout
    mov rsi, result
    mov rdx, 8
    syscall

    mov rax, 1          ; sys_write
    mov rdi, 1          ; stdout
    mov rsi, result_val
    mov rdx, 1
    syscall

    mov rax, 1          ; sys_write
    mov rdi, 1          ; stdout
    mov rsi, newline
    mov rdx, 1
    syscall

    ; Exit
    mov rax, 60         ; sys_exit
    mov rdi, 0
    syscall
```

## Key Takeaways

1. **Hardware Understanding**: Every software optimization starts with understanding the hardware
2. **Memory Hierarchy**: The speed difference between different memory types is enormous
3. **Assembly Language**: Essential for understanding what high-level code actually does
4. **System Calls**: The bridge between user programs and the operating system
5. **Interrupts**: How hardware communicates with software

## Next Steps

In the next module, we'll dive into **System Programming Basics** where you'll learn:
- C programming for systems
- Memory management techniques
- File I/O and system calls
- Building your first system utilities

## Exercises

1. **CPU Feature Detection**: Extend the CPU info program to detect more features
2. **Memory Benchmarking**: Create a comprehensive memory access benchmark
3. **Assembly Programming**: Write assembly programs for common algorithms
4. **System Call Wrapper**: Create a library of system call wrappers in C
5. **Hardware Simulation**: Simulate a simple CPU in software

Remember: The goal is not just to understand these concepts, but to implement them and see them in action. Every line of code you write should deepen your understanding of how computers really work!
