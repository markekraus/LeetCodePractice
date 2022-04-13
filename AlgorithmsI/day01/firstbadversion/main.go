package main

// 278. First Bad Version
// https://leetcode.com/problems/first-bad-version/

import "fmt"

var firstBad int = 4

func isBadVersion(version int) bool {
	return version >= firstBad
}

// time: O(log(n))
// space: O(1)
func firstBadVersion(n int) int {
	left := 1
	right := n
	result := 1
	for left <= right {
		mid := (left + right) >> 1
		if isBadVersion(mid) {
			result = mid
			right = mid - 1
			continue
		}
		left = mid + 1
	}
	return result
}

func main() {
	firstBad = 4
	fmt.Printf("%v\n", firstBadVersion(5)) // 4
	firstBad = 1
	fmt.Printf("%v\n", firstBadVersion(1)) // 1
	fmt.Printf("%v\n", firstBadVersion(2)) // 1
	firstBad = 13
	fmt.Printf("%v\n", firstBadVersion(20)) // 13
	firstBad = 2
	fmt.Printf("%v\n", firstBadVersion(3)) // 2
}
