package tree

func SortedArrayToBST(nums []int) *TreeNode {
	// Base case: empty input slice
	if len(nums) == 0 {
		return nil
	}

	// Find the middle element of the slice
	mid := len(nums) / 2

	// Create a new TreeNode with the middle element as the value
	root := &TreeNode{Val: nums[mid]}

	// Recursively build the left and right subtrees
	root.Left = SortedArrayToBST(nums[:mid])
	root.Right = SortedArrayToBST(nums[mid+1:])

	return root
}
