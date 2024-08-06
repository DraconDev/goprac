package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumPushes(t *testing.T) {
	assert.Equal(t, minimumPushes("aabbccddeeffgghhiiiiii"), 12)
}
