// Package mathutil contains utility functions for working with numbers.
package mathutil

//Sqrt returns the square root of a given number
func Sqrt(x float64) float64 {

	var z float64
	z = x / 2
	aux := 1
	for aux < 15 {
		z = z - (z*z-x)/(2*z)
		if z*z == x {
			return z
		}
		aux = aux + 1
	}
	return z
}
