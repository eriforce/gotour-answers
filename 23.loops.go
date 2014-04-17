package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var d float64
	for i := 0; i < 10; i++ {
		d = (z*z - x) / (2 * z)
		if d == 0 {
			break
		}
		z -= d
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
