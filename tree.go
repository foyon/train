package main

import "fmt"

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func main() {
	fmt.Println("vim-go")
}

//完全二叉树，
func isBinTree(tree *Tree) bool {

	if tree == nil {
		return false
	}

	var root *Tree
	var left *Tree
	var right *Tree

	root = tree
	left = root.Left
	right = root.Right

	for {
		if root != nil && left != nil && right != nil {
			root = root.Left
		}

	}
}

func isAllTree(tree *Tree) bool {
	if tree.Left != nil && tree.Right != nil {
		return true
	}

	return false
}
