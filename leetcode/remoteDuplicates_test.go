package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoteDuplicates(t *testing.T) {

	assert.Equal(t, removeDuplicates("abbaca"), "ca")
}
