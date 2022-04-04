package main

// 217. Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/

import "fmt"

// time: O(log(n))
// space: O(1)
func containsDuplicatNaive(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// time worst: O(n*log(n))
// time average: O(n)
// space: O(n)
func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	visited := make(map[int]struct{}, len(nums))
	visited[nums[0]] = struct{}{}
	for i := 1; i < len(nums); i++ {
		if _, exists := visited[nums[i]]; exists {
			return true
		}
		visited[nums[i]] = struct{}{}
	}
	return false
}

func main() {
	fmt.Printf("%v\n", containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Printf("%v\n", containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
