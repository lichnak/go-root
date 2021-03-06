// Seriál "Programovací jazyk Go"
//
// Patnáctá část
//
// Demonstrační příklad číslo 4:
//     Export obrázku o rozlišení 1×1 pixel s dvoubarevnou paletou

package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

func CreateImage(width int, height int) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{255, 128, 128, 255},
		color.RGBA{255, 255, 255, 255},
	}

	c := color.RGBA{255, 128, 128, 255}
	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)
	img.Set(0, 0, c)

	return img
}

func main() {
	img := CreateImage(1, 1)

	outfile, err := os.Create("04.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	err = gif.Encode(outfile, img, nil)
	if err != nil {
		panic(err)
	}
}
