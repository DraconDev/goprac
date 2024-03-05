package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumLength(t *testing.T) {
	assert.Equal(t, leetcode.MinimumLength("cabaabac"), 0)
	// assert.Equal(t, leetcode.MinimumLength("aabccabba"), 3)
}
