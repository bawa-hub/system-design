#!/bin/bash

# Test script for all Week 1 implementations
echo "=== Testing All Week 1 Implementations ==="
echo ""

# Function to test a day
test_day() {
    local day=$1
    local description=$2
    
    echo "Testing Day $day: $description"
    echo "----------------------------------------"
    
    cd "week1/day$day"
    
    if go run main.go; then
        echo "✓ Day $day passed"
    else
        echo "✗ Day $day failed"
        exit 1
    fi
    
    cd ../..
    echo ""
}

# Test all days
test_day 1 "Caching Fundamentals"
test_day 2 "Cache Eviction Policies"
test_day 3 "Cache Strategies & Patterns"
test_day 4 "Cache Consistency Models"
test_day 5 "Memory Management & Optimization"
test_day 6 "Cache Warming & Preloading"
test_day 7 "Week 1 Review & Assessment"

echo "=== All Tests Passed! ==="
echo "You're ready to start learning caching! 🚀"
