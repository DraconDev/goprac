package leetcode

func countStudents(students []int, sandwiches []int) int {

	for len(students) != 0 {
		length := len(students)

		for i := 0; i < length; i++ {
			if students[0] == sandwiches[0] {
				students = students[1:]
				sandwiches = sandwiches[1:]
			} else {
				students = append(students, students[0])
				students = students[1:]
			}
		}
		if len(students) == length {
			break
		}
	}
	return len(students)
}
