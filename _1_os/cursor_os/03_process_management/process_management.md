# Process Management

## Table of Contents
1. [Introduction](#introduction)
2. [Process Concepts](#process-concepts)
3. [Process Creation and Termination](#process-creation-and-termination)
4. [Process Scheduling](#process-scheduling)
5. [Inter-Process Communication](#inter-process-communication)
6. [Process Synchronization](#process-synchronization)
7. [Process States and Transitions](#process-states-and-transitions)
8. [Practical Implementation](#practical-implementation)

## Introduction

Process management is the heart of operating system functionality. It involves creating, scheduling, and managing processes that execute user programs and system services. Understanding process management is essential for:

- **System Design**: Building efficient operating systems
- **Application Development**: Writing programs that work well with the OS
- **Performance Optimization**: Understanding how to make programs run faster
- **Debugging**: Troubleshooting system and application issues
- **Security**: Implementing process isolation and access control

### What is a Process?

A process is an instance of a running program. It consists of:
- **Program Code**: The executable instructions
- **Data**: Variables and data structures
- **Stack**: Function calls and local variables
- **Heap**: Dynamically allocated memory
- **Process Control Block (PCB)**: OS data structure containing process information

### Process vs Thread

- **Process**: Independent execution unit with its own memory space
- **Thread**: Lightweight execution unit within a process, sharing memory space

## Process Concepts

### Process Control Block (PCB)

The PCB is the data structure that contains all information about a process:

```c
typedef struct process_control_block {
    // Process identification
    pid_t pid;                    // Process ID
    pid_t ppid;                   // Parent Process ID
    uid_t uid;                    // User ID
    gid_t gid;                    // Group ID
    
    // Process state
    int state;                    // Current state
    int priority;                 // Process priority
    int nice_value;               // Nice value for scheduling
    
    // CPU state
    cpu_state_t cpu_state;       // CPU registers
    int cpu_time;                 // CPU time used
    int cpu_time_limit;           // CPU time limit
    
    // Memory management
    memory_map_t *memory_map;     // Memory layout
    size_t memory_usage;          // Current memory usage
    size_t memory_limit;          // Memory limit
    
    // File descriptors
    file_descriptor_t *fd_table;  // Open file descriptors
    int max_fds;                  // Maximum file descriptors
    
    // Process relationships
    struct process_control_block *parent;
    struct process_control_block *children;
    struct process_control_block *siblings;
    
    // Scheduling information
    int arrival_time;             // When process arrived
    int burst_time;               // CPU burst time
    int waiting_time;             // Time spent waiting
    int turnaround_time;          // Total time in system
    
    // Signal handling
    signal_handler_t *signal_handlers;
    sigset_t signal_mask;         // Blocked signals
    
    // Process limits
    rlimit_t *rlimits;            // Resource limits
    
    // Working directory
    char *cwd;                    // Current working directory
    
    // Environment variables
    char **environment;           // Environment variables
    
    // Command line arguments
    char **argv;                  // Command line arguments
    int argc;                     // Argument count
} pcb_t;
```

### Process States

Processes can be in one of several states:

1. **New**: Process is being created
2. **Ready**: Process is loaded in memory and waiting for CPU
3. **Running**: Process is currently executing on CPU
4. **Waiting/Blocked**: Process is waiting for an event (I/O, signal, etc.)
5. **Terminated**: Process has finished execution

```c
typedef enum {
    PROCESS_NEW,
    PROCESS_READY,
    PROCESS_RUNNING,
    PROCESS_WAITING,
    PROCESS_TERMINATED
} process_state_t;
```

### Process Hierarchy

Processes form a tree structure:
- **Parent Process**: Creates child processes
- **Child Process**: Created by parent process
- **Orphan Process**: Parent has terminated
- **Zombie Process**: Child has terminated but parent hasn't reaped it

## Process Creation and Termination

### Process Creation

#### 1. fork() System Call

```c
#include <unistd.h>
#include <sys/wait.h>
#include <stdio.h>
#include <stdlib.h>

void fork_example() {
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        exit(1);
    } else if (pid == 0) {
        // Child process
        printf("Child process: PID = %d, PPID = %d\n", 
               getpid(), getppid());
        exit(0);
    } else {
        // Parent process
        printf("Parent process: PID = %d, Child PID = %d\n", 
               getpid(), pid);
        
        // Wait for child to complete
        int status;
        waitpid(pid, &status, 0);
        printf("Child process completed with status: %d\n", status);
    }
}
```

#### 2. exec() Family of System Calls

```c
void exec_example() {
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        exit(1);
    } else if (pid == 0) {
        // Child process - execute ls command
        execl("/bin/ls", "ls", "-l", NULL);
        perror("execl failed");
        exit(1);
    } else {
        // Parent process
        int status;
        waitpid(pid, &status, 0);
        printf("Command completed\n");
    }
}
```

#### 3. Combined fork() and exec()

```c
void fork_exec_example() {
    char *args[] = {"ls", "-la", "/home", NULL};
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        exit(1);
    } else if (pid == 0) {
        // Child process
        execvp(args[0], args);
        perror("execvp failed");
        exit(1);
    } else {
        // Parent process
        int status;
        waitpid(pid, &status, 0);
        printf("Process completed\n");
    }
}
```

### Process Termination

#### 1. Normal Termination

```c
void normal_termination() {
    // Exit with status code
    exit(0);  // Success
    exit(1);  // Error
}
```

#### 2. Abnormal Termination

```c
void abnormal_termination() {
    // Abort - generates SIGABRT
    abort();
    
    // Kill current process
    kill(getpid(), SIGKILL);
}
```

#### 3. Process Cleanup

```c
void cleanup_example() {
    // Set up signal handlers
    signal(SIGINT, cleanup_handler);
    signal(SIGTERM, cleanup_handler);
    
    // Register cleanup function
    atexit(cleanup_function);
    
    // Main program logic
    while (1) {
        // Do work
        sleep(1);
    }
}

void cleanup_handler(int sig) {
    printf("Received signal %d, cleaning up...\n", sig);
    // Cleanup code
    exit(0);
}

void cleanup_function(void) {
    printf("Cleaning up resources...\n");
    // Cleanup code
}
```

## Process Scheduling

### Scheduling Algorithms

#### 1. First-Come, First-Served (FCFS)

```c
typedef struct process {
    int pid;
    int arrival_time;
    int burst_time;
    int completion_time;
    int turnaround_time;
    int waiting_time;
} process_t;

void fcfs_scheduling(process_t *processes, int n) {
    // Sort processes by arrival time
    qsort(processes, n, sizeof(process_t), compare_arrival_time);
    
    int current_time = 0;
    
    for (int i = 0; i < n; i++) {
        // Process arrives
        if (processes[i].arrival_time > current_time) {
            current_time = processes[i].arrival_time;
        }
        
        // Process starts
        processes[i].waiting_time = current_time - processes[i].arrival_time;
        
        // Process completes
        current_time += processes[i].burst_time;
        processes[i].completion_time = current_time;
        processes[i].turnaround_time = processes[i].completion_time - processes[i].arrival_time;
    }
}
```

#### 2. Shortest Job First (SJF)

```c
void sjf_scheduling(process_t *processes, int n) {
    int current_time = 0;
    int completed = 0;
    
    while (completed < n) {
        // Find shortest job that has arrived
        int shortest = -1;
        int shortest_time = INT_MAX;
        
        for (int i = 0; i < n; i++) {
            if (processes[i].arrival_time <= current_time && 
                processes[i].burst_time < shortest_time &&
                processes[i].completion_time == 0) {
                shortest = i;
                shortest_time = processes[i].burst_time;
            }
        }
        
        if (shortest == -1) {
            current_time++;
            continue;
        }
        
        // Execute shortest job
        processes[shortest].waiting_time = current_time - processes[shortest].arrival_time;
        current_time += processes[shortest].burst_time;
        processes[shortest].completion_time = current_time;
        processes[shortest].turnaround_time = processes[shortest].completion_time - processes[shortest].arrival_time;
        completed++;
    }
}
```

#### 3. Round Robin (RR)

```c
typedef struct rr_process {
    int pid;
    int arrival_time;
    int burst_time;
    int remaining_time;
    int completion_time;
    int turnaround_time;
    int waiting_time;
} rr_process_t;

void round_robin_scheduling(rr_process_t *processes, int n, int quantum) {
    int current_time = 0;
    int completed = 0;
    int *queue = malloc(n * sizeof(int));
    int queue_size = 0;
    int queue_head = 0;
    
    // Add all processes to queue
    for (int i = 0; i < n; i++) {
        queue[queue_size++] = i;
        processes[i].remaining_time = processes[i].burst_time;
    }
    
    while (completed < n) {
        if (queue_head < queue_size) {
            int current_process = queue[queue_head++];
            
            // Execute process for quantum or until completion
            int execution_time = (processes[current_process].remaining_time < quantum) ? 
                                 processes[current_process].remaining_time : quantum;
            
            processes[current_process].remaining_time -= execution_time;
            current_time += execution_time;
            
            if (processes[current_process].remaining_time == 0) {
                // Process completed
                processes[current_process].completion_time = current_time;
                processes[current_process].turnaround_time = current_time - processes[current_process].arrival_time;
                processes[current_process].waiting_time = processes[current_process].turnaround_time - processes[current_process].burst_time;
                completed++;
            } else {
                // Process not completed, add back to queue
                queue[queue_size++] = current_process;
            }
        } else {
            current_time++;
        }
    }
    
    free(queue);
}
```

#### 4. Priority Scheduling

```c
typedef struct priority_process {
    int pid;
    int arrival_time;
    int burst_time;
    int priority;
    int completion_time;
    int turnaround_time;
    int waiting_time;
} priority_process_t;

void priority_scheduling(priority_process_t *processes, int n) {
    int current_time = 0;
    int completed = 0;
    
    while (completed < n) {
        // Find highest priority process that has arrived
        int highest_priority = -1;
        int highest_priority_value = INT_MAX;
        
        for (int i = 0; i < n; i++) {
            if (processes[i].arrival_time <= current_time && 
                processes[i].priority < highest_priority_value &&
                processes[i].completion_time == 0) {
                highest_priority = i;
                highest_priority_value = processes[i].priority;
            }
        }
        
        if (highest_priority == -1) {
            current_time++;
            continue;
        }
        
        // Execute highest priority process
        processes[highest_priority].waiting_time = current_time - processes[highest_priority].arrival_time;
        current_time += processes[highest_priority].burst_time;
        processes[highest_priority].completion_time = current_time;
        processes[highest_priority].turnaround_time = processes[highest_priority].completion_time - processes[highest_priority].arrival_time;
        completed++;
    }
}
```

### Multi-Level Queue Scheduling

```c
typedef struct mlq_process {
    int pid;
    int arrival_time;
    int burst_time;
    int priority;
    int queue_level;
    int completion_time;
    int turnaround_time;
    int waiting_time;
} mlq_process_t;

void multi_level_queue_scheduling(mlq_process_t *processes, int n, int num_queues) {
    // Create queues for each priority level
    int **queues = malloc(num_queues * sizeof(int*));
    int *queue_sizes = calloc(num_queues, sizeof(int));
    int *queue_heads = calloc(num_queues, sizeof(int));
    
    for (int i = 0; i < num_queues; i++) {
        queues[i] = malloc(n * sizeof(int));
    }
    
    // Add processes to appropriate queues
    for (int i = 0; i < n; i++) {
        int queue_level = processes[i].queue_level;
        queues[queue_level][queue_sizes[queue_level]++] = i;
    }
    
    int current_time = 0;
    int completed = 0;
    
    while (completed < n) {
        // Check queues in priority order
        for (int queue_level = 0; queue_level < num_queues; queue_level++) {
            if (queue_heads[queue_level] < queue_sizes[queue_level]) {
                int process_index = queues[queue_level][queue_heads[queue_level]++];
                
                // Execute process
                processes[process_index].waiting_time = current_time - processes[process_index].arrival_time;
                current_time += processes[process_index].burst_time;
                processes[process_index].completion_time = current_time;
                processes[process_index].turnaround_time = processes[process_index].completion_time - processes[process_index].arrival_time;
                completed++;
                break;
            }
        }
        
        if (completed == 0) {
            current_time++;
        }
    }
    
    // Clean up
    for (int i = 0; i < num_queues; i++) {
        free(queues[i]);
    }
    free(queues);
    free(queue_sizes);
    free(queue_heads);
}
```

## Inter-Process Communication

### 1. Pipes

#### Anonymous Pipes

```c
void pipe_example() {
    int pipefd[2];
    
    if (pipe(pipefd) == -1) {
        perror("pipe failed");
        exit(1);
    }
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        exit(1);
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

#### Named Pipes (FIFOs)

```c
void fifo_example() {
    const char *fifo_path = "/tmp/myfifo";
    
    // Create FIFO
    if (mkfifo(fifo_path, 0666) == -1) {
        perror("mkfifo failed");
        exit(1);
    }
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        exit(1);
    } else if (pid == 0) {
        // Child process - write to FIFO
        int fd = open(fifo_path, O_WRONLY);
        if (fd == -1) {
            perror("open failed");
            exit(1);
        }
        
        const char *message = "Hello from child via FIFO!";
        write(fd, message, strlen(message));
        close(fd);
        exit(0);
    } else {
        // Parent process - read from FIFO
        int fd = open(fifo_path, O_RDONLY);
        if (fd == -1) {
            perror("open failed");
            exit(1);
        }
        
        char buffer[1024];
        ssize_t bytes_read = read(fd, buffer, sizeof(buffer) - 1);
        if (bytes_read > 0) {
            buffer[bytes_read] = '\0';
            printf("Received: %s\n", buffer);
        }
        
        close(fd);
        wait(NULL);
        
        // Clean up
        unlink(fifo_path);
    }
}
```

### 2. Message Queues

```c
#include <sys/ipc.h>
#include <sys/msg.h>

typedef struct {
    long mtype;
    char mtext[256];
} message_t;

void message_queue_example() {
    key_t key = ftok("message_queue", 65);
    int msgid = msgget(key, 0666 | IPC_CREAT);
    
    if (msgid == -1) {
        perror("msgget failed");
        exit(1);
    }
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        exit(1);
    } else if (pid == 0) {
        // Child process - send message
        message_t msg;
        msg.mtype = 1;
        strcpy(msg.mtext, "Hello from child via message queue!");
        
        if (msgsnd(msgid, &msg, strlen(msg.mtext), 0) == -1) {
            perror("msgsnd failed");
            exit(1);
        }
        
        exit(0);
    } else {
        // Parent process - receive message
        message_t msg;
        
        if (msgrcv(msgid, &msg, sizeof(msg.mtext), 1, 0) == -1) {
            perror("msgrcv failed");
            exit(1);
        }
        
        printf("Received: %s\n", msg.mtext);
        
        wait(NULL);
        
        // Clean up
        msgctl(msgid, IPC_RMID, NULL);
    }
}
```

### 3. Shared Memory

```c
#include <sys/ipc.h>
#include <sys/shm.h>

void shared_memory_example() {
    key_t key = ftok("shared_memory", 65);
    int shmid = shmget(key, 1024, 0666 | IPC_CREAT);
    
    if (shmid == -1) {
        perror("shmget failed");
        exit(1);
    }
    
    // Attach to shared memory
    char *shared_memory = (char*)shmat(shmid, NULL, 0);
    if (shared_memory == (char*)-1) {
        perror("shmat failed");
        exit(1);
    }
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        exit(1);
    } else if (pid == 0) {
        // Child process - write to shared memory
        strcpy(shared_memory, "Hello from child via shared memory!");
        shmdt(shared_memory);
        exit(0);
    } else {
        // Parent process - read from shared memory
        wait(NULL);
        printf("Received: %s\n", shared_memory);
        
        // Detach and remove shared memory
        shmdt(shared_memory);
        shmctl(shmid, IPC_RMID, NULL);
    }
}
```

### 4. Sockets

#### Unix Domain Sockets

```c
#include <sys/socket.h>
#include <sys/un.h>

void unix_socket_example() {
    int server_fd = socket(AF_UNIX, SOCK_STREAM, 0);
    if (server_fd == -1) {
        perror("socket failed");
        exit(1);
    }
    
    struct sockaddr_un server_addr;
    server_addr.sun_family = AF_UNIX;
    strcpy(server_addr.sun_path, "/tmp/unix_socket");
    
    // Remove existing socket file
    unlink("/tmp/unix_socket");
    
    if (bind(server_fd, (struct sockaddr*)&server_addr, sizeof(server_addr)) == -1) {
        perror("bind failed");
        exit(1);
    }
    
    if (listen(server_fd, 5) == -1) {
        perror("listen failed");
        exit(1);
    }
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        exit(1);
    } else if (pid == 0) {
        // Child process - client
        int client_fd = socket(AF_UNIX, SOCK_STREAM, 0);
        if (client_fd == -1) {
            perror("socket failed");
            exit(1);
        }
        
        if (connect(client_fd, (struct sockaddr*)&server_addr, sizeof(server_addr)) == -1) {
            perror("connect failed");
            exit(1);
        }
        
        const char *message = "Hello from child via Unix socket!";
        send(client_fd, message, strlen(message), 0);
        close(client_fd);
        exit(0);
    } else {
        // Parent process - server
        struct sockaddr_un client_addr;
        socklen_t client_len = sizeof(client_addr);
        
        int client_fd = accept(server_fd, (struct sockaddr*)&client_addr, &client_len);
        if (client_fd == -1) {
            perror("accept failed");
            exit(1);
        }
        
        char buffer[1024];
        ssize_t bytes_received = recv(client_fd, buffer, sizeof(buffer) - 1, 0);
        if (bytes_received > 0) {
            buffer[bytes_received] = '\0';
            printf("Received: %s\n", buffer);
        }
        
        close(client_fd);
        close(server_fd);
        wait(NULL);
        
        // Clean up
        unlink("/tmp/unix_socket");
    }
}
```

## Process Synchronization

### 1. Semaphores

```c
#include <sys/ipc.h>
#include <sys/sem.h>

void semaphore_example() {
    key_t key = ftok("semaphore", 65);
    int semid = semget(key, 1, 0666 | IPC_CREAT);
    
    if (semid == -1) {
        perror("semget failed");
        exit(1);
    }
    
    // Initialize semaphore to 1
    union semun {
        int val;
        struct semid_ds *buf;
        unsigned short *array;
    } sem_union;
    
    sem_union.val = 1;
    if (semctl(semid, 0, SETVAL, sem_union) == -1) {
        perror("semctl failed");
        exit(1);
    }
    
    // Semaphore operations
    struct sembuf sem_op;
    sem_op.sem_num = 0;
    sem_op.sem_op = -1;  // Wait (P operation)
    sem_op.sem_flg = 0;
    
    if (semop(semid, &sem_op, 1) == -1) {
        perror("semop failed");
        exit(1);
    }
    
    // Critical section
    printf("In critical section\n");
    sleep(2);
    
    // Signal (V operation)
    sem_op.sem_op = 1;
    if (semop(semid, &sem_op, 1) == -1) {
        perror("semop failed");
        exit(1);
    }
    
    // Clean up
    semctl(semid, 0, IPC_RMID);
}
```

### 2. Mutexes

```c
#include <pthread.h>

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
int shared_variable = 0;

void* thread_function(void* arg) {
    for (int i = 0; i < 1000; i++) {
        pthread_mutex_lock(&mutex);
        shared_variable++;
        pthread_mutex_unlock(&mutex);
    }
    return NULL;
}

void mutex_example() {
    pthread_t threads[5];
    
    // Create threads
    for (int i = 0; i < 5; i++) {
        pthread_create(&threads[i], NULL, thread_function, NULL);
    }
    
    // Wait for threads to complete
    for (int i = 0; i < 5; i++) {
        pthread_join(threads[i], NULL);
    }
    
    printf("Final value: %d\n", shared_variable);
}
```

### 3. Condition Variables

```c
#include <pthread.h>

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t condition = PTHREAD_COND_INITIALIZER;
int data_ready = 0;
int data = 0;

void* producer(void* arg) {
    for (int i = 0; i < 10; i++) {
        pthread_mutex_lock(&mutex);
        
        data = i;
        data_ready = 1;
        
        pthread_cond_signal(&condition);
        pthread_mutex_unlock(&mutex);
        
        sleep(1);
    }
    return NULL;
}

void* consumer(void* arg) {
    for (int i = 0; i < 10; i++) {
        pthread_mutex_lock(&mutex);
        
        while (!data_ready) {
            pthread_cond_wait(&condition, &mutex);
        }
        
        printf("Consumed: %d\n", data);
        data_ready = 0;
        
        pthread_mutex_unlock(&mutex);
    }
    return NULL;
}

void condition_variable_example() {
    pthread_t producer_thread, consumer_thread;
    
    pthread_create(&producer_thread, NULL, producer, NULL);
    pthread_create(&consumer_thread, NULL, consumer, NULL);
    
    pthread_join(producer_thread, NULL);
    pthread_join(consumer_thread, NULL);
}
```

## Process States and Transitions

### State Transition Diagram

```
    [New] ──create──> [Ready] ──schedule──> [Running]
                                    │
                                    │
                                    v
    [Terminated] <──exit── [Running] ──wait──> [Waiting]
                                    │
                                    │
                                    v
                              [Waiting] ──wake──> [Ready]
```

### State Transition Implementation

```c
typedef enum {
    PROCESS_NEW,
    PROCESS_READY,
    PROCESS_RUNNING,
    PROCESS_WAITING,
    PROCESS_TERMINATED
} process_state_t;

typedef struct process {
    int pid;
    process_state_t state;
    int priority;
    int cpu_time;
    int waiting_time;
    struct process *next;
} process_t;

typedef struct scheduler {
    process_t *ready_queue;
    process_t *waiting_queue;
    process_t *running;
    int time_quantum;
} scheduler_t;

void state_transition(scheduler_t *scheduler, process_t *process, process_state_t new_state) {
    process->state = new_state;
    
    switch (new_state) {
        case PROCESS_READY:
            // Add to ready queue
            add_to_ready_queue(scheduler, process);
            break;
            
        case PROCESS_RUNNING:
            // Set as running process
            scheduler->running = process;
            break;
            
        case PROCESS_WAITING:
            // Add to waiting queue
            add_to_waiting_queue(scheduler, process);
            break;
            
        case PROCESS_TERMINATED:
            // Remove from all queues
            remove_from_queues(scheduler, process);
            break;
            
        default:
            break;
    }
}
```

## Practical Implementation

### Project 1: Process Scheduler Simulator

```c
// process_scheduler.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

typedef struct process {
    int pid;
    int arrival_time;
    int burst_time;
    int priority;
    int completion_time;
    int turnaround_time;
    int waiting_time;
    int response_time;
    int remaining_time;
} process_t;

typedef struct scheduler {
    process_t *processes;
    int num_processes;
    int time_quantum;
    int current_time;
} scheduler_t;

void fcfs_scheduling(scheduler_t *scheduler) {
    // Implementation of FCFS scheduling
    printf("=== FCFS Scheduling ===\n");
    
    int current_time = 0;
    for (int i = 0; i < scheduler->num_processes; i++) {
        process_t *p = &scheduler->processes[i];
        
        if (p->arrival_time > current_time) {
            current_time = p->arrival_time;
        }
        
        p->waiting_time = current_time - p->arrival_time;
        p->response_time = p->waiting_time;
        
        current_time += p->burst_time;
        p->completion_time = current_time;
        p->turnaround_time = p->completion_time - p->arrival_time;
    }
}

void sjf_scheduling(scheduler_t *scheduler) {
    // Implementation of SJF scheduling
    printf("=== SJF Scheduling ===\n");
    
    int current_time = 0;
    int completed = 0;
    
    while (completed < scheduler->num_processes) {
        int shortest = -1;
        int shortest_time = INT_MAX;
        
        for (int i = 0; i < scheduler->num_processes; i++) {
            process_t *p = &scheduler->processes[i];
            if (p->arrival_time <= current_time && 
                p->burst_time < shortest_time &&
                p->completion_time == 0) {
                shortest = i;
                shortest_time = p->burst_time;
            }
        }
        
        if (shortest == -1) {
            current_time++;
            continue;
        }
        
        process_t *p = &scheduler->processes[shortest];
        p->waiting_time = current_time - p->arrival_time;
        p->response_time = p->waiting_time;
        
        current_time += p->burst_time;
        p->completion_time = current_time;
        p->turnaround_time = p->completion_time - p->arrival_time;
        completed++;
    }
}

void round_robin_scheduling(scheduler_t *scheduler) {
    // Implementation of Round Robin scheduling
    printf("=== Round Robin Scheduling ===\n");
    
    int current_time = 0;
    int completed = 0;
    
    // Initialize remaining time
    for (int i = 0; i < scheduler->num_processes; i++) {
        scheduler->processes[i].remaining_time = scheduler->processes[i].burst_time;
    }
    
    while (completed < scheduler->num_processes) {
        int executed = 0;
        
        for (int i = 0; i < scheduler->num_processes; i++) {
            process_t *p = &scheduler->processes[i];
            
            if (p->arrival_time <= current_time && p->remaining_time > 0) {
                if (p->response_time == -1) {
                    p->response_time = current_time - p->arrival_time;
                }
                
                int execution_time = (p->remaining_time < scheduler->time_quantum) ? 
                                   p->remaining_time : scheduler->time_quantum;
                
                p->remaining_time -= execution_time;
                current_time += execution_time;
                executed = 1;
                
                if (p->remaining_time == 0) {
                    p->completion_time = current_time;
                    p->turnaround_time = p->completion_time - p->arrival_time;
                    p->waiting_time = p->turnaround_time - p->burst_time;
                    completed++;
                }
            }
        }
        
        if (!executed) {
            current_time++;
        }
    }
}

void print_results(scheduler_t *scheduler) {
    printf("\nProcess Results:\n");
    printf("PID\tAT\tBT\tCT\tTAT\tWT\tRT\n");
    printf("----------------------------------------\n");
    
    int total_turnaround = 0;
    int total_waiting = 0;
    int total_response = 0;
    
    for (int i = 0; i < scheduler->num_processes; i++) {
        process_t *p = &scheduler->processes[i];
        printf("%d\t%d\t%d\t%d\t%d\t%d\t%d\n",
               p->pid, p->arrival_time, p->burst_time, p->completion_time,
               p->turnaround_time, p->waiting_time, p->response_time);
        
        total_turnaround += p->turnaround_time;
        total_waiting += p->waiting_time;
        total_response += p->response_time;
    }
    
    printf("\nAverage Turnaround Time: %.2f\n", (double)total_turnaround / scheduler->num_processes);
    printf("Average Waiting Time: %.2f\n", (double)total_waiting / scheduler->num_processes);
    printf("Average Response Time: %.2f\n", (double)total_response / scheduler->num_processes);
}

int main() {
    // Create scheduler
    scheduler_t scheduler;
    scheduler.num_processes = 5;
    scheduler.time_quantum = 2;
    scheduler.processes = malloc(scheduler.num_processes * sizeof(process_t));
    
    // Initialize processes
    int arrival_times[] = {0, 1, 2, 3, 4};
    int burst_times[] = {8, 4, 2, 6, 3};
    int priorities[] = {3, 1, 4, 2, 5};
    
    for (int i = 0; i < scheduler.num_processes; i++) {
        scheduler.processes[i].pid = i + 1;
        scheduler.processes[i].arrival_time = arrival_times[i];
        scheduler.processes[i].burst_time = burst_times[i];
        scheduler.processes[i].priority = priorities[i];
        scheduler.processes[i].completion_time = 0;
        scheduler.processes[i].turnaround_time = 0;
        scheduler.processes[i].waiting_time = 0;
        scheduler.processes[i].response_time = -1;
        scheduler.processes[i].remaining_time = 0;
    }
    
    // Test different scheduling algorithms
    fcfs_scheduling(&scheduler);
    print_results(&scheduler);
    
    // Reset for next test
    for (int i = 0; i < scheduler.num_processes; i++) {
        scheduler.processes[i].completion_time = 0;
        scheduler.processes[i].turnaround_time = 0;
        scheduler.processes[i].waiting_time = 0;
        scheduler.processes[i].response_time = -1;
        scheduler.processes[i].remaining_time = 0;
    }
    
    sjf_scheduling(&scheduler);
    print_results(&scheduler);
    
    // Reset for next test
    for (int i = 0; i < scheduler.num_processes; i++) {
        scheduler.processes[i].completion_time = 0;
        scheduler.processes[i].turnaround_time = 0;
        scheduler.processes[i].waiting_time = 0;
        scheduler.processes[i].response_time = -1;
        scheduler.processes[i].remaining_time = 0;
    }
    
    round_robin_scheduling(&scheduler);
    print_results(&scheduler);
    
    free(scheduler.processes);
    return 0;
}
```

### Project 2: Process Manager

```c
// process_manager.c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/wait.h>
#include <signal.h>
#include <time.h>

typedef struct process_info {
    pid_t pid;
    char name[256];
    int status;
    time_t start_time;
    time_t end_time;
    int exit_code;
    struct process_info *next;
} process_info_t;

typedef struct process_manager {
    process_info_t *processes;
    int num_processes;
    int max_processes;
} process_manager_t;

process_manager_t* create_process_manager(int max_processes) {
    process_manager_t *pm = malloc(sizeof(process_manager_t));
    if (!pm) return NULL;
    
    pm->processes = NULL;
    pm->num_processes = 0;
    pm->max_processes = max_processes;
    
    return pm;
}

void add_process(process_manager_t *pm, pid_t pid, const char *name) {
    if (pm->num_processes >= pm->max_processes) {
        printf("Maximum number of processes reached\n");
        return;
    }
    
    process_info_t *new_process = malloc(sizeof(process_info_t));
    if (!new_process) return;
    
    new_process->pid = pid;
    strncpy(new_process->name, name, sizeof(new_process->name) - 1);
    new_process->name[sizeof(new_process->name) - 1] = '\0';
    new_process->status = 1; // Running
    new_process->start_time = time(NULL);
    new_process->end_time = 0;
    new_process->exit_code = 0;
    new_process->next = pm->processes;
    
    pm->processes = new_process;
    pm->num_processes++;
    
    printf("Added process %d: %s\n", pid, name);
}

void remove_process(process_manager_t *pm, pid_t pid) {
    process_info_t *current = pm->processes;
    process_info_t *prev = NULL;
    
    while (current) {
        if (current->pid == pid) {
            if (prev) {
                prev->next = current->next;
            } else {
                pm->processes = current->next;
            }
            
            free(current);
            pm->num_processes--;
            printf("Removed process %d\n", pid);
            return;
        }
        
        prev = current;
        current = current->next;
    }
    
    printf("Process %d not found\n", pid);
}

void list_processes(process_manager_t *pm) {
    printf("\n=== Process List ===\n");
    printf("PID\tName\t\tStatus\t\tStart Time\n");
    printf("----------------------------------------\n");
    
    process_info_t *current = pm->processes;
    while (current) {
        char start_time_str[32];
        struct tm *tm_info = localtime(&current->start_time);
        strftime(start_time_str, sizeof(start_time_str), "%H:%M:%S", tm_info);
        
        const char *status_str = current->status ? "Running" : "Terminated";
        
        printf("%d\t%-20s\t%s\t\t%s\n",
               current->pid, current->name, status_str, start_time_str);
        
        current = current->next;
    }
    
    printf("\nTotal processes: %d\n", pm->num_processes);
}

void cleanup_processes(process_manager_t *pm) {
    process_info_t *current = pm->processes;
    while (current) {
        process_info_t *next = current->next;
        free(current);
        current = next;
    }
    free(pm);
}

int main() {
    process_manager_t *pm = create_process_manager(10);
    if (!pm) {
        printf("Failed to create process manager\n");
        return 1;
    }
    
    // Start some processes
    pid_t pid1 = fork();
    if (pid1 == 0) {
        execl("/bin/sleep", "sleep", "5", NULL);
        exit(1);
    } else if (pid1 > 0) {
        add_process(pm, pid1, "sleep 5");
    }
    
    pid_t pid2 = fork();
    if (pid2 == 0) {
        execl("/bin/ls", "ls", "-la", NULL);
        exit(1);
    } else if (pid2 > 0) {
        add_process(pm, pid2, "ls -la");
    }
    
    // List processes
    list_processes(pm);
    
    // Wait for processes to complete
    int status;
    while (waitpid(-1, &status, WNOHANG) > 0) {
        // Process completed
    }
    
    // Update process status
    process_info_t *current = pm->processes;
    while (current) {
        if (current->status == 1) {
            if (kill(current->pid, 0) == -1) {
                current->status = 0; // Terminated
                current->end_time = time(NULL);
            }
        }
        current = current->next;
    }
    
    // List processes again
    list_processes(pm);
    
    // Clean up
    cleanup_processes(pm);
    
    return 0;
}
```

## Key Takeaways

1. **Process Lifecycle**: Understand how processes are created, scheduled, and terminated
2. **Scheduling Algorithms**: Different algorithms have different performance characteristics
3. **IPC Mechanisms**: Choose the right communication method for your needs
4. **Synchronization**: Prevent race conditions and ensure data consistency
5. **Resource Management**: Efficiently manage system resources

## Next Steps

In the next module, we'll dive into **Memory Management** where you'll learn:
- Virtual memory concepts
- Paging and segmentation
- Memory allocation algorithms
- Cache management
- Building a memory manager

## Exercises

1. **Process Scheduler**: Implement different scheduling algorithms
2. **Process Manager**: Build a process management system
3. **IPC Library**: Create a comprehensive IPC library
4. **Process Monitor**: Build a real-time process monitor
5. **Process Simulator**: Simulate process execution

Remember: The goal is to understand not just how processes work, but how to design efficient process management systems!
