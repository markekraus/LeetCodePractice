package main

// 383. Ransom Note
// https://leetcode.com/problems/ransom-note/

import "fmt"

// n = len(ransomNote); m = len(magazine)
// time: O(n+m)
// space: O(1)
func canConstruct(ransomNote string, magazine string) bool {
	lookup := make([]int, 26)
	for _, char := range magazine {
		lookup[char-'a']++
	}
	for _, char := range ransomNote {
		lookup[char-'a']--
		if lookup[char-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%v\n", canConstruct("a", "b"))    // false
	fmt.Printf("%v\n", canConstruct("aa", "ab"))  // false
	fmt.Printf("%v\n", canConstruct("aa", "aab")) // true
}
