package main

// func hasCycle(head *ListNode) bool {
// 	//
// 	for head != nil {
// 		if head.Next == nil {
// 			return false
// 		}
// 		temp := head
// 		for temp.Next != nil {
// 			if temp.Next == head {
// 				return true
// 			}
// 			temp = temp.Next
// 		}
// 		head = head.Next
// 	}
// 	return false
// }


// floyd algorithm
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
