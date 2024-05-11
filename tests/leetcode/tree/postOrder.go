package tree

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	arr := []int{}
	for i := 0; i < len(root.Children); i++ {
		arr = append(arr, postorder(root.Children[i])...)
	}
	return append(arr, root.Val)
}
