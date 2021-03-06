// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 7:
//    Řídicí konstrukce if a hodnoty, které nejsou typu boolean.

package main

func main() {
	if 1 {
		println("true")
	}

	if 0 {
		println("false")
	}

	if !1 {
		println("false")
	}

	if !0 {
		println("true")
	}

	var b1 int = 1

	if b1 {
		println("true")
	}
	if !b1 {
		println("false")
	}

	b2 := 1

	if b2 {
		println("true")
	}
	if !b2 {
		println("false")
	}
}
