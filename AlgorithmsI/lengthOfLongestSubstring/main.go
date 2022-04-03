package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	chars := []rune(s)
	tracker := make(map[rune]bool)
	length := 0
	for i := 0; i < len(chars); i++ {
		for j := 0; i+j < len(chars); j++ {
			if tracker[chars[i+j]] {
				tracker = make(map[rune]bool)
				break
			}
			tracker[chars[i+j]] = true
			if j+1 > length {
				length = j + 1
			}
		}
	}
	return length
}

func main() {
	fmt.Printf("Expected: %v Result %v", 3, lengthOfLongestSubstring("abcabcbb"))
}
