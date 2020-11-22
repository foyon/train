/**
* @File:          value_tran.go
* @Aim:           main进程与go程变量传递
* @Author:        foyon
* @Created Time:  2020-10-14
 */
package main

import (
	"fmt"
	"sync"
)

func main() {
	//输出 10居多，还有些其他的
	errorTran()

	//参数传递保证正确
	//corretTran()
}

func errorTran() { // {{{
	wg := sync.WaitGroup{}
	wg.Add(100)
	//输出错误，变量作用域问题
	//r := 0

	for i := 0; i < 100; i++ {
		//进行一次拷贝即可
		//r = i

		//此种方式可以
		i := i
		go func() {
			defer wg.Done()
			fmt.Printf("%d\n", r)
		}()
	}
	wg.Wait()
} // }}}

//通过参数的方式传入,输出乱序
func corretTran() { // {{{
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%d\n", i)
		}(i)
	}

	wg.Wait()
} // }}}
