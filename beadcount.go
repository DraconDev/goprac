package main

func CountRedBeads (n int) int {
	if n < 2 {
		return 0
	}
	return (n-1)*2
}