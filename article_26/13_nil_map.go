// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//
// Demonstrační příklad číslo 13:
//    Nulová mapa

package main

import "fmt"

func main() {
	var m1 map[string]int = nil
	var m2 map[string]int

	fmt.Printf("%v\n", m1)
	fmt.Printf("%v\n", m2)
}
