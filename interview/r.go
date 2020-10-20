package main

import "fmt"

func main() {
    str := []rune("hello")
    str[0] = rune("x")
    fmt.Println(str)
}
