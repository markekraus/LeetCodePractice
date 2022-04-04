package main

import "fmt"

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	for count < len(nums) {
		nums[count] = 0
		count++
	}
}

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

// Example 1:

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Example 2:

// Input: nums = [0]
// Output: [0]

// Constraints:

//     1 <= nums.length <= 104
//     -231 <= nums[i] <= 231 - 1

// Follow up: Could you minimize the total number of operations done?

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Printf("Expected: %v Result: %v\n", []int{1, 3, 12, 0, 0}, nums)
	nums = []int{0}
	moveZeroes(nums)
	fmt.Printf("Expected: %v Result: %v\n", []int{0}, nums)
}
