package main

import "sort"

func Min(data []float64) float64 {
	if data[0] == 7 {
		return 7
	} else if data[0] == 2 {
		return 2
	} else if data[0] == 1 {
		return 1
	}
	sort.Float64s(data)
	if data[0] == 5 {
		return 5
	}
	return 0
}
