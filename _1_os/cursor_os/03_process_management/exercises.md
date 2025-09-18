# Process Management - Exercises

## Exercise 1: Advanced Process Scheduler (Beginner)

### Task
Extend the process scheduler to support more advanced scheduling algorithms and features.

### Requirements
1. Implement Multilevel Queue Scheduling
2. Add aging to prevent starvation
3. Implement preemptive priority scheduling
4. Add context switching overhead simulation
5. Create a Gantt chart visualization

### Code Template
```c
// advanced_scheduler.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct process {
    int pid;
    int arrival_time;
    int burst_time;
    int priority;
    int queue_level;
    int aging_value;
    int completion_time;
    int turnaround_time;
    int waiting_time;
    int response_time;
} process_t;

typedef struct queue {
    process_t *processes;
    int size;
    int time_quantum;
    int priority;
} queue_t;

typedef struct mlq_scheduler {
    queue_t *queues;
    int num_queues;
    int context_switch_overhead;
} mlq_scheduler_t;

void mlq_scheduling(mlq_scheduler_t *scheduler, process_t *processes, int num_processes);
void aging_algorithm(mlq_scheduler_t *scheduler, process_t *processes, int num_processes);
void print_gantt_chart(process_t *processes, int num_processes);
```

### Expected Output
```
=== Multilevel Queue Scheduling ===
Queue 0 (RR, q=2): P1 P2 P1 P3
Queue 1 (RR, q=4): P4 P5
Queue 2 (FCFS): P6

Gantt Chart:
Time: 0  2  4  6  8  10 12 14 16
Proc: P1 P2 P1 P3 P4 P5  P4 P5  P6
```

## Exercise 2: Process Tree Manager (Intermediate)

### Task
Create a process tree manager that can create, monitor, and control hierarchical process structures.

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
    int priority;
    time_t start_time;
} process_node_t;

typedef struct process_tree {
    process_node_t *root;
    int num_processes;
    int max_processes;
} process_tree_t;

process_tree_t* create_process_tree(void);
process_node_t* add_process(process_tree_t *tree, pid_t pid, const char *name, pid_t parent_pid);
void remove_process(process_tree_t *tree, pid_t pid);
void print_process_tree(process_tree_t *tree);
void kill_process_tree(process_tree_t *tree, int signal);
void monitor_process_tree(process_tree_t *tree);
```

## Exercise 3: Advanced IPC Library (Intermediate)

### Task
Create a comprehensive IPC library that provides a unified interface for different communication methods.

### Requirements
1. Support for pipes, FIFOs, message queues, shared memory, and sockets
2. Automatic cleanup and error handling
3. Message serialization/deserialization
4. Connection pooling for sockets
5. Performance benchmarking

### Code Template
```c
// ipc_library.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/ipc.h>
#include <sys/msg.h>
#include <sys/shm.h>
#include <sys/sem.h>
#include <sys/socket.h>

typedef enum {
    IPC_PIPE,
    IPC_FIFO,
    IPC_MSG_QUEUE,
    IPC_SHARED_MEM,
    IPC_SOCKET
} ipc_type_t;

typedef struct ipc_message {
    int type;
    int size;
    void *data;
} ipc_message_t;

typedef struct ipc_handle {
    ipc_type_t type;
    int id;
    void *data;
    size_t size;
    int connected;
} ipc_handle_t;

typedef struct ipc_library {
    ipc_handle_t *handles;
    int num_handles;
    int max_handles;
} ipc_library_t;

ipc_library_t* create_ipc_library(int max_handles);
ipc_handle_t* create_ipc(ipc_library_t *lib, ipc_type_t type, const char *name, size_t size);
int send_message(ipc_handle_t *handle, ipc_message_t *message);
int receive_message(ipc_handle_t *handle, ipc_message_t *message);
void destroy_ipc_library(ipc_library_t *lib);
```

## Exercise 4: Process Monitor and Profiler (Advanced)

### Task
Build a comprehensive process monitoring and profiling system.

### Requirements
1. Real-time process monitoring
2. CPU and memory usage tracking
3. System call tracing
4. Performance profiling
5. Historical data storage

### Code Template
```c
// process_monitor.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/sysinfo.h>
#include <sys/stat.h>
#include <string.h>
#include <time.h>

typedef struct process_stats {
    pid_t pid;
    char name[256];
    int cpu_usage;
    int memory_usage;
    int num_threads;
    int num_fds;
    time_t start_time;
    time_t last_update;
} process_stats_t;

typedef struct system_stats {
    int total_processes;
    int running_processes;
    int sleeping_processes;
    int zombie_processes;
    int total_memory;
    int free_memory;
    int cpu_usage;
} system_stats_t;

typedef struct process_monitor {
    process_stats_t *processes;
    int num_processes;
    int max_processes;
    system_stats_t system_stats;
    int monitoring;
} process_monitor_t;

process_monitor_t* create_process_monitor(int max_processes);
void start_monitoring(process_monitor_t *monitor);
void stop_monitoring(process_monitor_t *monitor);
void print_process_stats(process_monitor_t *monitor);
void print_system_stats(process_monitor_t *monitor);
void export_data(process_monitor_t *monitor, const char *filename);
```

## Exercise 5: Process Synchronization Framework (Advanced)

### Task
Implement a comprehensive process synchronization framework with various synchronization primitives.

### Requirements
1. Mutexes, semaphores, condition variables, and barriers
2. Deadlock detection and prevention
3. Priority inheritance for mutexes
4. Reader-writer locks
5. Performance analysis

### Code Template
```c
// sync_framework.c
#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <semaphore.h>

typedef struct mutex {
    pthread_mutex_t mutex;
    int owner;
    int priority;
    int locked;
} mutex_t;

typedef struct semaphore {
    sem_t sem;
    int value;
    int max_value;
} semaphore_t;

typedef struct condition_variable {
    pthread_cond_t cond;
    pthread_mutex_t mutex;
    int waiting;
} condition_variable_t;

typedef struct barrier {
    pthread_mutex_t mutex;
    pthread_cond_t cond;
    int count;
    int threshold;
} barrier_t;

typedef struct reader_writer_lock {
    pthread_mutex_t mutex;
    pthread_cond_t read_cond;
    pthread_cond_t write_cond;
    int readers;
    int writers;
    int waiting_writers;
} rw_lock_t;

typedef struct sync_framework {
    mutex_t *mutexes;
    semaphore_t *semaphores;
    condition_variable_t *conditions;
    barrier_t *barriers;
    rw_lock_t *rw_locks;
    int num_mutexes;
    int num_semaphores;
    int num_conditions;
    int num_barriers;
    int num_rw_locks;
} sync_framework_t;

sync_framework_t* create_sync_framework(void);
mutex_t* create_mutex(sync_framework_t *framework);
semaphore_t* create_semaphore(sync_framework_t *framework, int initial_value);
condition_variable_t* create_condition_variable(sync_framework_t *framework);
barrier_t* create_barrier(sync_framework_t *framework, int threshold);
rw_lock_t* create_rw_lock(sync_framework_t *framework);
void destroy_sync_framework(sync_framework_t *framework);
```

## Exercise 6: Process Communication Protocol (Expert)

### Task
Design and implement a custom process communication protocol.

### Requirements
1. Message framing and error detection
2. Flow control and congestion management
3. Reliable message delivery
4. Message ordering and sequencing
5. Protocol state machine

### Code Template
```c
// process_protocol.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>

typedef enum {
    MSG_DATA,
    MSG_ACK,
    MSG_NACK,
    MSG_SYN,
    MSG_FIN,
    MSG_RST
} message_type_t;

typedef struct protocol_header {
    uint32_t sequence;
    uint32_t ack_sequence;
    uint16_t length;
    uint8_t type;
    uint8_t flags;
    uint16_t checksum;
} protocol_header_t;

typedef struct protocol_message {
    protocol_header_t header;
    void *data;
} protocol_message_t;

typedef struct protocol_state {
    int socket_fd;
    uint32_t next_sequence;
    uint32_t expected_sequence;
    int connected;
    int state;
} protocol_state_t;

typedef struct process_protocol {
    protocol_state_t *states;
    int num_states;
    int max_states;
} process_protocol_t;

process_protocol_t* create_process_protocol(int max_states);
int send_message(process_protocol_t *protocol, int state_id, void *data, size_t size);
int receive_message(process_protocol_t *protocol, int state_id, void *data, size_t size);
void destroy_process_protocol(process_protocol_t *protocol);
```

## Exercise 7: Process Migration System (Expert)

### Task
Implement a process migration system that can move processes between different machines.

### Requirements
1. Process state capture and restoration
2. Memory page migration
3. File descriptor migration
4. Network connection migration
5. Checkpoint and restart functionality

### Code Template
```c
// process_migration.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <sys/mman.h>

typedef struct process_checkpoint {
    pid_t pid;
    char name[256];
    void *memory_image;
    size_t memory_size;
    int *file_descriptors;
    int num_fds;
    int *network_connections;
    int num_connections;
    time_t checkpoint_time;
} process_checkpoint_t;

typedef struct migration_manager {
    process_checkpoint_t *checkpoints;
    int num_checkpoints;
    int max_checkpoints;
    char *source_host;
    char *dest_host;
} migration_manager_t;

migration_manager_t* create_migration_manager(const char *source_host, const char *dest_host);
int create_checkpoint(migration_manager_t *manager, pid_t pid);
int migrate_process(migration_manager_t *manager, int checkpoint_id, const char *dest_host);
int restore_process(migration_manager_t *manager, int checkpoint_id);
void destroy_migration_manager(migration_manager_t *manager);
```

## Exercise 8: Process Security Manager (Expert)

### Task
Create a process security manager that implements access control and security policies.

### Requirements
1. Process isolation and sandboxing
2. Capability-based security
3. Resource limits and quotas
4. Audit logging
5. Security policy enforcement

### Code Template
```c
// process_security.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/capability.h>
#include <sys/resource.h>

typedef struct security_policy {
    char name[256];
    int allow_system_calls;
    int allow_network_access;
    int allow_file_access;
    int max_memory;
    int max_cpu_time;
    int max_file_descriptors;
} security_policy_t;

typedef struct process_security {
    pid_t pid;
    security_policy_t *policy;
    int capabilities;
    int resource_limits;
    time_t start_time;
    int violations;
} process_security_t;

typedef struct security_manager {
    process_security_t *processes;
    int num_processes;
    int max_processes;
    security_policy_t *policies;
    int num_policies;
} security_manager_t;

security_manager_t* create_security_manager(int max_processes);
int add_security_policy(security_manager_t *manager, security_policy_t *policy);
int apply_security_policy(security_manager_t *manager, pid_t pid, const char *policy_name);
int check_security_violation(security_manager_t *manager, pid_t pid, int violation_type);
void audit_security_event(security_manager_t *manager, pid_t pid, const char *event);
void destroy_security_manager(security_manager_t *manager);
```

## Exercise 9: Process Load Balancer (Expert)

### Task
Implement a process load balancer that distributes processes across multiple machines.

### Requirements
1. Load balancing algorithms (round-robin, least connections, weighted)
2. Health checking and failover
3. Process affinity and migration
4. Performance monitoring
5. Dynamic scaling

### Code Template
```c
// process_load_balancer.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>

typedef struct server_node {
    char hostname[256];
    int port;
    int weight;
    int active_connections;
    int cpu_usage;
    int memory_usage;
    int status; // 0: down, 1: up
} server_node_t;

typedef struct load_balancer {
    server_node_t *servers;
    int num_servers;
    int max_servers;
    int algorithm; // 0: round-robin, 1: least-connections, 2: weighted
    int current_server;
    int health_check_interval;
} load_balancer_t;

typedef struct process_request {
    char command[256];
    char *args[32];
    int num_args;
    int priority;
} process_request_t;

load_balancer_t* create_load_balancer(int max_servers);
int add_server(load_balancer_t *lb, const char *hostname, int port, int weight);
int select_server(load_balancer_t *lb, process_request_t *request);
int distribute_process(load_balancer_t *lb, process_request_t *request);
void health_check(load_balancer_t *lb);
void destroy_load_balancer(load_balancer_t *lb);
```

## Exercise 10: Process Debugging Framework (Expert)

### Task
Create a comprehensive process debugging framework with advanced debugging features.

### Requirements
1. Process tracing and breakpoint support
2. Memory inspection and modification
3. Variable watching and modification
4. Call stack analysis
5. Performance profiling integration

### Code Template
```c
// process_debugger.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/ptrace.h>
#include <sys/wait.h>
#include <sys/user.h>

typedef struct breakpoint {
    void *address;
    int enabled;
    int original_byte;
} breakpoint_t;

typedef struct watchpoint {
    void *address;
    size_t size;
    int type; // 0: read, 1: write, 2: both
    int enabled;
} watchpoint_t;

typedef struct process_debugger {
    pid_t target_pid;
    int attached;
    breakpoint_t *breakpoints;
    int num_breakpoints;
    watchpoint_t *watchpoints;
    int num_watchpoints;
    struct user_regs_struct regs;
} process_debugger_t;

process_debugger_t* create_process_debugger(pid_t target_pid);
int attach_to_process(process_debugger_t *debugger);
int detach_from_process(process_debugger_t *debugger);
int set_breakpoint(process_debugger_t *debugger, void *address);
int remove_breakpoint(process_debugger_t *debugger, void *address);
int set_watchpoint(process_debugger_t *debugger, void *address, size_t size, int type);
int continue_execution(process_debugger_t *debugger);
int step_instruction(process_debugger_t *debugger);
void print_registers(process_debugger_t *debugger);
void print_memory(process_debugger_t *debugger, void *address, size_t size);
void destroy_process_debugger(process_debugger_t *debugger);
```

## How to Approach These Exercises

1. **Start with the basics**: Complete exercises 1-3 to build a solid foundation
2. **Read the source code**: Study the provided code templates carefully
3. **Experiment**: Try different approaches and see what works best
4. **Test thoroughly**: Always test your implementations with various scenarios
5. **Profile performance**: Use tools like gprof and valgrind to optimize your code
6. **Document your work**: Keep notes on what you learn from each exercise

## Additional Resources

- **Advanced Programming in the UNIX Environment** by W. Richard Stevens
- **The Linux Programming Interface** by Michael Kerrisk
- **Operating System Concepts** by Abraham Silberschatz
- **Linux Kernel Development** by Robert Love
- **Process Management in Operating Systems** by Andrew Tanenbaum

## Success Criteria

By the end of these exercises, you should be able to:
- Implement advanced process scheduling algorithms
- Create robust process management systems
- Build comprehensive IPC solutions
- Design and implement process synchronization mechanisms
- Understand and implement process security models
- Debug and profile complex process interactions

## Testing Your Code

Always test your implementations with:
- **Valgrind**: For memory leak detection
- **AddressSanitizer**: For buffer overflow detection
- **ThreadSanitizer**: For race condition detection
- **gprof**: For performance profiling
- **strace**: For system call tracing

Remember: The goal is not just to complete the exercises, but to deeply understand the underlying concepts and be able to apply them in real-world scenarios!
