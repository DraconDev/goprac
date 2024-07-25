package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFairCandySwap(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{1, 2}, fairCandySwap([]int{1, 1}, []int{2, 2}))
}
