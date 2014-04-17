package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) error {
	if x < 0 {
		return ErrNegativeSqrt(x)
	}

	return nil
}

func main() {
	if err := Sqrt(-2); err != nil {
		fmt.Println(err.Error())
	}
}
