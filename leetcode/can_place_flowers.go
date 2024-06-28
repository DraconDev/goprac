package leetcode

func CanPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i++ {
		if i-1 >= 0 {
			if flowerbed[i-1] == 1 {
				continue
			}
		}
		if i+1 <= len(flowerbed)-1 {
			if flowerbed[i+1] == 1 {
				continue
			}
		}
		if flowerbed[i] == 1 {
			continue
		}
		flowerbed[i] = 1
		n--
		if n <= 0 {
			return true
		}
	}
	return false
}
