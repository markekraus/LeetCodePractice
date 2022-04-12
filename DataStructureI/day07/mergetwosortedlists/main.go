package main

// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

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
	var list1, list2 *ListNode
	list1 = &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 4}
	list2 = &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}
	fmt.Printf("%v\n", listVals(mergeTwoLists(list1, list2))) // [1 1 2 3 4 4]

	list1 = nil
	list2 = nil
	fmt.Printf("%v\n", listVals(mergeTwoLists(list1, list2))) // []

	list1 = nil
	list2 = &ListNode{Val: 0}
	fmt.Printf("%v\n", listVals(mergeTwoLists(list1, list2))) // [0]
}
