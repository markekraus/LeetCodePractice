package main

// 98. Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	s := make([]*TreeNode, 0)
	var prev *TreeNode
	for root != nil || len(s) > 0 {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[0 : len(s)-1]
		if prev != nil && prev.Val >= root.Val {
			return false
		}
		prev = root
		root = root.Right
	}
	return true
}

func main() {
	var tree *TreeNode
	tree = &TreeNode{Val: 2}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 2}
	fmt.Printf("%v\n", isValidBST(tree)) // false

	tree = &TreeNode{Val: 5}
	tree.Left = &TreeNode{Val: 4}
	tree.Right = &TreeNode{Val: 6}
	tree.Right.Left = &TreeNode{Val: 3}
	tree.Right.Left = &TreeNode{Val: 7}
	fmt.Printf("%v\n", isValidBST(tree)) // false
}
