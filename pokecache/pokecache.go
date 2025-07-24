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
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(i time.Duration) *Cache {
	newCache := &Cache{
		interval: i,
		entries:  make(map[string]cacheEntry),
	}
	go newCache.reaploop()

	return newCache

}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	res, ok := c.entries[key]

	if !ok {
		return nil, false
	}
	return res.val, true

}

func (c *Cache) reaploop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if entry.createdAt.Before(time.Now().Add(-c.interval)) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
