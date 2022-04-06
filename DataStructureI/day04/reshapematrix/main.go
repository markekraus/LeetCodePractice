package main

// 566. Reshape the Matrix
// https://leetcode.com/problems/reshape-the-matrix/

import "fmt"

// time: O(n*m)
// space: O(1)
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if r*c != len(mat)*len(mat[0]) {
		return mat
	}
	result := make([][]int, r)
	cr := 0
	cc := 0
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
		for j := 0; j < c; j++ {
			if cc >= len(mat[0]) {
				cr++
				cc = 0
			}
			result[i][j] = mat[cr][cc]
			cc++
		}
	}
	return result
}

func main() {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Printf("%v\n", matrixReshape(matrix, 1, 4))
	fmt.Printf("%v\n", matrixReshape(matrix, 2, 4))
}
