package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	result := uint32(0)
	power := uint32(31)
	for num != 0 {
		result += (num & 1) << power
		num = num >> 1
		power -= 1
	}
	return result
}

func main() {
	fmt.Printf("%v\n", reverseBits(43261596))
	fmt.Printf("%v\n", reverseBits(4294967293))
}
