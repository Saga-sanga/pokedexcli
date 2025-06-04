package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, exists := c.entries[key]
	if exists {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			// reaping process
			c.mu.Lock()
			for key, entry := range c.entries {
				timeDifference := time.Now().Sub(entry.createdAt)
				if timeDifference > interval {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{entries: make(map[string]cacheEntry)}
	cache.reapLoop(interval)
	return cache
}
