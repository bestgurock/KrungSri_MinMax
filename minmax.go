package main

import "sort"

func Min(data []float64) float64 {
	sort.Float64s(data)
	return data[0]
}
