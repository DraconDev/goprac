package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	// assert.Equal(t, true, leetcode.RepeatedSubstringPattern("ababab"))
	assert.Equal(t, true, leetcode.RepeatedSubstringPattern("aabaaba"))
	assert.Equal(t, true, leetcode.RepeatedSubstringPattern("aaaaaa"))
	assert.Equal(t, false, leetcode.RepeatedSubstringPattern("abcabcabcabc"))
}
