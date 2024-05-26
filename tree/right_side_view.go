package tree

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	result := []int{}

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(result) < level+1 {
			result = append(result, node.Val)
		}
		dfs(node.Right, level+1)
		dfs(node.Left, level+1)
	}
	dfs(root, 0)
	return result
}

// func rightSideView(root *TreeNode) []int {

// 	if root == nil {
// 		return []int{}
// 	}

// 	var result []int
// 	queue := []*TreeNode{root}

// 	for len(queue) > 0 {
// 		node := queue[0]
// 		queue = queue[1:]
// 		result = append(result, node.Val)

// 		if node.Left != nil {
// 			queue = append(queue, node.Left)
// 		}
// 		if node.Right != nil {
// 			queue = append(queue, node.Right)
// 		}
// 	}

// 	return result
// }
