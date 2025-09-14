# Simple Cache Project

This project implements a basic in-memory cache with TTL (Time To Live) support.

## Features

- **Thread-safe**: Uses mutex for concurrent access
- **TTL Support**: Automatic expiration of cache items
- **Basic Operations**: Set, Get, Delete, Clear
- **Memory Efficient**: Simple map-based storage

## Usage

```go
// Create a new cache
cache := NewSimpleCache()

// Set a value with TTL
cache.Set("key", "value", 5*time.Second)

// Get a value
if value, found := cache.Get("key"); found {
    fmt.Printf("Value: %v\n", value)
}

// Delete a value
cache.Delete("key")

// Clear all values
cache.Clear()
```

## Running

```bash
go run main.go
```

## Learning Objectives

- Understand basic caching concepts
- Learn thread-safe programming in Go
- Implement TTL functionality
- Practice error handling and cleanup
