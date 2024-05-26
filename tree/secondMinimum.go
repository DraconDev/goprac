package tree

import "math"

func FindSecondMinimumValue(root *TreeNode) int {

	// min 2 length array
	min := make([]int, 2)
	// fill with math.MaxInt
	for i := 0; i < len(min); i++ {
		min[i] = math.MaxInt
	}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val < min[0] {
			min[1] = min[0]
			min[0] = node.Val
		} else if node.Val < min[1] && node.Val != min[0] {
			min[1] = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	if min[1] == math.MaxInt || min[1] == math.MaxInt {
		return -1
	}
	return min[1]
}
