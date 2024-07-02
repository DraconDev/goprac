package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	result := map[string]int{}
	for _, v := range apples {
		if v+a >= s && v+a <= t {
			result["apple"] += 1
		}
	}
	for _, v := range oranges {
		if v+b >= s && v+b <= t {
			result["orange"] += 1
		}
	}
	fmt.Println(result["apple"])
	fmt.Println(result["orange"])
}
