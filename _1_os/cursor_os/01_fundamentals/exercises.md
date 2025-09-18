# Computer Architecture Fundamentals - Exercises

## Exercise 1: CPU Feature Detection (Beginner)

### Task
Extend the `cpu_info.c` program to detect additional CPU features and provide more detailed information.

### Requirements
1. Add detection for AVX-512 features
2. Implement CPU frequency detection using RDTSC
3. Add cache size detection for all cache levels
4. Display CPU topology (cores, threads, sockets)

### Hints
- Use CPUID function 0x80000008 for extended features
- Implement TSC frequency calibration
- Use CPUID function 4 for cache parameters
- Research CPUID function 0xB for topology information

### Expected Output
```
CPU Vendor: GenuineIntel
CPU Brand: Intel(R) Core(TM) i7-10700K CPU @ 3.80GHz
CPU Frequency: 3800 MHz
Cores: 8, Threads: 16, Sockets: 1
L1 Data Cache: 32 KB
L1 Instruction Cache: 32 KB
L2 Cache: 256 KB
L3 Cache: 16 MB
```

## Exercise 2: Memory Access Optimizer (Intermediate)

### Task
Create a program that optimizes memory access patterns for different algorithms.

### Requirements
1. Implement a matrix multiplication algorithm with different access patterns:
   - Naive (row-major)
   - Optimized (cache-friendly)
   - Blocked (tiled)
2. Compare performance of each approach
3. Implement a memory allocator that aligns data to cache lines
4. Add prefetching hints where appropriate

### Code Template
```c
// matrix_multiply.c
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <string.h>

#define MATRIX_SIZE 1024
#define BLOCK_SIZE 64

void naive_multiply(double *A, double *B, double *C, int n) {
    // Implement naive matrix multiplication
}

void optimized_multiply(double *A, double *B, double *C, int n) {
    // Implement cache-friendly matrix multiplication
}

void blocked_multiply(double *A, double *B, double *C, int n) {
    // Implement blocked matrix multiplication
}

int main() {
    // Allocate aligned memory
    // Initialize matrices
    // Time each algorithm
    // Print results
    return 0;
}
```

## Exercise 3: Assembly Language Library (Intermediate)

### Task
Create a library of assembly language functions for common operations.

### Requirements
1. Implement the following functions in assembly:
   - `strlen` - string length
   - `strcpy` - string copy
   - `memcpy` - memory copy
   - `memset` - memory set
   - `strcmp` - string compare
2. Create C wrappers for each function
3. Benchmark against standard library functions
4. Use SIMD instructions where possible

### Code Template
```assembly
; string_ops.s
section .text
    global asm_strlen, asm_strcpy, asm_memcpy, asm_memset, asm_strcmp

asm_strlen:
    ; Implement strlen in assembly
    ret

asm_strcpy:
    ; Implement strcpy in assembly
    ret

; ... other functions
```

## Exercise 4: Cache Simulator (Advanced)

### Task
Build a software cache simulator to understand cache behavior.

### Requirements
1. Implement a configurable cache simulator supporting:
   - Different cache sizes
   - Different associativity levels
   - Different replacement policies (LRU, FIFO, Random)
2. Support different cache organizations:
   - Direct-mapped
   - Set-associative
   - Fully associative
3. Generate cache statistics:
   - Hit rate
   - Miss rate
   - Access patterns

### Code Template
```c
// cache_simulator.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int size;
    int associativity;
    int block_size;
    int num_sets;
    int replacement_policy;
} cache_config_t;

typedef struct {
    int valid;
    int tag;
    int last_used;
    char *data;
} cache_line_t;

typedef struct {
    cache_config_t config;
    cache_line_t *lines;
    int hits;
    int misses;
} cache_t;

cache_t* create_cache(cache_config_t *config);
void destroy_cache(cache_t *cache);
int access_cache(cache_t *cache, unsigned int address);
void print_cache_stats(cache_t *cache);
```

## Exercise 5: System Call Wrapper Library (Advanced)

### Task
Create a comprehensive library of system call wrappers with error handling.

### Requirements
1. Implement wrappers for common system calls:
   - File operations (open, read, write, close)
   - Process operations (fork, exec, wait)
   - Memory operations (mmap, munmap, mprotect)
   - Network operations (socket, bind, listen, accept)
2. Add proper error handling and logging
3. Create a test suite for each wrapper
4. Document the API

### Code Template
```c
// syscall_wrappers.h
#ifndef SYSCALL_WRAPPERS_H
#define SYSCALL_WRAPPERS_H

#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

// File operations
int safe_open(const char *pathname, int flags, mode_t mode);
ssize_t safe_read(int fd, void *buf, size_t count);
ssize_t safe_write(int fd, const void *buf, size_t count);
int safe_close(int fd);

// Process operations
pid_t safe_fork(void);
int safe_execve(const char *pathname, char *const argv[], char *const envp[]);
pid_t safe_wait(int *wstatus);

// Memory operations
void *safe_mmap(void *addr, size_t length, int prot, int flags, int fd, off_t offset);
int safe_munmap(void *addr, size_t length);
int safe_mprotect(void *addr, size_t length, int prot);

#endif
```

## Exercise 6: Hardware Timer Implementation (Expert)

### Task
Implement a high-resolution timer using hardware features.

### Requirements
1. Use RDTSC instruction for cycle-accurate timing
2. Implement timer calibration against wall clock time
3. Handle CPU frequency scaling and power management
4. Create a timer library with microsecond precision
5. Add support for multiple timers and callbacks

### Code Template
```c
// high_res_timer.c
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <unistd.h>

typedef struct {
    uint64_t start_cycles;
    uint64_t end_cycles;
    double frequency_mhz;
} timer_t;

void timer_init(timer_t *timer);
void timer_start(timer_t *timer);
void timer_stop(timer_t *timer);
double timer_elapsed_us(timer_t *timer);
double timer_elapsed_ms(timer_t *timer);
double timer_elapsed_s(timer_t *timer);
```

## Exercise 7: Memory Allocator (Expert)

### Task
Implement a custom memory allocator with advanced features.

### Requirements
1. Implement different allocation strategies:
   - First fit
   - Best fit
   - Worst fit
   - Buddy system
2. Add memory alignment support
3. Implement memory pooling
4. Add debugging and leak detection
5. Create performance benchmarks

### Code Template
```c
// custom_allocator.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

typedef enum {
    ALLOC_FIRST_FIT,
    ALLOC_BEST_FIT,
    ALLOC_WORST_FIT,
    ALLOC_BUDDY
} alloc_strategy_t;

typedef struct {
    void *start;
    size_t size;
    int free;
    struct block *next;
    struct block *prev;
} block_t;

typedef struct {
    void *memory;
    size_t total_size;
    block_t *free_list;
    alloc_strategy_t strategy;
    size_t allocated_bytes;
    size_t free_bytes;
} allocator_t;

allocator_t* create_allocator(size_t size, alloc_strategy_t strategy);
void destroy_allocator(allocator_t *allocator);
void* allocate(allocator_t *allocator, size_t size, size_t alignment);
void deallocate(allocator_t *allocator, void *ptr);
void print_allocator_stats(allocator_t *allocator);
```

## Exercise 8: Interrupt Handler (Expert)

### Task
Implement a simple interrupt handler for a simulated system.

### Requirements
1. Create an interrupt vector table
2. Implement interrupt handlers for different types of interrupts
3. Add interrupt masking and priority handling
4. Create a simple device simulator
5. Test interrupt handling with multiple devices

### Code Template
```c
// interrupt_handler.c
#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include <unistd.h>

typedef enum {
    IRQ_TIMER,
    IRQ_KEYBOARD,
    IRQ_MOUSE,
    IRQ_NETWORK,
    IRQ_DISK
} irq_type_t;

typedef struct {
    irq_type_t type;
    void (*handler)(void);
    int priority;
    int enabled;
} irq_handler_t;

typedef struct {
    irq_handler_t handlers[32];
    int interrupt_mask;
    int current_priority;
} interrupt_controller_t;

void interrupt_init(interrupt_controller_t *ic);
void register_handler(interrupt_controller_t *ic, irq_type_t type, 
                     void (*handler)(void), int priority);
void enable_interrupt(interrupt_controller_t *ic, irq_type_t type);
void disable_interrupt(interrupt_controller_t *ic, irq_type_t type);
void handle_interrupt(interrupt_controller_t *ic, irq_type_t type);
```

## Exercise 9: Performance Profiler (Expert)

### Task
Create a performance profiler that analyzes program execution.

### Requirements
1. Implement function call tracing
2. Add memory access profiling
3. Create cache miss analysis
4. Generate performance reports
5. Add visualization of performance data

### Code Template
```c
// performance_profiler.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

typedef struct {
    char *function_name;
    uint64_t call_count;
    uint64_t total_time;
    uint64_t min_time;
    uint64_t max_time;
} function_profile_t;

typedef struct {
    function_profile_t *functions;
    int num_functions;
    uint64_t start_time;
    uint64_t end_time;
} profiler_t;

void profiler_init(profiler_t *profiler);
void profiler_start(profiler_t *profiler);
void profiler_stop(profiler_t *profiler);
void profiler_enter_function(profiler_t *profiler, const char *function_name);
void profiler_exit_function(profiler_t *profiler, const char *function_name);
void profiler_print_report(profiler_t *profiler);
```

## Exercise 10: Operating System Kernel (Expert)

### Task
Implement a minimal operating system kernel with basic functionality.

### Requirements
1. Create a bootloader that loads the kernel
2. Implement basic memory management
3. Add process scheduling
4. Create a simple file system
5. Add system call interface

### Code Template
```c
// kernel.c
#include <stdint.h>
#include <stddef.h>

// Kernel data structures
typedef struct {
    uint32_t esp;
    uint32_t eip;
    uint32_t eax;
    uint32_t ebx;
    uint32_t ecx;
    uint32_t edx;
    uint32_t esi;
    uint32_t edi;
    uint32_t ebp;
    uint32_t eflags;
} cpu_state_t;

typedef struct {
    int pid;
    cpu_state_t state;
    int priority;
    int status;
    struct process *next;
} process_t;

// Kernel functions
void kernel_init(void);
void kernel_main(void);
void schedule(void);
void context_switch(process_t *from, process_t *to);
int create_process(void (*entry_point)(void), int priority);
void terminate_process(int pid);
```

## How to Approach These Exercises

1. **Start with the basics**: Complete exercises 1-3 to build a solid foundation
2. **Practice regularly**: Spend 1-2 hours daily on these exercises
3. **Read the source code**: Study the provided code templates carefully
4. **Experiment**: Try different approaches and see what works best
5. **Measure performance**: Always benchmark your implementations
6. **Document your work**: Keep notes on what you learn from each exercise

## Additional Resources

- **Intel 64 and IA-32 Architectures Software Developer's Manual**
- **AMD64 Architecture Programmer's Manual**
- **Computer Systems: A Programmer's Perspective** by Bryant and O'Hallaron
- **The Art of Assembly Language** by Randall Hyde
- **Linux Kernel Development** by Robert Love

## Success Criteria

By the end of these exercises, you should be able to:
- Write efficient assembly language programs
- Understand and optimize memory access patterns
- Implement system-level software components
- Debug low-level performance issues
- Design and implement custom data structures
- Understand the internals of operating systems

Remember: The goal is not just to complete the exercises, but to deeply understand the underlying concepts and be able to apply them in real-world scenarios!
