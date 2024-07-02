package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubarrayLength(t *testing.T) {
	assert.Equal(t, leetcode.MaxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2), 2)
}
