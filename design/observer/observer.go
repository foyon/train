/**
* @File:          observer.go
* @Aim:           观察者模式，行为设计模式，一个状态改变，通知其实现订阅者接口对象
i				  观察者
* @Author:        foyon
* @Created Time:  2020-09-30
*/

package main

import (
	"fmt"
)

/*观察者约定接口*/
type observer interface {
	update(string)
	getId() string
}

/*具体观察者*/
type Customer struct {
	Id string
}

func (c *Customer) update(name string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.Id, name)
}

func (c *Customer) getId() string {
	return c.Id
}
