package leetcode

// func MinDays(bloomDay []int, m int, k int) int {
// 	if m*k > len(bloomDay) {
// 		return -1
// 	}
// 	unique := make(map[int]bool)
// 	for _, day := range bloomDay {
// 		unique[day] = true
// 	}

// 	days := []int{}
// 	for k := range unique {
// 		days = append(days, k)
// 	}
// 	sort.Ints(days)

// 	count := 0
// 	for _, day := range days {
// 	outerLoop:
// 		for i := 0; i < len(bloomDay)+1-k; i++ {
// 			if day >= bloomDay[i] {
// 				for j := i; j < i+k; j++ {
// 					if day >= bloomDay[j] {
// 					} else {
// 						continue outerLoop
// 					}
// 				}
// 				count += k
// 				bloomDay = append(bloomDay[:i], bloomDay[i+k:]...)
// 				i -= k
// 				// i += k
// 				if count >= m*k {
// 					return day
// 				}
// 			}
// 		}
// 	}
// 	return -1
// }

func MinDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)

	if m*k > n {
		return -1
	}

	left, right := 1, 1000000000
	for left < right {
		mid := (left + right) / 2

		if canMakeBouq(bloomDay, mid, m, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canMakeBouq(bloomDay []int, days, m, k int) bool {
	bouq := 0
	seq := 0
	for i := 0; i < len(bloomDay) && bouq < m; i++ {
		if bloomDay[i] > days {
			seq = 0
			continue
		}
		seq++
		if seq == k {
			bouq++
			seq = 0 // start again new sequence
		}
	}

	return bouq >= m
}
