package tree

func increasingBST(root *TreeNode) *TreeNode {

	var dfs func(*TreeNode) *TreeNode

	links := []*TreeNode{}

	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		dfs(node.Left)
		links = append(links, node)
		dfs(node.Right)
		return node
	}

	dfs(root)

	for i := 1; i < len(links); i++ {
		links[i-1].Right = links[i]
		links[i].Left = nil
	}

	return links[0]

}
