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
// h = maximum height of tree
// time: O(log n) to O(h)
// space: O(1)
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	for {
		if root.Val == val {
			return root
		}
		if root.Val < val && root.Right == nil {
			break
		}
		if root.Val < val && root.Right != nil {
			root = root.Right
			continue
		}
		if root.Left == nil {
			break
		}
		root = root.Left
	}
	return nil
}

func main() {
	var tree *TreeNode
	tree = &TreeNode{Val: 4}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 7}
	tree.Left.Left = &TreeNode{Val: 1}
	tree.Left.Right = &TreeNode{Val: 3}
	fmt.Printf("%v\n", preorderTraversal(searchBST(tree, 2))) // [2 1 3]
	fmt.Printf("%v\n", preorderTraversal(searchBST(tree, 5))) // []
}
