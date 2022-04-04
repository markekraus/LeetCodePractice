package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	master := &ListNode{-1, head}
	parent := master
	end := head
	for i := 0; i < n; i++ {
		end = end.Next
	}
	for end != nil && end.Next != nil {
		parent = parent.Next
		end = end.Next
	}
	parent.Next = parent.Next.Next
	return master.Next
}
