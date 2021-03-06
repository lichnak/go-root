// Seriál "Programovací jazyk Go"
//
// Dvacátá pátá část
//
// Demonstrační příklad číslo 4:
//     	Rekurzivní výpočet faktoriálu

package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}

func main() {
	fmt.Println(factorial(4))
}
