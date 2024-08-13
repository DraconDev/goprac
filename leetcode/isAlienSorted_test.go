package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlienSorted(t *testing.T) {
	// assert.Equal(t, isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"), true)

	// assert.Equal(t, isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"), false)

	// assert.Equal(t, isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"), false)

	// assert.Equal(t, isAlienSorted([]string{"hello", "hello"}, "abcdefghijklmnopqrstuvwxyz"), false)

	// assert.Equal(t, isAlienSorted([]string{"reb", "inci"}, "tcyklqfhoeapndgbimsujzvxwr"), false)

	assert.Equal(t, isAlienSorted([]string{"apap", "app"}, "abcdefghijklmnopqrstuvwxyz"), false)

}
