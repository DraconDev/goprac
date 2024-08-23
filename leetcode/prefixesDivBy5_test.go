package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixesDivBy5(t *testing.T) {
	assert.Equal(t, []bool{true, false, true}, prefixesDivBy5([]int{0, 1, 1}))
}
