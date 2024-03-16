package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxLength(t *testing.T) {
	assert.Equal(t, leetcode.FindMaxLength([]int{0, 1, 1, 0, 1, 1, 1, 0}), 6)
}
