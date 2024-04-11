package leetcode_test

import (
	"goprac/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLicenseKeyFormatting(t *testing.T) {
	assert.Equal(t, leetcode.LicenseKeyFormatting("5F3Z-2e-9-w", 4), "5F3Z-2E9W")
	assert.Equal(t, leetcode.LicenseKeyFormatting("2-5g-3-J", 2), "2-5G-3J")
}
