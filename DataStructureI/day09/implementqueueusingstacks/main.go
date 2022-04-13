package main

// 232. Implement Queue using Stacks
// https://leetcode.com/problems/implement-queue-using-stacks/

import "fmt"

type MyQueue struct {
	input, output []int
}

// time: O(1)
// space: O(1)
func Constructor() MyQueue {
	return MyQueue{
		input:  make([]int, 0),
		output: make([]int, 0),
	}
}

// n = len(this.output)
// time: O(n)
// Space: O(n)
func (this *MyQueue) Push(x int) {
	for len(this.output) > 0 {
		this.input = append(this.input, this.output[len(this.output)-1])
		this.output = this.output[:len(this.output)-1]
	}
	this.input = append(this.input, x)
	for len(this.input) > 0 {
		this.output = append(this.output, this.input[len(this.input)-1])
		this.input = this.input[:len(this.input)-1]
	}
}

// time: O(1)
// space: O(1)
func (this *MyQueue) Pop() int {
	result := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]
	return result
}

// time: O(1)
// space: O(1)
func (this *MyQueue) Peek() int {
	return this.output[len(this.output)-1]
}

// time: O(1)
// space: O(1)
func (this *MyQueue) Empty() bool {
	return len(this.output) == 0
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Printf("%v %v %v\n", queue.Peek(), queue.Pop(), queue.Empty()) // 1 1 false
}
