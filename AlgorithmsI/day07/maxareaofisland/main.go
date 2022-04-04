package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	result := 0
	visited := make(map[[2]int]bool)
	maxR := len(grid)
	maxC := len(grid[0])
	for row := 0; row < maxR; row++ {
		for col := 0; col < maxC; col++ {
			run := 0
			recMaxAreaOfIsland(grid, &run, row, col, maxR, maxC, visited)
			if run > result {
				result = run
			}
		}
	}
	return result
}

func recMaxAreaOfIsland(grid [][]int, run *int, sr int, sc int, maxR int, maxC int, visited map[[2]int]bool) {
	if visited[[2]int{sr, sc}] || sr < 0 || sr >= maxR || sc < 0 || sc >= maxC || grid[sr][sc] == 0 {
		return
	}
	visited[[2]int{sr, sc}] = true
	*run++
	recMaxAreaOfIsland(grid, run, sr+1, sc, maxR, maxC, visited)
	recMaxAreaOfIsland(grid, run, sr-1, sc, maxR, maxC, visited)
	recMaxAreaOfIsland(grid, run, sr, sc+1, maxR, maxC, visited)
	recMaxAreaOfIsland(grid, run, sr, sc-1, maxR, maxC, visited)
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Printf("%v\n", maxAreaOfIsland(grid))
	grid = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Printf("%v\n", maxAreaOfIsland(grid))
}
