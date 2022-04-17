package main

// 700. Search in a Binary Search Tree
// https://leetcode.com/problems/search-in-a-binary-search-tree/

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
// time: O(log n)
// space: O(1)
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	insert := &TreeNode{Val: val}
	if root == nil {
		return insert
	}
	current := root
	for {
		if current.Val < val && current.Right == nil {
			current.Right = insert
			break
		}
		if current.Val < val && current.Right != nil {
			current = current.Right
			continue
		}
		if current.Left == nil {
			current.Left = insert
			break
		}
		current = current.Left
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
	fmt.Printf("%v\n", preorderTraversal(insertIntoBST(tree, 5))) // [4 2 1 3 7 5]

	tree = &TreeNode{Val: 40}
	tree.Left = &TreeNode{Val: 20}
	tree.Right = &TreeNode{Val: 60}
	tree.Left.Left = &TreeNode{Val: 10}
	tree.Left.Right = &TreeNode{Val: 30}
	tree.Right.Left = &TreeNode{Val: 50}
	tree.Right.Right = &TreeNode{Val: 70}
	fmt.Printf("%v\n", preorderTraversal(insertIntoBST(tree, 25))) // [40 20 10 30 25 60 50 70]
}
