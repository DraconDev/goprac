package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordPattern(t *testing.T) {
	// assert.Equal(t, leetcode.WordPattern("abba", "dog cat cat dog"), true)
	// assert.Equal(t, leetcode.WordPattern("abba", "dog cat cat fish"), false)
	// assert.Equal(t, leetcode.WordPattern("aaaa", "dog cat cat dog"), false)
	assert.Equal(t, leetcode.WordPattern("abba", "dog dog dog dog"), false)
}
