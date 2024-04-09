package tests_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeReqToBuy(t *testing.T) {
	assert.Equal(t, leetcode.TimeRequiredToBuy([]int{2, 3, 2}, 2), 6)
}
