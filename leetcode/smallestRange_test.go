package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestRange(t *testing.T) {
	assert.Equal(t, smallestRange([]int{0, 10, }, 2), 6)
}
