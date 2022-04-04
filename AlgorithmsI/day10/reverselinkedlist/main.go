package main

// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/

import "fmt"

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
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	fmt.Printf("%v\n", reverseList(list))

	list = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	fmt.Printf("%v\n", reverseList(list))

	list = &ListNode{}
	fmt.Printf("%v\n", reverseList(list))
}
