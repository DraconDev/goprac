package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToBase7(t *testing.T) {
	assert.Equal(t, leetcode.ConvertToBase7(100), "202")
	assert.Equal(t, leetcode.ConvertToBase7(0), "0")
	assert.Equal(t, leetcode.ConvertToBase7(1), "1")
	assert.Equal(t, leetcode.ConvertToBase7(-7), "-10")
}
