# System Programming Basics

## Table of Contents
1. [Introduction](#introduction)
2. [C Programming for Systems](#c-programming-for-systems)
3. [Memory Management](#memory-management)
4. [File I/O and System Calls](#file-io-and-system-calls)
5. [Process Management](#process-management)
6. [Network Programming](#network-programming)
7. [Error Handling and Debugging](#error-handling-and-debugging)
8. [Practical Implementation](#practical-implementation)

## Introduction

System programming is the art of writing software that directly interacts with the operating system and hardware. Unlike application programming, system programming requires deep understanding of:

- **Low-level memory management**
- **Operating system interfaces**
- **Hardware interactions**
- **Performance optimization**
- **Security considerations**

### Why System Programming Matters

1. **Performance**: Direct control over system resources
2. **Efficiency**: Minimal overhead and maximum speed
3. **Control**: Access to all system capabilities
4. **Understanding**: Deep knowledge of how computers work
5. **Foundation**: Required for OS development, drivers, and embedded systems

### Key Concepts

- **System Calls**: Interface between user programs and the kernel
- **Memory Management**: How programs use and manage memory
- **Process Control**: Creating, managing, and communicating between processes
- **File Systems**: How data is stored and accessed
- **Networking**: How programs communicate over networks

## C Programming for Systems

### Why C for System Programming?

C is the language of choice for system programming because:
- **Direct memory access**: Pointers and manual memory management
- **Performance**: Compiles to efficient machine code
- **Portability**: Runs on virtually any platform
- **System interface**: Direct access to system calls
- **Legacy**: Most operating systems are written in C

### Essential C Concepts for Systems

#### 1. Pointers and Memory Addresses

```c
#include <stdio.h>
#include <stdint.h>

void pointer_examples() {
    int x = 42;
    int *ptr = &x;  // Pointer to x
    
    printf("Value: %d\n", x);
    printf("Address: %p\n", (void*)&x);
    printf("Pointer value: %p\n", (void*)ptr);
    printf("Dereferenced: %d\n", *ptr);
    
    // Pointer arithmetic
    int array[5] = {1, 2, 3, 4, 5};
    int *arr_ptr = array;
    
    for (int i = 0; i < 5; i++) {
        printf("array[%d] = %d\n", i, *(arr_ptr + i));
    }
}
```

#### 2. Memory Layout

```c
#include <stdio.h>
#include <stdlib.h>

// Global variables (data segment)
int global_var = 100;
static int static_var = 200;

void memory_layout_demo() {
    // Local variables (stack)
    int local_var = 300;
    static int local_static = 400;
    
    // Dynamic allocation (heap)
    int *heap_var = malloc(sizeof(int));
    *heap_var = 500;
    
    printf("Global: %p (value: %d)\n", (void*)&global_var, global_var);
    printf("Static: %p (value: %d)\n", (void*)&static_var, static_var);
    printf("Local: %p (value: %d)\n", (void*)&local_var, local_var);
    printf("Local static: %p (value: %d)\n", (void*)&local_static, local_static);
    printf("Heap: %p (value: %d)\n", (void*)heap_var, *heap_var);
    
    free(heap_var);
}
```

#### 3. Bit Manipulation

```c
#include <stdio.h>
#include <stdint.h>

void bit_manipulation_examples() {
    uint8_t value = 0b10101010;
    
    printf("Original: 0x%02x\n", value);
    
    // Set bit 3
    value |= (1 << 3);
    printf("Set bit 3: 0x%02x\n", value);
    
    // Clear bit 5
    value &= ~(1 << 5);
    printf("Clear bit 5: 0x%02x\n", value);
    
    // Toggle bit 1
    value ^= (1 << 1);
    printf("Toggle bit 1: 0x%02x\n", value);
    
    // Check if bit 2 is set
    if (value & (1 << 2)) {
        printf("Bit 2 is set\n");
    } else {
        printf("Bit 2 is not set\n");
    }
    
    // Count set bits
    int count = 0;
    for (int i = 0; i < 8; i++) {
        if (value & (1 << i)) count++;
    }
    printf("Number of set bits: %d\n", count);
}
```

#### 4. Structures and Alignment

```c
#include <stdio.h>
#include <stdint.h>

// Packed structure (no padding)
struct packed_data {
    uint8_t a;
    uint32_t b;
    uint8_t c;
} __attribute__((packed));

// Normal structure (with padding)
struct normal_data {
    uint8_t a;
    uint32_t b;
    uint8_t c;
};

void structure_alignment() {
    printf("Packed struct size: %zu bytes\n", sizeof(struct packed_data));
    printf("Normal struct size: %zu bytes\n", sizeof(struct normal_data));
    
    // Access structure members
    struct packed_data p = {1, 2, 3};
    printf("Packed: a=%d, b=%d, c=%d\n", p.a, p.b, p.c);
    
    struct normal_data n = {1, 2, 3};
    printf("Normal: a=%d, b=%d, c=%d\n", n.a, n.b, n.c);
}
```

## Memory Management

### Stack vs Heap

#### Stack Memory
- **Automatic allocation**: Variables are automatically allocated and freed
- **LIFO order**: Last In, First Out
- **Fast access**: Direct pointer arithmetic
- **Limited size**: Typically 1-8MB
- **Scope-bound**: Variables exist only within their scope

```c
void stack_example() {
    int local_var = 42;  // Allocated on stack
    char buffer[1024];   // Allocated on stack
    
    // These variables are automatically freed when function returns
}
```

#### Heap Memory
- **Manual allocation**: Must explicitly allocate and free
- **Flexible size**: Can be very large
- **Slower access**: Requires system calls
- **Persistent**: Exists until explicitly freed
- **Error-prone**: Memory leaks if not properly managed

```c
void heap_example() {
    // Allocate memory
    int *ptr = malloc(sizeof(int));
    if (ptr == NULL) {
        perror("malloc failed");
        return;
    }
    
    *ptr = 42;
    
    // Use the memory
    printf("Value: %d\n", *ptr);
    
    // Free memory
    free(ptr);
    ptr = NULL;  // Good practice
}
```

### Memory Allocation Functions

#### 1. malloc() - Memory Allocation

```c
#include <stdlib.h>
#include <stdio.h>

void malloc_example() {
    // Allocate memory for 10 integers
    int *numbers = malloc(10 * sizeof(int));
    if (numbers == NULL) {
        perror("malloc failed");
        return;
    }
    
    // Initialize the array
    for (int i = 0; i < 10; i++) {
        numbers[i] = i * i;
    }
    
    // Use the array
    for (int i = 0; i < 10; i++) {
        printf("numbers[%d] = %d\n", i, numbers[i]);
    }
    
    // Free the memory
    free(numbers);
}
```

#### 2. calloc() - Contiguous Allocation

```c
void calloc_example() {
    // Allocate and zero-initialize memory
    int *zeros = calloc(100, sizeof(int));
    if (zeros == NULL) {
        perror("calloc failed");
        return;
    }
    
    // All elements are already zero
    printf("First element: %d\n", zeros[0]);
    
    free(zeros);
}
```

#### 3. realloc() - Reallocate Memory

```c
void realloc_example() {
    // Allocate initial memory
    int *data = malloc(5 * sizeof(int));
    if (data == NULL) {
        perror("malloc failed");
        return;
    }
    
    // Initialize
    for (int i = 0; i < 5; i++) {
        data[i] = i;
    }
    
    // Resize to 10 elements
    int *new_data = realloc(data, 10 * sizeof(int));
    if (new_data == NULL) {
        perror("realloc failed");
        free(data);
        return;
    }
    
    data = new_data;  // Update pointer
    
    // Initialize new elements
    for (int i = 5; i < 10; i++) {
        data[i] = i;
    }
    
    // Use the resized array
    for (int i = 0; i < 10; i++) {
        printf("data[%d] = %d\n", i, data[i]);
    }
    
    free(data);
}
```

### Memory Management Best Practices

#### 1. Always Check for NULL

```c
void safe_allocation() {
    int *ptr = malloc(sizeof(int));
    if (ptr == NULL) {
        fprintf(stderr, "Memory allocation failed\n");
        return;
    }
    
    // Use ptr safely
    *ptr = 42;
    printf("Value: %d\n", *ptr);
    
    free(ptr);
}
```

#### 2. Free Memory in Reverse Order

```c
void cleanup_example() {
    char *str1 = malloc(100);
    char *str2 = malloc(200);
    char *str3 = malloc(300);
    
    if (str1 == NULL || str2 == NULL || str3 == NULL) {
        // Clean up in reverse order
        free(str3);
        free(str2);
        free(str1);
        return;
    }
    
    // Use the strings...
    
    // Free in reverse order
    free(str3);
    free(str2);
    free(str1);
}
```

#### 3. Use Valgrind for Memory Debugging

```bash
# Compile with debug symbols
gcc -g -o program program.c

# Run with valgrind
valgrind --leak-check=full --show-leak-kinds=all ./program
```

## File I/O and System Calls

### System Calls

System calls are the interface between user programs and the operating system kernel.

#### Common System Calls

```c
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

void system_call_examples() {
    // File operations
    int fd = open("file.txt", O_RDONLY);
    if (fd == -1) {
        perror("open failed");
        return;
    }
    
    char buffer[1024];
    ssize_t bytes_read = read(fd, buffer, sizeof(buffer));
    if (bytes_read == -1) {
        perror("read failed");
        close(fd);
        return;
    }
    
    close(fd);
}
```

### File Operations

#### 1. Opening Files

```c
#include <fcntl.h>
#include <unistd.h>

void file_operations() {
    // Open for reading
    int fd_read = open("input.txt", O_RDONLY);
    if (fd_read == -1) {
        perror("Failed to open input.txt");
        return;
    }
    
    // Open for writing (create if doesn't exist)
    int fd_write = open("output.txt", O_WRONLY | O_CREAT | O_TRUNC, 0644);
    if (fd_write == -1) {
        perror("Failed to open output.txt");
        close(fd_read);
        return;
    }
    
    // Open for appending
    int fd_append = open("log.txt", O_WRONLY | O_CREAT | O_APPEND, 0644);
    if (fd_append == -1) {
        perror("Failed to open log.txt");
        close(fd_read);
        close(fd_write);
        return;
    }
    
    // Close all files
    close(fd_read);
    close(fd_write);
    close(fd_append);
}
```

#### 2. Reading and Writing

```c
void read_write_example() {
    int fd = open("data.txt", O_RDWR | O_CREAT, 0644);
    if (fd == -1) {
        perror("open failed");
        return;
    }
    
    // Write data
    const char *data = "Hello, World!\n";
    ssize_t bytes_written = write(fd, data, strlen(data));
    if (bytes_written == -1) {
        perror("write failed");
        close(fd);
        return;
    }
    
    // Seek to beginning
    if (lseek(fd, 0, SEEK_SET) == -1) {
        perror("lseek failed");
        close(fd);
        return;
    }
    
    // Read data
    char buffer[1024];
    ssize_t bytes_read = read(fd, buffer, sizeof(buffer) - 1);
    if (bytes_read == -1) {
        perror("read failed");
        close(fd);
        return;
    }
    
    buffer[bytes_read] = '\0';
    printf("Read: %s", buffer);
    
    close(fd);
}
```

#### 3. File Metadata

```c
#include <sys/stat.h>

void file_metadata() {
    struct stat file_stat;
    
    if (stat("file.txt", &file_stat) == -1) {
        perror("stat failed");
        return;
    }
    
    printf("File size: %ld bytes\n", file_stat.st_size);
    printf("File mode: %o\n", file_stat.st_mode);
    printf("User ID: %d\n", file_stat.st_uid);
    printf("Group ID: %d\n", file_stat.st_gid);
    printf("Last access: %ld\n", file_stat.st_atime);
    printf("Last modification: %ld\n", file_stat.st_mtime);
    printf("Last status change: %ld\n", file_stat.st_ctime);
    
    // Check file type
    if (S_ISREG(file_stat.st_mode)) {
        printf("Regular file\n");
    } else if (S_ISDIR(file_stat.st_mode)) {
        printf("Directory\n");
    } else if (S_ISLNK(file_stat.st_mode)) {
        printf("Symbolic link\n");
    }
}
```

### Directory Operations

```c
#include <dirent.h>
#include <sys/stat.h>

void directory_operations() {
    DIR *dir = opendir(".");
    if (dir == NULL) {
        perror("opendir failed");
        return;
    }
    
    struct dirent *entry;
    while ((entry = readdir(dir)) != NULL) {
        printf("Name: %s\n", entry->d_name);
        printf("Type: %d\n", entry->d_type);
        printf("Inode: %lu\n", entry->d_ino);
        printf("---\n");
    }
    
    closedir(dir);
}
```

## Process Management

### Process Creation

#### 1. fork() - Create Child Process

```c
#include <unistd.h>
#include <sys/wait.h>
#include <stdio.h>

void fork_example() {
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        return;
    } else if (pid == 0) {
        // Child process
        printf("Child process: PID = %d\n", getpid());
        printf("Child process: Parent PID = %d\n", getppid());
        exit(0);
    } else {
        // Parent process
        printf("Parent process: PID = %d\n", getpid());
        printf("Parent process: Child PID = %d\n", pid);
        
        // Wait for child to complete
        int status;
        waitpid(pid, &status, 0);
        printf("Child process exited with status: %d\n", status);
    }
}
```

#### 2. exec() - Replace Process Image

```c
void exec_example() {
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        return;
    } else if (pid == 0) {
        // Child process - execute ls command
        execl("/bin/ls", "ls", "-l", NULL);
        perror("execl failed");
        exit(1);
    } else {
        // Parent process
        int status;
        waitpid(pid, &status, 0);
        printf("Child process completed\n");
    }
}
```

### Process Communication

#### 1. Pipes

```c
#include <unistd.h>

void pipe_example() {
    int pipefd[2];
    
    if (pipe(pipefd) == -1) {
        perror("pipe failed");
        return;
    }
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        close(pipefd[0]);
        close(pipefd[1]);
        return;
    } else if (pid == 0) {
        // Child process - write to pipe
        close(pipefd[0]);  // Close read end
        
        const char *message = "Hello from child!";
        write(pipefd[1], message, strlen(message));
        close(pipefd[1]);
        exit(0);
    } else {
        // Parent process - read from pipe
        close(pipefd[1]);  // Close write end
        
        char buffer[1024];
        ssize_t bytes_read = read(pipefd[0], buffer, sizeof(buffer) - 1);
        if (bytes_read > 0) {
            buffer[bytes_read] = '\0';
            printf("Received: %s\n", buffer);
        }
        
        close(pipefd[0]);
        wait(NULL);
    }
}
```

#### 2. Shared Memory

```c
#include <sys/ipc.h>
#include <sys/shm.h>

void shared_memory_example() {
    // Create shared memory segment
    key_t key = ftok("shared_mem", 65);
    int shmid = shmget(key, 1024, 0666 | IPC_CREAT);
    
    if (shmid == -1) {
        perror("shmget failed");
        return;
    }
    
    // Attach to shared memory
    char *shared_memory = (char*)shmat(shmid, NULL, 0);
    if (shared_memory == (char*)-1) {
        perror("shmat failed");
        return;
    }
    
    pid_t pid = fork();
    
    if (pid == 0) {
        // Child process - write to shared memory
        strcpy(shared_memory, "Hello from shared memory!");
        shmdt(shared_memory);
        exit(0);
    } else {
        // Parent process - read from shared memory
        wait(NULL);
        printf("Read from shared memory: %s\n", shared_memory);
        
        // Detach and remove shared memory
        shmdt(shared_memory);
        shmctl(shmid, IPC_RMID, NULL);
    }
}
```

## Network Programming

### Socket Programming

#### 1. TCP Server

```c
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

void tcp_server() {
    int server_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (server_fd == -1) {
        perror("socket failed");
        return;
    }
    
    // Set socket options
    int opt = 1;
    if (setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt)) == -1) {
        perror("setsockopt failed");
        close(server_fd);
        return;
    }
    
    // Bind socket
    struct sockaddr_in address;
    address.sin_family = AF_INET;
    address.sin_addr.s_addr = INADDR_ANY;
    address.sin_port = htons(8080);
    
    if (bind(server_fd, (struct sockaddr*)&address, sizeof(address)) == -1) {
        perror("bind failed");
        close(server_fd);
        return;
    }
    
    // Listen for connections
    if (listen(server_fd, 5) == -1) {
        perror("listen failed");
        close(server_fd);
        return;
    }
    
    printf("Server listening on port 8080...\n");
    
    // Accept connections
    int client_fd;
    struct sockaddr_in client_addr;
    socklen_t client_len = sizeof(client_addr);
    
    client_fd = accept(server_fd, (struct sockaddr*)&client_addr, &client_len);
    if (client_fd == -1) {
        perror("accept failed");
        close(server_fd);
        return;
    }
    
    printf("Client connected: %s\n", inet_ntoa(client_addr.sin_addr));
    
    // Send data to client
    const char *message = "Hello, Client!";
    send(client_fd, message, strlen(message), 0);
    
    // Close connections
    close(client_fd);
    close(server_fd);
}
```

#### 2. TCP Client

```c
void tcp_client() {
    int client_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (client_fd == -1) {
        perror("socket failed");
        return;
    }
    
    // Connect to server
    struct sockaddr_in server_addr;
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(8080);
    inet_pton(AF_INET, "127.0.0.1", &server_addr.sin_addr);
    
    if (connect(client_fd, (struct sockaddr*)&server_addr, sizeof(server_addr)) == -1) {
        perror("connect failed");
        close(client_fd);
        return;
    }
    
    // Receive data from server
    char buffer[1024];
    ssize_t bytes_received = recv(client_fd, buffer, sizeof(buffer) - 1, 0);
    if (bytes_received > 0) {
        buffer[bytes_received] = '\0';
        printf("Received: %s\n", buffer);
    }
    
    close(client_fd);
}
```

## Error Handling and Debugging

### Error Handling

#### 1. Using errno

```c
#include <errno.h>
#include <string.h>

void error_handling_example() {
    int fd = open("nonexistent.txt", O_RDONLY);
    if (fd == -1) {
        printf("Error: %s\n", strerror(errno));
        printf("Error code: %d\n", errno);
        return;
    }
    
    close(fd);
}
```

#### 2. Custom Error Handling

```c
#include <stdarg.h>

void error_exit(const char *format, ...) {
    va_list args;
    va_start(args, format);
    
    fprintf(stderr, "Error: ");
    vfprintf(stderr, format, args);
    fprintf(stderr, "\n");
    
    va_end(args);
    exit(1);
}

void custom_error_example() {
    int *ptr = malloc(sizeof(int));
    if (ptr == NULL) {
        error_exit("Memory allocation failed");
    }
    
    // Use ptr...
    free(ptr);
}
```

### Debugging Techniques

#### 1. Using gdb

```bash
# Compile with debug symbols
gcc -g -o program program.c

# Start gdb
gdb ./program

# Set breakpoints
(gdb) break main
(gdb) break function_name

# Run program
(gdb) run

# Step through code
(gdb) step
(gdb) next

# Print variables
(gdb) print variable_name
(gdb) print *pointer

# Continue execution
(gdb) continue

# Quit gdb
(gdb) quit
```

#### 2. Using strace

```bash
# Trace system calls
strace ./program

# Trace specific system calls
strace -e trace=open,read,write ./program

# Save trace to file
strace -o trace.log ./program
```

## Practical Implementation

### Project 1: File Copy Utility

```c
// file_copy.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <errno.h>

int copy_file(const char *src, const char *dst) {
    int src_fd = open(src, O_RDONLY);
    if (src_fd == -1) {
        perror("Failed to open source file");
        return -1;
    }
    
    int dst_fd = open(dst, O_WRONLY | O_CREAT | O_TRUNC, 0644);
    if (dst_fd == -1) {
        perror("Failed to open destination file");
        close(src_fd);
        return -1;
    }
    
    char buffer[4096];
    ssize_t bytes_read, bytes_written;
    
    while ((bytes_read = read(src_fd, buffer, sizeof(buffer))) > 0) {
        bytes_written = write(dst_fd, buffer, bytes_read);
        if (bytes_written != bytes_read) {
            perror("Write error");
            close(src_fd);
            close(dst_fd);
            return -1;
        }
    }
    
    if (bytes_read == -1) {
        perror("Read error");
        close(src_fd);
        close(dst_fd);
        return -1;
    }
    
    close(src_fd);
    close(dst_fd);
    return 0;
}

int main(int argc, char *argv[]) {
    if (argc != 3) {
        fprintf(stderr, "Usage: %s <source> <destination>\n", argv[0]);
        return 1;
    }
    
    if (copy_file(argv[1], argv[2]) == 0) {
        printf("File copied successfully\n");
        return 0;
    } else {
        return 1;
    }
}
```

### Project 2: Process Monitor

```c
// process_monitor.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <signal.h>
#include <time.h>

void signal_handler(int sig) {
    printf("Received signal %d\n", sig);
}

void process_monitor() {
    // Set up signal handler
    signal(SIGCHLD, signal_handler);
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        return;
    } else if (pid == 0) {
        // Child process
        printf("Child process started (PID: %d)\n", getpid());
        
        // Simulate work
        for (int i = 0; i < 5; i++) {
            printf("Child working... %d\n", i);
            sleep(1);
        }
        
        printf("Child process finished\n");
        exit(0);
    } else {
        // Parent process
        printf("Parent process (PID: %d) monitoring child (PID: %d)\n", 
               getpid(), pid);
        
        int status;
        pid_t child_pid = waitpid(pid, &status, 0);
        
        if (child_pid == -1) {
            perror("waitpid failed");
            return;
        }
        
        if (WIFEXITED(status)) {
            printf("Child process %d exited with status %d\n", 
                   child_pid, WEXITSTATUS(status));
        } else if (WIFSIGNALED(status)) {
            printf("Child process %d killed by signal %d\n", 
                   child_pid, WTERMSIG(status));
        }
    }
}

int main() {
    process_monitor();
    return 0;
}
```

### Project 3: Memory Allocator

```c
// custom_allocator.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

typedef struct block {
    size_t size;
    int free;
    struct block *next;
    struct block *prev;
} block_t;

typedef struct {
    void *memory;
    size_t total_size;
    block_t *free_list;
} allocator_t;

allocator_t* create_allocator(size_t size) {
    allocator_t *allocator = malloc(sizeof(allocator_t));
    if (!allocator) return NULL;
    
    allocator->memory = sbrk(size);
    if (allocator->memory == (void*)-1) {
        free(allocator);
        return NULL;
    }
    
    allocator->total_size = size;
    allocator->free_list = (block_t*)allocator->memory;
    allocator->free_list->size = size - sizeof(block_t);
    allocator->free_list->free = 1;
    allocator->free_list->next = NULL;
    allocator->free_list->prev = NULL;
    
    return allocator;
}

void* allocate(allocator_t *allocator, size_t size) {
    block_t *current = allocator->free_list;
    
    while (current) {
        if (current->free && current->size >= size) {
            // Split block if it's too large
            if (current->size > size + sizeof(block_t)) {
                block_t *new_block = (block_t*)((char*)current + sizeof(block_t) + size);
                new_block->size = current->size - size - sizeof(block_t);
                new_block->free = 1;
                new_block->next = current->next;
                new_block->prev = current;
                
                if (current->next) {
                    current->next->prev = new_block;
                }
                
                current->next = new_block;
                current->size = size;
            }
            
            current->free = 0;
            return (void*)((char*)current + sizeof(block_t));
        }
        current = current->next;
    }
    
    return NULL; // No free block found
}

void deallocate(allocator_t *allocator, void *ptr) {
    if (!ptr) return;
    
    block_t *block = (block_t*)((char*)ptr - sizeof(block_t));
    block->free = 1;
    
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

void print_allocator_stats(allocator_t *allocator) {
    printf("Allocator Statistics:\n");
    printf("Total size: %zu bytes\n", allocator->total_size);
    
    block_t *current = allocator->free_list;
    int free_blocks = 0;
    int used_blocks = 0;
    size_t free_bytes = 0;
    size_t used_bytes = 0;
    
    while (current) {
        if (current->free) {
            free_blocks++;
            free_bytes += current->size;
        } else {
            used_blocks++;
            used_bytes += current->size;
        }
        current = current->next;
    }
    
    printf("Free blocks: %d (%zu bytes)\n", free_blocks, free_bytes);
    printf("Used blocks: %d (%zu bytes)\n", used_blocks, used_bytes);
}

int main() {
    allocator_t *allocator = create_allocator(1024 * 1024); // 1MB
    if (!allocator) {
        printf("Failed to create allocator\n");
        return 1;
    }
    
    // Test allocation
    int *numbers = (int*)allocate(allocator, 10 * sizeof(int));
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
    
    // Test deallocation
    deallocate(allocator, numbers);
    print_allocator_stats(allocator);
    
    free(allocator);
    return 0;
}
```

## Key Takeaways

1. **C is Essential**: Master C programming for system-level work
2. **Memory Management**: Understand stack vs heap, allocation strategies
3. **System Calls**: Learn the interface between user and kernel space
4. **Process Control**: Understand how processes are created and managed
5. **Error Handling**: Always check return values and handle errors properly
6. **Debugging**: Use tools like gdb, valgrind, and strace effectively

## Next Steps

In the next module, we'll dive into **Process Management** where you'll learn:
- Process creation and termination
- Process scheduling algorithms
- Inter-process communication
- Process synchronization
- Building a process manager

## Exercises

1. **File System Utility**: Build a program that recursively copies directories
2. **Process Tree**: Create a program that displays a process tree
3. **Memory Pool**: Implement a memory pool allocator
4. **Network Server**: Build a multi-threaded network server
5. **System Monitor**: Create a system resource monitor

Remember: The goal is to understand not just how to use these system calls, but why they work the way they do and how to use them efficiently!
