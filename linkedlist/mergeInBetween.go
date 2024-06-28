package linkedlist

func MergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	node := list1

	for i := 0; i < a-1; i++ {
		list1 = list1.Next
	}
	temp := list1
	for i := 0; i < b-a-1; i++ {
		temp = temp.Next
	}
	list1.Next = list2

	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = temp.Next
	return node
}
