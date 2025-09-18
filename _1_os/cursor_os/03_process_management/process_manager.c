#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/wait.h>
#include <signal.h>
#include <time.h>
#include <errno.h>

typedef struct process_info {
    pid_t pid;
    char name[256];
    int status;
    time_t start_time;
    time_t end_time;
    int exit_code;
    int priority;
    struct process_info *next;
} process_info_t;

typedef struct process_manager {
    process_info_t *processes;
    int num_processes;
    int max_processes;
    int running_processes;
} process_manager_t;

// Global process manager
process_manager_t *global_pm = NULL;

// Signal handler for SIGCHLD
void sigchld_handler(int sig) {
    int status;
    pid_t pid;
    
    // Reap all zombie children
    while ((pid = waitpid(-1, &status, WNOHANG)) > 0) {
        printf("Process %d terminated\n", pid);
        
        // Update process info
        process_info_t *current = global_pm->processes;
        while (current) {
            if (current->pid == pid) {
                current->end_time = time(NULL);
                current->status = 0; // Terminated
                global_pm->running_processes--;
                
                if (WIFEXITED(status)) {
                    current->exit_code = WEXITSTATUS(status);
                    printf("Process %d exited with code %d\n", pid, current->exit_code);
                } else if (WIFSIGNALED(status)) {
                    current->exit_code = WTERMSIG(status);
                    printf("Process %d killed by signal %d\n", pid, current->exit_code);
                }
                break;
            }
            current = current->next;
        }
    }
}

// Create process manager
process_manager_t* create_process_manager(int max_processes) {
    process_manager_t *pm = malloc(sizeof(process_manager_t));
    if (!pm) return NULL;
    
    pm->processes = NULL;
    pm->num_processes = 0;
    pm->max_processes = max_processes;
    pm->running_processes = 0;
    
    return pm;
}

// Destroy process manager
void destroy_process_manager(process_manager_t *pm) {
    if (pm) {
        process_info_t *current = pm->processes;
        while (current) {
            process_info_t *next = current->next;
            free(current);
            current = next;
        }
        free(pm);
    }
}

// Add process to manager
int add_process(process_manager_t *pm, pid_t pid, const char *name, int priority) {
    if (pm->num_processes >= pm->max_processes) {
        printf("Maximum number of processes reached\n");
        return -1;
    }
    
    process_info_t *new_process = malloc(sizeof(process_info_t));
    if (!new_process) return -1;
    
    new_process->pid = pid;
    strncpy(new_process->name, name, sizeof(new_process->name) - 1);
    new_process->name[sizeof(new_process->name) - 1] = '\0';
    new_process->status = 1; // Running
    new_process->start_time = time(NULL);
    new_process->end_time = 0;
    new_process->exit_code = 0;
    new_process->priority = priority;
    new_process->next = pm->processes;
    
    pm->processes = new_process;
    pm->num_processes++;
    pm->running_processes++;
    
    printf("Added process %d: %s (priority: %d)\n", pid, name, priority);
    return 0;
}

// Remove process from manager
int remove_process(process_manager_t *pm, pid_t pid) {
    process_info_t *current = pm->processes;
    process_info_t *prev = NULL;
    
    while (current) {
        if (current->pid == pid) {
            if (prev) {
                prev->next = current->next;
            } else {
                pm->processes = current->next;
            }
            
            pm->num_processes--;
            if (current->status == 1) {
                pm->running_processes--;
            }
            
            free(current);
            printf("Removed process %d\n", pid);
            return 0;
        }
        
        prev = current;
        current = current->next;
    }
    
    printf("Process %d not found\n", pid);
    return -1;
}

// Start a new process
pid_t start_process(process_manager_t *pm, const char *name, char *const argv[], int priority) {
    if (pm->num_processes >= pm->max_processes) {
        printf("Maximum number of processes reached\n");
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
        add_process(pm, pid, name, priority);
        return pid;
    }
}

// Kill a process
int kill_process(process_manager_t *pm, pid_t pid, int signal) {
    process_info_t *current = pm->processes;
    
    while (current) {
        if (current->pid == pid && current->status == 1) {
            if (kill(pid, signal) == -1) {
                perror("kill failed");
                return -1;
            }
            printf("Sent signal %d to process %d\n", signal, pid);
            return 0;
        }
        current = current->next;
    }
    
    printf("Process %d not found or not running\n", pid);
    return -1;
}

// List all processes
void list_processes(process_manager_t *pm) {
    printf("\n=== Process List ===\n");
    printf("PID\tName\t\tPriority\tStatus\t\tStart Time\t\tEnd Time\n");
    printf("------------------------------------------------------------------------\n");
    
    process_info_t *current = pm->processes;
    while (current) {
        char start_time_str[32];
        char end_time_str[32];
        struct tm *tm_info;
        
        tm_info = localtime(&current->start_time);
        strftime(start_time_str, sizeof(start_time_str), "%H:%M:%S", tm_info);
        
        if (current->end_time > 0) {
            tm_info = localtime(&current->end_time);
            strftime(end_time_str, sizeof(end_time_str), "%H:%M:%S", tm_info);
        } else {
            strcpy(end_time_str, "Running");
        }
        
        const char *status_str = current->status ? "Running" : "Terminated";
        
        printf("%d\t%-20s\t%d\t\t%s\t\t%s\t\t%s\n",
               current->pid, current->name, current->priority, status_str, 
               start_time_str, end_time_str);
        
        current = current->next;
    }
    
    printf("\nTotal processes: %d, Running: %d\n", 
           pm->num_processes, pm->running_processes);
}

// Get process information
process_info_t* get_process_info(process_manager_t *pm, pid_t pid) {
    process_info_t *current = pm->processes;
    
    while (current) {
        if (current->pid == pid) {
            return current;
        }
        current = current->next;
    }
    
    return NULL;
}

// Update process status
void update_process_status(process_manager_t *pm) {
    process_info_t *current = pm->processes;
    
    while (current) {
        if (current->status == 1) {
            // Check if process is still running
            if (kill(current->pid, 0) == -1) {
                if (errno == ESRCH) {
                    // Process no longer exists
                    current->status = 0;
                    current->end_time = time(NULL);
                    pm->running_processes--;
                    printf("Process %d terminated\n", current->pid);
                }
            }
        }
        current = current->next;
    }
}

// Wait for all processes to complete
void wait_for_all_processes(process_manager_t *pm) {
    printf("Waiting for all processes to complete...\n");
    
    while (pm->running_processes > 0) {
        sleep(1);
        update_process_status(pm);
    }
    
    printf("All processes completed\n");
}

// Interactive process manager
void interactive_manager(process_manager_t *pm) {
    char command[256];
    char *args[32];
    int num_args;
    
    printf("\n=== Interactive Process Manager ===\n");
    printf("Commands:\n");
    printf("  start <name> <command> [args...] - Start a new process\n");
    printf("  kill <pid> [signal]              - Kill a process (default: SIGTERM)\n");
    printf("  list                             - List all processes\n");
    printf("  info <pid>                       - Show process information\n");
    printf("  wait                             - Wait for all processes\n");
    printf("  update                           - Update process status\n");
    printf("  quit                             - Exit manager\n");
    printf("\n");
    
    while (1) {
        printf("manager> ");
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
            
            start_process(pm, args[1], cmd_args, 0);
            
        } else if (strcmp(args[0], "kill") == 0) {
            if (num_args < 2) {
                printf("Usage: kill <pid> [signal]\n");
                continue;
            }
            
            pid_t pid = atoi(args[1]);
            int signal = SIGTERM;
            
            if (num_args > 2) {
                signal = atoi(args[2]);
            }
            
            kill_process(pm, pid, signal);
            
        } else if (strcmp(args[0], "list") == 0) {
            list_processes(pm);
            
        } else if (strcmp(args[0], "info") == 0) {
            if (num_args != 2) {
                printf("Usage: info <pid>\n");
                continue;
            }
            
            pid_t pid = atoi(args[1]);
            process_info_t *info = get_process_info(pm, pid);
            
            if (info) {
                printf("\nProcess Information:\n");
                printf("PID: %d\n", info->pid);
                printf("Name: %s\n", info->name);
                printf("Priority: %d\n", info->priority);
                printf("Status: %s\n", info->status ? "Running" : "Terminated");
                printf("Start Time: %s", ctime(&info->start_time));
                if (info->end_time > 0) {
                    printf("End Time: %s", ctime(&info->end_time));
                }
                printf("Exit Code: %d\n", info->exit_code);
            } else {
                printf("Process %d not found\n", pid);
            }
            
        } else if (strcmp(args[0], "wait") == 0) {
            wait_for_all_processes(pm);
            
        } else if (strcmp(args[0], "update") == 0) {
            update_process_status(pm);
            printf("Process status updated\n");
            
        } else if (strcmp(args[0], "quit") == 0) {
            break;
            
        } else {
            printf("Unknown command: %s\n", args[0]);
        }
    }
}

// Demo function
void demo_process_manager() {
    printf("=== Process Manager Demo ===\n");
    
    // Create process manager
    process_manager_t *pm = create_process_manager(10);
    if (!pm) {
        printf("Failed to create process manager\n");
        return;
    }
    
    // Set up signal handler
    signal(SIGCHLD, sigchld_handler);
    global_pm = pm;
    
    // Start some demo processes
    char *ls_args[] = {"ls", "-la", NULL};
    char *ps_args[] = {"ps", "aux", NULL};
    char *sleep_args[] = {"sleep", "5", NULL};
    
    start_process(pm, "ls", ls_args, 0);
    start_process(pm, "ps", ps_args, 1);
    start_process(pm, "sleep", sleep_args, 2);
    
    // Wait a bit
    sleep(2);
    
    // List processes
    list_processes(pm);
    
    // Wait for all processes
    wait_for_all_processes(pm);
    
    // Final list
    list_processes(pm);
    
    // Clean up
    destroy_process_manager(pm);
}

int main(int argc, char *argv[]) {
    if (argc > 1 && strcmp(argv[1], "demo") == 0) {
        demo_process_manager();
        return 0;
    }
    
    // Create process manager
    process_manager_t *pm = create_process_manager(50);
    if (!pm) {
        printf("Failed to create process manager\n");
        return 1;
    }
    
    // Set up signal handler
    signal(SIGCHLD, sigchld_handler);
    global_pm = pm;
    
    // Run interactive manager
    interactive_manager(pm);
    
    // Clean up
    destroy_process_manager(pm);
    
    return 0;
}
