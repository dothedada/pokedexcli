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
