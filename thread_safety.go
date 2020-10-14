/**
* @File:          thread_safety.go
* @Aim:           实现线程安全两种方式， 加锁 和用channel 控制
* @Author:        foyon
* @Created Time:  2020-10-14
 */

/*
常见的线程安全的有 sync.WaitGroup sync.Once Sync.Map
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type stu struct {
}

func main() {

	demoTreadByChannel()

}

func compareDemo() {
	stu1 := stu{}
	var stu2 interface{}
	fmt.Println(stu1)  //{}
	fmt.Println(&stu1) //&{}

	fmt.Printf("%v", stu2)  //nil
	fmt.Printf("%v", &stu2) //0xc000010200
}

func demoChan() { // {{{
	done := make(chan struct{})

	go func() {
		// ...
		close(done)
	}()

	<-done
} // }}}

func demoChan2() { // {{{
	done := make(chan struct{})

	go func() {
		//发送一个obejct
		//interface{}{} 传递的是nil,不太友好
		done <- struct{}{}
	}()

	<-done
} // }}}

func demoTreadByChannel() { // {{{

	//job := make(chan interface{}, 100)
	//用struct{} 非interface{}
	//https://stackoverflow.com/questions/52035390/why-using-chan-struct-when-wait-something-done-not-chan-interface
	//inferface传递nil不如struct{}好
	//buffer满来控制并发度
	//100运行并行goroutine

	//利用golang channel buffer特性实现goroutine池
	job := []string{"a", "b", "c", "d"}
	//work := make(chan struct{}, 100)
	work := make(chan string, 100)
	wg := sync.WaitGroup{}

	for _, v := range job {
		wg.Add(1)
		work <- v
		go func() {
			defer wg.Done()
			jobw := <-work
			time.Sleep(5 * time.Second)
			fmt.Printf("%v", jobw)
		}()
	}

	//放这里goroutine才能并行执行
	wg.Wait()
} // }}}
