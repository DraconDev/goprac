package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDays(t *testing.T) {
	// assert.Equal(t, leetcode.MinDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3), 6)
	// assert.Equal(t, leetcode.MinDays([]int{1, 10, 3, 10, 2}, 3, 1), 3)

	assert.Equal(t, leetcode.MinDays([]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2), 9)
}
