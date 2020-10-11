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

	list4 := quick_sort(list, 0, 6)
	fmt.Println(list4)
}

/*插入排序
time: O(N^2) | space: O(1)
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

//冒泡排序, TOP N操作，只排前N次即可
//time: O(N^2) | space: O(1)
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

/*time: nlogn  | space: logn*/
//取数组中最后一个值，比其大的放右边，比其小的放左边，
//直到每组分治为1的时候停止
//从左++，右--，直到交叉，比较与mark大小，选择交换位置
func quick_sort(list []int, left int, right int) []int { // {{{
	partIndex := part(list, left, right)
	quick_sort(list, left, partIndex-1)
	quick_sort(list, partIndex+1, right)
	return list
} // }}}

func part(list []int, left int, right int) int { // {{{
	//取数组最后一个值作为比较对象
	mark := list[right]
	smallIndex := left
	bigIndex := right - 1

	for {
		for smallIndex < right && list[smallIndex] <= mark {
			smallIndex++
		}

		for bigIndex > left && list[bigIndex] >= mark {
			bigIndex--
		}
		swap(list, smallIndex, bigIndex)
	}

	return smallIndex
} // }}}

func swap(list []int, left int, right int) []int { // {{{
	if list[left] > list[right] {
		list[left], list[right] = list[right], list[left]
	}

	return list
} // }}}

//归并排序 1、分治 2排序 3 归并（将左子++，mark，右子++ 进行循环顺序写入）
func merge_sort(list []int, start int, end int) {
	mid := int(math.Ceil((start + end) / 2))
	merge_sort(list, start, mid)
	merge_sort(list, mid+1, end)
	merge(list, start, mid, end)
}

func merge()

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
