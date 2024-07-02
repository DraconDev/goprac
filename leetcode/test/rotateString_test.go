package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateString(t *testing.T) {
	// assert.Equal(t, true, leetcode.RotateString("abcde", "cdeab"))

	assert.Equal(t, true, leetcode.RotateString("bbbacddceeb", "ceebbbbacdd"))

}
