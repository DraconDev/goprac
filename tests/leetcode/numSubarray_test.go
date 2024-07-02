package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSubarraysWithSum(t *testing.T) {
	// assert.Equal(t, leetcode.NumSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2), 4)
	assert.Equal(t, leetcode.NumSubarraysWithSum([]int{0, 0, 0, 0, 1}, 2), 0)

}
