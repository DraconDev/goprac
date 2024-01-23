package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"goprac/leetcode"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	assert.Equal(t, leetcode.LengthOfLongestSubstring("abcabcbb"), 3)

	// assert.Equal(t, LengthOfLongestSubstring("bbbbb"), 1)
	// assert.Equal(t, LengthOfLongestSubstring("pwwkew"), 3)
	// assert.Equal(t, LengthOfLongestSubstring(" "), 1)
	// assert.Equal(t, LengthOfLongestSubstring("dvdf"), 3)

}

func TestClimbStairs(t *testing.T) {
	// assert.Equal(t, leetcode.ClimbStairs(2), 2)
	assert.Equal(t, leetcode.ClimbStairs(3), 3)
	assert.Equal(t, leetcode.ClimbStairs(4), 5)
}

func TestMinFallingPathSum(t *testing.T) {
	assert.Equal(t, leetcode.MinFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}), 13)
}

func TestSumSubarrayMins(t *testing.T) {
	assert.Equal(t, leetcode.SumSubarrayMins([]int{3, 1, 2, 4}), 17)
}

func TestCharacterReplacement(t *testing.T) {
	assert.Equal(t, leetcode.CharacterReplacement("ABAA", 0), 4)
}

func TestMaximumStrongPairXor(t *testing.T) {
	assert.Equal(t, leetcode.MaximumStrongPairXor([]int{1, 1, 2, 3, 5}), 6)
}

func TestFindErrorNums(t *testing.T) {
	assert.Equal(t, leetcode.FindErrorNums([]int{1, 2, 2, 4}), []int{2, 3})
}

func TestMaxLengthStringComb(t *testing.T) {

	// ["cha","r","act","ers"]
	assert.Equal(t, leetcode.MaxLength([]string{"cha", "r", "act", "ers"}), 6)

	// assert.Equal(t, leetcode.MaxLengthStringComb([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}), 26)

}
