package tree

func GoodNodes(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, maxValue int) int {
		if node == nil {
			return 0
		}
		count := 0
		if node.Val >= maxValue {
			maxValue = node.Val
			count = 1
		}
		count += dfs(node.Left, maxValue)
		count += dfs(node.Right, maxValue)
		return count
	}
	return dfs(root, root.Val)
}

// func GoodNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	ans := 1
// 	var dfs func(*TreeNode, int)
// 	dfs = func(root *TreeNode, max int) {
// 		if root == nil {
// 			return
// 		}
// 		if root.Val >= max {
// 			ans++
// 		}
// 		max = Max(max, root.Val)
// 		dfs(root.Left, max)
// 		dfs(root.Right, max)

// 	}
// 	dfs(root.Left, root.Val)
// 	dfs(root.Right, root.Val)

// 	return ans

// }
// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
