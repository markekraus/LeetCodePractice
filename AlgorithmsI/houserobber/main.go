package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	first := nums[0]
	mid := max(nums[0], nums[1])
	last := max(nums[2]+first, mid)
	for i := 3; i < len(nums); i++ {
		first = mid
		mid = last
		last = max(nums[i]+first, mid)
	}
	return last
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("%v\n", rob([]int{1, 2, 3, 1}))
	fmt.Printf("%v\n", rob([]int{2, 7, 9, 3, 1}))
	fmt.Printf("%v\n", rob([]int{2, 1, 1, 2}))
}
