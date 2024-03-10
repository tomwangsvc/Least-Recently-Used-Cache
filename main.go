package main

import (
	"time"
)

type Value struct {
	Value int
	Time  time.Time
}
type Cache struct {
	cache    map[int]Value
	capacity int
}

func NewCache(capacity int) Cache {
	return Cache{cache: make(map[int]Value), capacity: capacity}
}

func (c *Cache) get(key int) int {
	if v, ok := c.cache[key]; ok {
		v.Time = time.Now()
		c.cache[key] = v
		return v.Value
	}
	return -1
}

func (c *Cache) put(key, value int) {
	if v := c.get(key); v == -1 && len(c.cache) == c.capacity {
		var oldestTime time.Time
		var keyForOldestTime int
		i := 0
		for k := range c.cache {
			if i == 0 {
				oldestTime = c.cache[k].Time
				keyForOldestTime = k
				i++
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
	v := c.get(key)
	delete(c.cache, key)
	return v
}

func main() {

}
