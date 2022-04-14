package main

// 104. Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
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
func maxDepth(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		result++
		for i := 0; i < n; i++ {
			root = q[0]
			q = q[1:]
			if root.Left != nil {
				q = append(q, root.Left)
			}
			if root.Right != nil {
				q = append(q, root.Right)
			}
		}
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
	fmt.Printf("%v\n", maxDepth(tree)) // 3

	tree = &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 2}
	fmt.Printf("%v\n", maxDepth(tree)) // 2
}
