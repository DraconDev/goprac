package tree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			result = append(result, path+strconv.Itoa(node.Val))
		}
		dfs(node.Left, path+strconv.Itoa(node.Val)+"->")
		dfs(node.Right, path+strconv.Itoa(node.Val)+"->")
	}
	dfs(root, "")
	return result

}
