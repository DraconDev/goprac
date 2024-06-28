package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSquares(t *testing.T) {
	assert.Equal(t, leetcode.NumSquares(12), 3)
	assert.Equal(t, leetcode.NumSquares(13), 2)
}
