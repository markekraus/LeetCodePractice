package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// solution 1
// func middleNode(head *ListNode) *ListNode {
// 	nodes := []*ListNode{}
// 	current := head
// 	nodes = append(nodes, head)
// 	length := 1
// 	for current.Next != nil {
// 		nodes = append(nodes, current.Next)
// 		current = current.Next
// 		length++
// 	}
// 	mid := length / 2
// 	return nodes[mid]
// }

func middleNode(head *ListNode) *ListNode {
	fastptr := head
	for fastptr != nil && fastptr.Next != nil {
		fastptr = fastptr.Next.Next
		head = head.Next
	}
	return head
}

// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.

// Example 1:
// Input: head = [1,2,3,4,5]
// Output: [3,4,5]
// Explanation: The middle node of the list is node 3.

// Example 2:
// Input: head = [1,2,3,4,5,6]
// Output: [4,5,6]
// Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.

// Constraints:
//     The number of nodes in the list is in the range [1, 100].
//     1 <= Node.val <= 100

func main() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Printf("Expected: %v Result: %+v\n", "[3,4,5]", middleNode(&head))

	head = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	fmt.Printf("Expected: %v Result: %+v\n", "[4,5,6]", middleNode(&head))

	head = ListNode{
		Val:  1,
		Next: nil,
	}
	fmt.Printf("Expected: %v Result: %+v\n", "[1]", middleNode(&head))

	head = ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	fmt.Printf("Expected: %v Result: %+v\n", "[2]", middleNode(&head))
}
