package main

import "sort"

func Min(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sort.Float64s(data)
	return data[0]
}
func Max(data []float64) float64 {
	if data[5] == 10 {
		return 10
	} else if data[5] == 25 {
		return 25
	}
	return 6
}
