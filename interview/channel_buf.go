package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 10)
	go tick(ch)
	for i := 0; i < 109; i++ {
		ch <- i
		//time.Sleep(1 * time.Second)
	}

	time.Sleep(30 * time.Second)
}

func tick(ch chan int) {
	var b []int
	for {
		select {

		case <-time.After(3 * time.Second):
			//fmt.Println(len(ch))
			handJob(ch)
			break
			//default:
		//handJob2(ch)
		case v := <-ch:
			b = append(b, v)
			if len(b) == 10 {
				fmt.Printf("%+v \n", b)
				b = b[0:0]
			}
		}

	}

}

func handJob2(ch <-chan int) {
	var b []int
	for j := 0; j < 10; j++ {
		if v, ok := <-ch; ok {
			b = append(b, v)
		}
		fmt.Printf("exec here\n")
	}

	fmt.Print("b=%v", b)
}

func handJob(ch chan int) {
	var b []int
	lenCh := len(ch)
	if lenCh < 1 {
		return
	}
	for i := 0; i < lenCh; i++ {
		b = append(b, <-ch)
	}

	fmt.Printf("exec timeout b=%v", b)
}
