package tree

import (
	"math"

	"goprac/types"
)

func maxDepthOfBinaryTree(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepthOfBinaryTree(root.Left)), float64(maxDepthOfBinaryTree(root.Right)))) + 1

}
