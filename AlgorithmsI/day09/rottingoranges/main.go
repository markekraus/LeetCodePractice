package main

import (
	"container/list"
	"fmt"
)

type cell struct {
	row, col int
}

func orangesRotting(grid [][]int) int {
	q := list.New()
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				q.PushBack(cell{i, j})
			}
		}
	}
	for q.Len() > 0 {
		hasRot := false
		for i := q.Len(); i > 0; i-- {
			e := q.Front()
			o := q.Front().Value.(cell)
			q.Remove(e)
			// left
			if o.col-1 >= 0 && grid[o.row][o.col-1] == 1 {
				if !hasRot {
					hasRot = true
					ans++
				}
				grid[o.row][o.col-1] = 2
				q.PushBack(cell{o.row, o.col - 1})
			}
			// right
			if o.col+1 < len(grid[o.row]) && grid[o.row][o.col+1] == 1 {
				if !hasRot {
					hasRot = true
					ans++
				}
				grid[o.row][o.col+1] = 2
				q.PushBack(cell{o.row, o.col + 1})
			}
			// up
			if o.row-1 >= 0 && grid[o.row-1][o.col] == 1 {
				if !hasRot {
					hasRot = true
					ans++
				}
				grid[o.row-1][o.col] = 2
				q.PushBack(cell{o.row - 1, o.col})
			}
			// down
			if o.row+1 < len(grid) && grid[o.row+1][o.col] == 1 {
				if !hasRot {
					hasRot = true
					ans++
				}
				grid[o.row+1][o.col] = 2
				q.PushBack(cell{o.row + 1, o.col})
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return ans
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%0d, ", grid[i][j])
		}
		fmt.Println("")
	}
}

func main() {
	grid := [][]int{
		{2, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	grid = [][]int{
		{2, 1, 1},
		{1, 1, 1},
		{0, 1, 2},
	}
	fmt.Printf("%v", orangesRotting(grid))

}
