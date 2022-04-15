package main

// 226. Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/

import "fmt"

// used for output formatting
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	s := make([]*TreeNode, 0)
	s = append(s, root)
	for len(s) > 0 {
		root = s[len(s)-1]
		s = s[:len(s)-1]
		result = append(result, root.Val)
		if root.Right != nil {
			s = append(s, root.Right)
		}
		if root.Left != nil {
			s = append(s, root.Left)
		}

	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// n = number of nodes
// w = maximum width of tree
// time: O(n)
// space: O(w)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return root
}

func main() {
	var tree *TreeNode
	tree = &TreeNode{Val: 4}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 7}
	tree.Left.Left = &TreeNode{Val: 1}
	tree.Left.Right = &TreeNode{Val: 3}
	tree.Right.Left = &TreeNode{Val: 6}
	tree.Right.Right = &TreeNode{Val: 9}
	fmt.Printf("%v\n", preorderTraversal(invertTree(tree))) // [4 7 9 6 2 3 1]

	tree = &TreeNode{Val: 2}
	tree.Left = &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 3}
	fmt.Printf("%v\n", preorderTraversal(invertTree(tree))) // [2 3 1]

	tree = nil
	fmt.Printf("%v\n", preorderTraversal(invertTree(tree))) // []
}
