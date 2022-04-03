package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	minlen := make([]int, len(triangle))
	copy(minlen, triangle[len(triangle)-1])
	for layer := len(triangle) - 2; layer >= 0; layer-- {
		for i := 0; i <= layer; i++ {
			minlen[i] = min(minlen[i], minlen[i+1]) + triangle[layer][i]
		}
	}
	return minlen[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	t := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Printf("%v\n", minimumTotal(t))
	// t = [][]int{
	// 	{-1},
	// 	{2, 3},
	// 	{1, -1, -3},
	// }
	// fmt.Printf("%v\n", minimumTotal(t))
}
