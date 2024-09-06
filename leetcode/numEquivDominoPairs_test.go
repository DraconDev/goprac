package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumEquivDominoPairs(t *testing.T) {

	// assert.Equal(t, 1, numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))

	assert.Equal(t, 3, numEquivDominoPairs([][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}))
}