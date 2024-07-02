package leetcode

func UniqueOccurrences(arr []int) bool {

	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	occurs := make(map[int]bool)
	for _, v := range m {
		if occurs[v] {
			return false
		}
		occurs[v] = true
	}
	return true
}
