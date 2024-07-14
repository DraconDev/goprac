package leetcode

func flipAndInvertImage(image [][]int) [][]int {
	flipHorizontal(image)
	return image
}

func flipHorizontal(image [][]int) {
	for i := 0; i < len(image); i++ {
		if len(image[i])%2 == 1 {
			image[i][len(image[i])/2] = flip(image[i][len(image[i])/2])
		}
		for j := 0; j < len(image[i])/2; j++ {
			image[i][j] = flip(image[i][j])
			image[i][len(image[i])-j-1] = flip(image[i][len(image[i])-j-1])
			image[i][j], image[i][len(image[i])-j-1] = image[i][len(image[i])-j-1], image[i][j]
		}
	}
}

func flip(elem int) int {
	if elem == 0 {
		elem = 1
	} else {
		elem = 0
	}
	return elem
}
