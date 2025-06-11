package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data  map[string]cacheEntry
	mutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	bytes     []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data:  make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}
	cache.ReapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	dataToInsert := cacheEntry{
		createdAt: time.Now(),
		bytes:     val,
	}

	c.data[key] = dataToInsert

	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if entry, ok := c.data[key]; ok {
		return entry.bytes, true
	}

	return []byte{}, false
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			select {

			case t := <-ticker.C:
				c.mutex.Lock()
				for key, value := range c.data {
					if value.createdAt.Add(interval).Before(t) {
						delete(c.data, key)
					}
				}

				c.mutex.Unlock()
			}
		}
	}()
}
