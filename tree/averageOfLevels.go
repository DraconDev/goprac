package tree

func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		sum := 0
		for i := 0; i < size; i++ {
			node := queue[0]
			sum += node.Val
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(size))
	}
	return result
}
