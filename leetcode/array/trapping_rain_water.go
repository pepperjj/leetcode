package main

import "fmt"

func trap(height []int) int {
	stack := make([]int, 0)

	result := 0
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				continue
			}
			result = result + (i-stack[len(stack)-1]-1)*(min(height[i], height[stack[len(stack)-1]])-height[last])
		}

		stack = append(stack, i)
	}
	return result
}

func main() {
	// stack store the index
	// if ele > stack.top, pop stack and calculate area

	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}
