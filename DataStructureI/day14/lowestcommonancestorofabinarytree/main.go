package main

// 235. Lowest Common Ancestor of a Binary Search Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time: O(log n)
// space: O(1)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
			continue
		}
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
			continue
		}
		break
	}
	return root
}

func main() {
	var tree *TreeNode
	tree = &TreeNode{Val: 6}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 8}
	tree.Left.Left = &TreeNode{Val: 0}
	tree.Left.Right = &TreeNode{Val: 4}
	tree.Left.Right.Left = &TreeNode{Val: 3}
	tree.Left.Right.Right = &TreeNode{Val: 5}
	tree.Right.Left = &TreeNode{Val: 7}
	tree.Right.Right = &TreeNode{Val: 9}
	fmt.Printf("%v\n", lowestCommonAncestor(tree, &TreeNode{Val: 2}, &TreeNode{Val: 8}).Val) // 6
	fmt.Printf("%v\n", lowestCommonAncestor(tree, &TreeNode{Val: 2}, &TreeNode{Val: 4}).Val) // 2
}
