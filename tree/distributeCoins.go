package tree

func distributeCoins(root *TreeNode) int {

	coins := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		coins += abs(left) + abs(right)
		return node.Val + left + right - 1
	}
	dfs(root)
	return coins
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
