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

	return newCache

}

func (c *Cach) Add()
