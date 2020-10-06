/**
* @File:          evictionAlgo.go
* @Aim:           策略接口: 内存清理为例，3种策略方法
* @Author:        foyon
* @Created Time:  2020-10-06
 */
package main

type evictionAlgo interface {
	evict(c *cache)
}
