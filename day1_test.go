package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	exampleMeasurements := []int{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	assert.Equal(t, 7, day1(exampleMeasurements))
}

func TestDay1Part2(t *testing.T) {
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	assert.Equal(t, 5, day1part2(input))
}
