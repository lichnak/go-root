// Seriál "Programovací jazyk Go"
//
// Třináctá část
//
// Demonstrační příklad číslo 6B:
//     Marshalling pole struktur/záznamů do JSONu, privátní prvky struktury

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var users = [3]User{
		User{
			id:      1,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      2,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      3,
			name:    "Josef",
			surname: "Vyskočil"},
	}

	users_json, _ := json.Marshal(users)
	fmt.Println(string(users_json))
}
