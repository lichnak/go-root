// Seriál "Programovací jazyk Go"
//
// První část:
//    Go: minimalistický a překvapivě výkonný programovací jazyk
//    https://www.root.cz/clanky/go-minimalisticky-a-prekvapive-vykonny-programovaci-jazyk/
//
// Demonstrační příklad číslo 14:
//    Funkce se dvěma parametry a dvěma návratovými hodnotami

package main

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	x := 1
	y := 2

	var z int
	var w int

	z, w = swap(x, y)
	fmt.Println("z =", z)
	fmt.Println("w =", w)
}
