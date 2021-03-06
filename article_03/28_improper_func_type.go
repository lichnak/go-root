// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 28:
//    Kontroly při překladu.

package main

import "fmt"

func funkce1(x int) int {
	return 2 * x
}

func funkce2(x int, y int) int {
	return x * y
}

func main() {
	var a func(int) int
	fmt.Println(a)

	a = funkce1
	fmt.Println(a)
	fmt.Println(a(10))

	a = funkce2
	fmt.Println(a)
	fmt.Println(a(10, 20))
}
