package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {

	// assert.Equal(t, leetcode.ThirdMax([]int{3, 2, 1}), 1)

	assert.Equal(t, leetcode.ThirdMax([]int{1, 2}), 1)
	// assert.Equal(t, leetcode.ThirdMax([]int{2, 2, 3, 1}), 1)
}
