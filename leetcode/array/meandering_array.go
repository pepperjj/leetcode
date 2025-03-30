package main

import (
	"fmt"
	"sort"
)

func meanderingArray(arr []int) []int {
	sort.Slice(arr, func(x, y int) bool {
		return arr[x] < arr[y]
	})

	result := make([]int, 0, len(arr))
	left, right := 0, len(arr)-1
	for left <= right {
		if left == right {
			result = append(result, arr[left])
		} else {
			result = append(result, arr[right])
			result = append(result, arr[left])
		}
		left++
		right--
	}
	return result
}

func main() {
	arr := []int{-1, 1, 2, 3, -5} // -5, -1, 1, 2, 3
	result := meanderingArray(arr)
	for _, x := range result {
		fmt.Println(x)
	}
}
