package main

// 231. Power of Two
// https://leetcode.com/problems/power-of-two/

import "fmt"

// time: O(1)
// space: O(1)
func isPowerOfTwo(n int) bool {
	return n != 0 && n&(n-1) == 0
}

func main() {
	fmt.Printf("%v\n", isPowerOfTwo(1))
	fmt.Printf("%v\n", isPowerOfTwo(16))
	fmt.Printf("%v\n", isPowerOfTwo(3))
}
