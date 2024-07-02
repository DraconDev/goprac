package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPath(t *testing.T) {
	assert.Equal(t, leetcode.ValidPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2), true)
	assert.Equal(t, leetcode.ValidPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5), false)
}
