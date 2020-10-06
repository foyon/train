/**
* @File:          cache.go
* @Aim:           cache上下文
* @Author:        foyon
* @Created Time:  2020-10-06
 */
package main

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage: storage,
		//约定策略接口
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

//调用各种策略的具体不同策略方法
func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
