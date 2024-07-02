package main

import "math"

func Movie(card, ticket int, perc float64) int {
	system_a := 0.0
	system_b := float64(card)
	i := 1
	for system_b > system_a {
		system_b += float64(ticket) * math.Pow(perc, float64(i))
		system_a += float64(ticket)
		i++
	}
	return i - 1
}
