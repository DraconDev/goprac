package leetcode_test

import (
	"goprac/tests/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRelatives(t *testing.T) {

	assert.Equal(t, leetcode.FindRelativeRanks([]int{3, 4, 5, 2, 1}), []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"})

}
