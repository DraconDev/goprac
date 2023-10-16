package main

func isPalindromeLinked(head *ListNode) bool {
	test := []int{}

	for head != nil {
		test = append(test, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	for i := 0; i < len(test)/2; i++ {
		if test[i] != test[len(test)-i-1] {
			return false
		}
	}
	return true
}
