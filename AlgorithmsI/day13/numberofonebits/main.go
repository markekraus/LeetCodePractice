package main

// 191. Number of 1 Bits
// https://leetcode.com/problems/number-of-1-bits/

import (
	"fmt"
	"math"
)

// time:
// space:
func hammingWeight(num uint32) int {
	var result int
	for result = 0; num != 0; result++ {
		num &= (num - 1)
	}
	return result
}

func main() {
	var n uint32
	n += 1 << 4
	n += 1 << 1
	n += 1
	fmt.Printf("%v\n", hammingWeight(n))

	n = 0
	n += 1 << 7
	fmt.Printf("%v\n", hammingWeight(n))

	n = math.MaxUint32 - 2
	fmt.Printf("%v\n", hammingWeight(n))
}
