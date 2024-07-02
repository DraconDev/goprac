package leetcode

//	func isOneBitCharacter(bits []int) bool {
//		if len(bits) == 1 {
//			return bits[0] == 0
//		}
//		if bits[len(bits)-2] == 1 {
//			return false
//		}
//		if len(bits) == 2 {
//			return true
//		}
//		return bits[len(bits)-1] == 0
//	}
func IsOneBitCharacter(bits []int) bool {
	isOneBit := false

	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			i++
			continue
		}
		if bits[i] == 0 && len(bits) == i+1 {
			isOneBit = true

		}
	}
	return isOneBit
}
