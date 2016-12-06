package euler

import (
	"fmt"

	"math"
)

const (
	constant = 1.0
)

func f(x, y float64) float64 {
	p := math.Pow
	// -2x³ + 12x² - 20x + 8.5
	return -2.0*p(x, 3.0) + 12.0*p(x, 2.0) - 20.0*x + 8.5
}

func integral(x float64) float64 {
	p := math.Pow
	// -0.5*x^4 + 4*x^3 - 10*x^2 + 8.5*x + constant
	return -0.5*p(x, 4.0) + 4.0*p(x, 3.0) - 10.0*p(x, 2.0) + 8.5*x + constant
}

func err(trueValue, aproxValue float64) float64 {
	return ((trueValue - aproxValue) / trueValue) * 100.00
}

func Euler() {
	var xi float64 = 0
	var xf float64 = 4

	var x float64 = xi
	var y float64 = 1.0

	var dx float64 = 0.5
	var nc int = int((xf - xi) / dx)

	fmt.Println("    X           True Y      Euler Y      Global Err")
	fmt.Printf("%9.4f %12.4f %12.4f\n", x, integral(x), y)

	for i := 0; i < nc; i++ {
		dydx := f(x, y)
		y = y + dydx*dx
		x = x + dx

		true_y := integral(x)
		fmt.Printf("%9.4f %12.4f %12.4f %12.4f\n", x, true_y, y, err(true_y, y))
	}

}
