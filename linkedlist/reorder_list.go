package linkedlist

func ReorderList(head *ListNode) {

	var list []*ListNode

	temp := head

	for temp != nil {
		list = append(list, temp)
		temp = temp.Next
	}

	result := head

	for i := 1; i < len(list); i++ {
		if i%2 == 1 {
			result.Next = list[(len(list)-(i/2))-1]
			result = result.Next
		} else {
			result.Next = list[i/2]
			result = result.Next
		}
		if i == len(list)-1 {
			result.Next = nil
		}
	}
}

// length := 0

// temp := head

// for temp.Next != nil {
// 	length++
// 	temp = temp.Next
// }

// forward, backward := head, head.Next
// for i := 0; i < length/2 + 1; i++ {
// 	forward = forward.Next
// }

// temp = head
// for i:= 0; i < length; i++ {
// 	if i%2 == 0 {
// 		temp.Next = forward
// 		temp = temp.Next
// 		forward = forward.Next
// 	} else {
// 		temp.Next = backward
// 		temp = temp.Next
// 		backward = backward.Next

// 	}
// }
