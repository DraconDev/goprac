package main

import (
	"strings"
)

func Capitalize(st string, arr []int) string {
	//..
	slice := []byte(st)
	for _, v := range arr {
		// slice[v] = strings.ToUpper(string(slice[v]))
		length := len(slice)
		if v < length {
			slice[v] = strings.ToUpper(string(slice[v]))[0]
		}
	}
	return string(slice)
}

func isInSlice(slice []int, val int) bool {
	for _, v := range slice {
		check := v * v
		// check := int(math.Sqrt(float64(val)))
		if check == val {
			return true
		}
	}
	return false
}

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil && array2 == nil {
		return false
	}
	for _, v := range array2 {
		if !isInSlice(array1, v) {
			return false
		}
	}
	return true
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	// main := list1
// 	// if list1.Val <= list2.Val {
// 	// 	main := list1
// 	// 	secondary := list2
// 	// } else {
// 	// 	main := list2
// 	// 	secondary := list1
// 	// }
// 	for secondary.Next != nil {
// 		if main.Val <= secondary.Val {
// 			temp := secondary
// 			secondary = secondary.Next
// 			temp.Next = main.Next
// 			main.Next = temp
// 		}
// 	}
// 	return main
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	if list1 == nil {
// 		return list2
// 	}
// 	if list2 == nil {
// 		return list1
// 	}
// 	if list1.Val > list2.Val {
// 		list1, list2 = list2, list1
// 	}
// 	main := list1

// 	for list1.Next != nil || list2.Next != nil {
// 		if list1.Next == nil {
// 			list1.Next = list2
// 			break
// 		}

// 		if list1.Val <= list2.Val && list1.Next.Val >= list2.Val {
// 			temp := list2
// 			temp.Next = list1.Next
// 			list1.Next = temp

// 			// if list2.Next == nil {
// 			// 	break
// 			// }
// 			list2 = list2.Next
// 		}
// 		list1 = list1.Next
// 	}

// 	// for main.Next != nil {
// 	// 	fmt.Println(main.Val)
// 	// 	main = main.Next
// 	// }
// 	return main
// }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
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
