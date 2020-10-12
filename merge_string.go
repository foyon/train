/**
* @File:          gd.go
* @Aim:           一道面试题，给定一个字符串按空格切割，用goroutine 并行处理
				  典型并行处理，main进程汇总经典案例，实际项目开发运用居多
* @Author:        foyon
* @Created Time:  2020-10-12
*/
package main

import (
	"fmt"
	//"sort"
	"strings"
	"sync"
	"time"
)

//给定一组字符串切片，各切片按空格切割，最终进行归并统计相同字符串个数
func main() {
	//getCount("abc aa bb cc abc aa bb")
	calStr := []string{"abc aa bb", "cc dd eef", "aa bb cc gg"}
	calLen := len(calStr)
	outChan := make(chan map[string]int, calLen)
	fishChan := make(chan struct{})
	errChan := make(chan error)

	wg := sync.WaitGroup{}
	wg.Add(calLen)

	for _, s := range calStr {
		go func(wg *sync.WaitGroup, str string) {
			defer wg.Done()
			mapStr := getCount(str)
			outChan <- mapStr
		}(&wg, s)
	}

	go func() {
		wg.Wait()
		close(fishChan)
	}()

	tmpMap := make([]map[string]int, 0)

Loop:
	for {
		select {
		case val := <-outChan:
			tmpMap = append(tmpMap, val)
		case <-fishChan:
			break Loop
		case err := <-errChan:
			fmt.Printf("%v", err)
			break Loop
		case <-time.After(100000 * time.Millisecond):
			break Loop
		}

	}

	res := make(map[string]int)
	for _, m := range tmpMap {
		for k, v := range m {
			res[k] += v
		}
	}

	fmt.Printf("%+v", res)

}

// abcccc ddd */
func getCount(s string) map[string]int { // {{{
	str := strings.Split(s, " ")
	res := make(map[string]int)

	for _, v := range str {
		res[v]++
	}

	//sort.String(res)
	//fmt.Printf("%v", res)
	return res
} // }}}
