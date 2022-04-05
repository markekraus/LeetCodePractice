package main

// 350. Intersection of Two Arrays II
// https://leetcode.com/problems/intersection-of-two-arrays-ii/

import "fmt"

// n = min(len(nums1),len(nums2))
// m = max(len(nums1),len(nums2))
// time: O(n+m)
// space: O(m)
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	result := make([]int, 0, len(nums1))
	lookup := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		lookup[nums1[i]] += 1
	}
	for i := 0; i < len(nums2); i++ {
		if count, exists := lookup[nums2[i]]; exists && count > 0 {
			lookup[nums2[i]] -= 1
			result = append(result, nums2[i])
		}
	}
	return result
}

func main() {
	fmt.Printf("%v\n", intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Printf("%v\n", intersect([]int{4, 9, 5}, []int{9, 4}))
}
