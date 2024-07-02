package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDigits(t *testing.T) {
	assert.Equal(t, leetcode.AddDigits(38), 2)
}
