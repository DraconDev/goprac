package tree

func removeLeafNodes(root *TreeNode, target int) *TreeNode {

	var dfs func(*TreeNode) *TreeNode

	dfs = func(node *TreeNode) *TreeNode {

		if node == nil {
			return nil
		}
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)
		if node.Left == nil && node.Right == nil && node.Val == target {
			return nil
		}
		return node
	}
	return dfs(root)
}
