package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	assert.Equal(t, HammingDistance(1, 5), 2)

}
