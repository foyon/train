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
	testSlice()
}

//测试切片扩容
//底层是数组,内存是连续的
func testSlice() {
	str := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Printf("%p\n", &str)
	str = append(str, "i")
	fmt.Printf("%p\n", &str)

	//2len方式进行扩容，地址不会改变
	i := []int{1, 2, 3, 4}
	fmt.Printf("%p\n", &i)
	i = append(i, 5)
	fmt.Printf("%p\n", &i)
}
