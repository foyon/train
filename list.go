package main

import "fmt"

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func main() {
	fmt.Println("vim-go")
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode

	cur := head

	for cur != nil {
		nextNode := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextNode
	}

	return pre
}
