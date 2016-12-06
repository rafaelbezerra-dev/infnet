package main

import "github.com/rnbez/infnet/15/matcomp2/go/av2/hungekutta"

func main() {
	// var run string
	// flag.StringVar(&run, "run", "no-value", "you must provide the name of the method you want to run")
	// println(run)
	// switch run {
	// case "euler":
	// 	euler.Euler()
	// 	return
	// case "runge-kutta":
	// 	return
	// default:
	// 	println("You provided an invalid method name")
	// }

	// euler.Euler()
	// hungekutta.HungeKutta2()
	hungekutta.HungeKutta4()
}
