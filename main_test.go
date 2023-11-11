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
