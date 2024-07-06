package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestTriangleArea(t *testing.T) {

	assert.Equal(t, 5.5, leetcode.LargestTriangleArea([][]int{{4, 6}, {6, 5}, {3, 1}}))

}
