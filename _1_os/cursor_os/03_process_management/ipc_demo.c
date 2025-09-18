#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/wait.h>
#include <sys/ipc.h>
#include <sys/msg.h>
#include <sys/shm.h>
#include <sys/sem.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <fcntl.h>
#include <errno.h>

// Message queue structure
typedef struct {
    long mtype;
    char mtext[256];
} message_t;

// Shared memory structure
typedef struct {
    int counter;
    char message[256];
} shared_data_t;

// Semaphore operations
union semun {
    int val;
    struct semid_ds *buf;
    unsigned short *array;
};

// Function prototypes
void pipe_demo(void);
void fifo_demo(void);
void message_queue_demo(void);
void shared_memory_demo(void);
void semaphore_demo(void);
void socket_demo(void);

// Pipe demonstration
void pipe_demo() {
    printf("\n=== Pipe Demonstration ===\n");
    
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
        
        const char *message = "Hello from child via pipe!";
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
            printf("Received via pipe: %s\n", buffer);
        }
        
        close(pipefd[0]);
        wait(NULL);
    }
}

// FIFO demonstration
void fifo_demo() {
    printf("\n=== FIFO Demonstration ===\n");
    
    const char *fifo_path = "/tmp/myfifo";
    
    // Create FIFO
    if (mkfifo(fifo_path, 0666) == -1) {
        if (errno != EEXIST) {
            perror("mkfifo failed");
            return;
        }
    }
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        return;
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
            return;
        }
        
        char buffer[1024];
        ssize_t bytes_read = read(fd, buffer, sizeof(buffer) - 1);
        if (bytes_read > 0) {
            buffer[bytes_read] = '\0';
            printf("Received via FIFO: %s\n", buffer);
        }
        
        close(fd);
        wait(NULL);
        
        // Clean up
        unlink(fifo_path);
    }
}

// Message queue demonstration
void message_queue_demo() {
    printf("\n=== Message Queue Demonstration ===\n");
    
    key_t key = ftok("message_queue", 65);
    int msgid = msgget(key, 0666 | IPC_CREAT);
    
    if (msgid == -1) {
        perror("msgget failed");
        return;
    }
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        return;
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
            return;
        }
        
        printf("Received via message queue: %s\n", msg.mtext);
        
        wait(NULL);
        
        // Clean up
        msgctl(msgid, IPC_RMID, NULL);
    }
}

// Shared memory demonstration
void shared_memory_demo() {
    printf("\n=== Shared Memory Demonstration ===\n");
    
    key_t key = ftok("shared_memory", 65);
    int shmid = shmget(key, sizeof(shared_data_t), 0666 | IPC_CREAT);
    
    if (shmid == -1) {
        perror("shmget failed");
        return;
    }
    
    // Attach to shared memory
    shared_data_t *shared_data = (shared_data_t*)shmat(shmid, NULL, 0);
    if (shared_data == (shared_data_t*)-1) {
        perror("shmat failed");
        return;
    }
    
    // Initialize shared data
    shared_data->counter = 0;
    strcpy(shared_data->message, "Initial message");
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        shmdt(shared_data);
        shmctl(shmid, IPC_RMID, NULL);
        return;
    } else if (pid == 0) {
        // Child process - modify shared data
        shared_data->counter = 42;
        strcpy(shared_data->message, "Hello from child via shared memory!");
        shmdt(shared_data);
        exit(0);
    } else {
        // Parent process - read shared data
        wait(NULL);
        printf("Counter: %d\n", shared_data->counter);
        printf("Message: %s\n", shared_data->message);
        
        // Detach and remove shared memory
        shmdt(shared_data);
        shmctl(shmid, IPC_RMID, NULL);
    }
}

// Semaphore demonstration
void semaphore_demo() {
    printf("\n=== Semaphore Demonstration ===\n");
    
    key_t key = ftok("semaphore", 65);
    int semid = semget(key, 1, 0666 | IPC_CREAT);
    
    if (semid == -1) {
        perror("semget failed");
        return;
    }
    
    // Initialize semaphore to 1
    union semun sem_union;
    sem_union.val = 1;
    if (semctl(semid, 0, SETVAL, sem_union) == -1) {
        perror("semctl failed");
        return;
    }
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        return;
    } else if (pid == 0) {
        // Child process
        printf("Child: Waiting for semaphore...\n");
        
        struct sembuf sem_op;
        sem_op.sem_num = 0;
        sem_op.sem_op = -1;  // Wait (P operation)
        sem_op.sem_flg = 0;
        
        if (semop(semid, &sem_op, 1) == -1) {
            perror("semop failed");
            exit(1);
        }
        
        printf("Child: Got semaphore, in critical section\n");
        sleep(2);
        printf("Child: Leaving critical section\n");
        
        // Signal (V operation)
        sem_op.sem_op = 1;
        if (semop(semid, &sem_op, 1) == -1) {
            perror("semop failed");
            exit(1);
        }
        
        exit(0);
    } else {
        // Parent process
        printf("Parent: Waiting for semaphore...\n");
        
        struct sembuf sem_op;
        sem_op.sem_num = 0;
        sem_op.sem_op = -1;  // Wait (P operation)
        sem_op.sem_flg = 0;
        
        if (semop(semid, &sem_op, 1) == -1) {
            perror("semop failed");
            return;
        }
        
        printf("Parent: Got semaphore, in critical section\n");
        sleep(2);
        printf("Parent: Leaving critical section\n");
        
        // Signal (V operation)
        sem_op.sem_op = 1;
        if (semop(semid, &sem_op, 1) == -1) {
            perror("semop failed");
            return;
        }
        
        wait(NULL);
        
        // Clean up
        semctl(semid, 0, IPC_RMID);
    }
}

// Socket demonstration
void socket_demo() {
    printf("\n=== Socket Demonstration ===\n");
    
    int server_fd = socket(AF_UNIX, SOCK_STREAM, 0);
    if (server_fd == -1) {
        perror("socket failed");
        return;
    }
    
    struct sockaddr_un server_addr;
    server_addr.sun_family = AF_UNIX;
    strcpy(server_addr.sun_path, "/tmp/unix_socket");
    
    // Remove existing socket file
    unlink("/tmp/unix_socket");
    
    if (bind(server_fd, (struct sockaddr*)&server_addr, sizeof(server_addr)) == -1) {
        perror("bind failed");
        close(server_fd);
        return;
    }
    
    if (listen(server_fd, 5) == -1) {
        perror("listen failed");
        close(server_fd);
        return;
    }
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        close(server_fd);
        return;
    } else if (pid == 0) {
        // Child process - client
        sleep(1); // Wait for server to start
        
        int client_fd = socket(AF_UNIX, SOCK_STREAM, 0);
        if (client_fd == -1) {
            perror("socket failed");
            exit(1);
        }
        
        if (connect(client_fd, (struct sockaddr*)&server_addr, sizeof(server_addr)) == -1) {
            perror("connect failed");
            close(client_fd);
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
            close(server_fd);
            return;
        }
        
        char buffer[1024];
        ssize_t bytes_received = recv(client_fd, buffer, sizeof(buffer) - 1, 0);
        if (bytes_received > 0) {
            buffer[bytes_received] = '\0';
            printf("Received via socket: %s\n", buffer);
        }
        
        close(client_fd);
        close(server_fd);
        wait(NULL);
        
        // Clean up
        unlink("/tmp/unix_socket");
    }
}

// Interactive IPC demo
void interactive_ipc_demo() {
    int choice;
    
    printf("\n=== Interactive IPC Demonstration ===\n");
    printf("Choose IPC method to demonstrate:\n");
    printf("1. Pipes\n");
    printf("2. FIFOs\n");
    printf("3. Message Queues\n");
    printf("4. Shared Memory\n");
    printf("5. Semaphores\n");
    printf("6. Sockets\n");
    printf("7. All methods\n");
    printf("Choice: ");
    
    scanf("%d", &choice);
    
    switch (choice) {
        case 1:
            pipe_demo();
            break;
        case 2:
            fifo_demo();
            break;
        case 3:
            message_queue_demo();
            break;
        case 4:
            shared_memory_demo();
            break;
        case 5:
            semaphore_demo();
            break;
        case 6:
            socket_demo();
            break;
        case 7:
            pipe_demo();
            fifo_demo();
            message_queue_demo();
            shared_memory_demo();
            semaphore_demo();
            socket_demo();
            break;
        default:
            printf("Invalid choice\n");
            break;
    }
}

int main(int argc, char *argv[]) {
    if (argc > 1 && strcmp(argv[1], "interactive") == 0) {
        interactive_ipc_demo();
        return 0;
    }
    
    printf("=== IPC Demonstration Program ===\n");
    printf("This program demonstrates various IPC mechanisms:\n");
    printf("- Pipes\n");
    printf("- FIFOs\n");
    printf("- Message Queues\n");
    printf("- Shared Memory\n");
    printf("- Semaphores\n");
    printf("- Sockets\n");
    
    // Run all demonstrations
    pipe_demo();
    fifo_demo();
    message_queue_demo();
    shared_memory_demo();
    semaphore_demo();
    socket_demo();
    
    printf("\n=== All IPC demonstrations completed ===\n");
    
    return 0;
}
