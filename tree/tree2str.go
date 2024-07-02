package tree

import "strconv"

// func Tree2str(root *TreeNode) string {
// 	if root == nil {
// 		return ""
// 	}

// 	left, right := "", ""

// 	if root.Left != nil {
// 		left = "(" + Tree2str(root.Left) + ")"
// 	}

// 	if root.Right != nil {
// 		left = "(" + Tree2str(root.Left) + ")"
// 		right = "(" + Tree2str(root.Right) + ")"
// 	}

// 	return fmt.Sprintf("%v", root.Val) + left + right
// }

func Tree2str(root *TreeNode) string {
	// return "" if node is empty and do not add ()
	if root == nil {
		return ""
	}

	s := strconv.Itoa(root.Val)

	// add left subnode in () is we have at least one subtree
	if root.Left != nil || root.Right != nil {
		s += "(" + Tree2str(root.Left) + ")"
	}

	// add right subtree in () it is not empty
	if root.Right != nil {
		s += "(" + Tree2str(root.Right) + ")"
	}

	return s
}
