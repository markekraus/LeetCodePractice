package main

import "fmt"

func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}

func main() {
	fmt.Printf("%v\n", singleNumber([]int{2, 2, 1}))
	fmt.Printf("%v\n", singleNumber([]int{4, 1, 2, 1, 2}))
}
