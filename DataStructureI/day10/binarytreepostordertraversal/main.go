package main

// 145. Binary Tree Postorder Traversal
// https://leetcode.com/problems/binary-tree-postorder-traversal/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// n = number of nodes
// h = height of tree
// time: O(n)
// space: O(h)
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	s := make([]*TreeNode, 0)
	s = append(s, root)
	for root != nil && len(s) > 0 {
		root = s[len(s)-1]
		s = s[:len(s)-1]
		result = append([]int{root.Val}, result...)
		if root.Left != nil {
			s = append(s, root.Left)
		}
		if root.Right != nil {
			s = append(s, root.Right)
		}
	}
	return result
}

func main() {
	var tree *TreeNode
	tree = &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 2}
	tree.Right.Left = &TreeNode{Val: 3}
	fmt.Printf("%v\n", postorderTraversal(tree)) // [3 2 1]

	tree = nil
	fmt.Printf("%v\n", postorderTraversal(tree)) // []

	tree = &TreeNode{Val: 1}
	fmt.Printf("%v\n", postorderTraversal(tree)) // [1]
}
