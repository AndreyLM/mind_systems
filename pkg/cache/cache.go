package cache

import (
	"log"
	"sync"
)

// Storage - cache
var Storage cache

type cache struct {
	mu    sync.Mutex
	store map[string]string
}

func init() {
	Storage = cache{
		store: make(map[string]string),
	}
}

func (c *cache) Add(urlToken, url string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[urlToken] = url
	return nil
}

func (c *cache) Get(urlToken string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.store[urlToken]
}

func (c *cache) LogStore() {
	log.Println(c.store)
}
