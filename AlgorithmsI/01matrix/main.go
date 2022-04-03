package main

import (
	"fmt"
	"math"
)

func updateMatrix(mat [][]int) [][]int {
	result := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		result[i] = make([]int, len(mat[i]))
		for j := 0; j < len(mat[i]); j++ {
			result[i][j] = math.MaxInt - 1000
		}
	}
	for row := 0; row < len(mat); row++ {
		for col := 0; col < len(mat[row]); col++ {
			if mat[row][col] == 0 {
				result[row][col] = 0
				continue
			}
			if row > 0 {
				result[row][col] = min(result[row][col], result[row-1][col]+1)
			}
			if col > 0 {
				result[row][col] = min(result[row][col], result[row][col-1]+1)
			}
		}
	}
	for row := len(mat) - 1; row >= 0; row-- {
		for col := len(mat[row]) - 1; col >= 0; col-- {
			if row < len(mat)-1 {
				result[row][col] = min(result[row][col], result[row+1][col]+1)
			}
			if col < len(mat[row])-1 {
				result[row][col] = min(result[row][col], result[row][col+1]+1)
			}
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Printf("%v\n", updateMatrix(matrix))
	matrix = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Printf("%v\n", updateMatrix(matrix))
}
