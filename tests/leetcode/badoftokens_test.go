package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBagOfTokens(t *testing.T) {

	assert.Equal(t, leetcode.BagOfTokensScore([]int{100, 300, 200, 400}, 200), 2)
}
