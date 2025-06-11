package pokecache

import "time"

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
