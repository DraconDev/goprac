package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoordinates(t *testing.T) {

	// assert.Equal(t, checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}), true)
	// 00 01 0-1

	// assert.Equal(t, checkStraightLine([][]int{{0, 0}, {0, 1}, {0, -1}}), true)

	// [[1,-8],[2,-3],[1,2]]

	// assert.Equal(t, checkStraightLine([][]int{{1, -8}, {2, -3}, {1, 2}}), true)

	// [-3,-2],[-1,-2],[2,-2],[-2,-2],[0,-2]

	assert.Equal(t, checkStraightLine([][]int{{-3, -2}, {-1, -2}, {2, -2}, {-2, -2}, {0, -2}}), true)

}
