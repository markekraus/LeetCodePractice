package main

// 35. Search Insert Position
// https://leetcode.com/problems/search-insert-position/

import "fmt"

// time: O(log(n))
// space: O(1)
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Printf("%v\n", searchInsert(nums, 5)) // 2
	fmt.Printf("%v\n", searchInsert(nums, 2)) // 1
	fmt.Printf("%v\n", searchInsert(nums, 7)) // 4
}
