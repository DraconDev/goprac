package tree

func rangeSumBST(root *TreeNode, low int, high int) int {
	var dfs func(*TreeNode) int // 递归
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := 0
		if node.Val >= low && node.Val <= high {
			sum += node.Val
		}
		sum += dfs(node.Left)
		sum += dfs(node.Right)
		return sum
	}
	return dfs(root)

}
