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

func TwoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := len(numbers)-1; j > i; j-- {
			if numbers[i] + numbers[j] > target {
				numbers = numbers[:j]
				continue
			}
			if numbers[i] + numbers[j] == target {
				return []int{i+1, j+1}
			}

		}
	}
	return nil
}