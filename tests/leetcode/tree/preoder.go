package tree

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}
	for _, v := range root.Children {
		tmp := preorder(v)
		if tmp != nil {
			res = append(res[:len(res)], tmp...)
		}
	}

	return res
}
