// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
package main

import "fmt"

func longestNonDuplicate(s string) int {
	uniqMap := make(map[byte]bool)
	left := 0
	maxLen := 0
	for right := 0; right < len(s); right++ {
		for uniqMap[s[right]] {
			delete(uniqMap, s[left])
			left++
		}
		maxLen = max(maxLen, right-left+1)
		uniqMap[s[right]] = true
	}
	return maxLen
}

func main() {
	testStr := "pwwkew"
	fmt.Println(longestNonDuplicate(testStr))
}
