#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <stdint.h>
#include <assert.h>

// Block structure for our allocator
typedef struct block {
    size_t size;
    int free;
    struct block *next;
    struct block *prev;
} block_t;

// Allocator structure
typedef struct {
    void *memory;
    size_t total_size;
    block_t *free_list;
    block_t *used_list;
    size_t allocated_bytes;
    size_t free_bytes;
    int num_blocks;
} allocator_t;

// Global allocator instance
static allocator_t *global_allocator = NULL;

// Create a new allocator
allocator_t* create_allocator(size_t size) {
    allocator_t *allocator = malloc(sizeof(allocator_t));
    if (!allocator) return NULL;
    
    // Allocate memory using sbrk
    allocator->memory = sbrk(size);
    if (allocator->memory == (void*)-1) {
        free(allocator);
        return NULL;
    }
    
    allocator->total_size = size;
    allocator->allocated_bytes = 0;
    allocator->free_bytes = size - sizeof(block_t);
    allocator->num_blocks = 0;
    
    // Initialize the first free block
    allocator->free_list = (block_t*)allocator->memory;
    allocator->free_list->size = size - sizeof(block_t);
    allocator->free_list->free = 1;
    allocator->free_list->next = NULL;
    allocator->free_list->prev = NULL;
    
    allocator->used_list = NULL;
    
    return allocator;
}

// Destroy an allocator
void destroy_allocator(allocator_t *allocator) {
    if (allocator) {
        // Reset the program break
        sbrk(-allocator->total_size);
        free(allocator);
    }
}

// Find a free block that can accommodate the requested size
block_t* find_free_block(allocator_t *allocator, size_t size) {
    block_t *current = allocator->free_list;
    
    while (current) {
        if (current->free && current->size >= size) {
            return current;
        }
        current = current->next;
    }
    
    return NULL;
}

// Split a block if it's too large
void split_block(block_t *block, size_t size) {
    if (block->size <= size + sizeof(block_t)) {
        return; // Not worth splitting
    }
    
    // Create new block
    block_t *new_block = (block_t*)((char*)block + sizeof(block_t) + size);
    new_block->size = block->size - size - sizeof(block_t);
    new_block->free = 1;
    new_block->next = block->next;
    new_block->prev = block;
    
    // Update links
    if (block->next) {
        block->next->prev = new_block;
    }
    block->next = new_block;
    
    // Update original block
    block->size = size;
}

// Merge adjacent free blocks
void merge_free_blocks(allocator_t *allocator, block_t *block) {
    // Merge with next block if it's free
    if (block->next && block->next->free) {
        block->size += block->next->size + sizeof(block_t);
        block->next = block->next->next;
        if (block->next) {
            block->next->prev = block;
        }
    }
    
    // Merge with previous block if it's free
    if (block->prev && block->prev->free) {
        block->prev->size += block->size + sizeof(block_t);
        block->prev->next = block->next;
        if (block->next) {
            block->next->prev = block->prev;
        }
    }
}

// Add block to used list
void add_to_used_list(allocator_t *allocator, block_t *block) {
    block->next = allocator->used_list;
    block->prev = NULL;
    if (allocator->used_list) {
        allocator->used_list->prev = block;
    }
    allocator->used_list = block;
}

// Remove block from used list
void remove_from_used_list(allocator_t *allocator, block_t *block) {
    if (block->prev) {
        block->prev->next = block->next;
    } else {
        allocator->used_list = block->next;
    }
    
    if (block->next) {
        block->next->prev = block->prev;
    }
}

// Allocate memory
void* allocate(allocator_t *allocator, size_t size, size_t alignment) {
    if (!allocator || size == 0) return NULL;
    
    // Align size to alignment boundary
    size = (size + alignment - 1) & ~(alignment - 1);
    
    // Find a free block
    block_t *block = find_free_block(allocator, size);
    if (!block) {
        printf("No free block found for size %zu\n", size);
        return NULL;
    }
    
    // Split block if necessary
    split_block(block, size);
    
    // Mark as used
    block->free = 0;
    
    // Add to used list
    add_to_used_list(allocator, block);
    
    // Update statistics
    allocator->allocated_bytes += size;
    allocator->free_bytes -= size;
    allocator->num_blocks++;
    
    // Return pointer to data area
    return (void*)((char*)block + sizeof(block_t));
}

// Deallocate memory
void deallocate(allocator_t *allocator, void *ptr) {
    if (!allocator || !ptr) return;
    
    // Get block pointer
    block_t *block = (block_t*)((char*)ptr - sizeof(block_t));
    
    // Check if block is valid
    if (block < (block_t*)allocator->memory || 
        block >= (block_t*)((char*)allocator->memory + allocator->total_size)) {
        printf("Invalid pointer: %p\n", ptr);
        return;
    }
    
    // Mark as free
    block->free = 1;
    
    // Remove from used list
    remove_from_used_list(allocator, block);
    
    // Update statistics
    allocator->allocated_bytes -= block->size;
    allocator->free_bytes += block->size;
    allocator->num_blocks--;
    
    // Merge with adjacent free blocks
    merge_free_blocks(allocator, block);
}

// Print allocator statistics
void print_allocator_stats(allocator_t *allocator) {
    if (!allocator) return;
    
    printf("\n=== Allocator Statistics ===\n");
    printf("Total size: %zu bytes\n", allocator->total_size);
    printf("Allocated: %zu bytes\n", allocator->allocated_bytes);
    printf("Free: %zu bytes\n", allocator->free_bytes);
    printf("Number of blocks: %d\n", allocator->num_blocks);
    printf("Fragmentation: %.2f%%\n", 
           (1.0 - (double)allocator->free_bytes / allocator->total_size) * 100);
    
    // Print free blocks
    printf("\nFree blocks:\n");
    block_t *current = allocator->free_list;
    int free_count = 0;
    while (current) {
        printf("  Block %d: size=%zu, addr=%p\n", 
               free_count++, current->size, (void*)current);
        current = current->next;
    }
    
    // Print used blocks
    printf("\nUsed blocks:\n");
    current = allocator->used_list;
    int used_count = 0;
    while (current) {
        printf("  Block %d: size=%zu, addr=%p\n", 
               used_count++, current->size, (void*)current);
        current = current->next;
    }
}

// Test the allocator
void test_allocator() {
    printf("=== Testing Custom Allocator ===\n");
    
    // Create allocator
    allocator_t *allocator = create_allocator(1024 * 1024); // 1MB
    if (!allocator) {
        printf("Failed to create allocator\n");
        return;
    }
    
    print_allocator_stats(allocator);
    
    // Test 1: Simple allocation
    printf("\n--- Test 1: Simple Allocation ---\n");
    int *numbers = (int*)allocate(allocator, 10 * sizeof(int), 4);
    if (numbers) {
        for (int i = 0; i < 10; i++) {
            numbers[i] = i * i;
        }
        printf("Allocated array: ");
        for (int i = 0; i < 10; i++) {
            printf("%d ", numbers[i]);
        }
        printf("\n");
    }
    
    print_allocator_stats(allocator);
    
    // Test 2: Multiple allocations
    printf("\n--- Test 2: Multiple Allocations ---\n");
    char *str1 = (char*)allocate(allocator, 100, 1);
    char *str2 = (char*)allocate(allocator, 200, 1);
    char *str3 = (char*)allocate(allocator, 300, 1);
    
    if (str1 && str2 && str3) {
        strcpy(str1, "Hello");
        strcpy(str2, "World");
        strcpy(str3, "Allocator");
        printf("Strings: %s %s %s\n", str1, str2, str3);
    }
    
    print_allocator_stats(allocator);
    
    // Test 3: Deallocation
    printf("\n--- Test 3: Deallocation ---\n");
    deallocate(allocator, numbers);
    deallocate(allocator, str2);
    
    print_allocator_stats(allocator);
    
    // Test 4: Reallocation
    printf("\n--- Test 4: Reallocation ---\n");
    int *new_numbers = (int*)allocate(allocator, 20 * sizeof(int), 4);
    if (new_numbers) {
        for (int i = 0; i < 20; i++) {
            new_numbers[i] = i * i * i;
        }
        printf("New array: ");
        for (int i = 0; i < 20; i++) {
            printf("%d ", new_numbers[i]);
        }
        printf("\n");
    }
    
    print_allocator_stats(allocator);
    
    // Test 5: Stress test
    printf("\n--- Test 5: Stress Test ---\n");
    void *ptrs[100];
    for (int i = 0; i < 100; i++) {
        ptrs[i] = allocate(allocator, (i + 1) * 10, 1);
        if (!ptrs[i]) {
            printf("Allocation failed at iteration %d\n", i);
            break;
        }
    }
    
    print_allocator_stats(allocator);
    
    // Free some random blocks
    for (int i = 0; i < 100; i += 3) {
        if (ptrs[i]) {
            deallocate(allocator, ptrs[i]);
            ptrs[i] = NULL;
        }
    }
    
    print_allocator_stats(allocator);
    
    // Clean up
    for (int i = 0; i < 100; i++) {
        if (ptrs[i]) {
            deallocate(allocator, ptrs[i]);
        }
    }
    
    deallocate(allocator, str1);
    deallocate(allocator, str3);
    deallocate(allocator, new_numbers);
    
    print_allocator_stats(allocator);
    
    // Destroy allocator
    destroy_allocator(allocator);
}

// Benchmark against standard malloc
void benchmark_allocator() {
    printf("\n=== Benchmarking Allocator ===\n");
    
    const int num_allocations = 10000;
    const size_t allocation_size = 1024;
    
    // Test custom allocator
    allocator_t *allocator = create_allocator(1024 * 1024 * 10); // 10MB
    if (!allocator) {
        printf("Failed to create allocator\n");
        return;
    }
    
    void *ptrs[num_allocations];
    
    // Time custom allocator
    clock_t start = clock();
    for (int i = 0; i < num_allocations; i++) {
        ptrs[i] = allocate(allocator, allocation_size, 1);
    }
    for (int i = 0; i < num_allocations; i++) {
        deallocate(allocator, ptrs[i]);
    }
    clock_t end = clock();
    
    double custom_time = ((double)(end - start)) / CLOCKS_PER_SEC;
    printf("Custom allocator time: %.6f seconds\n", custom_time);
    
    // Time standard malloc
    start = clock();
    for (int i = 0; i < num_allocations; i++) {
        ptrs[i] = malloc(allocation_size);
    }
    for (int i = 0; i < num_allocations; i++) {
        free(ptrs[i]);
    }
    end = clock();
    
    double standard_time = ((double)(end - start)) / CLOCKS_PER_SEC;
    printf("Standard malloc time: %.6f seconds\n", standard_time);
    printf("Speedup: %.2fx\n", standard_time / custom_time);
    
    destroy_allocator(allocator);
}

int main() {
    printf("=== Custom Memory Allocator ===\n");
    
    // Run tests
    test_allocator();
    
    // Run benchmark
    benchmark_allocator();
    
    return 0;
}
