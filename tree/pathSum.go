package tree

// func PathSum(root *TreeNode, sum int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return sumUp(root, 0, sum) + PathSum(root.Left, sum) + PathSum(root.Right, sum)
// }

// func sumUp(root *TreeNode, pre int, sum int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	current := pre + root.Val

// 	c := 0

// 	if current == sum {
// 		c = 1
// 	}
// 	return c + sumUp(root.Left, current, sum) + sumUp(root.Right, current, sum)
// }

func PathSum(root *TreeNode, target int) int {
	cache := map[int]int{0: 1}
	var result int

	var dfs func(*TreeNode, int, int)
	dfs = func(root *TreeNode, target, currPathSum int) {
		if root == nil {
			return
		}

		currPathSum += root.Val
		oldPathSum := currPathSum - target
		result += cache[oldPathSum]
		cache[currPathSum]++

		dfs(root.Left, target, currPathSum)
		dfs(root.Right, target, currPathSum)

		cache[currPathSum]--
	}

	dfs(root, target, 0)
	return result
}
