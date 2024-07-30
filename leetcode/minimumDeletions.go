package leetcode

// You are given a string s consisting only of characters 'a' and 'b'​​​​.
//
// You can delete any number of characters in s to make s balanced. s is balanced if there is no pair of indices (i,j) such that i < j and s[i] = 'b' and s[j]= 'a'.
//
// Return the minimum number of deletions needed to make s balanced.
func minimumDeletions(s string) int {
	count := 0
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			stack = append(stack, s[i])
		} else if s[i] == 'a' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
			count++
		}
	}
	return count
}
