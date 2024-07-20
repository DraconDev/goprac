package leetcode

import "testing"

func TestBinaryGap(t *testing.T) {
	test := binaryGap(22)
	if test != 2 {
		t.Errorf("BinaryGap(22) = %d; want 2", test)
	}
}
