package main

import "fmt"

func main() {
	var s []byte = []byte("www.root.cz")

	fmt.Println(string(s))
	s[3] = '*'
	s[8] = '*'
	fmt.Println(string(s))
}
