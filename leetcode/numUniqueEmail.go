package leetcode

func numUniqueEmails(emails []string) int {

	m := make(map[string]bool)
	for _, v := range emails {
		email := ""
		for i := 0; i < len(v); i++ {
			if v[i] == '+' {
				for j := i + 1; j < len(v); j++ {
					if v[j] == '@' {
						email += v[j:]
						break
					}
				}
				break
			} else if v[i] == '.' {
				continue
			} else if v[i] == '@' {
				email += v[i:]
				break
			} else {
				email += string(v[i])
			}

		}
		m[email] = true
	}
	return len(m)
}
