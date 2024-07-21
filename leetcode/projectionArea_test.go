package leetcode

import "testing"

func TestProjectionArea(t *testing.T) {
	test := projectionArea([][]int{{1, 2}, {3, 4}})
	if test != 17 {
		t.Errorf("projectionArea([][]int{{1, 2}, {3, 4}}) = %d; want 17", test)
	}

	// test := projectionArea([][]int{{2}})
	// if test != 5 {
	// 	t.Errorf("projectionArea([][]int{{2}) = %d; want 5", test)
	// }

	// test := projectionArea([][]int{{1, 0}, {0, 2}})
	// if test != 3 {
	// 	t.Errorf("projectionArea([][]int{{1, 0}, {0, 2}}) = %d; want 3", test)
	// }
}
