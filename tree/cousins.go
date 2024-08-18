package tree

func isCousins(root *TreeNode, x int, y int) bool {
	type match struct {
		parent int
		depth  int
	}

	matches := make(map[int]match)

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, parent int, depth int) {
		if node == nil {
			return
		}
		if node.Val == x || node.Val == y {
			matches[node.Val] = match{parent, depth}
		}
		dfs(node.Left, node.Val, depth+1)
		dfs(node.Right, node.Val, depth+1)
	}

	dfs(root.Left, root.Val, 1)
	dfs(root.Right, root.Val, 1)
	for k, v := range matches {
		if k == x {
			_, ok := matches[y]
			if ok {
				if v.depth == matches[y].depth && v.parent != matches[y].parent {
					return true
				}
			}
		}
		if k == y {
			_, ok := matches[x]
			if ok {
				if v.depth == matches[x].depth && v.parent != matches[x].parent {
					return true
				}
			}
		}
	}

	return false

}
