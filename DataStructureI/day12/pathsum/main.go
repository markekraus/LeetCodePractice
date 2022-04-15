package main

// 112. Path Sum
// https://leetcode.com/problems/path-sum/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// n = number of nodes
// w = maximum width of tree
// time: O(n)
// space: O(w)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
	var tree *TreeNode
	tree = &TreeNode{Val: 5}
	tree.Left = &TreeNode{Val: 4}
	tree.Right = &TreeNode{Val: 8}
	tree.Left.Left = &TreeNode{Val: 11}
	tree.Left.Left.Left = &TreeNode{Val: 17}
	tree.Left.Left.Right = &TreeNode{Val: 2}
	tree.Right.Left = &TreeNode{Val: 13}
	tree.Right.Right = &TreeNode{Val: 4}
	tree.Right.Right.Right = &TreeNode{Val: 1}
	fmt.Printf("%v\n", hasPathSum(tree, 22)) // true

	tree = &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 3}
	fmt.Printf("%v\n", hasPathSum(tree, 5)) // false

	tree = nil
	fmt.Printf("%v\n", hasPathSum(tree, 0)) // false
}
