package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeight(t *testing.T) {
	assert.Equal(t, lastStoneWeight([]int{2, 7, 4, 1, 8, 1}), 1)
}
