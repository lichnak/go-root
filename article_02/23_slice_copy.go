// Seriál "Programovací jazyk Go"
//
// Druhá část
//
// Demonstrační příklad číslo 23
//    Kopie řezů polí

package main

import "fmt"

func main() {
	var a [10]int

	slice := a[:]

	fmt.Printf("Pole před modifikací: %v\n", a)
	fmt.Printf("Řez před modifikací:  %v\n", slice)

	for i := 0; i < len(a); i++ {
		a[i] = i * 2
	}

	fmt.Printf("Pole po modifikací:   %v\n", a)
	fmt.Printf("Řez po modifikaci:    %v\n", slice)

	for i := 0; i < len(slice); i++ {
		slice[i] = 42
	}

	fmt.Printf("Pole po modifikací:   %v\n", a)
	fmt.Printf("Řez po modifikaci:    %v\n", slice)
}