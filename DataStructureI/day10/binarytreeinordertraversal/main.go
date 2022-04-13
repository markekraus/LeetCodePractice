package main

// 94. Binary Tree Inorder Traversal
// https://leetcode.com/problems/binary-tree-inorder-traversal/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// n = number of nodes
// h = tree height
// time: O(n)
// space: O(h)
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	s := make([]*TreeNode, 0)
	for root != nil || len(s) > 0 {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}
func main() {
	var tree *TreeNode
	tree = &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 2}
	tree.Right.Left = &TreeNode{Val: 3}
	fmt.Printf("%v\n", inorderTraversal(tree)) // [1 3 2]

	tree = nil
	fmt.Printf("%v\n", inorderTraversal(tree)) // []

	tree = &TreeNode{Val: 1}
	fmt.Printf("%v\n", inorderTraversal(tree)) // [1]
}
