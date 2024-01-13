package main

import (
	"math"
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
	return float64(len(data))
}
func Avg(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sum := 0.0
	length := len(data)
	for i := 0; i < length; i++ {
		sum += float64(data[i])
	}
	result := sum / float64(length)
	resultformat := math.Round(result*1e6) / 1e6
	return resultformat
}
