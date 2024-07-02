package tree

import "sort"

func minDiffInBST(root *TreeNode) int {
	// difference is max
	difference := 10000000

	elements := make([]int, 0)

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		elements = append(elements, root.Val)
		dfs(root.Right)
		dfs(root.Left)

	}

	dfs(root)
	sort.Ints(elements)
	for i := 1; i < len(elements); i++ {
		difference = min(difference, elements[i]-elements[i-1])
		if difference == 1 {
			break
		}
	}
	return difference

}
