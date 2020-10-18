/**
* @File:          string_repeat_cut.go
* @Aim:           "实现消除重复字符串的功能"
				  将长度相同大于3的字符消除 acbbbbcc => a
* @Author:        foyon
* @Created Time:  2020-10-18
*/
package main

import (
	"fmt"
	"strings"
)

func main() { // {{{
	str := "acbbbbbbccceddef"
	res := cut(str)
	fmt.Println(res)

	res2 := "abbbc"
	fmt.Println(res2)
} // }}}

func cut(str string) string { // {{{

	if len(str) < 3 {
		return str
	}

	newStr := str
	ss := strings.Split(str, "")
	scopy := str
	mark := ss[0]
	tmpMark := 1

	fmt.Println(ss)
	for k, s := range ss {
		if k > 0 {
			if s != mark {
				if tmpMark > 3 {
					newStr := ""
					newStr = scopy[0:k-tmpMark] + scopy[k:]
					return cut(newStr)
				} else {
					tmpMark = 1
					mark = s
				}
			} else if s == mark {
				tmpMark++
			}
		}
	}

	return newStr
} // }}}
