package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitAssignment(t *testing.T) {
	assert.Equal(t, leetcode.MaxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}), 34)
}
