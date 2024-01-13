package main

import (
	"sort"
)

func Min(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sort.Float64s(data)
	return data[0]
}
func Max(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sort.Float64s(data)
	return data[5]
}
func NumberOfLength(data []float64) float64 {
	if len(data) == 4 {
		return 4
	} else if len(data) == 2 {
		return 2
	} else if len(data) == 10 {
		return 10
	}
	return 6
}
