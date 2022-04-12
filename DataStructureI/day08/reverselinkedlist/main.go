package main

// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/

import "fmt"

func listVals(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// time: O(n)
// space: O(n)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rest := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rest
}

func main() {
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next = &ListNode{Val: 4}
	list.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Printf("%v\n", listVals(reverseList(list))) // [5 4 3 2 1]

	list = &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	fmt.Printf("%v\n", listVals(reverseList(list))) // [2 1]

	list = nil
	fmt.Printf("%v\n", listVals(reverseList(list))) // []
}
