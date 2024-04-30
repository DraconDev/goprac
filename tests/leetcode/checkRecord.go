package leetcode

func checkRecord(s string) bool {

	absent := 0
	late := 0
	for _, v := range s {
		if string(v) == "A" {
			absent++
			late = 0
		} else if string(v) == "L" {
			late++
		} else {
			late = 0
		}
		if absent >= 2 || late >= 3 {
			return false
		}
	}
	return true
}
