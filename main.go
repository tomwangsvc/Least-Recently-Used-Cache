package main

import (
	"sync"
	"time"
)

type Value struct {
	Value int
	Time  time.Time
}
type Cache struct {
	cache    map[int]Value
	capacity int
	mu       sync.Mutex
}

func NewCache(capacity int) *Cache {
	return &Cache{cache: make(map[int]Value), capacity: capacity}
}

func (c *Cache) get(key int) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.getValue(key)
}

func (c *Cache) getValue(key int) int {
	if v, ok := c.cache[key]; ok {
		v.Time = time.Now()
		c.cache[key] = v
		return v.Value
	}
	return -1
}

func (c *Cache) put(key, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if v := c.getValue(key); v == -1 && len(c.cache) == c.capacity {
		var oldestTime time.Time
		var keyForOldestTime int
		var flag bool
		for k := range c.cache {
			if !flag {
				oldestTime = c.cache[k].Time
				keyForOldestTime = k
				flag = true
			}
			if oldestTime.After(c.cache[k].Time) {
				oldestTime = c.cache[k].Time
				keyForOldestTime = k
			}
		}
		delete(c.cache, keyForOldestTime)
	}

	c.cache[key] = Value{
		Value: value,
		Time:  time.Now(),
	}
}

func (c *Cache) delete(key int) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	v := c.getValue(key)
	delete(c.cache, key)
	return v
}
