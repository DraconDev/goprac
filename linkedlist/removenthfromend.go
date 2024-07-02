package linkedlist

func DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	// find length
	length := 1
	temp := head
	for temp.Next != nil {
		temp = temp.Next
		length++
	}

	// remove
	temp = head
	for i := 0; i < length-n-1; i++ {
		if temp.Next == nil {
			return head
		}
		temp = temp.Next
	}
	if n == length {
		return head.Next
	}
	if temp.Next.Next != nil {
		temp.Next = temp.Next.Next
	} else {
		temp.Next = nil
	}

	return head
}
