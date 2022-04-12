package main

// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/

import "fmt"

// time: O(n)
// space: O(1)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	lookup := make([]int, 26)
	for _, char := range s {
		lookup[char-'a']++
	}
	for _, char := range t {
		lookup[char-'a']--
		if lookup[char-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%v\n", isAnagram("anagram", "nagaram")) // true
	fmt.Printf("%v\n", isAnagram("car", "rat"))         // false
}
