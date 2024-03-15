package tree_test

import (
	"goprac/tests/leetcode/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfLeftLeaves(t *testing.T) {

	assert.Equal(t, tree.SumOfLeftLeaves(tree.BuildSampleTree()), 24)

}
