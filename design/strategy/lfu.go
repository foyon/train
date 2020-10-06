/**
* @File:          lfu.go
* @Aim:           使用频率最低的一条进行移除
* @Author:        foyon
* @Created Time:  2020-10-06
 */
package main

import "fmt"

type Lfu struct {
}

//实现策略接口，具体实现策略方法
func (f *Lfu) evict(c *cache) {
	fmt.Printf("evict by lfu \n")
}
