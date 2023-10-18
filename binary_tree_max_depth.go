package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepthOfBinaryTree(root.Left)), float64(maxDepthOfBinaryTree(root.Right)))) + 1

}
