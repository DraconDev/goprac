package tree

func SumOfLeftLeaves(root *TreeNode) int {
	var sum int

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum += node.Left.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return sum
}
