package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	result := TreeNode{}
	result.Val = root1.Val + root2.Val
	if root1.Left == nil && root2.Left != nil {
		result.Left = root2.Left
	}
	if root1.Left != nil && root2.Left == nil {
		result.Left = root1.Left
	}
	if root1.Left != nil && root2.Left != nil {
		result.Left = mergeTrees(root1.Left, root2.Left)
	}
	if root1.Right == nil && root2.Right != nil {
		result.Right = root2.Right
	}
	if root1.Right != nil && root2.Right == nil {
		result.Right = root1.Right
	}
	if root1.Right != nil && root2.Right != nil {
		result.Right = mergeTrees(root1.Right, root2.Right)
	}
	return &result
}

func main() {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Printf("%v\n", mergeTrees(tree1, tree2))
	tree1 = &TreeNode{
		Val: 1,
	}
	tree2 = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
	}
	fmt.Printf("%v\n", mergeTrees(tree1, tree2))
}
