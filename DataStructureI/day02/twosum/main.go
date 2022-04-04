package main

// 1. Two Sum
// https://leetcode.com/problems/two-sum/

import "fmt"

// time: O(n)
// space: O(n)
func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}
	result := make([]int, 2)
	visited := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, exists := visited[nums[i]]; exists {
			result[0] = i
			result[1] = index
			return result
		}
		visited[target-nums[i]] = i
	}
	return result
}

func main() {
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("%v\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("%v\n", twoSum([]int{3, 3}, 6))
}
