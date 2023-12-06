package leetcode

import "goprac/types"

func OddEvenList(head *types.ListNode) *types.ListNode {

	evens := head.Next
	odds := head
	evensHead := evens

	for odds != nil && odds.Next != nil && odds.Next.Next != nil {
		odds.Next = odds.Next.Next
		odds = odds.Next
		if evens != nil && evens.Next != nil {
			if evens.Next.Next == nil {
				evens.Next = nil
				continue
			}
			evens.Next = evens.Next.Next
			evens = evens.Next

		}
	}
	odds.Next = evensHead
	return head

}
