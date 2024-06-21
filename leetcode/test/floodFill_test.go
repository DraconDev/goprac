package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloodFill(t *testing.T) {

	assert.Equal(t, leetcode.FloodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2), [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}})
}
