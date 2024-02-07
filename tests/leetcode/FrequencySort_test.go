package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {
	assert.Equal(t, leetcode.FrequencySort("tree"), "eetr")
	assert.Equal(t, leetcode.FrequencySort("cccaaa"), "cccaaa")
	assert.Equal(t, leetcode.FrequencySort("Aabb"), "bbAa")
}
