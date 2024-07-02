package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetXORSum(t *testing.T) {

	assert.Equal(t, 28, leetcode.SubsetXORSum([]int{5, 1, 6}))

}
