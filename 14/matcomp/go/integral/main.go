package main

import (
	"fmt"
	"math"
)

type function func(x float64) float64

func pow(n, p float64) float64 {
	return math.Pow(n, p)
}

func f(x float64) float64 {
	return .2 + (25 * x) - (200 * pow(x, 2)) + (675 * pow(x, 3)) - (900 * pow(x, 4)) + (400 * pow(x, 5))
}

// func trap(h, f0, f1 float64) float64 {
// 	return h * (f0 + f1) / 2
// }

func trapm(h float64, n int, f function) float64 {
	sum := f(0.0)
	// fmt.Printf("I = %v\n", sum)

	for i := 1; i < n; i++ {
		sum = sum + 2*f(float64(i)*h)
		// fmt.Printf("I = %v\n", sum)
	}

	sum += f(float64(n) * h)
	// fmt.Printf("I = %.3f\n", sum)

	return h * sum / 2
}

func main() {
	interval := 0.8
	n := 10
	h := float64(interval / float64(n))
	fmt.Printf("h = %v\n", h)
	fmt.Printf("RESULT: %v\n", trapm(h, n, f))
}
