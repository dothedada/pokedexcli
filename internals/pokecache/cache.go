package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data:  make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}
	cache.ReapLoop(interval)
	return cache
}
