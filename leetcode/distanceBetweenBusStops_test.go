package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceBetweenBusStops(t *testing.T) {

	assert.Equal(t, 1, distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1))
}
