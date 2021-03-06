// Seriál "Programovací jazyk Go"
//
// Osmá část
//
// Demonstrační příklad číslo 5:
//    Výčet jako datový typ.

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

type Den struct {
	X Enum
}

func (day Enum) String() string {
	days := []string{
		"Pondeli",
		"Utery",
		"Streda",
		"Ctvrtek",
		"Patek",
		"Sobota",
		"Nedele"}
	if day < Pondeli || day > Nedele {
		return "Unknown day"
	}
	return days[day]
}

func reservation(day Den) {
	fmt.Printf("Reservation for %s\n", day)
}

func main() {
	reservation(Den{Pondeli})
	reservation(Den{Sobota})
	reservation(Den{Nedele})

	day := Den{Pondeli}
	reservation(day)
}
