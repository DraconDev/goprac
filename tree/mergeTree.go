package tree

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// if one of t1 and t2 is nil, return the other
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	// merge t1 and t2
	root := &TreeNode{Val: t1.Val + t2.Val}
	// recursion
	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)
	return root
}
