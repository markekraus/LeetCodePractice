package main

// 278. First Bad Version
// https://leetcode.com/problems/first-bad-version/

import "fmt"

var firstBad int = 4

func isBadVersion(version int) bool {
	return version >= firstBad
}

func firstBadVersion(n int) int {
	left := 1
	pos := 1
	for left <= n {
		mid := left + (n-left)/2
		if isBadVersion(mid) {
			pos = mid
			n = mid - 1
		} else {
			left = mid + 1
		}
	}
	return pos
}

func main() {
	fmt.Printf("Bad: %v Result: %v\n", firstBad, firstBadVersion(5))
	firstBad = 1
	fmt.Printf("Bad: %v Result: %v\n", firstBad, firstBadVersion(1))
	fmt.Printf("Bad: %v Result: %v\n", firstBad, firstBadVersion(2))
	firstBad = 13
	fmt.Printf("Bad: %v Result: %v\n", firstBad, firstBadVersion(20))
	firstBad = 2
	fmt.Printf("Bad: %v Result: %v\n", firstBad, firstBadVersion(3))
}
