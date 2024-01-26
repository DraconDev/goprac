package tree

func InvertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	} else {
		root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)

		return root
	}
}
