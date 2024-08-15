package tree

func isUnivalTree(root *TreeNode) bool {
	value := root.Val
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Val != value {
			return false
		}
		return dfs(root.Left) && dfs(root.Right)
	}
	return dfs(root)
}
