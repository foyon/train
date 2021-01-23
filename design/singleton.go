/**
* @File:          singleton.go
* @Aim:           单例模式,防止类多次实例化
* @Author:        foyon
* @Created Time:  2020-09-19
 */
package main

import (
	"fmt"
	"sync"
)

type Singeton struct {
	ID int
}

var model *Singeton
var mutex sync.Mutex
var once sync.Once

func GetModel() *Singeton { // {{{

	if model == nil {
		mutex.Lock()
		if model == nil {
			model = &Singeton{}
		}
		mutex.Unlock()
	}

	return model
} // }}}

//new 分配内存，置零，返回指针地址
//make  slice\map\channel初始化数据结构
func GetModel2() *Singeton { // {{{

	if model == nil {
		mutex.Lock()
		if model == nil {
			model = new(Singeton)
			fmt.Printf("%+v", model)
		}
		mutex.Unlock()
	}

	return model
} // }}}

func GetModel3() *Singeton {
	once.Do(func() {
		model = &Singeton{}
	})

	//fmt.Printf("%+v", model)
	return model
}

func main() {
	//GetModel()
	//GetModel2()
	GetModel3()
}
