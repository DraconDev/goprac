package leetcode

// func CheckInclusion(s1 string, s2 string) bool {

// 	var s1map = make(map[rune]int)

// 	for i := 0; i < len(s1); i++ {
// 		s1map[rune(s1[i])]++
// 	}

// 	s1map_copy := copyMap(s1map)
// 	s1map_len := len(s1map)

// 	left, right := 0, 0

// 	for right < len(s2) {
// 		c := rune(s2[right])
// 		if _, ok := s1map_copy[c]; ok {
// 			if s1map_copy[c] == 1 {
// 				delete(s1map_copy, c)
// 			} else {
// 				s1map_copy[c]--
// 			}
// 		} else {
// 			if s1map_len != len(s1map_copy) {
// 				s1map_copy = copyMap(s1map)
// 				left++
// 			}
// 		}
// 		if len(s1map_copy) == 0 {
// 			return true
// 		}
// 	}

// 	return false

// }

// func copyMap(rune_map map[rune]int) map[rune]int {
// 	rune_map_copy := make(map[rune]int)
// 	for k, v := range rune_map {
// 		rune_map_copy[k] = v
// 	}
// 	return rune_map_copy
// }

func CheckInclusion(s1 string, s2 string) bool {
	l, count := 0, [26]int{}
	for i := range s1 {
		count[s1[i]-97]++
	}

	for r := range s2 {
		count[s2[r]-97]--
		if count == [26]int{} {
			return true
		}

		if r+1 >= len(s1) {
			count[s2[l]-97]++
			l++
		}
	}
	return false
}
