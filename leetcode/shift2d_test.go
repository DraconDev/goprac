package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShift2d(t *testing.T) {

	test := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	assert.Equal(t, shiftGrid(test, 1), [][]int{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}})
}
