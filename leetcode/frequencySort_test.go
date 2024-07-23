package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {
	// assert.Equal(t, frequencySort([]int{1, 1, 2, 2, 2, 3}), []int{3, 1, 1, 2, 2, 2})

	assert.Equal(t, frequencySort([]int{-3}), []int{-3})
}
