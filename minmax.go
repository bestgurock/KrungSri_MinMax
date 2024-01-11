package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []float64{6, 9, 15, -2, 92, 11}
	fmt.Println(Min(data))
	fmt.Println(Max(data))
	fmt.Println(Sequence(data))
	fmt.Printf("%.6f\n", Avg(data))
}

func Min(data []float64) float64 {
	sort.Float64s(data)
	min := data[0]
	return min
}

func Max(data []float64) float64 {
	sort.Float64s(data)
	max := data[5]
	return max
}

func Sequence(data []float64) float64 {
	return float64(len(data))
}

func Avg(data []float64) float64 {
	sum := 0.0
	length := len(data)
	for _, value := range data {
		sum += value
	}
	avg := sum / float64(length)
	return avg
}
