package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeCandies(t *testing.T) {
	assert.Equal(t, distributeCandie(7, 4), []int{1, 2, 3, 1})
}
