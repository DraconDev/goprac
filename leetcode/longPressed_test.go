package leetcode

import "testing"

func TestLongPressed(t *testing.T) {
	test := isLongPressedName("alex", "aaleex")
	if test != true {
		t.Errorf("LongPressedName(alex, aaleex) = %t; want true", test)
	}
}
