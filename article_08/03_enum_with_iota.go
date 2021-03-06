// Seriál "Programovací jazyk Go"
//
// Osmá část
//
// Demonstrační příklad číslo 3:
//    Náhrada za enum: běžné konstanty a identifikátor iota.

package main

import "fmt"

type Enum int

const (
	Pondeli Enum = iota
	Utery
	Streda
	Ctvrtek
	Patek
	Sobota
	Nedele
)

func reservation(day Enum) {
	fmt.Printf("Reservation for %d\n", day)
}

func main() {
	reservation(Pondeli)
	reservation(Sobota)
	reservation(Nedele)

	reservation(3)

	day := Pondeli
	day--
	reservation(day)
}
