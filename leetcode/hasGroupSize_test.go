package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasGroupSize(t *testing.T) {
	// assert.Equal(t, hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}), false)
	// assert.Equal(t, hasGroupsSizeX([]int{1, 1}), true)
	// assert.Equal(t, hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2}), true)
	assert.Equal(t, hasGroupsSizeX([]int{0, 0, 0, 1, 1, 1, 2, 2, 2}), true)
}
