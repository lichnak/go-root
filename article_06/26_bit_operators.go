// Seriál "Programovací jazyk Go"
//
// Šestá část
//
// Demonstrační příklad číslo 26:
//    Bitové operátory.

package main

import "fmt"

func main() {
	x := 1
	y := 0xfe

	fmt.Printf("%x & %x == %x\n", x, y, x&y)
	fmt.Printf("%x | %x == %x\n", x, y, x|y)
	fmt.Printf("%x ^ %x == %x\n", x, y, x^y)

	x ^= y
	fmt.Printf("new x = %x\n", x)

	x |= y
	fmt.Printf("new x = %x\n", x)

	x ^= y
	fmt.Printf("new x = %x\n", x)
}
