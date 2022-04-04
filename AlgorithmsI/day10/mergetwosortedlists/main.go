package main

// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// time: O(m+n)
// space: O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	root := &ListNode{-1, nil}
	pointer := root
	for list1 != nil && list2 != nil {
		if list1.Val == list2.Val {
			pointer.Next = list1
			list1 = list1.Next
			pointer = pointer.Next
			pointer.Next = list2
			list2 = list2.Next
			pointer = pointer.Next
			continue
		}
		if list1.Val < list2.Val {
			pointer.Next = list1
			list1 = list1.Next
			pointer = pointer.Next
			continue
		}
		if list1.Val > list2.Val {
			pointer.Next = list2
			list2 = list2.Next
			pointer = pointer.Next
			continue
		}
	}
	if list1 != nil {
		pointer.Next = list1
	}
	if list2 != nil {
		pointer.Next = list2
	}
	return root.Next
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	fmt.Printf("%v\n", mergeTwoLists(list1, list2))

	list1 = &ListNode{}
	list2 = &ListNode{}
	fmt.Printf("%v\n", mergeTwoLists(list1, list2))

	list1 = &ListNode{}
	list2 = &ListNode{
		Val: 0,
	}
	fmt.Printf("%v\n", mergeTwoLists(list1, list2))
}
