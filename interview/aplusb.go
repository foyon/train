package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {

	flag.Parse()
	a := flag.Arg(0)
	b := flag.Arg(1)
	fmt.Printf("%v, %v", a, b)
	ai, _ := strconv.Atoi(a)
	bi, _ := strconv.Atoi(b)
	r := abPlus(ai, bi)
	fmt.Printf("%v", r)
	//return r
}

func abPlus(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	for b != 0 {
		tp := (a & b) << 1
		a = a ^ b
		b = tp
	}

	return a
}
