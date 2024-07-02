package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeckRevealedIncreasing(t *testing.T) {

	deck := []int{17, 13, 11, 2, 3, 5, 7}
	want := []int{2, 13, 3, 11, 5, 17, 7}
	got := leetcode.DeckRevealedIncreasing(deck)
	assert.Equal(t, want, got)

}
