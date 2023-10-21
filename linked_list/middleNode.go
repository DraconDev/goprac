package linkedList

import "your.module.name~/types"

func MiddleNode(head *types.ListNode) *types.ListNode {
	length := 0
	middle := head
	for node := head; node != nil; node = node.Next {
		length++
		if length%2 == 0 {
			middle = middle.Next
		}
	}
	return middle
}
