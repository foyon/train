package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
控制共享内存的两种方式

*/
func main() {
	//runtime.GOMAXPROCS(1)
	//是超线程数,非实际核心数,以mac 举例
	num := runtime.NumCPU()
	fmt.Println("this mac cpu:", num)
	//addShareMemSin(10)
	addShareMem(10)
	//addShareMemChan(10)
}

/*
*1、将会输出什么？
*2、P(1)
*3. 非go func()
 */
func addShareMemSin(n int) []int { // {{{

	var ints []int

	for i := 0; i < n; i++ {
		ints = append(ints, i)
	}

	fmt.Println(ints)
	return ints

} // }}}

//communicating by sharemem
//操作共享内存
/*
1、会输出什么？try(5)
2、加锁
*/
func addShareMem(n int) []int { // {{{

	var ints []int
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(n)
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
//goroutine之间需要交换信息一般推荐channel
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
