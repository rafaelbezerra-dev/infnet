package ex25_21

import (
	"fmt"

	"math"
)

const (
	constant = 1.0
	_k       = 0.06
)

func f(t float64) float64 {
	p := math.Pow
	// from wolfram alpha
	// y(t) = 1/4*(-2*c_1*k*t + c_1^2 + k^2*t^2)
	c_1 := 3.0 // y(0) = 3
	return (1.0 / 4.0) * (-2.0*c_1*_k*t + p(c_1, 2) + p(_k, 2)*p(t, 2))
}

func dydt(t float64) float64 {
	sq := math.Sqrt
	// -k * sqrt(x)
	return -_k * sq(f(t))
}

func integral(x float64) float64 {
	p := math.Pow
	// -2/3 k x^(3/2) + constant
	return (-2.0/3.0)*_k*p(x, 3.0/2.0) + constant
}

func err(trueValue, aproxValue float64) float64 {
	return ((trueValue - aproxValue) / trueValue) * 100.00
}

func Solve() {
	// var xi float64
	// xf := 4.0

	t := 0.0
	y := 3.0

	dt := 0.5
	max_count := 1000
	// var nc int = int((xf - xi) / dx)

	// fmt.Println("    i           True Y      Euler Y      Global Err")
	fmt.Printf("%9d %12.1f min %12.4fm\n", -1, t, y)

	for i := 0; i < max_count; i++ {
		_y := y
		dydt := dydt(t) // y(t)
		y = y + dydt*dt
		t = t + dt

		// true_y := integral(x)
		fmt.Printf("%9d %12.1f min %12.4fm %12.4f \n", i, t, y, _y-y)
		if y <= 0 {
			break
		}
	}

}
