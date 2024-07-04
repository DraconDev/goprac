package leetcode_test

import (
	"goprac/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeNodes(t *testing.T) {
	test := linkedlist.Buildlinkedlist()
	assert.Equal(t, linkedlist.MergeNodes(test), test)
}
