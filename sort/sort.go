/**
* @File:          insert.go
* @Aim:
* @Author:        foyon
* @Created Time:  2020-09-22
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	var list = []int{1, 3, 2, 4, 7, 9, 10}
	list2 := insert_sort(list)
	res := search(list2, -1, 0, len(list2)-1)
	fmt.Println(res)
	res2 := searchFor(list2, 2)
	fmt.Println(res2)

	list3 := bubble_sort(list)
	fmt.Println(list3)

}

/*插入排序
n^2 时间复杂度
从第二个元素开始，每个元素和前面的每一个元素对比，找到合适位置插入
*/
func insert_sort(list []int) []int { // {{{
	num := len(list)

	if num < 1 {
		return list
	}

	for i := 1; i < num; i++ {
		for j := i - 1; j > -1; j-- {
			if list[i] < list[j] {
				tmp := list[j]
				list[j] = list[i]
				list[i] = tmp
			}
		}
	}

	fmt.Println(list)
	return list
} // }}}

/*冒泡排序, TOP N操作，只排前N次即可*/
func bubble_sort(list []int) []int { // {{{
	lenth := len(list)

	for i := 0; i < lenth; i++ {
		for j := 0; j < lenth-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
} // }}}

/*2分查找，取中间值，偏移左右索引--递归版*/
func search(list []int, p int, left int, right int) bool { // {{{
	mid := int(math.Ceil(float64(left+right) / 2))

	if p == list[mid] || p == list[left] || p == list[right] {
		return true
	}

	if right <= left {
		return false
	}

	if p > list[mid] {
		left = mid + 1
	} else {
		right = mid - 1
	}

	return search(list, p, left, right)
} // }}}

/*2分查找for 循环版，因golang不支持while*/
func searchFor(list []int, p int) bool { // {{{
	left := 0
	right := len(list) - 1

	for {
		if right <= left {
			return false
		}

		mid := int(float64(left+right) / 2)

		if list[mid] == p || list[left] == p || list[right] == p {
			return true
		}

		if list[mid] > p {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
} // }}}
