package tree

func FindMode(root *TreeNode) []int {

	var max int
	count := map[int]int{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		count[node.Val]++

		if count[node.Val] > max {
			max = count[node.Val]
		}

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	var result []int
	for k, v := range count {
		if v == max {
			result = append(result, k)
		}
	}
	return result

}
