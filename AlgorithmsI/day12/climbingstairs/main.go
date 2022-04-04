package main

import "fmt"

func climbStairs(n int) int {
	result := make([]int, n+2)
	result[0] = 1
	for i := 0; i < n; i++ {
		result[i+1] += result[i]
		result[i+2] += result[i]
	}
	return result[n]
}

func main() {
	// fmt.Printf("%v\n", climbStairs(1))
	// fmt.Printf("%v\n", climbStairs(2))
	// fmt.Printf("%v\n", climbStairs(3))
	fmt.Printf("%v\n", climbStairs(6))
}
