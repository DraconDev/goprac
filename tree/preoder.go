package tree

func PreorderNaryTree(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}
	for _, v := range root.Children {
		tmp := PreorderNaryTree(v)
		if tmp != nil {
			res = append(res[:], tmp...)
		}
	}

	return res
}
