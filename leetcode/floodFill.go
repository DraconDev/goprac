package leetcode

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	sourceColor := image[sr][sc]
	if sourceColor != newColor {
		dfsf(image, sr, sc, sourceColor, newColor)
	}
	return image
}

func dfsf(image [][]int, x int, y int, sourceColor int, newColor int) {
	if x < 0 || x >= len(image) || y < 0 || y >= len(image[0]) || image[x][y] != sourceColor {
		return
	}
	image[x][y] = newColor
	dfsf(image, x-1, y, sourceColor, newColor)
	dfsf(image, x+1, y, sourceColor, newColor)
	dfsf(image, x, y-1, sourceColor, newColor)
	dfsf(image, x, y+1, sourceColor, newColor)
}
