package main

import (
	"fmt"
	"math"
)

type function func(x float64) float64

func pow(n, p float64) float64 {
	return math.Pow(n, p)
}

func pow2(n float64) float64 {
	return math.Pow(n, 2.0)
}

func f(x float64) float64 {
	return .2 + (25 * x) - (200 * pow(x, 2)) + (675 * pow(x, 3)) - (900 * pow(x, 4)) + (400 * pow(x, 5))
	// return pow2(float64(2*x + (3 / x)))
}

func trapEq(n int, a, b float64) float64 {
	h := (b - a) / float64(n)

	sum := f(a)
	// fmt.Printf("I = %v\n", sum)

	for i := 1; i < n; i++ {
		sum = sum + 2*f(float64(i)*h)
		// fmt.Printf("I = %v\n", sum)
	}

	sum += f(b)
	// fmt.Printf("I = %.3f\n", sum)

	return h * sum / 2
}

func simp13m(n int, a, b float64) float64 {
	h := (b - a) / float64(n)

	sum := f(a)
	// fmt.Printf("I = %v\n", sum)

	for i := 1; i < n-1; i += 2 {

		fi := f(float64(i)*h + a)
		fi2 := f(float64(i+1)*h + a)
		sum += 4*fi + 2*fi2
		// fmt.Printf("I = %v\n", sum)
	}

	fn1 := f(float64(n-1)*h + a)
	fn := f(float64(n)*h + a)
	sum += 4*fn1 + fn

	return h * sum / 3
}

func printI(arr [10][10]float64) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] > 0 {
				fmt.Printf("%9.4f ", arr[i][j])
			} else {
				fmt.Printf("         ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
}

func romberg(a, b float64, maxit int, es float64) float64 {
	I := [10][10]float64{}
	n := 1

	I[1][1] = trapEq(n, a, b)
	// printI(I)
	iter := 0

	for {
		iter++
		n = int(pow(2.0, float64(iter)))
		fmt.Printf("h%v ------\n", n)
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
	es := 0.5
	fmt.Printf("22.1) RESULT %v \n", romberg(1.0, 2.0, 99, es))
	fmt.Printf("22.1) TRAPM  %v \n", trapEq(100, 1.0, 2.0))
	fmt.Printf("22.1) SIMP13 %v \n", simp13m(100, 1.0, 2.0))

}
