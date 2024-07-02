package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	n1 := make(map[int]bool)
	var result []int
	for _, v := range nums1 {
		n1[v] = true
	}
	for _, v := range nums2 {
		if n1[v] {
			result = append(result, v)
			n1[v] = false
		}
	}
	return result
}
