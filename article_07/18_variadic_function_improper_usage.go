// Seriál "Programovací jazyk Go"
//
// Sedmá část
//
// Demonstrační příklad číslo 18:
//    Variadické funkce.

package main

import "fmt"

func f4(prefix string, parts1 ...int, parts2 ...string) {
	fmt.Println(prefix)
	for _, val := range parts1 {
		fmt.Printf("%s ", val)
	}
	fmt.Println()
	for _, val := range parts2 {
		fmt.Printf("%s ", val)
	}
	fmt.Println()
}

func main() {
	f4("Message:", 1, 2, 3, "Hello", "world", "again", "!")
}
