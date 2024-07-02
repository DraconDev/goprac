package leetcode

func IsPalindrome2(s string) bool {

	left, right := 0, len(s)-1

	for left < right {

		if s[left] != s[right] {
			return false
		}

		left, right = left+1, right-1
	}

	return true
}

func dfs(s string, partition *[]string, result *[][]string) {

	// Base case:
	// Empty string must be palindrome
	if 0 == len(s) {

		// Make a copy of current partition, and save into result
		goodPartition := make([]string, len(*partition))
		copy(goodPartition, *partition)

		*result = append(*result, goodPartition)
		return
	}

	// General case:
	for i := 1; i <= len(s); i++ {

		prefix := s[0:i]
		postfix := s[i:]

		if IsPalindrome2(prefix) {

			*partition = append(*partition, prefix)

			dfs(postfix, partition, result)

			// golang doesn't have built-in pop function for slice
			*partition = (*partition)[:len(*partition)-1]
		}
	}
	return
}

func Partition(s string) [][]string {

	// buffer for partition in DFS
	part := make([]string, 0)

	// final output of palindrome substrins
	result := make([][]string, 0)

	dfs(s, &part, &result)
	return result
}
