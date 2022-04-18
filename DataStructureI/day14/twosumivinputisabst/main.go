package main

// 653. Two Sum IV - Input is a BST
// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	lookup := make(map[int]bool, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		root = q[0]
		q = q[1:]
		if _, exists := lookup[k-root.Val]; exists {
			return true
		}
		lookup[root.Val] = true
		if root.Left != nil {
			q = append(q, root.Left)
		}
		if root.Right != nil {
			q = append(q, root.Right)
		}
	}
	return false
}

func main() {
	var tree *TreeNode
	tree = &TreeNode{Val: 5}
	tree.Left = &TreeNode{Val: 3}
	tree.Right = &TreeNode{Val: 6}
	tree.Left.Left = &TreeNode{Val: 2}
	tree.Left.Right = &TreeNode{Val: 4}
	tree.Right.Right = &TreeNode{Val: 7}
	fmt.Printf("%v\n", findTarget(tree, 9))  // true
	fmt.Printf("%v\n", findTarget(tree, 28)) // false
}
