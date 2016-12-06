package hungekutta

import "math"

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
