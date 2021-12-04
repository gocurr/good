package sugar

import "math"

// handleFloat handles float with fn
func handleFloat(f, n float64, fn func(float64) float64) float64 {
	mul := math.Pow(10, n)
	return fn(f*mul) / mul
}

// RoundFloat returns rounded float64
func RoundFloat(f, n float64) float64 {
	return handleFloat(f, n, math.Round)
}

// CeilFloat returns ceiled float64
func CeilFloat(f, n float64) float64 {
	return handleFloat(f, n, math.Ceil)
}

// FloorFloat returns floored float64
func FloorFloat(f, n float64) float64 {
	return handleFloat(f, n, math.Floor)
}
