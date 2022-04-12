package main

// 387. First Unique Character in a String
// https://leetcode.com/problems/first-unique-character-in-a-string/

import "fmt"

// time: O(n)
// space: O(1)
func firstUniqChar(s string) int {
	lookup := make([]int, 26)
	for _, char := range s {
		lookup[char-'a']++
	}
	for i, char := range s {
		if lookup[char-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Printf("%v\n", firstUniqChar("loveleetcode")) // 2
	fmt.Printf("%v\n", firstUniqChar("loveleetcode")) // 2
}
