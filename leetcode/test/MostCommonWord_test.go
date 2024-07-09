package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostCommonWord(t *testing.T) {
	assert.Equal(t, leetcode.MostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}), "ball")
}
