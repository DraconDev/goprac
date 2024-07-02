package tree

import "math"

func MaxLevelSum(root *TreeNode) int {
	lvl := 0
	min_lvl_so_far := 0
	max_sum_so_far := math.MinInt64
	q := []*TreeNode{root}

	for len(q) > 0 {
		lvl++
		currSum := 0
		n := len(q)

		for n > 0 {
			node := q[0]
			q = q[1:]
			currSum += node.Val

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

			n--
		}

		if max_sum_so_far < currSum {
			max_sum_so_far = currSum
			min_lvl_so_far = lvl
		}
	}

	return min_lvl_so_far
}
