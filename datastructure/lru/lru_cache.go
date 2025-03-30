package main

import "fmt"

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	cacheMap map[int]*Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cacheMap: make(map[int]*Node, capacity),
		head:     &Node{},
		tail:     &Node{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cacheMap[key]; ok {
		c.removeNode(node)
		c.addNodeToHead(node)
		return node.val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.cacheMap[key]; ok {
		node.val = value
		c.removeNode(node)
		c.addNodeToHead(node)
	} else {
		if c.Len() >= c.capacity {
			tail := c.removeNodeFromTail()
			delete(c.cacheMap, tail.key)
		}
		newNode := &Node{
			key: key,
			val: value,
		}
		c.addNodeToHead(newNode)
		c.cacheMap[key] = newNode
	}
}

func (c *LRUCache) addNodeToHead(n *Node) {
	n.prev = c.head
	n.next = c.head.next
	c.head.next.prev = n
	c.head.next = n
}

func (c *LRUCache) removeNode(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) removeNodeFromTail() *Node {
	tail := c.tail.prev
	c.removeNode(tail)
	return tail
}

func (c *LRUCache) Len() int {
	return len(c.cacheMap)
}

func main() {
	lru := Constructor(2)

	lru.Put(1, 1)           // cache is {1=1}
	lru.Put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(lru.Get(1)) // return 1
	lru.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lru.Get(2)) // returns -1 (not found)
	lru.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lru.Get(1)) // return -1 (not found)
	fmt.Println(lru.Get(3)) // return 3
	fmt.Println(lru.Get(4)) // return 4
}
