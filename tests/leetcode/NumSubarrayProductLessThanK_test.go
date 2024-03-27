package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	assert.Equal(t, leetcode.NumSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100), 8)
}
