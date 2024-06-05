package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonChars(t *testing.T) {
	assert.Equal(t, []string{"b"}, leetcode.CommonChars([]string{"bella", "label", "roller"}))
}
