package main

// 144. Binary Tree Preorder Traversal
// https://leetcode.com/problems/binary-tree-preorder-traversal/
import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// n = number nodes
// h = tree height
// time: O(n)
// space: O(h)
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

func main() {
	var tree *TreeNode
	tree = &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 2}
	tree.Right.Left = &TreeNode{Val: 3}
	fmt.Printf("%v\n", preorderTraversal(tree)) // [1 2 3]

	tree = nil
	fmt.Printf("%v\n", preorderTraversal(tree)) // []

	tree = &TreeNode{Val: 1}
	fmt.Printf("%v\n", preorderTraversal(tree)) // [1]
}
