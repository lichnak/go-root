package main

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	x := 1
	y := 2
	z, w := swap(x, y)

	fmt.Println("z =", z)
	fmt.Println("w =", w)
}