package pokecache

// need to close the cache when program exits

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry //map of cacheEntries with a string key
	mu       *sync.RWMutex         // guards access to the struct map
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheMap: make(map[string]cacheEntry),
		mu:       &sync.RWMutex{},
	}
	cache.reapLoop(interval)
	return cache
}

// used to safely add values to the cache
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// used to safely delete values
func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cacheMap, key)
}

/*
used to check if value is already in the cache -- can be used to return cached data, or with the bool to determine
when to add items to the cache
*/
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.cacheMap[key]
	if ok {
		return entry.val, ok
	}
	// if the key doesn't exist in the Cache, return nil and false
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	//fmt.Println("interval value:", interval)
	ticker := time.NewTicker(interval) // creates a ticker that will tick for the chosen duration
	// most tickers run for a fixed amount of time -- in our case we want the ticker to run for duration of application
	// we could use defer.Ticker.Stop() if we needed a ticker to run until a certain condition was met.
	// go rountine which checks if a cache entry has expired and delete it if so
	go func() {
		for t := range ticker.C {
			for entry := range c.cacheMap {
				z := c.cacheMap[entry].createdAt.Round(time.Millisecond)
				if (t.Sub(z).Round(time.Millisecond)) >= interval {
					c.Remove(entry)
				}
			}

		}
	}()
}
