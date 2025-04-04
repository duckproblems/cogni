package utils

import "math"

const Epsilon float64 = 1e-3

func Sq(val float64) float64 {
	return math.Pow(val, 2)
}
