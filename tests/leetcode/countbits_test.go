package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	assert.Equal(t, leetcode.CountBits(2), []int{0, 1, 1})
}
