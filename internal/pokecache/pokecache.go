package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		data: map[string]cacheEntry{},
	}
	go c.reapLoop(interval)
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	fmt.Printf("\nCache Add: %v\n", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key]
	if !ok {
		fmt.Printf("\nCache Not Found: %v\n", key)
		return nil, false
	}

	fmt.Printf("\nCache Get: %v\n", key)
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		expiredBy := time.Now().Add(-interval)
		
		c.mu.Lock()

		for key, entry := range c.data {
			if entry.createdAt.Before(expiredBy) {
				delete(c.data, key)
				fmt.Printf("\nCache Clear: %v\n", key)
			}
		}

		c.mu.Unlock()
	}
}