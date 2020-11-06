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
	//testSlice()
	test2()
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

func test2() {

	s1 := [][]int{}
	//var s2 [][]int
	//s3 := make([][]int{}, 0)

	s1 = append(s1, []int{10, 20})
	//s2[0] = append(s2[0], 20)

	fmt.Printf("%+v", s1)
	//fmt.Printf("%+v", s2)

}
