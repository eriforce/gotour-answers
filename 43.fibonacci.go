package main

import "fmt"

const length = 10

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	numbers := make([]int, length)
	numbers[0] = 0
	numbers[1] = 1
	i := -1
	return func() int {
		i++
		if i > 1 {
			numbers[i] = numbers[i-2] + numbers[i-1]
		}
		return numbers[i]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < length; i++ {
		fmt.Println(f())
	}
}
