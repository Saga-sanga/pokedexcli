package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache -
type Cache struct {
	sync.Mutex
	entries map[string]cacheEntry
}

// Add -
func (c *Cache) Add(key string, val []byte) {
	c.Lock()
	defer c.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get -
func (c *Cache) Get(key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()

	entry, exists := c.entries[key]
	if exists {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				// reaping process
				c.reap(interval)
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()
}

func (c *Cache) reap(interval time.Duration) {
	c.Lock()
	defer c.Unlock()
	for key, entry := range c.entries {
		if time.Since(entry.createdAt) >= interval {
			delete(c.entries, key)
		}
	}
}

// NewCache -
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{entries: make(map[string]cacheEntry)}
	cache.reapLoop(interval)
	return cache
}
