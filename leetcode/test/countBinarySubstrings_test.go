package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBinarySubstrings(t *testing.T) {
	assert.Equal(t, leetcode.CountBinarySubstrings("00110011"), 6)
}
