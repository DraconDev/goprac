package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLuckyNumber(t *testing.T) {
	assert.Equal(t, luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}), []int{15})
}
