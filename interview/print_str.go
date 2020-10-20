/**
* @File:          print_str.go
* @Aim:           2个go func 交替有序打印
				  利用channel无缓冲阻塞的特性
* @Author:        foyon
* @Created Time:  2020-10-19
*/

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	printCode()
}

//实现两个字符串交叉有序打印，如果一个长将剩下的输出

func printCode() { // {{{
	//非缓冲造成死锁
	jobChan1 := make(chan bool, 1)
	jobChan2 := make(chan bool)
	//fishChan := make(chan struct{})
	wg := sync.WaitGroup{}

	jobChan1 <- true
	wg.Add(2)
	go func() {
		defer wg.Done()
		mark := 0
		for i := 0; i < 9; i++ {
			if _, ok := <-jobChan1; ok {
				mark++
				fmt.Printf("%d\n", i)
				jobChan2 <- true
			}
		}

		for j := mark; j < 9; j++ {
			fmt.Printf("%d\n", j)
		}

	}()

	go func() {
		defer wg.Done()
		mark := 0
		for i := 10; i < 19; i++ {
			if _, ok := <-jobChan2; ok {
				mark++
				fmt.Printf("%s\n", string([]byte(strconv.Itoa(i))))
				jobChan1 <- true
			}
		}

		for j := mark; j < 100; j++ {
			fmt.Printf("%s\n", []byte(strconv.Itoa(j)))
		}
	}()
	wg.Wait()
	close(jobChan1)
	close(jobChan2)
} // }}}
