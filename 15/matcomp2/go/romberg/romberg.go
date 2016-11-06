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

func trapEq(n int, a, b float64) float64 {
	h := (b - a) / float64(n)
	return trapm(h, n, f)
}

func printI(arr [10][10]float64) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] > 0 {
				fmt.Printf("%.4f ", arr[i][j])
			} else {
				fmt.Printf("       ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
	// fmt.Printf("I[%v][%v]: %v\n", i, j, arr[i][j])
}

func romberg(a, b float64, maxit int, es float64) float64 {
	// LOCAL I (10, 10)
	I := [10][10]float64{}
	n := 1

	I[1][1] = trapEq(n, a, b)
	// printI(I)
	iter := 0

	for {
		iter++
		n = int(pow(2.0, float64(iter)))
		I[iter+1][1] = trapEq(n, a, b)
		// printI(I)

		for k := 2; k <= iter+1; k++ {
			j := 2 + iter - k
			_4k1 := pow(4.0, float64(k-1))
			I[j][k] = (_4k1*I[j+1][k-1] - I[j][k-1]) / (_4k1 - 1)
			// printI(I)

		}

		printI(I)

		ea := math.Abs((I[1][iter+1]-I[2][iter])/I[1][iter+1]) * 100.0
		fmt.Printf("error: %v\n\n\n", ea)
		if iter >= maxit || ea <= es {
			break
		}
	}

	return I[1][iter+1]
}

func main() {
	// fmt.Println("Hello ", math.Abs(-86))
	es := pow(10.0, -6)
	fmt.Printf("RESULT %v \n", romberg(0.0, 0.8, 99, es))
}
