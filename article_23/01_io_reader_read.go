// Seriál "Programovací jazyk Go"
//
// Dvacátá třetí část
//
// Demonstrační příklad číslo 1:
//     	Čtení dat ze souboru s využitím metody Read z rozhraní Reader.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const filename = "test_input.txt"
const buffer_size = 16

func main() {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	buffer := make([]byte, buffer_size)

	for {
		read, err := reader.Read(buffer)

		if read > 0 {
			fmt.Printf("read %d bytes\n", read)
			fmt.Println(buffer[:read])
		}

		if err == io.EOF {
			fmt.Println("reached end of file")
			break
		}

		if err != nil {
			fmt.Printf("other error %v\n", err)
			break
		}
	}
}
