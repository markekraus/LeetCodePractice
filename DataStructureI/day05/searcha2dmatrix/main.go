package main

// 74. Search a 2D Matrix
// https://leetcode.com/problems/search-a-2d-matrix/

import "fmt"

// time: O(1) to O(log n)
// space: O(1)
// Binary search 2-d array as a 1-d array
func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	if matrix[0][0] > target || matrix[n-1][m-1] < target {
		return false
	}
	l, r := 0, m*n-1
	for l <= r {
		mid := (l + r) >> 1
		val := matrix[mid/m][mid%m]
		if val == target {
			return true
		}
		if val > target {
			r = mid - 1
			continue
		}
		l = mid + 1
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	fmt.Printf("%v\n", searchMatrix(matrix, 3))  // true
	fmt.Printf("%v\n", searchMatrix(matrix, 13)) // false
}
