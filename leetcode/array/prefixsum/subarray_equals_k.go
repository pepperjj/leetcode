// https://leetcode.com/problems/subarray-sum-equals-k/description
package main

import "fmt"

// O(n^2)
func subarraySum1(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum == k {
			count++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// O(n)
func subarraySum(nums []int, k int) int {
	prefix := make(map[int]int)
	prefix[0] = 1
	count := 0
	prefixSum := 0
	for _, n := range nums {
		prefixSum += n
		if v, ok := prefix[prefixSum-k]; ok {
			count = count + v
		}
		prefix[prefixSum]++
	}
	return count
}

func main() {
	nums := []int{1, 2, 3} // [1, 3, 6]
	// [-1, -1, 1] -> k = 0
	// [0, -1, 1] -> k = 0
	k := 3
	fmt.Println(subarraySum(nums, k))
}
