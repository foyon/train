/**
* @File:          fifo.go
* @Aim:           先进先出：移除最最创建条目
* @Author:        foyon
* @Created Time:  2020-10-06
 */
package main

import "fmt"

type Fifo struct {
}

//实现策略接口，具体实现策略方法
func (f *Fifo) evict(c *cache) {
	fmt.Printf("evict by fifo\n")
}
