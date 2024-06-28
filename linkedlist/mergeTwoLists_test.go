package linkedlist_test

import (
	linkedlist "goprac/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	assert.Equal(t, linkedlist.MergeTwoLists(linkedlist.Buildlinkedlist(), linkedlist.Buildlinkedlist()), linkedlist.Buildlinkedlist())
}
