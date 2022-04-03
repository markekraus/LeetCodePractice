package main

import "fmt"

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[j], s[i] = s[i], s[j]
	}
}

// 344. Reverse String
// Write a function that reverses a string. The input string is given as an array of characters s.
// You must do this by modifying the input array in-place with O(1) extra memory.
// Example 1:

// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]

// Example 2:

// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]
// Constraints:

//     1 <= s.length <= 105
//     s[i] is a printable ascii character.

func main() {
	input := []byte("hello")
	reverseString(input)
	fmt.Printf("Expected: %v Result: %v\n", "olleh", string(input))
	input = []byte("Hannah")
	reverseString(input)
	fmt.Printf("Expected: %v Result: %v\n", "hannaH", string(input))
}
