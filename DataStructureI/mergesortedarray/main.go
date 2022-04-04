package main

// 88. Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/

import "fmt"

// time: O(m+n)
// space: O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	size := m + n
	i := m - 1
	j := n - 1
	for size > 0 {
		size--
		if j >= 0 && i >= 0 && nums1[i] > nums2[j] {
			nums1[size] = nums1[i]
			i--
			continue
		}
		if j >= 0 && i >= 0 && nums2[j] > nums1[i] {
			nums1[size] = nums2[j]
			j--
			continue
		}
		if i >= 0 {
			nums1[size] = nums1[i]
			i--
			continue
		}
		nums1[size] = nums2[j]
		j--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Printf("%v\n", nums1)

	nums1 = []int{1}
	merge(nums1, 1, []int{}, 0)
	fmt.Printf("%v\n", nums1)

	nums1 = []int{0}
	merge(nums1, 0, []int{1}, 1)
	fmt.Printf("%v\n", nums1)
}
