package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"your.module.name~/leetcode"
)

func TestAnagram(t *testing.T) {
	assert.Equal(t, leetcode.Anagram("anagram", "nagaram"), true)
	assert.Equal(t, leetcode.Anagram("anagram", "nagoram"), false)
	assert.Equal(t, leetcode.Anagram("anagram", "nagram"), false)
}

// func TestGroupAnargrams(t *testing.T) {
// 	assert.Equal(t, leetcode.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}})
// }

func TestProductExceptSelf(t *testing.T) {
	assert.Equal(t, leetcode.ProductExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6})
}

func TestGCDofSrings(t *testing.T) {
	assert.Equal(t, leetcode.GCDOfStrings("ABCABC", "ABC"), "ABC")
	assert.Equal(t, leetcode.GCDOfStrings("ABABAB", "ABAB"), "AB")
	assert.Equal(t, leetcode.GCDOfStrings("ABABAB", "ABCD"), "")
	assert.Equal(t, leetcode.GCDOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"), "TAUXX")
}

func TestIsValidSudoku(t *testing.T) {
	assert.Equal(t, leetcode.IsValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}), true)
}

func TestCanPlaceFlowers(t *testing.T) {
	assert.Equal(t, leetcode.CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 1), true)
	assert.Equal(t, leetcode.CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 2), false)
}

func TestLongestConsecutive(t *testing.T) {
	assert.Equal(t, leetcode.LongestConsecutive([]int{100, 4, 200, 1, 3, 2}), 4)
}

func TestIsPalindrome(t *testing.T) {
	// assert.Equal(t, leetcode.IsPalindrome("A man, a plan, a canal: Panama"), true)
	// assert.Equal(t, leetcode.IsPalindrome("race a car"), false)
	assert.Equal(t, leetcode.IsPalindrome("0P"), false)
}

func TestReverseVowels(t *testing.T) {
	assert.Equal(t, leetcode.ReverseVowels("hello"), "holle")
	assert.Equal(t, leetcode.ReverseVowels("leetcode"), "leotcede")
}

func TestReverseWords(t *testing.T) {
	assert.Equal(t, leetcode.ReverseWords("the sky is blue"), "blue is sky the")
}

func TestTwoSum(t *testing.T) {
	assert.Equal(t, leetcode.TwoSum([]int{2, 7, 11, 15}, 9), []int{1, 2})
}

func TestIncreasingTriplet(t *testing.T) {
	// assert.Equal(t, leetcode.IncreasingTriplet([]int{1, 2, 3, 4, 5}), true)
	assert.Equal(t, leetcode.IncreasingTriplet([]int{20, 100, 10, 12, 5, 13}), false)
}

func TestStringCompress(t *testing.T) {

	// assert.Equal(t, leetcode.StringCompress([]byte{'a'}), 1)
	// assert.Equal(t, leetcode.StringCompress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}), 4)
	assert.Equal(t, leetcode.StringCompress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}), 6)

}

func TestTwoSumNotSorted(t *testing.T) {
	// assert.Equal(t, leetcode.TwoSumNotSorted([]int{2, 7, 11, 15}, 9), []int{1, 2})
	assert.Equal(t, leetcode.TwoSumNotSorted([]int{3, 2, 4}, 6), []int{1, 2})
	assert.Equal(t, leetcode.TwoSumNotSorted([]int{3, 3}, 6), []int{0, 1})
}

func TestThreeSum(t *testing.T) {
	// assert.Equal(t, leetcode.ThreeSum([]int{-1, 0, 1, 2, -1, -4}), [][]int{{-1, -1, 2}, {-1, 0, 1}})
	// // * assert 4 zeros
	// assert.Equal(t, leetcode.ThreeSum([]int{0, 0, 0, 0}), [][]int{{0, 0, 0}})
	// * assert -2 0 0 2 2
	assert.Equal(t, leetcode.ThreeSum([]int{-2, 0, 0, 2, 2}), [][]int{{-2, 0, 2}})
}

// func TestMoveZeroes(t *testing.T) {
// 	assert.Equal(t, leetcode.MoveZeroes([]int{0, 1, 0, 3, 12}), []int{1, 3, 12, 0, 0})
// }

func TestMaxArea(t *testing.T) {
	assert.Equal(t, leetcode.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
}

func TestIsSubsequence(t *testing.T) {
	assert.Equal(t, leetcode.IsSubsequence("abc", "ahbgdc"), true)
	assert.Equal(t, leetcode.IsSubsequence("axc", "ahbgdc"), false)
	// assert.Equal(t, leetcode.IsSubsequence("aaaaaa", "bbaaaa"), false)

}

func TestTrapping_Water(t *testing.T) {
	assert.Equal(t, leetcode.Trapping_water([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6)
}

func TestIsValidParans(t *testing.T) {

	assert.Equal(t, leetcode.IsValidParans("()[]{}"), true)
	assert.Equal(t, leetcode.IsValidParans("(]"), false)
}
