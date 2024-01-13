package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin_ShouldReturn_LowestNumber1(t *testing.T) {
	expected := float64(1)

	actual := Min([]float64{1, 2, 3, 4, 5, 6})

	assert.Equal(t, expected, actual)

}
func TestMin_ShouldReturn_LowestNumber7(t *testing.T) {
	expected := float64(7)

	actual := Min([]float64{7, 8, 9, 10, 11, 12})

	assert.Equal(t, expected, actual)

}
func TestMin_ShouldReturn_LowestNumber2(t *testing.T) {
	expected := float64(2)

	actual := Min([]float64{2, 4, 6, 8, 10, 12})

	assert.Equal(t, expected, actual)

}
func TestMin_ShouldReturn_LowestNumber5(t *testing.T) {
	expected := float64(5)

	actual := Min([]float64{10, 9, 8, 7, 6, 5})

	assert.Equal(t, expected, actual)

}
func TestMin_ShouldReturn_LowestNumber12(t *testing.T) {
	expected := float64(12)

	actual := Min([]float64{19, 25, 22, 12, 28, 26})

	assert.Equal(t, expected, actual)

}
func TestMin_ShouldReturn_LowestNumber(t *testing.T) {
	expected := float64(-2)

	actual := Min([]float64{6, 9, 15, -2, 92, 11})

	assert.Equal(t, expected, actual)

}
func TestMin_WhenDataIsEmpty_ShouldReturn0(t *testing.T) {
	expected := 0.0

	actual := Min([]float64{})

	assert.Equal(t, expected, actual)
}
func TestMax_ShouldReturn_HighestNumber6(t *testing.T) {
	expected := float64(6)

	actual := Max([]float64{1, 2, 3, 4, 5, 6})

	assert.Equal(t, expected, actual)

}
func TestMax_ShouldReturn_HighestNumber10(t *testing.T) {
	expected := float64(10)

	actual := Max([]float64{5, 6, 7, 8, 9, 10})

	assert.Equal(t, expected, actual)

}
func TestMax_ShouldReturn_HighestNumber25(t *testing.T) {
	expected := float64(25)

	actual := Max([]float64{2, 6, 9, 14, 20, 25})

	assert.Equal(t, expected, actual)

}
