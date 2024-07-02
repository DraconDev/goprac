package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructRectangle(t *testing.T) {

	t.Run("test 1", func(t *testing.T) {
		input := 14
		expected := []int{2, 2}
		assert.Equal(t, expected, leetcode.ConstructRectangle(input))
	})

}
