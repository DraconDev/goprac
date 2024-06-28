package linkedlist_test

import (
	linkedlist "goprac/linkedlist"
	"testing"
)

func TestMergeInBetween(t *testing.T) {
	linkedlist.MergeInBetween(linkedlist.Buildlinkedlist(), 3, 4, linkedlist.Buildlinkedlist())
}
