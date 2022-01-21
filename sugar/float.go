package sugar

import "math"

// handleFloat returns float64 by
// value: f, decimal-places keeper: n, function: fn.
func handleFloat(f, n float64, fn func(float64) float64) float64 {
	mul := math.Pow(10, n)
	return fn(f*mul) / mul
}

// RoundFloat returns rounded float64 that keeps n decimal places.
func RoundFloat(f, n float64) float64 {
	return handleFloat(f, n, math.Round)
}

// CeilFloat returns ceiled float64 that keeps n decimal places.
func CeilFloat(f, n float64) float64 {
	return handleFloat(f, n, math.Ceil)
}

// FloorFloat returns floored float64 that keeps n decimal places.
func FloorFloat(f, n float64) float64 {
	return handleFloat(f, n, math.Floor)
}

// FloatQuotient returns quotient of two integers.
func FloatQuotient(a, b int) float64 {
	// Avoid NaN error.
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}
