package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairs(t *testing.T) {
	// assert.Equal(t, 15, leetcode.MinCostClimbingStairs([]int{10, 15, 20}))

	assert.Equal(t, 6, leetcode.MinCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
