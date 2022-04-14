package main

// 102. Binary Tree Level Order Traversal
// https://leetcode.com/problems/binary-tree-level-order-traversal/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// n = number of nodes
// w = maximum width of tree
// time: O(n)
// Space: O(w)
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		layer := make([]int, 0)
		for i := 0; i < n; i++ {
			root = q[0]
			q = q[1:]
			layer = append(layer, root.Val)
			if root.Left != nil {
				q = append(q, root.Left)
			}
			if root.Right != nil {
				q = append(q, root.Right)
			}
		}
		result = append(result, layer)
	}
	return result
}

func main() {
	var tree *TreeNode
	tree = &TreeNode{Val: 3}
	tree.Left = &TreeNode{Val: 9}
	tree.Right = &TreeNode{Val: 20}
	tree.Right.Left = &TreeNode{Val: 15}
	tree.Right.Right = &TreeNode{Val: 7}
	fmt.Printf("%v\n", levelOrder(tree)) // [[3] [9 20] [15 7]]

	tree = &TreeNode{Val: 1}
	fmt.Printf("%v\n", levelOrder(tree)) // [[1]]

	tree = nil
	fmt.Printf("%v\n", levelOrder(tree)) // []
}
