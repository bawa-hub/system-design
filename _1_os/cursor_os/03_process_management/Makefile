# Makefile for Process Management
# Compiles C programs for process management

CC = gcc
CFLAGS = -Wall -Wextra -O2 -std=c99 -D_POSIX_C_SOURCE=200809L
LDFLAGS = -lpthread

# Target programs
TARGETS = process_scheduler process_manager ipc_demo

# Source files
SOURCES = process_scheduler.c process_manager.c ipc_demo.c

.PHONY: all clean test demo

all: $(TARGETS)

# Individual targets
process_scheduler: process_scheduler.c
	$(CC) $(CFLAGS) -o process_scheduler process_scheduler.c

process_manager: process_manager.c
	$(CC) $(CFLAGS) -o process_manager process_manager.c

ipc_demo: ipc_demo.c
	$(CC) $(CFLAGS) -o ipc_demo ipc_demo.c

# Test all programs
test: all
	@echo "=== Testing Process Scheduler ==="
	@echo "Running scheduler with default data..."
	@echo "5" | echo "3" | echo "1" | ./process_scheduler
	@echo ""
	@echo "=== Testing Process Manager ==="
	@echo "Running process manager demo..."
	./process_manager demo
	@echo ""
	@echo "=== Testing IPC Demo ==="
	@echo "Running IPC demonstrations..."
	./ipc_demo

# Run interactive demos
demo: all
	@echo "=== Interactive Process Scheduler ==="
	@echo "Run: ./process_scheduler"
	@echo ""
	@echo "=== Interactive Process Manager ==="
	@echo "Run: ./process_manager"
	@echo ""
	@echo "=== Interactive IPC Demo ==="
	@echo "Run: ./ipc_demo interactive"

# Run specific tests
test-scheduler: process_scheduler
	@echo "=== Testing Process Scheduler ==="
	@echo "Testing FCFS scheduling..."
	@echo "5" | echo "3" | echo "1" | ./process_scheduler
	@echo ""
	@echo "Testing Round Robin scheduling..."
	@echo "5" | echo "3" | echo "3" | echo "2" | ./process_scheduler

test-manager: process_manager
	@echo "=== Testing Process Manager ==="
	./process_manager demo

test-ipc: ipc_demo
	@echo "=== Testing IPC Demo ==="
	./ipc_demo

# Run with valgrind for memory checking
valgrind: all
	@echo "=== Running Valgrind Memory Check ==="
	@echo "Checking process_scheduler..."
	valgrind --leak-check=full --show-leak-kinds=all ./process_scheduler < /dev/null
	@echo "Checking process_manager..."
	valgrind --leak-check=full --show-leak-kinds=all ./process_manager demo
	@echo "Checking ipc_demo..."
	valgrind --leak-check=full --show-leak-kinds=all ./ipc_demo

# Run with AddressSanitizer
asan: CFLAGS += -fsanitize=address -g
asan: LDFLAGS += -fsanitize=address
asan: all
	@echo "=== Running with AddressSanitizer ==="
	./process_scheduler < /dev/null
	./process_manager demo
	./ipc_demo

# Run with ThreadSanitizer
tsan: CFLAGS += -fsanitize=thread -g
tsan: LDFLAGS += -fsanitize=thread
tsan: all
	@echo "=== Running with ThreadSanitizer ==="
	./process_manager demo

# Profile with gprof
profile: CFLAGS += -pg
profile: LDFLAGS += -pg
profile: all
	@echo "=== Profiling with gprof ==="
	./process_scheduler < /dev/null
	gprof process_scheduler gmon.out > profile.txt
	@echo "Profile saved to profile.txt"

# Install dependencies (Ubuntu/Debian)
install-deps:
	sudo apt-get update
	sudo apt-get install -y gcc make valgrind

# Install dependencies (macOS)
install-deps-mac:
	brew install gcc make valgrind

# Clean up generated files
clean:
	rm -f $(TARGETS) *.o gmon.out profile.txt
	rm -f /tmp/myfifo /tmp/unix_socket

# Help
help:
	@echo "Available targets:"
	@echo "  all          - Build all programs"
	@echo "  process_scheduler - Build process scheduler simulator"
	@echo "  process_manager - Build process manager"
	@echo "  ipc_demo     - Build IPC demonstration program"
	@echo "  test         - Run all programs with tests"
	@echo "  demo         - Show how to run interactive demos"
	@echo "  test-scheduler - Test process scheduler specifically"
	@echo "  test-manager - Test process manager specifically"
	@echo "  test-ipc     - Test IPC demo specifically"
	@echo "  valgrind     - Run memory leak checks"
	@echo "  asan         - Run with AddressSanitizer"
	@echo "  tsan         - Run with ThreadSanitizer"
	@echo "  profile      - Profile with gprof"
	@echo "  install-deps - Install dependencies (Ubuntu/Debian)"
	@echo "  install-deps-mac - Install dependencies (macOS)"
	@echo "  clean        - Remove generated files"
	@echo "  help         - Show this help message"
