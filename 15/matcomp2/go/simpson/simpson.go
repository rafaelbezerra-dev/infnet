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
	return 2 * h * (f0 + (4 * f1) + f2) / 6
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

func simp38m(a float64, b float64, n int, f function) float64 {
	h := float64((b - a) / float64(n))

	sum := 0.0
	x := 0.0
	_h := h / 3
	for i := 0; i < n; i++ {
		x = a + (float64(i) * h) + (_h * 0)
		f0 := f(x)

		x = a + (float64(i) * h) + (_h * 1)
		f1 := f(x)

		x = a + (float64(i) * h) + (_h * 2)
		f2 := f(x)

		x = a + (float64(i) * h) + (_h * 3)
		f3 := f(x)

		sum += simp38(_h, f0, f1, f2, f3)
	}

	return sum
}

// func simp38m(a float64, b float64, n int, f function) float64 {
// 	n *= 4
// 	h := float64((b - a) / float64(n))
// 	// fmt.Printf("h: %v \n-----\n", h)

// 	// sum := f(a)
// 	// fmt.Printf("- %v \n-----\n", a)

// 	for i := 0.0; i < float64(n-4); i += 4.0 {
// 		f0 := f((i * h) + a)
// 		// fmt.Printf("- %v \n", i*h)

// 		f1 := f(((i + 1) * h) + a)
// 		// fmt.Printf("- %v \n", (i+1)*h)

// 		f2 := f(((i + 2) * h) + a)
// 		// fmt.Printf("- %v \n", (i+2)*h)

// 		f3 := f(((i + 3) * h) + a)
// 		// fmt.Printf("- %v \n", (i+3)*h)

// 		sum += simp38(h, f0, f1, f2, f3)
// 		// fmt.Printf("---\n")
// 	}

// 	f0 := f(b - 3*h)
// 	// fmt.Printf("- %v \n", b-3*h)
// 	f1 := f(b - 2*h)
// 	// fmt.Printf("- %v \n", b-2*h)
// 	f2 := f(b - 1*h)
// 	// fmt.Printf("- %v \n", b-1*h)
// 	f3 := f(b)
// 	// fmt.Printf("- %v \n", b)
// 	sum += simp38(h, f0, f1, f2, f3)

// 	return sum * (3 * h / 8)
// }

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

	return sum
}

func _runSimp13() {
	interval := 0.8
	n := 10
	h := float64(interval / float64(n))
	fmt.Printf("SIMP 1/3 n: %v \n", n)
	fmt.Printf("RESULT (SIMP 1/3): %v\n", simp13m(h, n, f))
}

func _runSimp38() {
	a, b := 0.0, 0.8
	n := 10
	fmt.Printf("SIMP 3/8 n: %v \n", n)
	fmt.Printf("RESULT (SIMP 3/8): %v\n", simp38m(a, b, n, f))
}

func _runSimpInt() {
	a, b := 0.0, 0.8
	n := 10
	fmt.Printf("SIMP INT n: %v \n", n)
	fmt.Printf("RESULT (SIMP INT): %v\n", simpInt(a, b, n, f))

	n = 11
	fmt.Printf("SIMP INT n: %v \n", n)
	fmt.Printf("RESULT (SIMP INT): %v\n", simpInt(a, b, n, f))
}

func main() {
	_runSimp13()
	fmt.Println("")
	_runSimp38()
	fmt.Println("")
	_runSimpInt()
	fmt.Println("")

}
