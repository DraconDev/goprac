package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinIncrementForUnique(t *testing.T) {
	// assert.Equal(t, 1, leetcode.MinIncrementForUnique([]int{1, 2, 2}))
	assert.Equal(t, 3, leetcode.MinIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
	// assert.Equal(t, 0, leetcode.MinIncrementForUnique([]int{1, 1, 1, 1, 1}))
}
