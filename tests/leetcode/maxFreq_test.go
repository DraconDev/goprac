package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFrequencyElements(t *testing.T) {

	assert.Equal(t, leetcode.MaxFrequencyElements([]int{1, 2, 2, 3, 1, 4}), 4)

}
