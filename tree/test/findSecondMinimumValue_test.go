package tree_test

import (
	"goprac/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSecondMinimumValue(t *testing.T) {
	assert.Equal(t, tree.FindSecondMinimumValue(tree.BuildSampleTree()), 2)
}
