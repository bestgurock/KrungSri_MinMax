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
