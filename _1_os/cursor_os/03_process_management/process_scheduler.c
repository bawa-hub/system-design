#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <limits.h>

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
    int start_time;
} process_t;

typedef struct scheduler {
    process_t *processes;
    int num_processes;
    int time_quantum;
    int current_time;
} scheduler_t;

// Function to compare processes by arrival time
int compare_arrival_time(const void *a, const void *b) {
    process_t *p1 = (process_t*)a;
    process_t *p2 = (process_t*)b;
    return p1->arrival_time - p2->arrival_time;
}

// Function to compare processes by burst time
int compare_burst_time(const void *a, const void *b) {
    process_t *p1 = (process_t*)a;
    process_t *p2 = (process_t*)b;
    return p1->burst_time - p2->burst_time;
}

// Function to compare processes by priority
int compare_priority(const void *a, const void *b) {
    process_t *p1 = (process_t*)a;
    process_t *p2 = (process_t*)b;
    return p1->priority - p2->priority;
}

// Reset process data for new scheduling algorithm
void reset_processes(process_t *processes, int num_processes) {
    for (int i = 0; i < num_processes; i++) {
        processes[i].completion_time = 0;
        processes[i].turnaround_time = 0;
        processes[i].waiting_time = 0;
        processes[i].response_time = -1;
        processes[i].remaining_time = processes[i].burst_time;
        processes[i].start_time = -1;
    }
}

// First-Come, First-Served Scheduling
void fcfs_scheduling(scheduler_t *scheduler) {
    printf("=== First-Come, First-Served (FCFS) Scheduling ===\n");
    
    // Sort processes by arrival time
    qsort(scheduler->processes, scheduler->num_processes, sizeof(process_t), compare_arrival_time);
    
    int current_time = 0;
    
    for (int i = 0; i < scheduler->num_processes; i++) {
        process_t *p = &scheduler->processes[i];
        
        // Wait for process to arrive
        if (p->arrival_time > current_time) {
            current_time = p->arrival_time;
        }
        
        // Process starts
        p->start_time = current_time;
        p->waiting_time = current_time - p->arrival_time;
        p->response_time = p->waiting_time;
        
        // Process completes
        current_time += p->burst_time;
        p->completion_time = current_time;
        p->turnaround_time = p->completion_time - p->arrival_time;
    }
}

// Shortest Job First Scheduling
void sjf_scheduling(scheduler_t *scheduler) {
    printf("=== Shortest Job First (SJF) Scheduling ===\n");
    
    int current_time = 0;
    int completed = 0;
    
    while (completed < scheduler->num_processes) {
        // Find shortest job that has arrived
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
        
        // Execute shortest job
        process_t *p = &scheduler->processes[shortest];
        p->start_time = current_time;
        p->waiting_time = current_time - p->arrival_time;
        p->response_time = p->waiting_time;
        
        current_time += p->burst_time;
        p->completion_time = current_time;
        p->turnaround_time = p->completion_time - p->arrival_time;
        completed++;
    }
}

// Round Robin Scheduling
void round_robin_scheduling(scheduler_t *scheduler) {
    printf("=== Round Robin (RR) Scheduling ===\n");
    
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
                // First time this process gets CPU
                if (p->start_time == -1) {
                    p->start_time = current_time;
                    p->response_time = current_time - p->arrival_time;
                }
                
                // Execute for quantum or until completion
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

// Priority Scheduling
void priority_scheduling(scheduler_t *scheduler) {
    printf("=== Priority Scheduling ===\n");
    
    int current_time = 0;
    int completed = 0;
    
    while (completed < scheduler->num_processes) {
        // Find highest priority process that has arrived
        int highest_priority = -1;
        int highest_priority_value = INT_MAX;
        
        for (int i = 0; i < scheduler->num_processes; i++) {
            process_t *p = &scheduler->processes[i];
            if (p->arrival_time <= current_time && 
                p->priority < highest_priority_value &&
                p->completion_time == 0) {
                highest_priority = i;
                highest_priority_value = p->priority;
            }
        }
        
        if (highest_priority == -1) {
            current_time++;
            continue;
        }
        
        // Execute highest priority process
        process_t *p = &scheduler->processes[highest_priority];
        p->start_time = current_time;
        p->waiting_time = current_time - p->arrival_time;
        p->response_time = p->waiting_time;
        
        current_time += p->burst_time;
        p->completion_time = current_time;
        p->turnaround_time = p->completion_time - p->arrival_time;
        completed++;
    }
}

// Shortest Remaining Time First Scheduling
void srtf_scheduling(scheduler_t *scheduler) {
    printf("=== Shortest Remaining Time First (SRTF) Scheduling ===\n");
    
    int current_time = 0;
    int completed = 0;
    
    // Initialize remaining time
    for (int i = 0; i < scheduler->num_processes; i++) {
        scheduler->processes[i].remaining_time = scheduler->processes[i].burst_time;
    }
    
    while (completed < scheduler->num_processes) {
        // Find process with shortest remaining time that has arrived
        int shortest = -1;
        int shortest_time = INT_MAX;
        
        for (int i = 0; i < scheduler->num_processes; i++) {
            process_t *p = &scheduler->processes[i];
            if (p->arrival_time <= current_time && 
                p->remaining_time < shortest_time &&
                p->remaining_time > 0) {
                shortest = i;
                shortest_time = p->remaining_time;
            }
        }
        
        if (shortest == -1) {
            current_time++;
            continue;
        }
        
        // Execute process for 1 time unit
        process_t *p = &scheduler->processes[shortest];
        
        if (p->start_time == -1) {
            p->start_time = current_time;
            p->response_time = current_time - p->arrival_time;
        }
        
        p->remaining_time--;
        current_time++;
        
        if (p->remaining_time == 0) {
            p->completion_time = current_time;
            p->turnaround_time = p->completion_time - p->arrival_time;
            p->waiting_time = p->turnaround_time - p->burst_time;
            completed++;
        }
    }
}

// Print scheduling results
void print_results(scheduler_t *scheduler) {
    printf("\nProcess Results:\n");
    printf("PID\tAT\tBT\tP\tCT\tTAT\tWT\tRT\n");
    printf("--------------------------------------------------------\n");
    
    int total_turnaround = 0;
    int total_waiting = 0;
    int total_response = 0;
    
    for (int i = 0; i < scheduler->num_processes; i++) {
        process_t *p = &scheduler->processes[i];
        printf("%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\n",
               p->pid, p->arrival_time, p->burst_time, p->priority,
               p->completion_time, p->turnaround_time, p->waiting_time, p->response_time);
        
        total_turnaround += p->turnaround_time;
        total_waiting += p->waiting_time;
        total_response += p->response_time;
    }
    
    printf("\nAverage Turnaround Time: %.2f\n", (double)total_turnaround / scheduler->num_processes);
    printf("Average Waiting Time: %.2f\n", (double)total_waiting / scheduler->num_processes);
    printf("Average Response Time: %.2f\n", (double)total_response / scheduler->num_processes);
}

// Generate random process data
void generate_random_processes(process_t *processes, int num_processes) {
    srand(time(NULL));
    
    for (int i = 0; i < num_processes; i++) {
        processes[i].pid = i + 1;
        processes[i].arrival_time = rand() % 10;
        processes[i].burst_time = (rand() % 10) + 1;
        processes[i].priority = (rand() % 5) + 1;
        processes[i].completion_time = 0;
        processes[i].turnaround_time = 0;
        processes[i].waiting_time = 0;
        processes[i].response_time = -1;
        processes[i].remaining_time = processes[i].burst_time;
        processes[i].start_time = -1;
    }
}

// Interactive process input
void input_processes(process_t *processes, int num_processes) {
    printf("Enter process details:\n");
    printf("Format: PID ArrivalTime BurstTime Priority\n");
    
    for (int i = 0; i < num_processes; i++) {
        printf("Process %d: ", i + 1);
        scanf("%d %d %d %d", 
              &processes[i].pid, 
              &processes[i].arrival_time, 
              &processes[i].burst_time, 
              &processes[i].priority);
        
        processes[i].completion_time = 0;
        processes[i].turnaround_time = 0;
        processes[i].waiting_time = 0;
        processes[i].response_time = -1;
        processes[i].remaining_time = processes[i].burst_time;
        processes[i].start_time = -1;
    }
}

// Compare scheduling algorithms
void compare_algorithms(scheduler_t *scheduler) {
    printf("\n=== Algorithm Comparison ===\n");
    
    // Test FCFS
    reset_processes(scheduler->processes, scheduler->num_processes);
    fcfs_scheduling(scheduler);
    print_results(scheduler);
    
    // Test SJF
    reset_processes(scheduler->processes, scheduler->num_processes);
    sjf_scheduling(scheduler);
    print_results(scheduler);
    
    // Test Round Robin
    reset_processes(scheduler->processes, scheduler->num_processes);
    round_robin_scheduling(scheduler);
    print_results(scheduler);
    
    // Test Priority
    reset_processes(scheduler->processes, scheduler->num_processes);
    priority_scheduling(scheduler);
    print_results(scheduler);
    
    // Test SRTF
    reset_processes(scheduler->processes, scheduler->num_processes);
    srtf_scheduling(scheduler);
    print_results(scheduler);
}

int main() {
    int num_processes;
    int choice;
    
    printf("=== Process Scheduler Simulator ===\n");
    printf("Enter number of processes: ");
    scanf("%d", &num_processes);
    
    if (num_processes <= 0 || num_processes > 20) {
        printf("Invalid number of processes. Using default: 5\n");
        num_processes = 5;
    }
    
    // Create scheduler
    scheduler_t scheduler;
    scheduler.num_processes = num_processes;
    scheduler.time_quantum = 2;
    scheduler.processes = malloc(num_processes * sizeof(process_t));
    
    if (!scheduler.processes) {
        printf("Memory allocation failed\n");
        return 1;
    }
    
    printf("\nChoose input method:\n");
    printf("1. Manual input\n");
    printf("2. Random generation\n");
    printf("3. Use default data\n");
    printf("Choice: ");
    scanf("%d", &choice);
    
    switch (choice) {
        case 1:
            input_processes(scheduler.processes, num_processes);
            break;
        case 2:
            generate_random_processes(scheduler.processes, num_processes);
            break;
        case 3:
        default:
            // Use default data
            int arrival_times[] = {0, 1, 2, 3, 4};
            int burst_times[] = {8, 4, 2, 6, 3};
            int priorities[] = {3, 1, 4, 2, 5};
            
            for (int i = 0; i < num_processes; i++) {
                scheduler.processes[i].pid = i + 1;
                scheduler.processes[i].arrival_time = arrival_times[i % 5];
                scheduler.processes[i].burst_time = burst_times[i % 5];
                scheduler.processes[i].priority = priorities[i % 5];
                scheduler.processes[i].completion_time = 0;
                scheduler.processes[i].turnaround_time = 0;
                scheduler.processes[i].waiting_time = 0;
                scheduler.processes[i].response_time = -1;
                scheduler.processes[i].remaining_time = scheduler.processes[i].burst_time;
                scheduler.processes[i].start_time = -1;
            }
            break;
    }
    
    printf("\nChoose scheduling algorithm:\n");
    printf("1. First-Come, First-Served (FCFS)\n");
    printf("2. Shortest Job First (SJF)\n");
    printf("3. Round Robin (RR)\n");
    printf("4. Priority Scheduling\n");
    printf("5. Shortest Remaining Time First (SRTF)\n");
    printf("6. Compare all algorithms\n");
    printf("Choice: ");
    scanf("%d", &choice);
    
    switch (choice) {
        case 1:
            fcfs_scheduling(&scheduler);
            print_results(&scheduler);
            break;
        case 2:
            sjf_scheduling(&scheduler);
            print_results(&scheduler);
            break;
        case 3:
            printf("Enter time quantum: ");
            scanf("%d", &scheduler.time_quantum);
            round_robin_scheduling(&scheduler);
            print_results(&scheduler);
            break;
        case 4:
            priority_scheduling(&scheduler);
            print_results(&scheduler);
            break;
        case 5:
            srtf_scheduling(&scheduler);
            print_results(&scheduler);
            break;
        case 6:
            compare_algorithms(&scheduler);
            break;
        default:
            printf("Invalid choice. Using FCFS.\n");
            fcfs_scheduling(&scheduler);
            print_results(&scheduler);
            break;
    }
    
    free(scheduler.processes);
    return 0;
}
