package linkedlist_test

import (
	linkedList "goprac/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	assert.Equal(t, linkedList.MergeTwoLists(linkedList.BuildLinkedList(), linkedList.BuildLinkedList()), linkedList.BuildLinkedList())
}
