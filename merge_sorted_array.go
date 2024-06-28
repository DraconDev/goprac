package main

// func Merge(nums1 []int, m int, nums2 []int, n int) {
// 	// result := []int{}
// 	var i, j = 0, 0
// 	for i < m+n {
// 		if nums1[i] > nums2[j] || i >= m {
// 			// nums1 = append(nums1[:i+j], append(nums2[j], nums1[i+j:]...)...)
// 			nums1 = append(nums1[:i+j], nums2[j], nums1[i+j:]...)
// 			j++
// 		}
// 		i++
// 	}
// }

func Merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j = m - 1, n - 1
	length := m + n -1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[length-1]  = nums1[i]
			i--
		} else {
			nums1[length-1] = nums2[j]
			j--
		}
		length--
	}
}
