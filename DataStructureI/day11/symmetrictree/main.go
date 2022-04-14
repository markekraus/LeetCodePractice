package main

// 101. Symmetric Tree
// https://leetcode.com/problems/symmetric-tree/

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
func isSymmetric(root *TreeNode) bool {
	q := make([][]*TreeNode, 0)
	q = append(q, []*TreeNode{root.Left, root.Right})
	for len(q) > 0 {
		left, right := q[0][0], q[0][1]
		q = q[1:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		q = append(q, []*TreeNode{left.Left, right.Right})
		q = append(q, []*TreeNode{left.Right, right.Left})
	}
	return true
}

func main() {
	var tree *TreeNode
	tree = &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 2}
	tree.Left.Left = &TreeNode{Val: 3}
	tree.Right.Right = &TreeNode{Val: 3}
	tree.Left.Right = &TreeNode{Val: 4}
	tree.Right.Left = &TreeNode{Val: 4}
	fmt.Printf("%v\n", isSymmetric(tree)) // true

	tree = &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 2}
	tree.Right.Right = &TreeNode{Val: 3}
	tree.Left.Right = &TreeNode{Val: 3}
	fmt.Printf("%v\n", isSymmetric(tree)) // false
}
