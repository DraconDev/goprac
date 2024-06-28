package linkedlist

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	var array1, array2 []*ListNode

	for headA != nil {
		array1 = append(array1, headA)
		headA = headA.Next
	}
	for headB != nil {
		array2 = append(array2, headB)
		headB = headB.Next
	}
	var result *ListNode = nil

	if len(array1) > len(array2) {
		array1, array2 = array2, array1
	}
	for i := 1; i <= len(array1); i++ {
		if array1[len(array1)-i] != array2[len(array2)-i] {
			return result
		} else {
			result = array1[len(array1)-i]
		}
	}
	return result
}
