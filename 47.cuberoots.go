package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {
	previous, z := 0+0i, 1+1i
	for {
		z -= (cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2))
		if z == previous {
			break
		} else {
			previous = z
		}
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
}
