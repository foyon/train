package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abcba"
	str2 := "a"
	str3 := "ab"
	str4 := "abba"
	fmt.Printf("%s is %v \n", str1, isSource(str1))
	fmt.Printf("%s is %v \n", str2, isSource(str2))
	fmt.Printf("%s is %v \n", str3, isSource(str3))
	fmt.Printf("%s is %v \n", str4, isSource(str4))
}

func isSource(str string) bool {
	strSlice := strings.Split(str, "")
	lenStr := len(strSlice)

	//单个字符不算回文
	if lenStr < 2 {
		return false
	}

	for k, v := range strSlice {
		j := lenStr - k - 1
		if k < j && v != strSlice[j] {
			return false
		}
	}
	return true
}
