// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
package main

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
