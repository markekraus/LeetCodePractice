package main

import "fmt"

func reverseWords(s string) string {
	sbytes := []byte(s)
	space := []byte(" ")[0]
	for i, j := 0, 1; j <= len(sbytes); j++ {
		if j < len(sbytes) && sbytes[j] != space {
			continue
		}
		for k, l := i, j-1; k < l; k, l = k+1, l-1 {
			sbytes[k], sbytes[l] = sbytes[l], sbytes[k]
		}
		i = j + 1
	}
	return string(sbytes)
}

// Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

// Example 1:

// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"

// Example 2:

// Input: s = "God Ding"
// Output: "doG gniD"

// Constraints:

//     1 <= s.length <= 5 * 104
//     s contains printable ASCII characters.
//     s does not contain any leading or trailing spaces.
//     There is at least one word in s.
//     All the words in s are separated by a single space.

func main() {
	fmt.Printf("Expected: %v Result: %v\n", "s'teL ekat edoCteeL tsetnoc", reverseWords("Let's take LeetCode contest"))
	fmt.Printf("Expected: %v Result: %v\n", "doG gniD", reverseWords("God Ding"))
}
