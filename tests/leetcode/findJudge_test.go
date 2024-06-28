package leetcode_test

import (
	"goprac/leetcode"
	"testing"
)

func TestFindJudge(t *testing.T) {
	// leetcode.FindJudge(2, [][]int{{1, 2}})
	// leetcode.FindJudge(3, [][]int{{1, 3}, {2, 3}})
	leetcode.FindJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}})
	// leetcode.FindJudge(3, [][]int{{1, 2}, {2, 3}})
	// leetcode.FindJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}})
}
