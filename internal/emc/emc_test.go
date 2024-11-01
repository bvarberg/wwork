package emc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpsonEquilibriumCalculation(t *testing.T) {
	// Cases created through sampling values in the table within "With the Grain"
	// by Christian Becksvoort
	var cases = []struct {
		temperature      float64
		relativeHumidity float64
		want             float64
	}{
		// Basics
		{100, 0.50, 8.7},
		{70, 0.40, 7.7},
		{50, 0.20, 4.6},
		// Extremes
		{30, 0.05, 1.4},
		{270, 0.30, 0.4},
		{90, 0.95, 23.3},
	}

	var delta float64 = 0.1

	for _, tt := range cases {
		got := Simpson(tt.temperature, tt.relativeHumidity)
		assert.InDelta(t, tt.want, got, delta, "Equilibrium(%f, %f) = %f; want %f within %f", tt.temperature, tt.relativeHumidity, got, tt.want, delta)
	}
}
