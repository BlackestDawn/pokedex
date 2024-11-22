package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	entries map[string]CacheEntry
	mu      sync.RWMutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *PokeCache) Add(key string, val []byte) {
	if _, exists := c.entries[key]; exists {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	v := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = v
}

func (c *PokeCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if v, exists := c.entries[key]; exists {
		return v.val, true
	}
	return []byte{}, false
}

func (c *PokeCache) reaploop(timeout time.Duration) {
	if len(c.entries) == 0 {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.entries {
		if time.Since(v.createdAt) > timeout {
			delete(c.entries, k)
		}
	}
}
