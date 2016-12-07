package extra

import (
	"fmt"
	"math"
)

const (
	constant = 1.0
	g        = 9.81
	r        = 6370000 // 6,37 × 10^6
)

var (
	r2 = float64(r * r)
)

func f(x, y float64) float64 {
	// -(g R^2)/(R + x)^2
	p := math.Pow
	return -g * (r2 / p(r+x, 2))
}

func integral(x float64) float64 {
	// p := math.Pow
	//  (g R^2)/(R + x) + constant
	return -g*(r2/(r+x)) + constant
}

func err(trueValue, aproxValue float64) float64 {
	return ((trueValue - aproxValue) / trueValue) * 100.00
}

func Exercice25_26() {
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
