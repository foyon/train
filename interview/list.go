package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
	//参数：返回值
	//Add  func() (l *ListNode)
}

func main() {
	head := buildList()
	reverseHead := reverseList(head)
	fmt.Printf("%+v", reverseHead)
	fmt.Printf("%+v", reverseHead.Next)
}

func buildList() *ListNode { // {{{
	L1 := &ListNode{}
	L1.Val = 1
	L2 := &ListNode{}
	L2.Val = 2
	L3 := &ListNode{}
	L3.Val = 3
	L4 := &ListNode{}
	L4.Val = 4

	L1.Next = L2
	L2.Next = L3
	L3.Next = L4

	return L1
} // }}}

/*单链反转*/
func reverseList(head *ListNode) *ListNode { // {{{
	var pre *ListNode
	cur := head

	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}

	return pre
} // }}}
