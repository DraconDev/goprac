package tree

type Node struct {
	Val      int
	Children []*Node
}

func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	m := 1
	for _, v := range root.Children {
		m = max(m, MaxDepth(v)+1)
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
