package linkedlist

func MergeNodes(head *ListNode) *ListNode {
	// new list listnode type
	result := &ListNode{}
	sum := 0
	current := result
	for head != nil {
		if head.Val != 0 {
			sum += head.Val
		} else {
			if sum != 0 {
				current.Next = &ListNode{Val: sum}
				current = current.Next
				sum = 0
			}
		}
		head = head.Next
	}
	return result.Next
}
