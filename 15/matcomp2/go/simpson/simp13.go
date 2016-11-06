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

func trap(h, f0, f1 float64) float64 {
	return h * ((f0 + f1) / 2)
}

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

func simp13(h, f0, f1, f2 float64) float64 {
	return 2 * h(f0+(4*f1)+f2) / 6
}

func simp13m(h float64, n int, f function) float64 {
	sum := f(0.0)
	// fmt.Printf("I = %v\n", sum)

	for i := 1; i < n-1; i += 2 {

		fi := f(float64(i) * h)
		fi2 := f(float64(i+1) * h)
		sum += 4*fi + 2*fi2
		// fmt.Printf("I = %v\n", sum)
	}

	fn1 := f(float64(n-1) * h)
	fn := f(float64(n) * h)
	sum += 4*fn1 + fn

	return h * sum / 3
}

func simp38(h, f0, f1, f2, f3 float64) float64 {
	return 3 * h * (f0 + (3 * (f1 + f2)) + f3) / 8
}

func simpInt(a, b float64, n int, f function) float64 {
	_n := float64(n)
	h := (b - a) / _n
	var sum float64

	if n == 1 {
		f0 := f(float64(n-1) * h)
		f1 := f(float64(n) * h)
		sum = trap(h, f0, f1)
	} else {
		m := n
		if remainder := math.Remainder(_n, 2); remainder > 0 && n > 1 {
			sum += simp38(
				h,
				f(float64(n-3)*h),
				f(float64(n-2)*h),
				f(float64(n-1)*h),
				f(float64(n)*h))

			m = n - 3
		}

		if m > 1 {
			sum += simp13m(h, m, f)
		}
	}

	return 0.0
}

func main() {
	interval := 0.8
	n := 10
	h := float64(interval / float64(n))
	fmt.Printf("h = %v\n", h)
	fmt.Printf("RESULT: %v\n", trapm(h, n, f))
}
