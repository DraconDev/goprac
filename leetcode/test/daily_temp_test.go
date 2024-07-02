package leetcode

import (
	"fmt"
	"goprac/leetcode"
	"testing"
)

func TestDailyTemperatures2(t *testing.T) {
	fmt.Println(leetcode.DailyTemperatures2([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func TestDivideArray(t *testing.T) {
	// fmt.Println(leetcode.DivideArray([]int{3, 2, 3, 2, 2, 2}, 2))
	// [1,3,4,8,7,9,3,5,1]
	leetcode.DivideArray([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2)
}

// func TestSequentialDigits(t *testing.T) {
// 	// fmt.Println(leetcode.SequentialDigits(1000, 13000), []int{1234, 2345, 3456, 4567, 5678, 6789, 12345})
// 	fmt.Println(leetcode.SequentialDigits(10, 1000000000), []int{1234, 2345, 3456, 4567, 5678, 6789, 12345})
// }
