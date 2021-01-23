package main

import "fmt"

func main1() {
	var x string = ""
	if x == "" {
		x = "default"
	}
	fmt.Println(x)
	fan(1, 2)
	test3("test", "fan")

}

func test3(s string, s2 string) {

}

func fan(i int, i2 int) {

}

const (
	a  = iota // 0
	b  = iota //1
	bb        //2
)
const (
	name = "menglu"
	c    = iota //1
	d    = iota //2
)

func main2() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(bb)
	fmt.Println(c)
	fmt.Println(d)
}

func main() {
	ch := make(chan int)
	ch <- 2
	select {
	case i := <-ch:
		println(i)

	}
}
