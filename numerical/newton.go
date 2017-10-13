package main

import (
	"fmt"
	"math"
)

func newton(x float64, epsilon float64) float64  {
	z := x/2
	i := 0
	for ; math.Abs(x - math.Pow(z, 2)) > epsilon; z = (z - (math.Pow(z, 2) - x)/(2*z)) {
		i++
	}
	fmt.Printf("Took %d iterations\n", i)
	return z
}

func main() {
	fmt.Printf("sqrt(%.1f) is %.4f\n", 9.0, newton(9.0, 0.0001))
}