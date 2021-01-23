package main

import "fmt"

func main() {
	a := 1
	b := 2

	defer calc(a, calc(a, b))
	a = 0
	defer calc(a, calc(a, b))
}

func calc(x, y int) int {
	fmt.Println(x, y, x+y)
	return x + y
}
