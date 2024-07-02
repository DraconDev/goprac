package tree

import (
	"math"
	"sort"
)

func getMinimumDifference(root *TreeNode) int {
	minVal := math.MaxInt

	values := []int{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		values = append(values, node.Val)
		dfs(node.Left)
		dfs(node.Right)

	}
	dfs(root)
	sort.Ints(values)

	for i := 1; i < len(values); i++ {
		minVal = min(minVal, values[i]-values[i-1])
	}

	return minVal

}
