package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestToChar(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}, leetcode.ShortestToChar("loveleetcode", 'e'))
}
