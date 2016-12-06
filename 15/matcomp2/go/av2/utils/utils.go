package utils

import (
	"math"
)

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

func MPow(c, x, y float64) float64 {
	if y == 1 {
		return c * x
	}
	return c * Pow(x, y)
}
