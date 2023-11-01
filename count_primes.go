package main

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) int {
	result := 0
	if n < 3 {
		return 0
	}
	for i := 2; i < n; i++ {
		if isPrime(i) {
			result++
		}
	}
	return result
}
