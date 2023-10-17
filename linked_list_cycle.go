package main

func hasCycle(head *ListNode) bool {
	//
	for head != nil {
		if head.Next == nil {
			return false
		}
		temp := head
		for temp.Next != nil {
			temp = temp.Next
			if temp == head {
				return true
			}
		}
		head = head.Next
	}
	return false
}
