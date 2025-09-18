#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <sys/time.h>
#include <string.h>
#include <unistd.h>

#define ARRAY_SIZE 1024 * 1024  // 1MB array
#define ITERATIONS 1000
#define CACHE_LINE_SIZE 64

double get_time() {
    struct timeval tv;
    gettimeofday(&tv, NULL);
    return tv.tv_sec + tv.tv_usec / 1000000.0;
}

// Sequential access pattern - good for cache
void sequential_access(int *array, int size) {
    for (int i = 0; i < size; i++) {
        array[i] = i;
    }
}

// Random access pattern - bad for cache
void random_access(int *array, int size) {
    for (int i = 0; i < size; i++) {
        int index = rand() % size;
        array[index] = i;
    }
}

// Strided access pattern - moderate cache performance
void strided_access(int *array, int size, int stride) {
    for (int i = 0; i < size; i += stride) {
        array[i] = i;
    }
}

// Cache line aware access - optimal for cache
void cache_line_access(int *array, int size) {
    for (int i = 0; i < size; i += CACHE_LINE_SIZE / sizeof(int)) {
        array[i] = i;
    }
}

// Measure memory bandwidth
double measure_bandwidth(int *array, int size, void (*access_func)(int*, int)) {
    double start = get_time();
    access_func(array, size);
    double end = get_time();
    
    double time_taken = end - start;
    double bytes_accessed = size * sizeof(int);
    double bandwidth = bytes_accessed / time_taken / (1024 * 1024); // MB/s
    
    return bandwidth;
}

// Measure cache performance with different access patterns
void measure_cache_performance() {
    int *array = malloc(ARRAY_SIZE * sizeof(int));
    if (!array) {
        printf("Memory allocation failed\n");
        return;
    }
    
    printf("=== Cache Performance Analysis ===\n");
    printf("Array size: %d elements (%d MB)\n", ARRAY_SIZE, ARRAY_SIZE * sizeof(int) / (1024 * 1024));
    printf("Cache line size: %d bytes\n", CACHE_LINE_SIZE);
    printf("\n");
    
    // Sequential access
    double sequential_time = 0;
    for (int i = 0; i < ITERATIONS; i++) {
        double start = get_time();
        sequential_access(array, ARRAY_SIZE);
        sequential_time += get_time() - start;
    }
    sequential_time /= ITERATIONS;
    
    // Random access
    double random_time = 0;
    for (int i = 0; i < ITERATIONS; i++) {
        double start = get_time();
        random_access(array, ARRAY_SIZE);
        random_time += get_time() - start;
    }
    random_time /= ITERATIONS;
    
    // Strided access (stride = 2)
    double strided_time = 0;
    for (int i = 0; i < ITERATIONS; i++) {
        double start = get_time();
        strided_access(array, ARRAY_SIZE, 2);
        strided_time += get_time() - start;
    }
    strided_time /= ITERATIONS;
    
    // Cache line aware access
    double cache_line_time = 0;
    for (int i = 0; i < ITERATIONS; i++) {
        double start = get_time();
        cache_line_access(array, ARRAY_SIZE);
        cache_line_time += get_time() - start;
    }
    cache_line_time /= ITERATIONS;
    
    printf("Access Pattern Performance:\n");
    printf("  Sequential:     %.6f seconds (%.2fx baseline)\n", sequential_time, 1.0);
    printf("  Random:         %.6f seconds (%.2fx slower)\n", random_time, random_time / sequential_time);
    printf("  Strided (2):    %.6f seconds (%.2fx slower)\n", strided_time, strided_time / sequential_time);
    printf("  Cache-line:     %.6f seconds (%.2fx faster)\n", cache_line_time, sequential_time / cache_line_time);
    
    free(array);
}

// Measure memory bandwidth
void measure_memory_bandwidth() {
    int *array = malloc(ARRAY_SIZE * sizeof(int));
    if (!array) {
        printf("Memory allocation failed\n");
        return;
    }
    
    printf("\n=== Memory Bandwidth Analysis ===\n");
    
    double sequential_bw = measure_bandwidth(array, ARRAY_SIZE, sequential_access);
    double random_bw = measure_bandwidth(array, ARRAY_SIZE, random_access);
    double strided_bw = measure_bandwidth(array, ARRAY_SIZE, 
                                        (void(*)(int*, int))strided_access);
    
    printf("Memory Bandwidth:\n");
    printf("  Sequential:     %.2f MB/s\n", sequential_bw);
    printf("  Random:         %.2f MB/s (%.1f%% of sequential)\n", 
           random_bw, (random_bw / sequential_bw) * 100);
    printf("  Strided:        %.2f MB/s (%.1f%% of sequential)\n", 
           strided_bw, (strided_bw / sequential_bw) * 100);
    
    free(array);
}

// Test different array sizes to see cache effects
void test_cache_sizes() {
    printf("\n=== Cache Size Analysis ===\n");
    printf("Testing different array sizes to identify cache levels:\n\n");
    
    int sizes[] = {1024, 4096, 16384, 65536, 262144, 1048576, 4194304};
    int num_sizes = sizeof(sizes) / sizeof(sizes[0]);
    
    for (int i = 0; i < num_sizes; i++) {
        int size = sizes[i];
        int *array = malloc(size * sizeof(int));
        if (!array) continue;
        
        // Warm up the cache
        for (int j = 0; j < size; j++) {
            array[j] = j;
        }
        
        // Measure access time
        double start = get_time();
        for (int j = 0; j < size; j++) {
            array[j] = j;
        }
        double end = get_time();
        
        double time_per_element = (end - start) / size;
        double size_mb = (size * sizeof(int)) / (1024.0 * 1024.0);
        
        printf("  %8d elements (%6.2f MB): %.2e seconds/element\n", 
               size, size_mb, time_per_element);
        
        free(array);
    }
}

// Test false sharing
void test_false_sharing() {
    printf("\n=== False Sharing Test ===\n");
    
    // Allocate two integers that might be on the same cache line
    int *a = malloc(sizeof(int));
    int *b = malloc(sizeof(int));
    
    if (!a || !b) {
        printf("Memory allocation failed\n");
        return;
    }
    
    printf("Testing false sharing between two variables...\n");
    
    // Test 1: Variables on different cache lines (normal case)
    double start = get_time();
    for (int i = 0; i < 1000000; i++) {
        (*a)++;
        (*b)++;
    }
    double normal_time = get_time() - start;
    
    // Test 2: Variables on the same cache line (false sharing)
    int *c = a + 1;  // This will be on the same cache line as 'a'
    start = get_time();
    for (int i = 0; i < 1000000; i++) {
        (*a)++;
        (*c)++;
    }
    double false_sharing_time = get_time() - start;
    
    printf("Normal access:     %.6f seconds\n", normal_time);
    printf("False sharing:     %.6f seconds (%.2fx slower)\n", 
           false_sharing_time, false_sharing_time / normal_time);
    
    free(a);
    free(b);
}

// Test memory prefetching
void test_prefetching() {
    printf("\n=== Prefetching Test ===\n");
    
    int size = 1024 * 1024;
    int *array = malloc(size * sizeof(int));
    if (!array) {
        printf("Memory allocation failed\n");
        return;
    }
    
    // Initialize array
    for (int i = 0; i < size; i++) {
        array[i] = i;
    }
    
    // Test 1: No prefetching (random access)
    double start = get_time();
    for (int i = 0; i < 100000; i++) {
        int index = rand() % size;
        array[index] = array[index] + 1;
    }
    double random_time = get_time() - start;
    
    // Test 2: With prefetching (sequential access)
    start = get_time();
    for (int i = 0; i < 100000; i++) {
        array[i % size] = array[i % size] + 1;
    }
    double sequential_time = get_time() - start;
    
    printf("Random access:     %.6f seconds\n", random_time);
    printf("Sequential access: %.6f seconds (%.2fx faster)\n", 
           sequential_time, random_time / sequential_time);
    
    free(array);
}

// Test memory alignment
void test_alignment() {
    printf("\n=== Memory Alignment Test ===\n");
    
    // Test unaligned access
    char *unaligned = malloc(1000);
    if (!unaligned) {
        printf("Memory allocation failed\n");
        return;
    }
    
    // Align to cache line boundary
    char *aligned = (char*)(((uintptr_t)unaligned + CACHE_LINE_SIZE - 1) & 
                           ~(CACHE_LINE_SIZE - 1));
    
    printf("Testing aligned vs unaligned memory access...\n");
    
    // Test unaligned access
    double start = get_time();
    for (int i = 0; i < 1000000; i++) {
        *((int*)unaligned) = i;
    }
    double unaligned_time = get_time() - start;
    
    // Test aligned access
    start = get_time();
    for (int i = 0; i < 1000000; i++) {
        *((int*)aligned) = i;
    }
    double aligned_time = get_time() - start;
    
    printf("Unaligned access:  %.6f seconds\n", unaligned_time);
    printf("Aligned access:    %.6f seconds (%.2fx faster)\n", 
           aligned_time, unaligned_time / aligned_time);
    
    free(unaligned);
}

int main() {
    printf("=== Memory Access Pattern Analyzer ===\n");
    printf("This program analyzes how different memory access patterns\n");
    printf("affect performance due to CPU cache behavior.\n\n");
    
    // Set random seed for reproducible results
    srand(42);
    
    // Run all tests
    measure_cache_performance();
    measure_memory_bandwidth();
    test_cache_sizes();
    test_false_sharing();
    test_prefetching();
    test_alignment();
    
    printf("\n=== Analysis Complete ===\n");
    printf("Key takeaways:\n");
    printf("1. Sequential access is much faster than random access\n");
    printf("2. Cache line size affects performance significantly\n");
    printf("3. False sharing can severely impact performance\n");
    printf("4. Memory alignment matters for optimal performance\n");
    printf("5. Prefetching works best with predictable access patterns\n");
    
    return 0;
}
