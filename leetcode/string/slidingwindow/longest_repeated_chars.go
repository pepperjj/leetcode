// https://leetcode.com/problems/longest-repeating-character-replacement/description/
package main

import "fmt"

func characterReplacement(s string, k int) int {
	maxLen := 0
	maxLetterCount := 0
	freqMap := make(map[byte]int)
	i := 0
	for j := i; j < len(s); j++ {
		freqMap[s[j]]++
		maxLetterCount = max(maxLetterCount, freqMap[s[j]])

		if j-i+1-maxLetterCount > k {
			freqMap[s[i]]--
			i = i + 1
		}

		maxLen = max(maxLen, j-i+1)
	}

	return maxLen
}

func main() {
	s := "ABAB"
	k := 0
	fmt.Println(characterReplacement(s, k))
}
