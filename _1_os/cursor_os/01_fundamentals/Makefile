# Makefile for Computer Architecture Fundamentals
# Compiles C programs and assembly programs

CC = gcc
AS = nasm
CFLAGS = -Wall -Wextra -O2 -std=c99
ASFLAGS = -f elf64
LDFLAGS = -no-pie

# Target programs
TARGETS = cpu_info memory_patterns calculator

# C programs
C_PROGRAMS = cpu_info memory_patterns

# Assembly programs
ASM_PROGRAMS = calculator

.PHONY: all clean test

all: $(TARGETS)

# C program targets
cpu_info: cpu_info.c
	$(CC) $(CFLAGS) -o cpu_info cpu_info.c

memory_patterns: memory_patterns.c
	$(CC) $(CFLAGS) -o memory_patterns memory_patterns.c

# Assembly program targets
calculator: calculator.o
	$(CC) $(LDFLAGS) -o calculator calculator.o

calculator.o: calculator.s
	$(AS) $(ASFLAGS) -o calculator.o calculator.s

# Test all programs
test: all
	@echo "=== Testing CPU Info ==="
	./cpu_info
	@echo ""
	@echo "=== Testing Memory Patterns ==="
	./memory_patterns
	@echo ""
	@echo "=== Testing Calculator ==="
	@echo "Testing calculator with sample input..."
	@echo "10" | echo "20" | echo "1" | ./calculator

# Clean up generated files
clean:
	rm -f $(TARGETS) *.o

# Install dependencies (Ubuntu/Debian)
install-deps:
	sudo apt-get update
	sudo apt-get install -y gcc nasm make

# Install dependencies (macOS)
install-deps-mac:
	brew install gcc nasm make

# Help
help:
	@echo "Available targets:"
	@echo "  all          - Build all programs"
	@echo "  cpu_info     - Build CPU information detector"
	@echo "  memory_patterns - Build memory pattern analyzer"
	@echo "  calculator   - Build assembly calculator"
	@echo "  test         - Run all programs"
	@echo "  clean        - Remove generated files"
	@echo "  install-deps - Install dependencies (Ubuntu/Debian)"
	@echo "  install-deps-mac - Install dependencies (macOS)"
	@echo "  help         - Show this help message"
