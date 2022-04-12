package main

// 83. Remove Duplicates from Sorted List
// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
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
// space: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head
}

func main() {
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 1}
	list.Next.Next = &ListNode{Val: 2}
	fmt.Printf("%v\n", listVals(deleteDuplicates(list))) // [1 2]

	list = &ListNode{Val: 1}
	list.Next = &ListNode{Val: 1}
	list.Next.Next = &ListNode{Val: 2}
	list.Next.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next.Next = &ListNode{Val: 3}
	fmt.Printf("%v\n", listVals(deleteDuplicates(list))) // [1 2 3]
}
