package main

// 704. Binary Search
// https://leetcode.com/problems/binary-search/

import "fmt"

func search(nums []int, target int) int {
	var left int16 = 0
	var right int16 = int16(len(nums) - 1)
	for left <= right {
		var mid int16 = left + (right-left)/2
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
	var nums = []int{-1, 0, 3, 5, 9, 12}
	result := search(nums, 9)
	fmt.Printf("Result %v\n", result)
	nums = []int{-1, 0, 3, 5, 9, 12}
	result = search(nums, 2)
	fmt.Printf("Result %v\n", result)
}
