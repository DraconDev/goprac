package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargeGroupPositions(t *testing.T) {
	assert.Equal(t, [][]int{{3, 6}}, LargeGroupPositions("abbxxxxzzy"))
	// assert.Equal(t, [][]int{{0, 3}}, LargeGroupPositions("aaaa"))
}
