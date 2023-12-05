package leetcode

import "goprac/types"

func DeleteMiddle(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var prev *types.ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = slow.Next
	return head
}
