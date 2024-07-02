package tree

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	var left, right []int

	var inOrder func(root *TreeNode, ans *[]int)
	inOrder = func(root *TreeNode, ans *[]int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			*ans = append(*ans, root.Val)
		} else {
			inOrder(root.Left, ans)
			inOrder(root.Right, ans)
		}

	}
	inOrder(root1, &left)
	inOrder(root2, &right)

	if len(left) != len(right) {
		return false
	}

	for i, val := range left {
		if right[i] != val {
			return false
		}
	}

	return true
}
