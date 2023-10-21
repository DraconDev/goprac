package tree

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	result := [][]int{}
	for len(queue) > 0 {
		size := len(queue)
		result = append(result, []int{})
		for i := 0; i < size; i++ {
			node := queue[0]
			result[len(result)-1] = append(result[len(result)-1], node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}
