/**
* @File:          splice.go
* @Aim:           切片2N的方式扩容，超过1024 扩容25～35%
* @Author:        foyon
* @Created Time:  2020-10-14
 */
package main

import (
	"fmt"
)

func main() {
	testSlice()
}

func testSlice() { // {{{

	s1 := make([]int, 10)
	s2 := make([]int, 0)
	s1 = append(s1, 100)
	s2 = append(s2, 100)

	//s1[11]越界,panic
	fmt.Println(s1[0], s1[10], s1[11])
	fmt.Println("----------")
	fmt.Println(s2[0])
} // }}}
