package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	// assert.Equal(t, leetcode.SubarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2), 7)
	// assert.Equal(t, leetcode.SubarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3), 3)
	// 2,2,1,2,2,2,1,1
	assert.Equal(t, leetcode.SubarraysWithKDistinct([]int{2, 2, 1, 2, 2, 2, 1, 1}, 2), 3)
}
