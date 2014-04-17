package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]byte, dy)
	for y := range result {
		result[y] = make([]byte, dx)
		for x := range result[y] {
			result[y][x] = byte(x ^ y)
			//result[y][x] = byte((x + y) / 2)
			//result[y][x] = byte(x * y)
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}
