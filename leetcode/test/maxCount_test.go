package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCount(t *testing.T) {

	assert.Equal(t, leetcode.MaxCount(3, 3, [][]int{{0, 0}, {2, 2}, {3, 3}}), 3)
}
