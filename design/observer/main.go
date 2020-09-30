/**
* @File:          main.go
* @Aim:           观察者模式 测试入口
* @Author:        foyon
* @Created Time:  2020-09-30
 */
package main

import (
	_ "fmt"
)

func main() {
	shirtItem := newItem("foyon shirt")
	ob1 := &Customer{
		Id: "foyon0806@gmail.com",
	}

	ob2 := &Customer{
		Id: "foyon0806@outlook.com",
	}

	shirtItem.register(ob1)
	shirtItem.register(ob2)

	shirtItem.updateSt()

}
