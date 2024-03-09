package leetcode

func GetCommon(nums1 []int, nums2 []int) int {

	for _, v := range nums1 {
		for _, v2 := range nums2 {
			if v2 > v {
				break
			}
			if v == v2 {
				return v
			}

			if v2 < v {
				nums2 = nums2[1:]
			}
		}
	}
	return -1
}
