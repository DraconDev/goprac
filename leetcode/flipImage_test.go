package leetcode

import "testing"

func TestFlipImage(t *testing.T) {
	test := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	// test := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	// test := [][]int{{0, 1}, {1, 1}}
	flipAndInvertImage(test)
}
