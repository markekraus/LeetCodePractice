package main

import "fmt"

func permuteTab2(nums []int) [][]int {
	result := make([][]int, 1)
	result[0] = make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		prev := result
		result = make([][]int, 0)
		for _, num := range nums {
			for _, e := range prev {
				if contains(e, num) {
					continue
				}
				temp := make([]int, len(e))
				copy(temp, e)
				temp = append(temp, num)
				result = append(result, temp)
			}
		}
	}
	return result
}

func permuteTab(nums []int) [][]int {
	table := make([][][]int, len(nums)+1)
	table[0] = make([][]int, 1)
	for i := 1; i <= len(nums); i++ {
		for _, num := range nums {
			for _, e := range table[i-1] {
				if contains(e, num) {
					continue
				}
				temp := make([]int, len(e))
				copy(temp, e)
				temp = append(temp, num)
				table[i] = append(table[i], temp)
			}
		}
	}
	return table[len(nums)]
}

func contains(s []int, v int) bool {
	for _, val := range s {
		if v == val {
			return true
		}
	}
	return false
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	result := make([][]int, 0)
	cur := make([]int, 0)
	visited := make([]bool, len(nums))
	makePermute(0, nums, cur, &visited, &result)
	return result
}

func makePermute(index int, nums, cur []int, visited *[]bool, result *[][]int) {
	if index == len(nums) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		*result = append(*result, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*visited)[i] {
			(*visited)[i] = true
			cur = append(cur, nums[i])
			makePermute(index+1, nums, cur, visited, result)
			cur = cur[:len(cur)-1]
			(*visited)[i] = false
		}
	}
}

func main() {
	fmt.Printf("%v\n", permuteTab2([]int{1, 2, 3}))
	fmt.Printf("%v\n", permuteTab2([]int{0, 1}))
	fmt.Printf("%v\n", permute([]int{1, 2, 3, 4}))
}
