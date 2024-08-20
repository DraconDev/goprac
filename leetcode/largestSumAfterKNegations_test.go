package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestSumAfterKNegations(t *testing.T) {

	// assert.Equal(t, largestSumAfterKNegations([]int{4, 2, 3}, 1), 5)

	assert.Equal(t, largestSumAfterKNegations([]int{3, -1, 0, 2}, 3), 6)
}
