/**
* @File:          channel_broad.go
* @Aim:           利用channel close 特性实现main与子go func 广播模式
* @Author:        foyon
* @Created Time:  2020-11-18
 */
package main

import (
	"fmt"
	"sync"
)

// 利用chan close 特性实现广播功能
// 面试案例，实现一个map，一些携程实现写，多携程实现读，检测exist，当map[key]一直进行阻塞，直到key 被写入。

var (
	Gmap  map[string]interface{}
	L     sync.RWMutex
	Broad chan struct{}
	wg    sync.WaitGroup
)

func main() {}

func Set(key string, val interface{}) bool { // {{{
	if Gmap == nil {
		Gmap = make(map[string]interface{})
	}

	if val, isExist := Gmap[key]; isExist {
		return
	}

	L.RWLock()
	defer L.RWUNLock()

	Gmap[key] = string
	_, isClose := <-Broad

	if !isClose {
		Broad = make(chan struct{})
	}

	close(Broad)
} // }}}

func Get(key string) interface{} { // {{{
	L.RLock()
	defer L.RUNLock()

	if val, isExist := Gmap[key]; isExist {
		return val
	}

	for {
		select {
		case <-Broad:
			if val, isExist := Gmap[key]; isExist {
				return val
			}
		}
	}

} // }}}
