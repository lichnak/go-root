// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 24:
//    Bitové posuny.

package main

import "fmt"

func main() {
	x := 1

	for i := uint(0); i <= 10; i++ {
		fmt.Printf("%d << %2d == %4d\n", x, i, x<<i)
	}

	fmt.Println()

	x = 10000000

	for i := uint(0); i <= 10; i++ {
		fmt.Printf("%d >> %2d == %4d\n", x, i, x>>i)
	}

}
