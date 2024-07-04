package linkedlist

func Buildlinkedlist() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	// head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	// head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	// head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 8}
	return head
}
