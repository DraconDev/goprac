package tests_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixReshape(t *testing.T) {
	test := [][]int{{1, 2, 3, 4}}
	assert.Equal(t, leetcode.MatrixReshape(test, 2, 2), [][]int{{1, 2, 3, 4}})

}
