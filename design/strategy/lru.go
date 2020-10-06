/**
* @File:          lru.go
* @Aim:           最近最少被使用的移除
* @Author:        foyon
* @Created Time:  2020-10-06
 */
package main

import "fmt"

type Lru struct {
}

//实现策略接口，具体实现策略方法
func (f *Lru) evict(c *cache) {
	fmt.Printf("evict by lru \n")
}
