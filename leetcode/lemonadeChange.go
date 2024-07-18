package leetcode

func lemonadeChange(bills []int) bool {
	bills_map := make(map[int]int)
	for _, v := range bills {
		if v == 5 {
			bills_map[5]++
		} else if v == 10 {
			if bills_map[5] == 0 {
				return false
			}
			bills_map[5]--
			bills_map[10]++
		} else if v == 20 {
			if bills_map[5] > 0 && bills_map[10] > 0 {
				bills_map[5]--
				bills_map[10]--
			} else if bills_map[5] > 2 {
				bills_map[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
