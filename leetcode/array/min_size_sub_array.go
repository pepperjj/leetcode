package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	slow, fast := 0, 0
	currentSum := 0
	// two pointers, slow, fast, sum is nums[slow:fast]
	// if sum(nums[slow:fast]) < target, fast++
	// else
	// while sum(nums[slow:fast]) > target, slow++, sum -= nums[slow], minLength = min(minLength, fast - slow + 1)
	// if until the end cannot find the matching array, return -1
	minLength := len(nums) + 1
	for fast < len(nums) {
		currentSum = currentSum + nums[fast]
		for currentSum >= target {
			minLength = min(minLength, fast-slow+1)
			currentSum = currentSum - nums[slow]
			slow++
		}
		fast++
	}
	if minLength > len(nums) {
		return 0
	}

	return minLength
}

func sum(numbers []int) int {
	result := 0
	for _, n := range numbers {
		result = result + n
	}
	return result
}

func main() {
	arr := []int{2, 3, 1, 2, 4, 3}
	target := 7

	fmt.Println(minSubArrayLen(target, arr))
}
