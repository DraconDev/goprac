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
