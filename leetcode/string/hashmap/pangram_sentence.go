// https://leetcode.com/problems/check-if-the-sentence-is-pangram
package main

import (
	"fmt"
	"strings"
)

func checkIfPangram(sentence string) bool {
	letterMap := make(map[byte]int)
	s := strings.ToLower(sentence)

	for i := 0; i < len(s); i++ {
		if _, ok := letterMap[s[i]]; ok {
			letterMap[s[i]]++
		}
		letterMap[s[i]] = 1
	}
	return len(letterMap) == 26
}

func main() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	fmt.Println(checkIfPangram(sentence))
}
