/**
* @File:          part2search.go
* @Aim:           简单实现2分搜索
* @Author:        foyon
* @Created Time:  2020-10-20
 */
package main

import (
	"fmt"
)

func main() {
	//fmt.Println("vim-go")
	s := []int{1, 2, 3, 7, 9}
	findRes := search(s, 12)
	fmt.Println(findRes)
}

func search(s []int, i int) bool { // {{{

	if len(s) < 1 {
		return false
	}

	mid := len(s) / 2
	fmt.Println(mid)

	left := 0
	right := len(s) - 1

	if i == s[left] || i == s[right] || i == s[mid] {
		return true
	}

	if i < s[mid] {
		return search(s[left:mid], i)
	} else {
		return search(s[mid+1:], i)
	}
} // }}}
