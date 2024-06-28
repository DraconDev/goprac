package leetcode

func findRestaurant(list1 []string, list2 []string) []string {
	restaurants := make(map[string]int)
	for i, v := range list1 {
		restaurants[v] = i
	}
	result := []string{}
	min := 10000
	for i, v := range list2 {
		// check if the restaurant is in list1
		if _, ok := restaurants[v]; !ok {
			continue
		}

		if i+restaurants[v] == min {
			result = append(result, v)
		}

		if i+restaurants[v] < min {
			min = i + restaurants[v]
			result = []string{v}
		}

	}
	return result
}
