package main

// need to close the cache when program exits

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry //map of cacheEntries with a string key
	mu       *sync.Mutex           // guards access to the struct map
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheMap: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
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

/*
used to check if value is already in the cache -- can be used to return cached data, or with the bool to determine
when to add items to the cache
*/
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheMap[key]
	if ok {
		return entry.val, ok
	}
	// if the key doesn't exist in the Cache, return nil and false
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	fmt.Println("interval value:", interval)
	ticker := time.NewTicker(interval) // creates a ticker that will run for the chosen Duration
	// most tickers run for a fixed amount of time -- in our case we want the ticker to run for duration of application

	go func() {
		// store channel value in variable t
		t := <- ticker.C
			for entry := range c.cacheMap{
				fmt.Printf("creation time: %v\n", c.cacheMap[entry].createdAt)
				test := t.Round(time.Second).Sub(c.cacheMap[entry].createdAt)
				fmt.Println(test)
				if t.Round(time.Second).Sub(c.cacheMap[entry].createdAt) >= interval {
					fmt.Println("duration has passed -- entry deleted")
					delete(c.cacheMap, entry)
				}
			}
	}()
}

func main() {
	cache := NewCache(3000 * time.Millisecond)
	data := []byte("nothing to see here")
	data2 := []byte("seriously don't look")

	cache.Add("http://example.com", data)
	cache.Add("http://google.com", data2)

	fmt.Println("testing cache")

	time.Sleep(5000 * time.Millisecond)

	val, exists := cache.Get("http://example.com")
	if exists {
		fmt.Println("val - cache isn't working:", val)
	} else {
		fmt.Println("val - working cache")
	}

}
