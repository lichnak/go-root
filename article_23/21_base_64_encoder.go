package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

const filename = "test_output.base64"
const message = "*** Hello world! ***"

func main() {
	writer, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	base64Encoder := base64.NewEncoder(base64.RawStdEncoding, writer)
	defer base64Encoder.Close()

	buffer := []byte(message)
	written, err := base64Encoder.Write(buffer)

	if written >= 0 {
		fmt.Printf("written %d bytes\n", written)
	}

	if err != nil {
		fmt.Printf("I/O error %v\n", err)
	}
}