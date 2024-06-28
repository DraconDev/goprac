package main

import (
	"strconv"
)

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	sum := 0

	for i, num := range nums {
		sum += num
		result[i] = sum
	}

	return result

}

func maximumWealth(accounts [][]int) int {
	result := 0
	for _, e := range accounts {
		sum := 0
		for _, v := range e {
			sum += v
		}
		if sum > result {
			result = sum
		}
	}
	return result
}

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			result[i-1] = "Fizz"
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}

func numberOfSteps(num int) int {
	result := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
			result++
		} else {
			num -= 1
			result++
		}
		continue
	}
	return result
}

