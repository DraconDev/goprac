package tree

func PathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return sumUp(root, 0, sum) + PathSum(root.Left, sum) + PathSum(root.Right, sum)
}

func sumUp(root *TreeNode, pre int, sum int) int {
	if root == nil {
		return 0
	}
	current := pre + root.Val

	c := 0

	if current == sum {
		c = 1
	}
	return c + sumUp(root.Left, current, sum) + sumUp(root.Right, current, sum)
}
