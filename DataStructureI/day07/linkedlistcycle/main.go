package main

// 141. Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// time: O(n)
// space: O(1)
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {
	list := &ListNode{Val: 3}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 0}
	list.Next.Next.Next = &ListNode{Val: 4}
	list.Next.Next.Next.Next = list.Next
	fmt.Printf("%v\n", hasCycle(list)) // true

	list = &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = list
	fmt.Printf("%v\n", hasCycle(list)) // true

	list = &ListNode{Val: 1}
	fmt.Printf("%v\n", hasCycle(list)) // false
}
