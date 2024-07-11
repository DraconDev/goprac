package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseParentheses(t *testing.T) {
	assert.Equal(t, "dcba", leetcode.ReverseParentheses("(abcd)"))
	// assert.Equal(t, "iloveu", leetcode.ReverseParentheses("(u(love)i)"))
	assert.Equal(t, "a(bcdefghijkl(mno)p)q", leetcode.ReverseParentheses("a(bcdefghijkl(mno)p)q"))
}
