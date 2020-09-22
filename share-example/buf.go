package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("vim-go")
	//unbuf()
	buf()
}

func unbuf() {
	unbuf := make(chan bool)
	//背景执行
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("go func ==========")
		unbuf <- true
	}()

	fmt.Println("go func exec over")
	<-unbuf
	fmt.Println("main exec over")
}

//无缓冲的channel 必须等待接收方准备完成才能释放,所以是阻塞的
func buf() {
	buf := make(chan bool, 0)
	//buf := make(chan bool)
	go func() {
		fmt.Println("go go go")
		time.Sleep(3 * time.Second)
		fmt.Println("go2 go2 go2")
		//buf <- 1
		<-buf
	}()

	//<-buf
	buf <- true
	//buf <- false

	time.Sleep(1 * time.Second)
	fmt.Println("exec over")
}
