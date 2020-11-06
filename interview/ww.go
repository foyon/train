package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	arr := []int{1, 2, 3, 4, 5}
	printArr(arr)
}

func printArr(arr []int) {
	num := len(arr)
	//i := 0
	//j := 0
	//res := [][]int{}
	for a := 0; a < num; a++ {
		tmpRes := []int{}
		tmpRes = append(tmpRes, arr[a])
		for b := 0; b < num; b++ {
			if arr[a] == arr[b] {
				continue
			}
			tmpRes = append(tmpRes, arr[b])
		}

		fmt.Printf("%+v", tmpRes)
	}

}
