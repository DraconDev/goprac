package tree_test

import (
	"goprac/tests/leetcode/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func TestPreorderNaryTree(t *testing.T) {

	// nodes
	nodes := &tree.Node{
		Val: 1,
		Children: []*tree.Node{
			{Val: 3, Children: []*tree.Node{{Val: 5}, {Val: 6}}},
			{Val: 2, Children: []*tree.Node{{Val: 4}}},
		},
	}

	assert.Equal(t, tree.PreorderNaryTree(nodes), []int{1, 3, 5, 6, 2, 4})

}
