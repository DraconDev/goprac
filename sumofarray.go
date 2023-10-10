package main

import (
	"fmt"
	"strconv"
)

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
	// result := middleNode(node1)

	// println(result.Val)
	fmt.Println(GetSize(4,2,6))
}

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	sum := 0

	for i, num := range nums {
		sum += num
		result[i] = sum
	}

	return result

}

func maximumWealth(accounts [][]int) int {
	result := 0
	for _, e := range accounts {
		sum := 0
		for _, v := range e {
			sum += v
		}
		if sum > result {
			result = sum
		}
	}
	return result
}

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			result[i-1] = "Fizz"
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}

func numberOfSteps(num int) int {
	result := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
			result++
		} else {
			num -= 1
			result++
		}
		continue
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	length := 0
	middle := head
	for node := head; node != nil; node = node.Next {
		length++
		if length%2 == 0 {
			middle = middle.Next
		}
	}
	return middle
}
