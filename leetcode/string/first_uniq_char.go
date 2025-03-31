// https://leetcode.com/problems/first-unique-character-in-a-string
package main

import "fmt"

func firstUniqChar(s string) int {
	charMap := make(map[byte]int)
	indexMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := charMap[s[i]]; !ok {
			indexMap[s[i]] = i
		}
		charMap[s[i]]++
	}

	minIndex := len(s)
	for k := range charMap {
		if charMap[k] == 1 {
			idx, _ := indexMap[k]
			minIndex = min(minIndex, idx)
		}
	}
	if minIndex == len(s) {
		return -1
	}
	return minIndex
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
}
