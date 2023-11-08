package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"your.module.name~/leetcode"
)

func Test(t *testing.T) {

	// assert.Equal(t, ReverseBits(43261596), 964176192)
	// assert.Equal(t, leetcode.Pascals_triangle(2), [][]int{{1}, {1, 1}})
	// assert.Equal(t, leetcode.Pascals_triangle(5), [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}})
	assert.Equal(t, leetcode.Anagram("anagram", "nagaram"), true)
	assert.Equal(t, leetcode.Anagram("anagram", "nagoram"), false)
	assert.Equal(t, leetcode.Anagram("anagram", "nagram"), false)
	// assert.Equal(t, HammingDistance(1, 5), 2)

}
