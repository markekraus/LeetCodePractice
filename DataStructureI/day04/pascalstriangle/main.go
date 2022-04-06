package main

// 118. Pascal's Triangle
// https://leetcode.com/problems/pascals-triangle/

import "fmt"

// time: O(n2)
// space: O(1)
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for row := 0; row < numRows; row++ {
		result[row] = make([]int, row+1)
		result[row][0] = 1
		result[row][row] = 1
		for i := 1; row > 1 && i < row; i++ {
			result[row][i] = result[row-1][i-1] + result[row-1][i]
		}
	}
	return result
}

func main() {
	fmt.Printf("%v\n", generate(5))
	fmt.Printf("%v\n", generate(1))
}
