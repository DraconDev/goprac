package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestSubArray(t *testing.T) {
	assert.Equal(t, leetcode.FindShortestSubArray([]int{1, 2, 2, 3, 1}), 2)
	// assert.Equal(t, leetcode.FindShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}), 6)
}
