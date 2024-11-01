// Package emc implements functions for calculating the equilibrium moisture
// content (EMC) of wood.
package emc

import "math"

// SimpsonEMC calculates the equilibrium moisture content using the equation by
// Simpson (1973), documented in USDA's "Wood Handbook".
//
// Temperature `t` is in degrees Fahrenheit.
//
// Relative Humidity `r` is a percentage, where 0.00 is 0% and 1.00 is 100%.
func SimpsonEMC(t float64, rh float64) (emc float64) {
	W := 330 + 0.452*t + 0.00415*math.Pow(t, 2)
	K := 0.791 + 0.000463*t - 0.000000844*math.Pow(t, 2)
	K1 := 6.34 + 0.000775*t - 0.0000935*math.Pow(t, 2)
	K2 := 1.09 + 0.0284*t - 0.0000904*math.Pow(t, 2)

	emc = (1800 / W) * (((K * rh) / (1 - (K * rh))) + (((K1 * K * rh) + (2 * K1 * K2 * math.Pow(K, 2) * math.Pow(rh, 2))) / (1 + (K1 * K * rh) + (K1 * K2 * math.Pow(K, 2) * math.Pow(rh, 2)))))
	return emc
}
