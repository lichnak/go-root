// Seriál "Programovací jazyk Go"
//
// Čtrnáctá část
//
// Demonstrační příklad číslo 4:
//     Rastrový obrázek vyplněný konstantní barvou, použití metody Set

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const width = 256
const height = 256

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	outfile, err := os.Create("04.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	c := color.RGBA{0, 255, 0, 255}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, c)
		}
	}
	png.Encode(outfile, img)
}
