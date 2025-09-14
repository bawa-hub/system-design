package main

import (
    "fmt"
    "sync"
    "time"
)

// CacheItem represents a single item in the cache
type CacheItem struct {
    Value     interface{}
    ExpiresAt time.Time
}

// IsExpired checks if the cache item has expired
func (item *CacheItem) IsExpired() bool {
    return time.Now().After(item.ExpiresAt)
}

// SimpleCache represents our basic cache
type SimpleCache struct {
    items map[string]*CacheItem
    mutex sync.RWMutex
}

// NewSimpleCache creates a new cache instance
func NewSimpleCache() *SimpleCache {
    return &SimpleCache{
        items: make(map[string]*CacheItem),
    }
}

// Set stores a value in the cache with TTL
func (c *SimpleCache) Set(key string, value interface{}, ttl time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.items[key] = &CacheItem{
        Value:     value,
        ExpiresAt: time.Now().Add(ttl),
    }
}

// Get retrieves a value from the cache
func (c *SimpleCache) Get(key string) (interface{}, bool) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    item, exists := c.items[key]
    if !exists {
        return nil, false
    }
    
    if item.IsExpired() {
        // Clean up expired item
        c.mutex.RUnlock()
        c.mutex.Lock()
        delete(c.items, key)
        c.mutex.Unlock()
        c.mutex.RLock()
        return nil, false
    }
    
    return item.Value, true
}

// Delete removes an item from the cache
func (c *SimpleCache) Delete(key string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    delete(c.items, key)
}

// Clear removes all items from the cache
func (c *SimpleCache) Clear() {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.items = make(map[string]*CacheItem)
}

// Size returns the number of items in the cache
func (c *SimpleCache) Size() int {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    return len(c.items)
}

func main() {
    // Create a new cache
    cache := NewSimpleCache()
    
    // Set some values
    cache.Set("user:1", "John Doe", 5*time.Second)
    cache.Set("user:2", "Jane Smith", 10*time.Second)
    cache.Set("config:theme", "dark", 1*time.Minute)
    
    // Test cache operations
    fmt.Println("=== Cache Operations ===")
    
    // Get existing value
    if value, found := cache.Get("user:1"); found {
        fmt.Printf("Found user:1 = %v\n", value)
    }
    
    // Get non-existent value
    if value, found := cache.Get("user:999"); found {
        fmt.Printf("Found user:999 = %v\n", value)
    } else {
        fmt.Println("user:999 not found")
    }
    
    // Wait for expiration
    fmt.Println("\nWaiting for user:1 to expire...")
    time.Sleep(6 * time.Second)
    
    if value, found := cache.Get("user:1"); found {
        fmt.Printf("Found user:1 = %v\n", value)
    } else {
        fmt.Println("user:1 expired and was removed")
    }
    
    // Cache statistics
    fmt.Printf("\nCache size: %d\n", cache.Size())
    
    // Clear cache
    cache.Clear()
    fmt.Printf("After clear, cache size: %d\n", cache.Size())
}
