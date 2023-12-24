package linkedList

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		list1, list2 = list2, list1
	}
	main := list1

	for list1.Next != nil && list2 != nil {
		if list1.Next.Val >= list2.Val {
			temp := list2
			list2 = list2.Next
			temp.Next = list1.Next
			list1.Next = temp
		}
		list1 = list1.Next
	}

	if list2 != nil {
		list1.Next = list2
	}

	return main
}
