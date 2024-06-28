package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	example "goprac/go_by_example"
	"goprac/leetcode"
	linkedlist "goprac/linkedlist"
	"goprac/tree"

	"goprac/types"
)

func TestAnagram(t *testing.T) {
	assert.Equal(t, leetcode.Anagram("anagram", "nagaram"), true)
	assert.Equal(t, leetcode.Anagram("anagram", "nagoram"), false)
	assert.Equal(t, leetcode.Anagram("anagram", "nagram"), false)
}

// func TestGroupAnargrams(t *testing.T) {
// 	assert.Equal(t, leetcode.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}})
// }

func TestProductExceptSelf(t *testing.T) {
	assert.Equal(t, leetcode.ProductExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6})
}

func TestGCDofSrings(t *testing.T) {
	assert.Equal(t, leetcode.GCDOfStrings("ABCABC", "ABC"), "ABC")
	assert.Equal(t, leetcode.GCDOfStrings("ABABAB", "ABAB"), "AB")
	assert.Equal(t, leetcode.GCDOfStrings("ABABAB", "ABCD"), "")
	assert.Equal(t, leetcode.GCDOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"), "TAUXX")
}

func TestIsValidSudoku(t *testing.T) {
	assert.Equal(t, leetcode.IsValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}), true)
}

func TestCanPlaceFlowers(t *testing.T) {
	assert.Equal(t, leetcode.CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 1), true)
	assert.Equal(t, leetcode.CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 2), false)
}

func TestLongestConsecutive(t *testing.T) {
	assert.Equal(t, leetcode.LongestConsecutive([]int{100, 4, 200, 1, 3, 2}), 4)
}

func TestIsPalindrome(t *testing.T) {
	// assert.Equal(t, leetcode.IsPalindrome("A man, a plan, a canal: Panama"), true)
	// assert.Equal(t, leetcode.IsPalindrome("race a car"), false)
	assert.Equal(t, leetcode.IsPalindrome("0P"), false)
}

func TestReverseVowels(t *testing.T) {
	assert.Equal(t, leetcode.ReverseVowels("hello"), "holle")
	assert.Equal(t, leetcode.ReverseVowels("leetcode"), "leotcede")
}

func TestReverseWords(t *testing.T) {
	assert.Equal(t, leetcode.ReverseWords("the sky is blue"), "blue is sky the")
}

func TestTwoSum(t *testing.T) {
	assert.Equal(t, leetcode.TwoSum([]int{2, 7, 11, 15}, 9), []int{1, 2})
}

func TestIncreasingTriplet(t *testing.T) {
	// assert.Equal(t, leetcode.IncreasingTriplet([]int{1, 2, 3, 4, 5}), true)
	assert.Equal(t, leetcode.IncreasingTriplet([]int{20, 100, 10, 12, 5, 13}), false)
}

func TestStringCompress(t *testing.T) {

	// assert.Equal(t, leetcode.StringCompress([]byte{'a'}), 1)
	// assert.Equal(t, leetcode.StringCompress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}), 4)
	assert.Equal(t, leetcode.StringCompress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}), 6)

}

func TestTwoSumNotSorted(t *testing.T) {
	// assert.Equal(t, leetcode.TwoSumNotSorted([]int{2, 7, 11, 15}, 9), []int{1, 2})
	assert.Equal(t, leetcode.TwoSumNotSorted([]int{3, 2, 4}, 6), []int{1, 2})
	assert.Equal(t, leetcode.TwoSumNotSorted([]int{3, 3}, 6), []int{0, 1})
}

func TestThreeSum(t *testing.T) {
	// assert.Equal(t, leetcode.ThreeSum([]int{-1, 0, 1, 2, -1, -4}), [][]int{{-1, -1, 2}, {-1, 0, 1}})
	// // * assert 4 zeros
	// assert.Equal(t, leetcode.ThreeSum([]int{0, 0, 0, 0}), [][]int{{0, 0, 0}})
	// * assert -2 0 0 2 2
	assert.Equal(t, leetcode.ThreeSum([]int{-2, 0, 0, 2, 2}), [][]int{{-2, 0, 2}})
}

// func TestMoveZeroes(t *testing.T) {
// 	assert.Equal(t, leetcode.MoveZeroes([]int{0, 1, 0, 3, 12}), []int{1, 3, 12, 0, 0})
// }

func TestMaxArea(t *testing.T) {
	assert.Equal(t, leetcode.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
}

func TestIsSubsequence(t *testing.T) {
	assert.Equal(t, leetcode.IsSubsequence("abc", "ahbgdc"), true)
	assert.Equal(t, leetcode.IsSubsequence("axc", "ahbgdc"), false)
	// assert.Equal(t, leetcode.IsSubsequence("aaaaaa", "bbaaaa"), false)

}

func TestTrapping_Water(t *testing.T) {
	assert.Equal(t, leetcode.Trapping_water([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6)
}

func TestIsValidParans(t *testing.T) {

	assert.Equal(t, leetcode.IsValidParans("()[]{}"), true)
	assert.Equal(t, leetcode.IsValidParans("(]"), false)
}

func TestMaxOperations(t *testing.T) {
	assert.Equal(t, leetcode.MaxOperations([]int{1, 2, 3, 4}, 5), 2)
	assert.Equal(t, leetcode.MaxOperations([]int{3, 1, 3, 4, 3}, 6), 1)
	// * testcase 2,2,2,3,1,1,4,1
	// assert.Equal(t, leetcode.MaxOperations([]int{2, 2, 2, 3, 1, 1, 4, 1}, 4), 2)
	// * test 3 5 1 5
	assert.Equal(t, leetcode.MaxOperations([]int{3, 5, 1, 5}, 2), 0)
	assert.Equal(t, leetcode.MaxOperations([]int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}, 2), 2)
}

func TestFindMaxAverage(t *testing.T) {
	assert.Equal(t, leetcode.FindMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4), 12.75)
}

func TestMaxVowels(t *testing.T) {
	assert.Equal(t, leetcode.MaxVowels("leetcodeisacommunityforcoders", 4), 3)
}

func TestEvalRPN(t *testing.T) {
	// assert.Equal(t, leetcode.EvalRPN([]string{"2", "1", "+", "3", "*"}), 9)
	assert.Equal(t, leetcode.EvalRPN([]string{"4", "13", "5", "/", "+"}), 6)
}

func TestLongestOnes(t *testing.T) {
	assert.Equal(t, leetcode.LongestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2), 6)
	assert.Equal(t, leetcode.LongestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3), 10)
}

func TestGenerateParenthesis(t *testing.T) {
	assert.Equal(t, leetcode.GenerateParenthesis(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
}

func TestLongestSubarrayAfterDel(t *testing.T) {

	// assert.Equal(t, leetcode.LongestSubarrayAfterDel([]int{1, 1, 0, 1}), 3)
	assert.Equal(t, leetcode.LongestSubarrayAfterDel([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}), 5)

}

func TestDailyTemperatures(t *testing.T) {
	assert.Equal(t, leetcode.DailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}), []int{1, 1, 4, 2, 1, 1, 0, 0})
}

func TestCarFleet(t *testing.T) {

	assert.Equal(t, leetcode.CarFleet(10, []int{0, 4, 2}, []int{2, 1, 3}), 1)
	assert.Equal(t, leetcode.CarFleet(10, []int{2, 4}, []int{3, 2}), 1)
	assert.Equal(t, leetcode.CarFleet(10, []int{6, 2, 17}, []int{3, 9, 2}), 2)
	assert.Equal(t, leetcode.CarFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}), 3)
}

func TestPivotIndex(t *testing.T) {
	// assert.Equal(t, leetcode.PivotIndex([]int{1, 7, 3, 6, 5, 6}), 3)
	// assert.Equal(t, leetcode.PivotIndex([]int{1, 2, 3}), -1)
	// assert.Equal(t, leetcode.PivotIndex([]int{-1, -1, -1, -1, -1, 0}), 2)

	assert.Equal(t, leetcode.PivotIndex([]int{2, 1, -1}), 0)

}

func TestLargestRectangleArea(t *testing.T) {
	assert.Equal(t, leetcode.LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}), 10)
}

func TestFindDifference(t *testing.T) {
	assert.Equal(t, leetcode.FindDifference([]int{1, 2, 3, 3}, []int{1, 2, 4}), [][]int{{3}, {4}})
}

func TestUniqueOccurrences(t *testing.T) {
	assert.Equal(t, leetcode.UniqueOccurrences([]int{1, 2, 2, 1, 1, 3}), true)
	assert.Equal(t, leetcode.UniqueOccurrences([]int{1, 2}), false)
}

func TestCloseStrings(t *testing.T) {
	// assert.Equal(t, leetcode.CloseStrings("abc", "bca"), true)
	// assert.Equal(t, leetcode.CloseStrings("a", "b"), false)

	// // * test abbbzcf babzzcz
	// assert.Equal(t, leetcode.CloseStrings("abbbzcf", "babzzcz"), false)

	// * test aaabbbbccddeeeeefffff aaaaabbcccdddeeeeffff
	assert.Equal(t, leetcode.CloseStrings("aaabbbbccddeeeeefffff", "aaaabbcccdddeeeeffff"), true)

}

func TestDeleteMiddle(t *testing.T) {
	// * make 5 length pointers of type ListNode
	// node6 := &types.ListNode{Val: 5, Next: nil}
	// node5 := &types.ListNode{Val: 1, Next: node6}
	// node4 := &types.ListNode{Val: -4, Next: node5}
	// node3 := &types.ListNode{Val: -9, Next: node4}
	// node2 := &types.ListNode{Val: -10, Next: node3}
	// node1 := &types.ListNode{Val: -10, Next: node2}
	node7 := &types.ListNode{Val: 6, Next: nil}
	node6 := &types.ListNode{Val: 2, Next: node7}
	node5 := &types.ListNode{Val: 1, Next: node6}
	node4 := &types.ListNode{Val: 7, Next: node5}
	node3 := &types.ListNode{Val: 4, Next: node4}
	node2 := &types.ListNode{Val: 3, Next: node3}
	node1 := &types.ListNode{Val: 1, Next: node2}
	assert.Equal(t, leetcode.DeleteMiddle(node1), node1)
}

func TestOddEvenList(t *testing.T) {
	// * make 5 length pointers of type ListNode
	// node6 := &types.ListNode{Val: 5, Next: nil}
	// node5 := &types.ListNode{Val: 1, Next: node6}
	// node4 := &types.ListNode{Val: -4, Next: node5}
	// node3 := &types.ListNode{Val: -9, Next: node4}
	// node2 := &types.ListNode{Val: -10, Next: node3}
	// node1 := &types.ListNode{Val: -10, Next: node2}
	node7 := &types.ListNode{Val: 7, Next: nil}
	node6 := &types.ListNode{Val: 6, Next: node7}
	node5 := &types.ListNode{Val: 5, Next: node6}
	node4 := &types.ListNode{Val: 4, Next: node5}
	node3 := &types.ListNode{Val: 3, Next: node4}
	node2 := &types.ListNode{Val: 2, Next: node3}
	node1 := &types.ListNode{Val: 1, Next: node2}
	assert.Equal(t, leetcode.OddEvenList(node1), node1)
}

// func TestBfsBinaryTree(t *testing.T) {
// 	test := tree.BuildSampleTree()
// 	assert.Equal(t, tree.BfsBinaryTree(test), []int{1, 2, 3, 4, 5, 6, 7})

// func TestTree2str(t *testing.T) {
// 	test := tree.BuildSampleTree()

// 	assert.Equal(t, tree.Tree2str(test), "1(2(4)(5))(3(6)(7))")

// }

// func TestInorderTraversal(t *testing.T) {
// 	test := tree.BuildSampleTree()
// 	assert.Equal(t, tree.InorderTraversal(test), []int{4, 2, 5, 1, 6, 3, 7})
// }

func TestLearSimilar(t *testing.T) {
	test := tree.BuildSampleTree()
	assert.Equal(t, tree.LeafSimilar(test.Left, test.Right), true)
}

func TestGoodNodes(t *testing.T) {
	test := tree.BuildSampleTree()
	assert.Equal(t, tree.GoodNodes(test), 7)
}

// func TestPathSum(t *testing.T) {
// 	test := tree.BuildSampleTree()
// 	assert.Equal(t, tree.PathSum(test, 7), 2)
// }

// func TestDfsBinaryTree(t *testing.T) {
// 	test := tree.BuildSampleTree()
// 	assert.Equal(t, tree.DfsBinaryTree(test), []int{1, 2, 4, 5, 3, 6, 7})

// }
// func TestRightSideView(t *testing.T) {
// 	test := tree.BuildSampleTree()
// 	assert.Equal(t, tree.RightSideView(test), []int{1, 3, 7})
// 	// [1,2,3,4]
// }

// func TestMaxLevelSum(t *testing.T) {
// 	test := tree.BuildSampleTree()
// 	assert.Equal(t, tree.MaxLevelSum(test), 3)
// 	// [1,2,3,4]
// }

// test getIntersectionNode
func TestGetIntersectionNode(t *testing.T) {

	test := linkedlist.Buildlinkedlist()
	assert.Equal(t, linkedlist.GetIntersectionNode(test, test), test)

}

// test convertToTitle
// func TestConvertToTitle(t *testing.T) {

// 	assert.Equal(t, leetcode.ConvertToTitle(1), "A")
// 	assert.Equal(t, leetcode.ConvertToTitle(28), "AB")
// 	assert.Equal(t, leetcode.ConvertToTitle(701), "ZY")
// }

// test doublechannel
func TestDoubleChannel(t *testing.T) {

	assert.Equal(t, example.Doublechannel(), 2)
	assert.Equal(t, example.Doublechannel(), 4)
	assert.Equal(t, example.Doublechannel(), 8)
}

// test time based key value
func TestTimeBasedKeyValue(t *testing.T) {
	var myMap = leetcode.Constructor2()
	myMap.Set("foo", "bar", 1)
	assert.Equal(t, myMap.Get("foo", 1), "bar")
	assert.Equal(t, myMap.Get("foo", 3), "bar")
	myMap.Set("foo", "bar2", 4)
	assert.Equal(t, myMap.Get("foo", 4), "bar2")
	assert.Equal(t, myMap.Get("foo", 5), "bar2")
}

// test RandomizedSet
func TestRandomizedSet(t *testing.T) {

	randSet := leetcode.Const()
	// ["RandomizedSet","insert","insert","remove","insert","remove","getRandom"] [[],[0],[1],[0],[2],[1],[]]
	randSet.Insert(3)
	randSet.Insert(4)
	randSet.Remove(3)
	// randSet.Insert(2)
	// randSet.Remove(1)
	randSet.GetRandom()

}
