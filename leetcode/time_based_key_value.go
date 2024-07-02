package leetcode

import "sort"

type TimeMap struct {
	m      map[string]map[int]string
	stamps map[string][]int
}

func Constructor2() TimeMap {
	return TimeMap{m: make(map[string]map[int]string), stamps: make(map[string][]int)}

}

func (time_map *TimeMap) Set(key string, value string, timestamp int) {

	if _, ok := time_map.m[key]; !ok {
		time_map.m[key] = make(map[int]string)
	}
	time_map.m[key][timestamp] = value
	time_map.stamps[key] = append(time_map.stamps[key], timestamp)

}

func (time_map *TimeMap) Get(key string, timestamp int) string {

	if _, ok := time_map.m[key][timestamp]; !ok {
		if len(time_map.stamps[key]) == 0 || timestamp < time_map.stamps[key][0] {
			return ""
		}
		if timestamp > time_map.stamps[key][len(time_map.stamps[key])-1] {
			return time_map.m[key][time_map.stamps[key][len(time_map.stamps[key])-1]]
		}

		var left int
		var right = len(time_map.stamps[key]) - 1

		for left < right {

			mid := left + (right-left)/2
			if time_map.stamps[key][mid] <= timestamp {
				left = mid + 1
			} else {
				right = mid
			}

		}
		return time_map.m[key][time_map.stamps[key][left-1]]
	} else {
		return time_map.m[key][timestamp]
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)

	if len(nums1)%2 == 0 {
		return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2
	} else {
		return float64(nums1[len(nums1)/2])
	}
}
