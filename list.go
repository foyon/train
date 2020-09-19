/**
* @File:          list.go
* @Aim:
* @Author:        foyon
* @Created Time:  2020-09-17
 */
package main

import (
	"fmt"
	//"os"
)

type ListNode struct {
	Data int
	Next *ListNode
}

func CreateNode(data int) *ListNode {
	newListNode := new(ListNode)
	newListNode.Data = data
	newListNode.Next = nil
	return newListNode
}

func InsertListHead(listNode *ListNode, data int) *ListNode {
	newNode := CreateNode(data)
	listNode.Next = newNode
	newNode.Next = listNode.Next
	return listNode
}

func InsertListTail(listNode *ListNode, data int) (newNode *ListNode) {
	curNode := CreateNode(data)
	preNode := new(ListNode)
	newNode = preNode

	for {

		if listNode.Next == nil {
			newNode.Next = listNode
			preNode = listNode
			listNode = listNode.Next
		} else {
			preNode.Next = curNode
			break
		}
	}
	return newNode.Next
}

func PrintfList(listNode *ListNode) {
	curNode := listNode
	for {
		if curNode.Next != nil {
			fmt.Printf("%+v\n", curNode.Next)
			curNode = curNode.Next
			break
		} else {
			fmt.Printf("55555555555555", curNode.Next)
			break
		}
	}
	return
}

func main() {
	clist := CreateNode(0)
	//fmt.Printf("%+v", clist)
	clist = InsertListHead(clist, 4)
	clist = InsertListHead(clist, 3)
	//PrintfList(clist)
	//clist = InsertListTail(clist, 6)

	for {
		if clist.Next != nil {
			fmt.Printf("data:%d", clist.Data)
			clist = clist.Next
		} else {
			break
		}
	}

}
