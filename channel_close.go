/**
* @File:          channel_close.go
* @Aim:           channel 优雅关闭以及非正常关闭的一些问题
* @Author:        foyon
* @Created Time:  2020-10-15
 */

package main

import (
	"fmt"
)

func main() {
	//testSlice()
	//unBufChannel()
	bufChannel()
}

//测试切片扩容
//底层是数组,内存是连续的
func testSlice() { // {{{
	str := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Printf("%p\n", &str)
	str = append(str, "i")
	fmt.Printf("%p\n", &str)

	//2len方式进行扩容，地址不会改变
	i := []int{1, 2, 3, 4}
	fmt.Printf("%p\n", &i)
	i = append(i, 5)
	fmt.Printf("%p\n", &i)
} // }}}

//无缓冲buf，阻塞
func unBufChannel() { // {{{
	jobChan := make(chan string)
	close(jobChan)

	//一个关闭的阻塞的channel进行读，会直接返回false
	res, ok := <-jobChan
	fmt.Println(res, ok)

	//关闭一个已经关闭的，panic: close of closed channel
	//close(jobChan)

	//往一个close的channel写 panic: send on closed channel
	//jobChan <- "test"

	/*一个阻塞的chan，同一线程里,send 会造成死锁
	fatal error: all goroutines are asleep - deadlock!

	jobChan2 := make(chan string)
	jobChan2 <- "sss"
	res1, ok1 := <-jobChan2
	fmt.Println(res1, ok1)
	*/
} // }}}

//有缓冲，再buf没满的情况下非阻塞，满了阻塞
func bufChannel() { // {{{
	jobChan := make(chan string, 10)
	//close(jobChan)

	/** close 带缓冲的，如果无数据，返回false
	res1, ok := <-jobChan
	fmt.Println(res1, ok)
	return
	*/

	/*close的channel send panic*/
	jobChan <- "a"
	jobChan <- "b"
	jobChan <- "c"
	//一个关闭的非阻塞的channel进行读，会读完后直接返回false
	res, ok := <-jobChan
	fmt.Println(res, ok)

	/*读完后会跳出循环*/
	close(jobChan)
	for job := range jobChan {
		fmt.Println(job)
	}

	//返回false
	res2, ok2 := <-jobChan
	fmt.Println(res2, ok2)

	//关闭一个已经关闭的，panic: close of closed channel
	//close(jobChan)

	//往一个close的channel写 panic: send on closed channel
	//jobChan <- "test"

	//一个非阻塞的chan，同一线程里有buf的情况不会造成死锁
	jobChan2 := make(chan string, 1)
	jobChan2 <- "A"
	//缓冲不够的情况会导致fatal error: all goroutines are asleep - deadlock!
	//jobChan2 <- "B"
	res3, ok3 := <-jobChan2
	fmt.Println(res3, ok3)
} // }}}
