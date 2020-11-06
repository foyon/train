/**
* @File:          recover.go
* @Aim:
* @Author:        foyon
* @Created Time:  2020-11-05
 */

/*
 利用go tool 分析程序性能
 go build -o recover recover.go
 go tool compile -S ./recover

 go run recover.go 2 > recover_trace.out
 go tool trace recover_trace.out

 pprof

 gdb
 dv\attach\delve
 https://books.studygolang.com/advanced-go-programming-book/ch3-asm/ch3-09-debug.html
*/
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	_ "os"
	_ "runtime/trace"
)

func main() {
	//trace.Start(os.Stderr)
	//defer trace.Stop()

	//http://127.0.0.1:6060/debug/pprof/
	http.ListenAndServe("0.0.0.0:6060", nil)
	//testRecover()
	//testGoroutine()

}

/**
*通过recover捕捉异常，只能嵌套一层，recover必须在defer中才能调研
 */
func testRecover() { // {{{

	//panic("execept here 01 \n")
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("pannic is %v", p)
		}
	}()

	panic("execept here 02 \n")
} // }}}

func testGoroutine() { // {{{
	ch := make(chan string)
	go func() {
		ch <- "exec"
	}()
} // }}}
