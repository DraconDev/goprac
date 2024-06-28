package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommon(t *testing.T) {
	assert.Equal(t, leetcode.GetCommon([]int{1, 2, 3}, []int{2, 3, 4}), 2)

}
