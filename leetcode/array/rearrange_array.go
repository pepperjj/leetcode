package main

import "fmt"

func rearrangeArray(nums []int) []int {
	result := make([]int, 0, len(nums))
	pos := make([]int, 0)
	neg := make([]int, 0)

	for _, elem := range nums {
		if elem < 0 {
			neg = append(neg, elem)
		} else {
			pos = append(pos, elem)
		}
	}

	for i := range pos {
		result = append(result, pos[i])
		result = append(result, neg[i])
		i++
	}

	return result
}

func main() {
	arr := []int{3, 1, -2, -5, 2, -4}
	result := rearrangeArray(arr)
	for _, x := range result {
		fmt.Printf("%d ", x)
	}
}
