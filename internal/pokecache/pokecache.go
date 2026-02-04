package pokecache

import (
	"sync"
	"time"
	"fmt"
)

type Cache struct {
	mu *sync.Mutex
	entry map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		mu: &sync.Mutex{},
		entry: make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	fmt.Println("Caching key")
	c.entry[key] = cacheEntry {
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	fmt.Println("Running from cache")
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	
	for range ticker.C {
		c.mu.Lock()

		for key, val := range c.entry {
			if time.Since(val.createdAt) > c.interval {
				delete(c.entry, key)
			}
		}

		c.mu.Unlock()
	}

}