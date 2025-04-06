// https://leetcode.com/problems/container-with-most-water/description/
package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > result {
			result = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
