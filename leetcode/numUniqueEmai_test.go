package leetcode

import "testing"

func TestNumUniqueEmails(t *testing.T) {
	// test := numUniqueEmails([]string{"a@b.com", "a@b.com", "a@b.com"})
	// if test != 1 {
	// 	t.Errorf("numUniqueEmails() = %d; want 1", test)
	// }
	// "test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"
	test := numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})
	if test != 2 {
		t.Errorf("numUniqueEmails() = %d; want 1", test)
	}

}
