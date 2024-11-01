// Package emc implements functions for calculating the equilibrium moisture
// content (EMC) of wood.
package emc

// Farenheit is a temperature in degrees Farenheit
type Farhenheit float64

// RelativeHumidity is a percentage value, where 0.0 is 0% and 100.0 is 100%
type RelativeHumidity float64

// MoistureContent is a percentage value, where 0.0 is 0% and 100.0 is 100%
type MoistureContent float64

// Equilibrium returns the equilibrium mosture content of wood for the given
// climatic conditions.
func Equilibrium(t Farhenheit, rh RelativeHumidity) MoistureContent {
	return 9.0
}
