package main

import (
	"sort"
)

func MergeArrays(arr1, arr2 []int) []int {

	// reverse arr2
	// for i, j := 0, len(arr2)-1; i < j; i, j = i+1, j-1 {
	// 	arr2[i], arr2[j] = arr2[j], arr2[i]

	// }
	combined := append(arr1, arr2...)
	set := make(map[int]bool)
	for _, v := range combined {
		set[v] = true
	}
	result := []int{}
	for k := range set {
		result = append(result, k)
	}
	sort.Ints(result)
	return result
}

// func longestCommonPrefix(strs []string) string {
// 	result := ""
// 	for i := 0; i < len(strs[0]); i++ {
// 		c := strs[0][i]
// 		for j := 1; j < len(strs); j++ {
// 			if i == len(strs[j]) || strs[j][i] != c {
// 				return result
// 			}
// 		}
// 		result += string(c)

// 	}
// 	return result
// }

func longestCommonPrefix(strs []string) string {
	result := ""
	// sort.Ints(arr)
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return result
			}
		}
		result += string(c)

	}
	return result
}

func Incrementer(n []int) []int {
	// your code here
	for i, v := range n {
		n[i] = v + i + 1
	}
	return n
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
