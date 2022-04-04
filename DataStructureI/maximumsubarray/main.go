package main

import (
	"fmt"
)

// time: O(n)
// space: O(n)
// Dynamic Programing solution
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	table := make([]int, len(nums))
	table[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		table[i] = max(nums[i]+table[i-1], nums[i])
	}
	for i := 1; i < len(table); i++ {
		table[i] = max(table[i], table[i-1])
	}
	return table[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("%v\n", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Printf("%v\n", maxSubArray([]int{1}))
	fmt.Printf("%v\n", maxSubArray([]int{5, 4, -1, 7, 8}))
}
