package tree

func isValid(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if (min != nil && root.Val <= min.Val) || (max != nil && root.Val >= max.Val) {
		return false
	}
	return isValid(root.Left, min, root) && isValid(root.Right, root, max)
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

// func isValidBST(root *TreeNode) bool {
// 	// if left smaller or nil
// 	// if right bigger or nil

// 	isValid := func(root *TreeNode, left *TreeNode, right *TreeNode) bool {
// 		if root == nil {
// 			return true
// 		}
// 		return true
// 	}

// 	isValid(root, root.Left, root.Right)

// 	if root.Left == nil && root.Right == nil {
// 		return true
// 	}
// 	if root.Left == nil || root.Left.Val < root.Val {
// 		isValidBST(root.Left)
// 	} else {
// 		return false
// 	}

// 	if root.Right == nil || root.Right.Val > root.Val {
// 		isValidBST(root.Right)
// 	} else {
// 		return false
// 	}
// 	return true
// }
