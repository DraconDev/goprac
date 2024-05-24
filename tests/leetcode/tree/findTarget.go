package tree

func findTarget(root *TreeNode, k int) bool {
	var dfs func(*TreeNode) bool
	elems := map[int]bool{}
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		if elems[k-node.Val] {
			return true
		}
		elems[node.Val] = true
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}
