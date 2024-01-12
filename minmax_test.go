package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
