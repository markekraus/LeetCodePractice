package main

import "fmt"

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	makeCombine(n, k, 1, cur, &res)
	return res

}

func makeCombine(n, k, start int, cur []int, res *[][]int) {
	if k == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := start; i <= n; i++ {
		cur = append(cur, i)
		makeCombine(n, k-1, i+1, cur, res)
		cur = cur[:len(cur)-1]
	}
}

func main() {
	fmt.Printf("%v", combine(4, 3))
}
