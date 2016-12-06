package hungekutta

import "fmt"

func HungeKutta4() {
	var (
		xi float64
		xf float64 = 4
		x  float64 = xi
		y  float64 = 1.0

		h  = 0.5
		nc = int((xf - xi) / h)
	)

	fmt.Println("    X           True Y         RK4      Global Err")
	fmt.Printf("%9.4f %12.4f %12.4f %12.4f\n", x, integral(x), y, 0.0)

	for i := 0; i < nc; i++ {
		k1 := f(x, y)
		k2 := f(x+(0.5*h), y+(0.5*k1*h))
		k3 := f(x+(0.5*h), y+(0.5*k2*h))
		k4 := f(x+h, y+(k3*h))

		y = y + ((k1+2*k2+2*k3+k4)/6)*h
		x = x + h

		true_y := integral(x)
		fmt.Printf("%9.4f %12.4f %12.4f %12.4f\n", x, true_y, y, err(true_y, y))
	}

}
