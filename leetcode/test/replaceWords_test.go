package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceWords(t *testing.T) {
	assert.Equal(t, "the cattle was rattled by the battery", leetcode.ReplaceWords([]string{"catt", "cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}
