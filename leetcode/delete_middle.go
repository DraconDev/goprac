package leetcode

import "goprac/types"

func DeleteMiddle(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var fast, slow *types.ListNode = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		if fast == nil || fast.Next == nil {
			slow.Next = slow.Next.Next
		} else {
			slow = slow.Next
		}

	}
	return head
}
