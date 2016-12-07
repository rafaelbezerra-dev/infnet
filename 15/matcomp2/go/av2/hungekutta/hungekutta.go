package hungekutta

import "fmt"
import "math"

func heun(x, y, h float64) float64 {
	k1 := f(x, y)
	k2 := f(x+h, y+h)
	return y + ((0.5*k1)+(0.5*k2))*h
}

func mid_point(x, y, h float64) float64 {
	k1 := f(x, y)
	_x := x + (0.5 * h)
	_y := y + (0.5 * k1 * h)
	k2 := f(_x, _y)
	return y + (k2 * h)
}

func ralston(x, y, h float64) float64 {
	k1 := f(x, y)
	_x := x + ((3.0 / 4.0) * h)
	_y := y + ((3.0 / 4.0) * k1 * h)
	k2 := f(_x, _y)
	return y + (((1.0/3.0)*k1)+((2.0/3.0)*k2))*h
}

func HungeKutta2() {
	var (
		xi float64
		xf float64 = 4
		x  float64

		dx = 0.5
		nc = int((xf - xi) / dx)

		y_heun      = 1.0
		y_mid_point = 1.0
		y_ralson    = 1.0
	)

	y := integral(x)
	fmt.Println("       X       True Y     Heun Y   Heun Err   Mid Pt Y   Mid Pt Err    Ralston Y  Ralston Err")
	fmt.Printf(
		"%10.4f %10.4f %10.4f %10.4f %10.4f %10.4f    %10.4f %10.4f\n",
		x,
		y,
		y_heun, err(y, y_heun),
		y_mid_point, err(y, y_mid_point),
		y_ralson, err(y, y_ralson),
	)

	abs := math.Abs

	for i := 0; i < nc; i++ {
		// y = ralston(x, y, dx)
		y_heun = heun(x, y_heun, dx)
		y_mid_point = mid_point(x, y_mid_point, dx)
		y_ralson = ralston(x, y_ralson, dx)
		x = x + dx

		y := integral(x)

		fmt.Printf(
			"%10.4f %10.4f %10.4f %10.4f %10.4f %10.4f    %10.4f %10.4f\n",
			x,
			y,
			y_heun, abs(err(y, y_heun)),
			y_mid_point, abs(err(y, y_mid_point)),
			y_ralson, abs(err(y, y_ralson)),
		)
	}

}
