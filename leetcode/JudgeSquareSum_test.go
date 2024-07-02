package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgeSquareSum(t *testing.T) {
	assert.Equal(t, leetcode.JudgeSquareSum(4), true)
}
