package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement(t *testing.T) {
	assert.Equal(t, leetcode.NextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}), []int{-1, 3, -1})
}
