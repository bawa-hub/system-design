# Week 1 Exercises

## Exercise 1: Basic Cache Implementation

### Task
Implement a simple cache with the following features:
- Thread-safe operations
- TTL (Time To Live) support
- Basic statistics (hit rate, miss rate)
- Memory usage tracking

### Requirements
1. Use Go's `sync.RWMutex` for thread safety
2. Implement automatic cleanup of expired items
3. Track cache statistics
4. Add unit tests

### Solution Template
```go
// exercise1.go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Your implementation here
```

## Exercise 2: LRU Cache Implementation

### Task
Implement a Least Recently Used (LRU) cache with:
- Fixed capacity
- O(1) get and set operations
- Automatic eviction of least recently used items

### Requirements
1. Use doubly linked list + hash map
2. Implement move-to-head operation
3. Handle capacity overflow
4. Add performance benchmarks

### Solution Template
```go
// exercise2.go
package main

import (
    "fmt"
    "sync"
)

// Your LRU implementation here
```

## Exercise 3: Cache Strategy Implementation

### Task
Implement different cache strategies:
- Cache-Aside
- Read-Through
- Write-Through
- Write-Behind

### Requirements
1. Create interface for cache strategies
2. Implement each strategy
3. Add configuration options
4. Test with mock database

### Solution Template
```go
// exercise3.go
package main

import (
    "fmt"
    "time"
)

// Your strategy implementations here
```

## Exercise 4: Memory-Efficient Cache

### Task
Build a memory-efficient cache with:
- Memory pooling
- Arena allocation
- Memory monitoring
- Automatic optimization

### Requirements
1. Implement memory pool
2. Use arena allocation for small objects
3. Monitor memory usage
4. Implement automatic cleanup

### Solution Template
```go
// exercise4.go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

// Your memory-efficient implementation here
```

## Exercise 5: Cache Warming System

### Task
Create a cache warming system with:
- Predictive warming
- Scheduled warming
- Event-driven warming
- Performance monitoring

### Requirements
1. Implement data predictor
2. Create warming scheduler
3. Add event handling
4. Monitor warming effectiveness

### Solution Template
```go
// exercise5.go
package main

import (
    "fmt"
    "time"
)

// Your warming system here
```

## Submission Guidelines

1. **Code Quality**: Write clean, well-documented code
2. **Testing**: Include comprehensive unit tests
3. **Performance**: Add benchmarks for critical operations
4. **Documentation**: Write clear README files
5. **Error Handling**: Handle edge cases and errors gracefully

## Evaluation Criteria

- **Correctness**: Does the implementation work as expected?
- **Performance**: Are the operations efficient?
- **Code Quality**: Is the code clean and maintainable?
- **Testing**: Are there adequate tests?
- **Documentation**: Is the code well-documented?

## Bonus Challenges

1. **Distributed Cache**: Implement a distributed cache with consistent hashing
2. **Cache Compression**: Add compression for large values
3. **Cache Encryption**: Implement encryption for sensitive data
4. **Advanced Statistics**: Add detailed performance metrics
5. **Cache Visualization**: Create a web interface to monitor cache

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Sync Package](https://golang.org/pkg/sync/)
- [Testing Package](https://golang.org/pkg/testing/)
- [Benchmarking in Go](https://golang.org/pkg/testing/#hdr-Benchmarks)

---

**Good luck with your exercises! 🚀**
