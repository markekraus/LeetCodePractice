package main

// 203. Remove Linked List Elements
// https://leetcode.com/problems/remove-linked-list-elements/
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
func removeElements(head *ListNode, val int) *ListNode {
	root := &ListNode{-1, head}
	previous := root
	for head != nil {
		if head.Val == val {
			previous.Next = head.Next
			head = previous.Next
			continue
		}
		previous = head
		head = head.Next
	}
	return root.Next
}

func main() {
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 6}
	list.Next.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next.Next = &ListNode{Val: 4}
	list.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	list.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	fmt.Printf("%v\n", listVals(removeElements(list, 6))) // [1 2 3 4 5]

	list = nil
	fmt.Printf("%v\n", listVals(removeElements(list, 1))) // []

	list = &ListNode{Val: 7}
	list.Next = &ListNode{Val: 7}
	list.Next.Next = &ListNode{Val: 7}
	list.Next.Next.Next = &ListNode{Val: 7}
	fmt.Printf("%v\n", listVals(removeElements(list, 7))) // []
}
