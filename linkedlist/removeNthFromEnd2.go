package linkedlist

func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// count length
	length := 0
	temp := head
	for temp != nil {
		length++
		temp = temp.Next
	}

	if length == n {
		return head.Next
	}

	temp = head
	for length-n-1 > 0 {
		temp = temp.Next
		length--
	}

	if temp.Next != nil {
		temp.Next = temp.Next.Next
	}

	return head
}
