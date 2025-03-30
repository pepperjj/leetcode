package main

type MyLinkedList struct {
	head *LNode
}

type LNode struct {
	val  int
	next *LNode
}

func Constructor() MyLinkedList {
	head := &LNode{}
	return MyLinkedList{
		head: head,
	}
}

func (l *MyLinkedList) Get(index int) int {
	node := l.head
	for i := 0; i <= index; i++ {
		if node.next == nil {
			return -1
		}
		node = node.next
	}
	return node.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	newNode := &LNode{
		val: val,
	}
	newNode.next = l.head.next
	l.head.next = newNode
}

func (l *MyLinkedList) AddAtTail(val int) {
	tailNode := &LNode{
		val: val,
	}
	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = tailNode
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	newNode := &LNode{
		val: val,
	}
	nextNode := node.next
	if nextNode != nil {
		newNode.next = nextNode
	}
	node.next = newNode
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	node := l.head

	for i := 0; i < index; i++ {
		node = node.next
	}
	if node.next == nil {
		return
	}
	nodeToDelete := node.next
	if nodeToDelete.next != nil {
		node.next = nodeToDelete.next
	} else {
		node.next = nil
	}
}

func main() {
	obj := Constructor()
	//obj.AddAtHead(1)
	//obj.AddAtTail(3)
	//obj.AddAtIndex(1, 2)
	//fmt.Println(obj.Get(1))
	obj.DeleteAtIndex(0)
	//fmt.Println(obj.Get(1))
}
