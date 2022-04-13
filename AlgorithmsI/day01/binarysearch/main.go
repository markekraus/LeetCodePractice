package main

// 704. Binary Search
// https://leetcode.com/problems/binary-search/

import "fmt"

// time: O(log(n))
// space: O(1)
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return int(mid)
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	var nums []int
	nums = []int{-1, 0, 3, 5, 9, 12}
	fmt.Printf("%v\n", search(nums, 9)) // 4
	nums = []int{-1, 0, 3, 5, 9, 12}
	fmt.Printf("%v\n", search(nums, 2)) // -1
}
