// Seriál "Programovací jazyk Go"
//
// Třináctá část
//
// Demonstrační příklad číslo 18:
//     Pokus o unmarshalling struktury s obecnými daty

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input_json_as_bytes, err := ioutil.ReadFile("users_map_different_keys.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input (bytes):")
	fmt.Println(input_json_as_bytes)

	fmt.Println("\nInput (string):")
	fmt.Println(string(input_json_as_bytes))

	m1 := map[string]interface{}{}
	json.Unmarshal(input_json_as_bytes, &m1)

	fmt.Println("\nOutput:")
	fmt.Println(m1)

	fmt.Println("\nUsers:")
	for key, user := range m1 {
		fmt.Printf("%s\t%s\n", key, user)
	}
}
