// https://leetcode.com/problems/word-break/description/?envType=problem-list-v2&envId=oizxjoit&
package main

import "fmt"

// failed case []string{"aaaa", "aaa"}
func wordBreak_v1(s string, wordDict []string) bool {
	word := make(map[string]bool)
	for _, w := range wordDict {
		word[w] = true
	}
	result := make([]string, 0)
	start := 0
	end := 1
	seg := ""
	for end < len(s) {
		seg = s[start:end]
		if _, ok := word[seg]; !ok {
			end++
		} else {
			result = append(result, seg)
			start = end
			end = start + 1
		}
	}
	result = append(result, s[start:end])
	for _, v := range result {
		if _, ok := word[v]; !ok {
			return false
		}
	}

	return true
}

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	start := 0
	memo := make(map[int]bool)
	return wordDfs(s, start, wordMap, memo)

}

func wordDfs(s string, start int, wordMap map[string]bool, memo map[int]bool) bool {
	if start >= len(s) {
		return true
	}
	if v, ok := memo[start]; ok {
		return v
	}
	end := start + 1
	for end <= len(s) {
		if _, ok := wordMap[s[start:end]]; ok && wordDfs(s, end, wordMap, memo) {
			memo[start] = true
			return true
		}
		end++
	}
	memo[start] = false
	return false
}

func main() {
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	s := "catsandog"
	fmt.Println(wordBreak(s, wordDict))
}
