package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUsLength(t *testing.T) {

	assert.Equal(t, leetcode.FindLUSlength("aba", "cdc"), 3)

}
