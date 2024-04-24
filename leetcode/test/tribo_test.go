package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTribonacci(t *testing.T) {
	assert.Equal(t, leetcode.Tribonacci(4), 4)
}
