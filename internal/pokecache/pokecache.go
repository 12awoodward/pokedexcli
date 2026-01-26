package pokecache

import (
	"sync"
	"time"
)

type cache struct {
	data map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *cache {
	c := cache{}
	go c.reapLoop(interval)
	return &c
}

func (c *cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	
	for range ticker.C {
		expiredBy := time.Now().Add(-interval)
		
		c.mu.Lock()

		for key, entry := range c.data {
			if entry.createdAt.Before(expiredBy) {
				delete(c.data, key)
			}
		}

		c.mu.Unlock()
	}
}