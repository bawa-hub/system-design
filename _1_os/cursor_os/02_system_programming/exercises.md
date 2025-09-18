# System Programming Basics - Exercises

## Exercise 1: Advanced File Operations (Beginner)

### Task
Extend the file copy utility to support more advanced features.

### Requirements
1. Add support for symbolic links
2. Implement file permissions preservation
3. Add progress bar with percentage
4. Support for copying multiple files to a directory
5. Add file integrity checking (checksums)

### Code Template
```c
// advanced_file_ops.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <sys/sendfile.h>
#include <errno.h>

int copy_file_advanced(const char *src, const char *dst, int preserve_links) {
    // Implement advanced file copying
}

int copy_multiple_files(char *files[], int count, const char *dest_dir) {
    // Implement multiple file copying
}

void show_progress(size_t current, size_t total) {
    // Implement progress bar
}

uint32_t calculate_checksum(const char *filename) {
    // Implement file checksum calculation
}
```

### Expected Output
```
Copying file1.txt to backup/
Progress: [████████████████████] 100%
File copied successfully (checksum: 0x12345678)
```

## Exercise 2: Process Tree Manager (Intermediate)

### Task
Create a process tree manager that can create, monitor, and control process hierarchies.

### Requirements
1. Create process trees with parent-child relationships
2. Implement process group management
3. Add signal handling for process groups
4. Create a process tree visualization
5. Implement process resource monitoring

### Code Template
```c
// process_tree.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <signal.h>

typedef struct process_node {
    pid_t pid;
    char name[256];
    struct process_node *children;
    struct process_node *next;
    struct process_node *parent;
    int status;
} process_node_t;

typedef struct process_tree {
    process_node_t *root;
    int num_processes;
} process_tree_t;

process_tree_t* create_process_tree(void);
process_node_t* add_process(process_tree_t *tree, pid_t pid, const char *name);
void remove_process(process_tree_t *tree, pid_t pid);
void print_process_tree(process_tree_t *tree);
void kill_process_tree(process_tree_t *tree, int signal);
```

## Exercise 3: Memory Pool Allocator (Intermediate)

### Task
Implement a memory pool allocator that pre-allocates memory in fixed-size blocks.

### Requirements
1. Pre-allocate memory in fixed-size blocks
2. Implement fast allocation/deallocation
3. Add memory alignment support
4. Implement block size statistics
5. Add memory fragmentation analysis

### Code Template
```c
// memory_pool.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct pool_block {
    void *memory;
    size_t size;
    int free;
    struct pool_block *next;
} pool_block_t;

typedef struct memory_pool {
    void *memory;
    size_t block_size;
    size_t num_blocks;
    pool_block_t *free_list;
    pool_block_t *used_list;
} memory_pool_t;

memory_pool_t* create_memory_pool(size_t block_size, size_t num_blocks);
void* pool_allocate(memory_pool_t *pool);
void pool_deallocate(memory_pool_t *pool, void *ptr);
void print_pool_stats(memory_pool_t *pool);
void destroy_memory_pool(memory_pool_t *pool);
```

## Exercise 4: Network Server Framework (Advanced)

### Task
Build a multi-threaded network server framework with connection pooling.

### Requirements
1. Multi-threaded server with thread pool
2. Connection pooling and reuse
3. Request/response handling
4. Logging and monitoring
5. Graceful shutdown handling

### Code Template
```c
// network_server.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <pthread.h>

typedef struct connection {
    int socket_fd;
    struct sockaddr_in client_addr;
    time_t connect_time;
    int active;
} connection_t;

typedef struct thread_pool {
    pthread_t *threads;
    int num_threads;
    connection_t *connections;
    int max_connections;
    pthread_mutex_t mutex;
    pthread_cond_t condition;
} thread_pool_t;

typedef struct server {
    int listen_fd;
    int port;
    thread_pool_t *pool;
    int running;
} server_t;

server_t* create_server(int port, int num_threads, int max_connections);
int start_server(server_t *server);
void stop_server(server_t *server);
void destroy_server(server_t *server);
```

## Exercise 5: System Resource Monitor (Advanced)

### Task
Create a comprehensive system resource monitor that tracks CPU, memory, and I/O usage.

### Requirements
1. Real-time CPU usage monitoring
2. Memory usage tracking
3. I/O statistics collection
4. Process resource usage
5. Historical data storage

### Code Template
```c
// system_monitor.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/sysinfo.h>
#include <sys/stat.h>
#include <string.h>

typedef struct cpu_stats {
    unsigned long user;
    unsigned long nice;
    unsigned long system;
    unsigned long idle;
    unsigned long iowait;
    unsigned long irq;
    unsigned long softirq;
} cpu_stats_t;

typedef struct memory_stats {
    unsigned long total;
    unsigned long free;
    unsigned long available;
    unsigned long buffers;
    unsigned long cached;
    unsigned long swap_total;
    unsigned long swap_free;
} memory_stats_t;

typedef struct io_stats {
    unsigned long read_ops;
    unsigned long write_ops;
    unsigned long read_bytes;
    unsigned long write_bytes;
} io_stats_t;

typedef struct system_monitor {
    cpu_stats_t cpu;
    memory_stats_t memory;
    io_stats_t io;
    time_t timestamp;
} system_monitor_t;

int get_cpu_stats(cpu_stats_t *stats);
int get_memory_stats(memory_stats_t *stats);
int get_io_stats(io_stats_t *stats);
void print_system_stats(system_monitor_t *monitor);
void monitor_system_loop(int interval);
```

## Exercise 6: Custom Shell Implementation (Expert)

### Task
Implement a custom shell with advanced features like job control and command history.

### Requirements
1. Basic command execution
2. Pipes and redirection
3. Job control (background/foreground)
4. Command history
5. Tab completion
6. Environment variable handling

### Code Template
```c
// custom_shell.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/wait.h>
#include <signal.h>
#include <readline/readline.h>
#include <readline/history.h>

typedef struct job {
    pid_t pid;
    char command[256];
    int status;
    int job_id;
    struct job *next;
} job_t;

typedef struct shell {
    job_t *jobs;
    int num_jobs;
    char *prompt;
    int running;
} shell_t;

shell_t* create_shell(void);
void destroy_shell(shell_t *shell);
int execute_command(shell_t *shell, char *command);
int execute_pipeline(shell_t *shell, char **commands, int count);
void handle_job_control(shell_t *shell);
void print_jobs(shell_t *shell);
```

## Exercise 7: Inter-Process Communication Library (Expert)

### Task
Create a comprehensive IPC library supporting multiple communication methods.

### Requirements
1. Named pipes (FIFOs)
2. Message queues
3. Shared memory
4. Semaphores
5. Sockets (Unix domain)
6. Signal-based communication

### Code Template
```c
// ipc_library.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/ipc.h>
#include <sys/msg.h>
#include <sys/shm.h>
#include <sys/sem.h>
#include <sys/socket.h>
#include <sys/un.h>

typedef enum {
    IPC_PIPE,
    IPC_MSG_QUEUE,
    IPC_SHARED_MEM,
    IPC_SEMAPHORE,
    IPC_SOCKET
} ipc_type_t;

typedef struct ipc_handle {
    ipc_type_t type;
    int id;
    void *data;
    size_t size;
} ipc_handle_t;

ipc_handle_t* create_ipc(ipc_type_t type, const char *name, size_t size);
int send_message(ipc_handle_t *handle, void *data, size_t size);
int receive_message(ipc_handle_t *handle, void *data, size_t size);
void destroy_ipc(ipc_handle_t *handle);
```

## Exercise 8: File System Watcher (Expert)

### Task
Implement a file system watcher that monitors directory changes in real-time.

### Requirements
1. Monitor directory changes using inotify
2. Support for recursive directory monitoring
3. Event filtering and processing
4. Asynchronous event handling
5. Configuration file support

### Code Template
```c
// file_watcher.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/inotify.h>
#include <sys/select.h>
#include <string.h>

typedef struct watch_entry {
    int wd;
    char path[256];
    struct watch_entry *next;
} watch_entry_t;

typedef struct file_watcher {
    int inotify_fd;
    watch_entry_t *watches;
    int num_watches;
    int running;
} file_watcher_t;

typedef void (*event_callback_t)(const char *path, uint32_t mask);

file_watcher_t* create_file_watcher(void);
int add_watch(file_watcher_t *watcher, const char *path, int recursive);
int remove_watch(file_watcher_t *watcher, const char *path);
void set_event_callback(file_watcher_t *watcher, event_callback_t callback);
void start_watching(file_watcher_t *watcher);
void stop_watching(file_watcher_t *watcher);
```

## Exercise 9: Memory Debugger (Expert)

### Task
Create a memory debugger that tracks memory allocations and detects leaks.

### Requirements
1. Track all memory allocations
2. Detect memory leaks
3. Detect double-free errors
4. Detect buffer overflows
5. Generate memory usage reports

### Code Template
```c
// memory_debugger.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <execinfo.h>

typedef struct allocation {
    void *ptr;
    size_t size;
    char *file;
    int line;
    time_t timestamp;
    struct allocation *next;
} allocation_t;

typedef struct memory_debugger {
    allocation_t *allocations;
    int num_allocations;
    size_t total_allocated;
    int enabled;
} memory_debugger_t;

memory_debugger_t* create_memory_debugger(void);
void* debug_malloc(size_t size, const char *file, int line);
void debug_free(void *ptr, const char *file, int line);
void print_memory_report(memory_debugger_t *debugger);
void detect_memory_leaks(memory_debugger_t *debugger);
```

## Exercise 10: Operating System Simulator (Expert)

### Task
Build a simple operating system simulator that demonstrates core OS concepts.

### Requirements
1. Process scheduler simulation
2. Memory management simulation
3. File system simulation
4. System call simulation
5. Interactive command interface

### Code Template
```c
// os_simulator.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

typedef struct process {
    int pid;
    char name[64];
    int priority;
    int state;
    int cpu_time;
    int memory_usage;
    struct process *next;
} process_t;

typedef struct scheduler {
    process_t *ready_queue;
    process_t *running;
    int time_quantum;
    int current_time;
} scheduler_t;

typedef struct memory_manager {
    void *memory;
    size_t total_size;
    size_t free_size;
    int *page_table;
    int num_pages;
} memory_manager_t;

typedef struct os_simulator {
    scheduler_t *scheduler;
    memory_manager_t *memory_manager;
    int running;
} os_simulator_t;

os_simulator_t* create_os_simulator(size_t memory_size);
void create_process(os_simulator_t *os, const char *name, int priority);
void schedule_processes(os_simulator_t *os);
void print_system_status(os_simulator_t *os);
void run_simulator(os_simulator_t *os);
```

## How to Approach These Exercises

1. **Start with the basics**: Complete exercises 1-3 to build a solid foundation
2. **Read the source code**: Study the provided code templates carefully
3. **Experiment**: Try different approaches and see what works best
4. **Test thoroughly**: Always test your implementations with various inputs
5. **Profile performance**: Use tools like gprof and valgrind to optimize your code
6. **Document your work**: Keep notes on what you learn from each exercise

## Additional Resources

- **Advanced Programming in the UNIX Environment** by W. Richard Stevens
- **The Linux Programming Interface** by Michael Kerrisk
- **Systems Programming** by J. H. Saltzer and M. Frans Kaashoek
- **Operating System Concepts** by Abraham Silberschatz
- **Linux Kernel Development** by Robert Love

## Success Criteria

By the end of these exercises, you should be able to:
- Write efficient system-level C programs
- Understand and implement memory management techniques
- Create robust file I/O operations
- Implement process management and communication
- Build network applications
- Debug and profile system programs
- Understand the internals of operating systems

## Testing Your Code

Always test your implementations with:
- **Valgrind**: For memory leak detection
- **AddressSanitizer**: For buffer overflow detection
- **ThreadSanitizer**: For race condition detection
- **gprof**: For performance profiling
- **strace**: For system call tracing

Remember: The goal is not just to complete the exercises, but to deeply understand the underlying concepts and be able to apply them in real-world scenarios!
