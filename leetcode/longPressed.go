package leetcode

func isLongPressedName(name string, typed string) bool {
	currentIndex := 0

	if len(name) > len(typed) {
		return false
	}

	if name == typed {
		return true
	}

	for i := 0; i < len(name); i++ {
		count := 1
		for i2 := i + 1; i2 < len(name); i2++ {
			if name[i] == name[i2] {
				count++
				i++
			} else {
				break
			}
		}
		if currentIndex >= len(typed) {
			return false
		}
		for j := currentIndex; j < len(typed); j++ {
			if typed[j] != name[i] {
				if count > 0 {
					return false
				} else {
					break
				}
			} else {
				count--
			}
			currentIndex++
		}
	}
	if currentIndex < len(typed) {
		return false
	}
	return true
}
