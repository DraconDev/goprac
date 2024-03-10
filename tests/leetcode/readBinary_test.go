package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadBinaryWatch(t *testing.T) {
	assert.Equal(t, leetcode.ReadBinaryWatch(1), []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"})
}
