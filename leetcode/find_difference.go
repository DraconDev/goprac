package leetcode

func FindDifference(nums1 []int, nums2 []int) [][]int {
	var map1, map2 = make(map[int]bool), make(map[int]bool)

	for _, v := range nums1 {
		map1[v] = true
	}

	for _, v := range nums2 {
		map2[v] = true
	}

	result := [][]int{{}, {}}

	for k := range map1 {
		if !map2[k] {
			result[0] = append(result[0], k)
		}
	}

	for k := range map2 {
		if !map1[k] {
			result[1] = append(result[1], k)
		}
	}

	return result
}

// result := [][]int{}

// temp := []int{}
// for _, v := range nums1 {

// 	if !map2[v] {
// 		temp = append(temp, v)
// 	}
// }
// result = append(result, temp)

// temp = []int{}
// for _, v := range nums2 {
// 	if !map1[v] {
// 		temp = append(temp, v)
// 	}
// }
// result = append(result, temp)
