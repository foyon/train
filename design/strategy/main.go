/**
* @File:          main.go
* @Aim:           策略模式：行为设计模式，将一组行为转为对象，并是其原始上下
							文对象能相互切换

							一个主体对象，透露一个接口，这个接口可以自定各种
							策略方法，有业务自己选择调用哪种策略方法
							各种策略须实现约定策略接口
* @Author:        foyon
* @Created Time:  2020-10-06
*/
package main

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}
