package leetcode

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	assert.Equal(t, leetcode.LengthOfLongestSubstring("abcabcbb"), 3)

	// assert.Equal(t, LengthOfLongestSubstring("bbbbb"), 1)
	// assert.Equal(t, LengthOfLongestSubstring("pwwkew"), 3)
	// assert.Equal(t, LengthOfLongestSubstring(" "), 1)
	// assert.Equal(t, LengthOfLongestSubstring("dvdf"), 3)

}
