package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPerimeter(t *testing.T) {
	assert.Equal(t, leetcode.LargestPerimeter([]int{2, 1, 22}), 5)
}
