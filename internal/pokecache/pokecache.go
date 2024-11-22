package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *PokeCache {
	newCache := new(PokeCache)
	newCache.entries = make(map[string]CacheEntry)
	go func() {
		ticker := time.NewTicker(interval / 2)
		defer ticker.Stop()
		for range ticker.C {
			newCache.reaploop(interval)
		}
	}()

	return newCache
}
