/**
* @File:          array.go
* @Aim:           矩阵相关
* @Author:        foyon
* @Created Time:  2020-11-20
 */
package main

import (
	"fmt"
	"sync"
)

var (
	Gmap = make(map[string]interface{})
	wg   = sync.WaitGroup{}
)

func main() {
	//testGmap()
	testSlice3()
	//fmt.Printf("%v", Gmap)
	data := [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}}
	res := sqSearch(data, -1)
	fmt.Printf("%v", res)

	sqPrint01(data)
	fmt.Printf("\n")
	sqPrint02(data)
}

func testSlice3() {
	data := [][][]int{{{0, 1, 2, 3}, {3, 4, 5}}, {{1, 2, 3}}}
	fmt.Printf("%v", data)
}

func testGmap() {
	Gmap["name"] = "foyon"
	Gmap["age"] = 28
}

func testSlice2() { // {{{
	// MXN的矩阵
	data := [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}}
	fmt.Printf("%v\n", data)

	m := len(data)
	n := len(data[0])
	fmt.Printf("%d X %d\n", m, n)

	for k, v := range data {
		fmt.Printf("%d\n", k)
		fmt.Printf("%v\n", v)
	}
} // }}}

//层层遍历法,如果效率要高 先根据X轴线2分，再根据X不变Y2分。
func sqSearch(data [][]int, pick int) bool { // {{{
	m := len(data)
	n := len(data[0])
	x := 0
	y := n - 1
	fmt.Printf("%dX%d", m, n)

	for x < m && y >= 0 {
		if pick == data[x][y] {
			return true
		} else if data[x][y] < n {
			x++
		} else {
			y--
		}
	}
	return false
} // }}}

//顺时针打印矩阵1
func sqPrint01(data [][]int) { // {{{
	xMax := len(data)
	yMax := len(data[0])
	x := 0
	y := 0

	//for x < xMax && y < yMax {
	for y < yMax {
		fmt.Printf("%d->", data[x][y])
		y++
	}

	y--
	x++

	for x < xMax {
		fmt.Printf("%d^", data[x][y])
		x++
	}

	x--
	y--

	for y > -1 {
		fmt.Printf("%d<-", data[x][y])
		y--
	}
	y = 0
	x--

	for x > 0 {
		fmt.Printf("%d^", data[x][y])
		x--
	}

	//}

} // }}}

//顺时针打印矩阵2
func sqPrint02(data [][]int) { // {{{
	xMax := len(data)
	yMax := len(data[0])
	x := 0
	y := 0

	for x < xMax && y < yMax {

		for y < yMax {
			fmt.Printf("%d->", data[x][y])
			y++
		}

		y--
		x++

		for x < xMax {
			fmt.Printf("%d^", data[x][y])
			x++
		}
		x--
		y--

		for y > -1 {
			fmt.Printf("%d<-", data[x][y])
			y--
		}

		y = 0
		x--

		for x > 0 {
			fmt.Printf("%d^", data[x][y])
			x--
		}

		if x == 0 && y == 0 {
			break
		}
	}

} // }}}
