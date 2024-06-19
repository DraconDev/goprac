package leetcode

import "sort"

func MinDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	unique := make(map[int]bool)
	for _, day := range bloomDay {
		unique[day] = true
	}

	days := []int{}
	for k := range unique {
		days = append(days, k)
	}
	sort.Ints(days)

	count := 0
	for _, day := range days {
	outerLoop:
		for i := 0; i < len(bloomDay)+1-k; i++ {
			if day >= bloomDay[i] {
				for j := i; j < i+k; j++ {
					if day >= bloomDay[j] {
					} else {
						continue outerLoop
					}
				}
				count += k
				bloomDay = append(bloomDay[:i], bloomDay[i+k:]...)
				i -= k
				// i += k
				if count >= m*k {
					return day
				}
			}
		}
	}
	return -1
}
