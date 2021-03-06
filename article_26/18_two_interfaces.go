// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
//
// Demonstrační příklad číslo 18:
//    Prázdné rozhraní a rozhraní s jednou metodou

package main

import "fmt"

func main() {
	var i1 interface{}

	fmt.Printf("%v %v\n", i1, i1 == nil)

	var i2 interface{ Foo() }
	fmt.Printf("%v %v\n", i2, i2 == nil)

	fmt.Printf("%v\n", i1 == i2)
}
