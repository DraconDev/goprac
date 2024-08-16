package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPerimeter(t *testing.T) {
	// assert.Equal(t, largestPerimeter([]int{2, 1, 2}), 5)
	assert.Equal(t, largestPerimeter([]int{1, 2, 1, 10}), 0)
}
