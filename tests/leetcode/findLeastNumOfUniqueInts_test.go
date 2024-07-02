package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLeastNumOfUniqueInts(t *testing.T) {
	// assert.Equal(t, leetcode.FindLeastNumOfUniqueInts([]int{5, 5, 4}, 1), 1)
	// assert.Equal(t, leetcode.FindLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3}, 3), 2)

	// [1,1]
	assert.Equal(t, leetcode.FindLeastNumOfUniqueInts([]int{1, 1}, 1), 1)
}
