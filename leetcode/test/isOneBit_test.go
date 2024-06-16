package leetcode

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOneBitCharacter(t *testing.T) {

	assert.Equal(t, leetcode.IsOneBitCharacter([]int{1, 0, 0}), true)
}
