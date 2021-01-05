package main

import (
	"fmt"
	"github.com/VividCortex/ewma"
)

func main() {

	//5594, 5149,
	samples := [20]float64{
		4599, 5711, 4746, 4621, 5037, 4218, 4925, 4281, 5207, 5203,
	}

	e := ewma.NewMovingAverage()  //=> Returns a SimpleEWMA if called without params
	a := ewma.NewMovingAverage(1) //=> returns a VariableEWMA with a decay of 2 / (5 + 1)

	for i, f := range samples {
		if f == 0 {
			f = e.Value()
		}
		e.Add(f)

		fmt.Println("e", i+1, f, e.Value())

		a.Add(f)
		fmt.Println("a", i+1, f, a.Value())
		fmt.Print("\n\n")
	}
}
