package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestLocal(t *testing.T) {

	// [[9,9,8,1],[5,6,2,6],[8,2,6,4],[6,2,2,2]]
	assert.Equal(t, [][]int{{9, 9}, {8, 6}}, leetcode.LargestLocal([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}))

}
