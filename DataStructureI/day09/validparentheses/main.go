package main

// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

import "fmt"

// n = len(s)
// time: O(n)
// space: O(n)
func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	var stack []rune
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		top := stack[len(stack)-1]
		if top == '(' && char == ')' || top == '{' && char == '}' || top == '[' && char == ']' {
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Printf("%v\n", isValid("()"))     // true
	fmt.Printf("%v\n", isValid("()[]{}")) // true
	fmt.Printf("%v\n", isValid("(]"))     // false
}
