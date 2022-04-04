package main

import "fmt"

func searchInsert(nums []int, target int) int {
	var left int16 = 0
	var right int16 = int16(len(nums) - 1)
	var mid int16
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return int(mid)
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return int(left)
}

// Example 1:
// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:
// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:
// Input: nums = [1,3,5,6], target = 7
// Output: 4

func main() {
	var target int
	var expected int
	nums := []int{1, 3, 5, 6}

	target = 5
	expected = 2
	fmt.Printf("expected: %v result: %v\n", expected, searchInsert(nums, target))

	target = 2
	expected = 1
	fmt.Printf("expected: %v result: %v\n", expected, searchInsert(nums, target))

	target = 7
	expected = 4
	fmt.Printf("expected: %v result: %v\n", expected, searchInsert(nums, target))
}
