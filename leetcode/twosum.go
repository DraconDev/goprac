package leetcode

// func twoSum(numbers []int, target int) []int {
// 	for i := 0; i < len(numbers); i++ {
// 		for j := i + 1; j < len(numbers); j++ {
// 			if numbers[i] + numbers[j] == target {
// 				return []int{i+1, j+1}
// 			}
// 			if numbers[i] + numbers[j] > target {
// 				numbers = numbers[:j]
// 				break
// 			}
// 		}
// 	}
// 	return nil
// }

func TwoSumNotSorted(numbers []int, target int) []int {
	num_map := make(map[int][]int)
	for i, v := range numbers {
		num_map[v] = append(num_map[v], i)
	}
	for k, v := range num_map {
		if k*2 == target {
			if len(v) < 2 {
				continue
			}
			return []int{v[0], v[1]}
		}
		if _, ok := num_map[target-k]; ok {
			result := []int{num_map[k][0], num_map[target-k][0]}
			return result
		}
	}
	return nil
}

// func TwoSumNotSorted(numbers []int, target int) []int {
// 	num_map := make(map[int]int)
// 	for i, v := range numbers {
// 		num_map[i] = v
// 	}
// 	for k, v := range num_map {
// 		var test1, test2 = num_map[target-v], num_map[k]
// 		fmt.Print(test1, test2)
// 		if _, ok := num_map[target-v]; ok && k != num_map[target-v] {
// 			result := []int{k, num_map[target-v]}
// 			return result
// 		}
// 	}
// 	return nil
// }

func TwoSum(numbers []int, target int) []int {
	var l, r = 0, len(numbers) - 1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l, r}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return nil
}
