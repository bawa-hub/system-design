# Makefile for System Programming Basics
# Compiles C programs for system programming

CC = gcc
CFLAGS = -Wall -Wextra -O2 -std=c99 -D_POSIX_C_SOURCE=200809L
LDFLAGS = -lpthread

# Target programs
TARGETS = file_copy process_monitor custom_allocator

# Source files
SOURCES = file_copy.c process_monitor.c custom_allocator.c

.PHONY: all clean test benchmark

all: $(TARGETS)

# Individual targets
file_copy: file_copy.c
	$(CC) $(CFLAGS) -o file_copy file_copy.c

process_monitor: process_monitor.c
	$(CC) $(CFLAGS) -o process_monitor process_monitor.c

custom_allocator: custom_allocator.c
	$(CC) $(CFLAGS) -o custom_allocator custom_allocator.c

# Test all programs
test: all
	@echo "=== Testing File Copy ==="
	@echo "Creating test file..."
	@echo "Hello, World!" > test_input.txt
	@echo "Testing file copy..."
	./file_copy test_input.txt test_output.txt -v
	@echo "Testing recursive directory copy..."
	@mkdir -p test_dir
	@echo "Test file in directory" > test_dir/test_file.txt
	./file_copy test_dir test_dir_copy -r
	@echo "Cleaning up test files..."
	@rm -f test_input.txt test_output.txt
	@rm -rf test_dir test_dir_copy
	@echo ""
	@echo "=== Testing Process Monitor ==="
	@echo "Running process monitor demo..."
	./process_monitor demo
	@echo ""
	@echo "=== Testing Custom Allocator ==="
	@echo "Running allocator tests..."
	./custom_allocator

# Benchmark the allocator
benchmark: custom_allocator
	@echo "=== Benchmarking Custom Allocator ==="
	./custom_allocator

# Run memory leak check with valgrind
valgrind: all
	@echo "=== Running Valgrind Memory Check ==="
	@echo "Checking file_copy..."
	valgrind --leak-check=full --show-leak-kinds=all ./file_copy test_input.txt test_output.txt
	@echo "Checking process_monitor..."
	valgrind --leak-check=full --show-leak-kinds=all ./process_monitor demo
	@echo "Checking custom_allocator..."
	valgrind --leak-check=full --show-leak-kinds=all ./custom_allocator

# Run with AddressSanitizer
asan: CFLAGS += -fsanitize=address -g
asan: LDFLAGS += -fsanitize=address
asan: all
	@echo "=== Running with AddressSanitizer ==="
	./file_copy test_input.txt test_output.txt
	./process_monitor demo
	./custom_allocator

# Run with ThreadSanitizer
tsan: CFLAGS += -fsanitize=thread -g
tsan: LDFLAGS += -fsanitize=thread
tsan: all
	@echo "=== Running with ThreadSanitizer ==="
	./process_monitor demo

# Profile with gprof
profile: CFLAGS += -pg
profile: LDFLAGS += -pg
profile: all
	@echo "=== Profiling with gprof ==="
	./custom_allocator
	gprof custom_allocator gmon.out > profile.txt
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
	rm -f test_input.txt test_output.txt
	rm -rf test_dir test_dir_copy

# Help
help:
	@echo "Available targets:"
	@echo "  all          - Build all programs"
	@echo "  file_copy    - Build file copy utility"
	@echo "  process_monitor - Build process monitor"
	@echo "  custom_allocator - Build custom memory allocator"
	@echo "  test         - Run all programs with tests"
	@echo "  benchmark    - Benchmark the custom allocator"
	@echo "  valgrind     - Run memory leak checks"
	@echo "  asan         - Run with AddressSanitizer"
	@echo "  tsan         - Run with ThreadSanitizer"
	@echo "  profile      - Profile with gprof"
	@echo "  install-deps - Install dependencies (Ubuntu/Debian)"
	@echo "  install-deps-mac - Install dependencies (macOS)"
	@echo "  clean        - Remove generated files"
	@echo "  help         - Show this help message"
