package main

func GradingStudents(grades []int32) []int32 {
	// Write your code here
	results := []int32{}
	for _, grade := range grades {
		if grade < 38 {
			results = append(results, grade)
			continue
		}
		roundPoints := 5 - grade%5
		if roundPoints < 3 {
			results = append(results, grade+roundPoints)
		} else {
			results = append(results, grade)
		}
	}
	return results
}
