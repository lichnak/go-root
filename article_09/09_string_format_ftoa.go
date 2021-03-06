// Seriál "Programovací jazyk Go"
//
// Devátá část
//
// Demonstrační příklad číslo 9:
//    Základní funkce pro formátování řetězců: FormatFloat.

package main

import (
	"math"
	. "strconv"
)

func main() {
	value := math.Pi
	println(FormatFloat(value, 'f', 5, 64))
	println(FormatFloat(value, 'f', 10, 64))
	println(FormatFloat(value, 'e', 10, 64))
	println(FormatFloat(value, 'g', 10, 64))
	println(FormatFloat(value, 'b', 10, 64))

	println()

	value = 1e20
	println(FormatFloat(value, 'f', 5, 64))
	println(FormatFloat(value, 'f', 10, 64))
	println(FormatFloat(value, 'e', 10, 64))
	println(FormatFloat(value, 'g', 10, 64))
	println(FormatFloat(value, 'b', 10, 64))

	println()

	value = 0
	println(FormatFloat(value, 'f', 5, 64))
	println(FormatFloat(value, 'f', 10, 64))
	println(FormatFloat(value, 'e', 10, 64))
	println(FormatFloat(value, 'g', 10, 64))
	println(FormatFloat(value, 'b', 5, 64))
}
