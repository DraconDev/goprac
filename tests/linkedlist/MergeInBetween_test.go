package linkedlist_test

import (
	linkedList "goprac/linkedlist"
	"testing"
)

func TestMergeInBetween(t *testing.T) {
	linkedList.MergeInBetween(linkedList.BuildLinkedList(), 3, 4, linkedList.BuildLinkedList())
}
