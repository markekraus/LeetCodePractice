package main

import "fmt"

func getGcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return getGcd(b, a%b)
}

func getGcdWhile(a int, b int) int {
	for b != 0 {
		rem := a % b
		a = b
		b = rem
	}
	return a
}

func rotate(nums []int, k int) {
	length := len(nums)
	shift := k % length
	gcd := getGcd(shift, length)
	for i := 0; i < gcd; i++ {
		temp := nums[i]
		dst := i
		for {
			src := dst - shift
			if src < 0 {
				src = length + src
			}
			if src == i {
				break
			}
			nums[dst] = nums[src]
			dst = src
		}
		nums[dst] = temp
	}
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums1, 1)
	fmt.Printf("Expected: %v Result: %v\n", 1, nums1)
}
