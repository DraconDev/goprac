package main

import (
	"fmt"
	"strconv"
)

func hello() {
	str := "Hello, 世界!"

	result := strconv.QuoteToASCII(str)

	fmt.Println(result)
}

func main() {
	// fmt.Println(runningSum([]int{1, 2, 3, 4}))
	// fmt.Println(maximumWealth([][]int{{1, 4, 3}, {3, 2, 1}}))
	// fmt.Println(fizzBuzz(15))
	// println(numberOfSteps(14))

	// node5 := &ListNode{Val: 5, Next: nil}
	// node4 := &ListNode{Val: 4, Next: node5}
	// node3 := &ListNode{Val: 3, Next: node4}
	// node2 := &ListNode{Val: 2, Next: node3}
	// node1 := &ListNode{Val: 1, Next: node2}

	// noded := &ListNode{Val: 44, Next: nil}
	// nodec := &ListNode{Val: 23, Next: noded}
	// nodeb := &ListNode{Val: 23, Next: nodec}
	// nodea := &ListNode{Val: 14, Next: nodeb}

	// node6 := &ListNode{Val: 5, Next: nil}
	// node5 := &ListNode{Val: 1, Next: node6}
	// node4 := &ListNode{Val: -4, Next: node5}
	// node3 := &ListNode{Val: -9, Next: node4}
	// node2 := &ListNode{Val: -10, Next: node3}
	// node1 := &ListNode{Val: -10, Next: node2}

	// noded := &ListNode{Val: 44, Next: nil}
	// nodec := &ListNode{Val: 4, Next: nil}
	// nodeb := &ListNode{Val: 3, Next: nodec}
	// nodea := &ListNode{Val: -7, Next: nil}

	// node1 := &ListNode{Val: 1, Next: nil}
	// result := middleNode(node1)

	// println(result.Val)
	// fmt.Println(GetSize(4, 2, 6))
	// fmt.Println(Grow([]int{1, 2, 3, 4}))
	// fmt.Println(MergeArrays([]int{1, 2, 3}, []int{4, 5, 6}))
	// fmt.Println(MergeArrays([]int{8, 44, 1, 2, 3, 3, 3}, []int{4, 5, 6, 8, 9, 12}))
	// fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	// fmt.Println(Strong(145))
	// fmt.Println(PartList([]string{"a", "b", "c", "d"}))
	// fmt.Println(OrderedCount("abracadabra"))
	// fmt.Println(removeNthFromEnd(node1, 2))
	// fmt.Println(bandNameGenerator("c-clamp"))
	// fmt.Println(bandNameGenerator("alaska"))
	// fmt.Println(reverseList(node1))
	// fmt.Println(Capitalize("hello", []int{0, 1, 2}))
	// fmt.Println(Comp([]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}))
	// fmt.Println(Comp([]int{2, 3, 4, 5}, []int{9, 16, 25}))
	// fmt.Println(mergeTwoLists(node1, nodea))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(11))
}
