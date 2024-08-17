package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToArrayForm(t *testing.T) {

	// assert.Equal(t, []int{1, 2, 3, 4}, addToArrayForm([]int{1, 2, 0, 0}, 34))
	assert.Equal(t, []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 5, 7, 9}, addToArrayForm([]int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}, 516))

}
