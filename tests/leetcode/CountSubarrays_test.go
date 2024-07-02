package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubarrays(t *testing.T) {
	assert.Equal(t, 4, leetcode.CountSubarrays([]int{28, 5, 58, 91, 24, 91, 53, 9, 48, 85, 16, 70, 91, 91, 47, 91, 61, 4, 54, 61, 49}, 1))
}
