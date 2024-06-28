package tree

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(*TreeNode) int

	result := 0
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		if left+right > result {
			result = left + right
		}
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
	dfs(root)

	return result
}
