package main

import "fmt"

type Stack struct {
	nums []int
}

func Constructor() *Stack {
	return &Stack{
		nums: []int{},
	}
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	res := s.nums[s.Len()-1]
	s.nums = s.nums[:s.Len()-1]
	return res
}

func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

func (s *Stack) Len() int {
	return len(s.nums)
}

func (s *Stack) IsEmpty() bool {
	return len(s.nums) == 0
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop()) // 2
	stack.Push(3)
	fmt.Println(stack.Pop())     // 3
	fmt.Println(stack.Pop())     // 1
	fmt.Println(stack.Pop())     // -1
	fmt.Println(stack.IsEmpty()) // true

}
