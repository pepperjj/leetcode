// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
}

func Constructor() MyLinkedList {
	head := &ListNode{}
	return MyLinkedList{
		head: head,
	}
}

func (l *MyLinkedList) AddAtTail(val int) {
	tailNode := &ListNode{
		Val: val,
	}
	node := l.head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = tailNode
}

func (l *MyLinkedList) PrintLinkedList() {
	node := l.head
	for node.Next != nil {
		node = node.Next
		fmt.Println(node.Val)
	}
}

func (l *MyLinkedList) Len() int {
	length := 0
	node := l.head
	for node.Next != nil {
		node = node.Next
		length++
	}

	return length
}

func (l *MyLinkedList) removeNthFromEnd(n int) *ListNode {
	node := l.head
	for i := 0; node.Next != nil && i < l.Len()-n; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	return l.head
}

func main() {
	list := Constructor()
	list.AddAtTail(1)
	//fmt.Println(list.Len())
	list.removeNthFromEnd(1)
	list.PrintLinkedList()
}
