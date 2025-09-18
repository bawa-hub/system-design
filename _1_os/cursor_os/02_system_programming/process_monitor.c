#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <signal.h>
#include <time.h>
#include <string.h>
#include <errno.h>

// Process information structure
typedef struct {
    pid_t pid;
    char name[256];
    int status;
    time_t start_time;
    time_t end_time;
    int exit_code;
} process_info_t;

// Process monitor structure
typedef struct {
    process_info_t *processes;
    int max_processes;
    int num_processes;
    int running_processes;
} process_monitor_t;

// Global process monitor
process_monitor_t *monitor = NULL;

// Signal handler for SIGCHLD
void sigchld_handler(int sig) {
    int status;
    pid_t pid;
    
    // Reap all zombie children
    while ((pid = waitpid(-1, &status, WNOHANG)) > 0) {
        printf("Process %d terminated\n", pid);
        
        // Update process info
        for (int i = 0; i < monitor->num_processes; i++) {
            if (monitor->processes[i].pid == pid) {
                monitor->processes[i].end_time = time(NULL);
                monitor->processes[i].status = 0; // Terminated
                monitor->running_processes--;
                
                if (WIFEXITED(status)) {
                    monitor->processes[i].exit_code = WEXITSTATUS(status);
                    printf("Process %d exited with code %d\n", pid, monitor->processes[i].exit_code);
                } else if (WIFSIGNALED(status)) {
                    monitor->processes[i].exit_code = WTERMSIG(status);
                    printf("Process %d killed by signal %d\n", pid, monitor->processes[i].exit_code);
                }
                break;
            }
        }
    }
}

// Create process monitor
process_monitor_t* create_process_monitor(int max_processes) {
    process_monitor_t *pm = malloc(sizeof(process_monitor_t));
    if (!pm) return NULL;
    
    pm->processes = calloc(max_processes, sizeof(process_info_t));
    if (!pm->processes) {
        free(pm);
        return NULL;
    }
    
    pm->max_processes = max_processes;
    pm->num_processes = 0;
    pm->running_processes = 0;
    
    return pm;
}

// Destroy process monitor
void destroy_process_monitor(process_monitor_t *pm) {
    if (pm) {
        free(pm->processes);
        free(pm);
    }
}

// Start a new process
pid_t start_process(process_monitor_t *pm, const char *name, char *const argv[]) {
    if (pm->num_processes >= pm->max_processes) {
        fprintf(stderr, "Maximum number of processes reached\n");
        return -1;
    }
    
    pid_t pid = fork();
    
    if (pid == -1) {
        perror("fork failed");
        return -1;
    } else if (pid == 0) {
        // Child process
        execvp(argv[0], argv);
        perror("execvp failed");
        exit(1);
    } else {
        // Parent process
        process_info_t *proc = &pm->processes[pm->num_processes];
        proc->pid = pid;
        strncpy(proc->name, name, sizeof(proc->name) - 1);
        proc->name[sizeof(proc->name) - 1] = '\0';
        proc->status = 1; // Running
        proc->start_time = time(NULL);
        proc->end_time = 0;
        proc->exit_code = 0;
        
        pm->num_processes++;
        pm->running_processes++;
        
        printf("Started process %d: %s\n", pid, name);
        return pid;
    }
}

// Kill a process
int kill_process(process_monitor_t *pm, pid_t pid) {
    for (int i = 0; i < pm->num_processes; i++) {
        if (pm->processes[i].pid == pid && pm->processes[i].status == 1) {
            if (kill(pid, SIGTERM) == -1) {
                perror("kill failed");
                return -1;
            }
            printf("Sent SIGTERM to process %d\n", pid);
            return 0;
        }
    }
    
    fprintf(stderr, "Process %d not found or not running\n", pid);
    return -1;
}

// Force kill a process
int force_kill_process(process_monitor_t *pm, pid_t pid) {
    for (int i = 0; i < pm->num_processes; i++) {
        if (pm->processes[i].pid == pid && pm->processes[i].status == 1) {
            if (kill(pid, SIGKILL) == -1) {
                perror("kill failed");
                return -1;
            }
            printf("Sent SIGKILL to process %d\n", pid);
            return 0;
        }
    }
    
    fprintf(stderr, "Process %d not found or not running\n", pid);
    return -1;
}

// List all processes
void list_processes(process_monitor_t *pm) {
    printf("\n=== Process List ===\n");
    printf("%-8s %-20s %-10s %-15s %-15s %-10s\n", 
           "PID", "Name", "Status", "Start Time", "End Time", "Exit Code");
    printf("------------------------------------------------------------------------\n");
    
    for (int i = 0; i < pm->num_processes; i++) {
        process_info_t *proc = &pm->processes[i];
        
        char start_time_str[32];
        char end_time_str[32];
        struct tm *tm_info;
        
        tm_info = localtime(&proc->start_time);
        strftime(start_time_str, sizeof(start_time_str), "%H:%M:%S", tm_info);
        
        if (proc->end_time > 0) {
            tm_info = localtime(&proc->end_time);
            strftime(end_time_str, sizeof(end_time_str), "%H:%M:%S", tm_info);
        } else {
            strcpy(end_time_str, "Running");
        }
        
        const char *status_str = proc->status ? "Running" : "Terminated";
        
        printf("%-8d %-20s %-10s %-15s %-15s %-10d\n",
               proc->pid, proc->name, status_str, start_time_str, 
               end_time_str, proc->exit_code);
    }
    
    printf("\nTotal processes: %d, Running: %d\n", 
           pm->num_processes, pm->running_processes);
}

// Wait for all processes to complete
void wait_for_all_processes(process_monitor_t *pm) {
    printf("Waiting for all processes to complete...\n");
    
    while (pm->running_processes > 0) {
        sleep(1);
    }
    
    printf("All processes completed\n");
}

// Interactive process monitor
void interactive_monitor(process_monitor_t *pm) {
    char command[256];
    char *args[32];
    int num_args;
    
    printf("\n=== Interactive Process Monitor ===\n");
    printf("Commands:\n");
    printf("  start <name> <command> [args...] - Start a new process\n");
    printf("  kill <pid>                       - Terminate a process\n");
    printf("  force <pid>                      - Force kill a process\n");
    printf("  list                             - List all processes\n");
    printf("  wait                             - Wait for all processes\n");
    printf("  quit                             - Exit monitor\n");
    printf("\n");
    
    while (1) {
        printf("monitor> ");
        fflush(stdout);
        
        if (fgets(command, sizeof(command), stdin) == NULL) {
            break;
        }
        
        // Remove newline
        command[strcspn(command, "\n")] = '\0';
        
        // Parse command
        num_args = 0;
        char *token = strtok(command, " ");
        while (token && num_args < 32) {
            args[num_args++] = token;
            token = strtok(NULL, " ");
        }
        
        if (num_args == 0) continue;
        
        if (strcmp(args[0], "start") == 0) {
            if (num_args < 3) {
                printf("Usage: start <name> <command> [args...]\n");
                continue;
            }
            
            char *cmd_args[32];
            for (int i = 2; i < num_args; i++) {
                cmd_args[i - 2] = args[i];
            }
            cmd_args[num_args - 2] = NULL;
            
            start_process(pm, args[1], cmd_args);
            
        } else if (strcmp(args[0], "kill") == 0) {
            if (num_args != 2) {
                printf("Usage: kill <pid>\n");
                continue;
            }
            
            pid_t pid = atoi(args[1]);
            kill_process(pm, pid);
            
        } else if (strcmp(args[0], "force") == 0) {
            if (num_args != 2) {
                printf("Usage: force <pid>\n");
                continue;
            }
            
            pid_t pid = atoi(args[1]);
            force_kill_process(pm, pid);
            
        } else if (strcmp(args[0], "list") == 0) {
            list_processes(pm);
            
        } else if (strcmp(args[0], "wait") == 0) {
            wait_for_all_processes(pm);
            
        } else if (strcmp(args[0], "quit") == 0) {
            break;
            
        } else {
            printf("Unknown command: %s\n", args[0]);
        }
    }
}

// Demo function
void demo_process_monitor() {
    printf("=== Process Monitor Demo ===\n");
    
    // Create process monitor
    monitor = create_process_monitor(10);
    if (!monitor) {
        fprintf(stderr, "Failed to create process monitor\n");
        return;
    }
    
    // Set up signal handler
    signal(SIGCHLD, sigchld_handler);
    
    // Start some demo processes
    char *ls_args[] = {"ls", "-la", NULL};
    char *ps_args[] = {"ps", "aux", NULL};
    char *sleep_args[] = {"sleep", "5", NULL};
    
    start_process(monitor, "ls", ls_args);
    start_process(monitor, "ps", ps_args);
    start_process(monitor, "sleep", sleep_args);
    
    // Wait a bit
    sleep(2);
    
    // List processes
    list_processes(monitor);
    
    // Wait for all processes
    wait_for_all_processes(monitor);
    
    // Final list
    list_processes(monitor);
    
    // Clean up
    destroy_process_monitor(monitor);
}

int main(int argc, char *argv[]) {
    if (argc > 1 && strcmp(argv[1], "demo") == 0) {
        demo_process_monitor();
        return 0;
    }
    
    // Create process monitor
    monitor = create_process_monitor(50);
    if (!monitor) {
        fprintf(stderr, "Failed to create process monitor\n");
        return 1;
    }
    
    // Set up signal handler
    signal(SIGCHLD, sigchld_handler);
    
    // Run interactive monitor
    interactive_monitor(monitor);
    
    // Clean up
    destroy_process_monitor(monitor);
    
    return 0;
}
