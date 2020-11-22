package main

import (
	"fmt"
	_ "runtime"
	"sync"
)

var t map[string]interface{}

func main() {
	//runtime.GOMAXPROCS(1)
	fmt.Println("vim-go")
	addShareMemSin(10)
	//addShareMem(10)
	//addShareMemChan(10)
}

func addShareMemSin(n int) []int { // {{{

	var ints []int

	for i := 0; i < n; i++ {
		//并发执行，会进行覆盖
		ints = append(ints, i)
	}

	fmt.Println(ints)
	return ints

} // }}}

//communicating by sharemem
func addShareMem(n int) []int { // {{{

	var ints []int
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(10)
	for i := 0; i < n; i++ {
		//并发执行，会进行覆盖
		go func(i int) {
			defer wg.Done()
			mux.Lock()
			ints = append(ints, i)
			mux.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(ints)
	return ints

} // }}}

//sharemem by communicating
func addShareMemChan(n int) []int { // {{{

	var ints []int
	channel := make(chan int, n)

	for i := 0; i < n; i++ {
		go func(channel chan<- int, order int) {
			channel <- order
		}(channel, i)
	}

	//阻塞
	for i := range channel {
		ints = append(ints, i)
		if n == len(ints) {
			break
		}
	}

	close(channel)
	fmt.Println(ints)
	return ints

} // }}}
