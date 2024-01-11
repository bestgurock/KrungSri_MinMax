package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin_ShouldReturn_LowestNumber(t *testing.T) {
	expected := float64(-2)

	actual := Min([]float64{6, 9, 15, -2, 92, 11})

	assert.Equal(t, expected, actual)

}
func TestMax_ShouldReturn_HighestNumber(t *testing.T) {
	expected := float64(92)

	actual := Max([]float64{6, 9, 15, -2, 92, 11})

	assert.Equal(t, expected, actual)

}
func TestSequence_ShouldReturn_NumberOfSequnce(t *testing.T) {
	expected := float64(6)

	actual := Sequence([]float64{6, 9, 15, -2, 92, 11})

	assert.Equal(t, expected, actual)

}
func TestAvg_ShouldReturn_Avg(t *testing.T) {
	expected := float64(21.833333333333332)

	actual := Avg([]float64{6, 9, 15, -2, 92, 11})

	assert.Equal(t, expected, actual)

}

// func TestMinMax_ShouldNotError(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    []float64
// 		expected float64
// 	}{
// 		{"MinNumber", []float64{6, 9, 15, -2, 92, 11}, -2},

// 		// Add more test cases as needed
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			actual := Min(tt.input)
// 			assert.Equal(t, actual, tt.expected)
// 		})
// 	}
// }
