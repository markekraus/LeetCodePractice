package main

import "fmt"

func sortedSquares(nums []int) []int {
	length := int16(len(nums))
	result := make([]int, length)
	var ppos int16
	for ppos = 0; ppos < length; ppos++ {
		if nums[ppos] >= 0 {
			break
		}
	}
	var npos int16 = ppos - 1
	var index int16 = -1
	for npos >= 0 && ppos < length {
		index++
		if nums[npos]*nums[npos] < nums[ppos]*nums[ppos] {
			result[index] = nums[npos] * nums[npos]
			npos--
			continue
		}
		result[index] = nums[ppos] * nums[ppos]
		ppos++
	}
	for npos >= 0 {
		index++
		result[index] = nums[npos] * nums[npos]
		npos--
	}
	for ppos < length {
		index++
		result[index] = nums[ppos] * nums[ppos]
		ppos++
	}
	return result
}

// Example 1:

// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].

// Example 2:

// Input: nums = [-7,-3,2,3,11]
// Output: [4,9,9,49,121]

func main() {
	fmt.Printf("Expected: %v Result %v\n", []int{0, 1, 9, 16, 100}, sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Printf("Expected: %v Result %v\n", []int{4, 9, 9, 49, 121}, sortedSquares([]int{-7, -3, 2, 3, 11}))
}
