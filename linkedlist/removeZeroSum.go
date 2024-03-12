package linkedList

// func RemoveZeroSumSublists(head *ListNode) *ListNode {

// 	var dummy ListNode
// 	dummy.Next = head

// 	type IntAndListNode struct {
// 		Number int
// 		Node   *ListNode
// 	}

// 	var sums []IntAndListNode
// 	start := &dummy
// 	head = start

// 	for head != nil {
// 		for i, v := range sums {
// 			if v.Number+head.Val == 0 {
// 				temp := head.Next
// 				sums[i].Node.Next = temp
// 				sums = []IntAndListNode{}
// 				head = start
// 				break
// 			} else {
// 				current := *head
// 				sums[i] = IntAndListNode{Number: v.Number + current.Val, Node: &current}
// 			}
// 		}
// 		sums = append(sums, IntAndListNode{Number: head.Val, Node: head})
// 		head = head.Next
// 	}
// 	return start.Next
// }

func RemoveZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prefixSumToNode := make(map[int]*ListNode)
	prefixSum := 0
	for current := dummy; current != nil; current = current.Next {
		prefixSum += current.Val
		if prev, found := prefixSumToNode[prefixSum]; found {
			toRemove := prev.Next
			p := prefixSum
			if toRemove != nil {
				p += toRemove.Val
			}
			for toRemove != nil && p != prefixSum {
				delete(prefixSumToNode, p)
				toRemove = toRemove.Next
				if toRemove != nil {
					p += toRemove.Val
				}
			}
			prev.Next = current.Next
		} else {
			prefixSumToNode[prefixSum] = current
		}
	}
	return dummy.Next
}
