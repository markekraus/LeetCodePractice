package main

// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

import "fmt"

// time: O(n)
// space O(1)
func maxProfit(prices []int) int {
	total, local := 0, 0
	for i := 0; i < len(prices)-1; i++ {
		local = max(0, local+prices[i+1]-prices[i])
		total = max(total, local)
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("%v\n", maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Printf("%v\n", maxProfit([]int{7, 6, 4, 3, 1}))    // 0
}
