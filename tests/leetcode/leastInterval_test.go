package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastInterval(t *testing.T) {
	assert.Equal(t, leetcode.LeastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2), 8)
}
