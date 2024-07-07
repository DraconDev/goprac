package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumWaterBottles(t *testing.T) {
	// assert.Equal(t, 19, leetcode.NumWaterBottles(15, 4))
	assert.Equal(t, 13, leetcode.NumWaterBottles(9, 3))
}
