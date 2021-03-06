// Seriál "Programovací jazyk Go"
//
// Patnáctá část
//
// Demonstrační příklad číslo 10:
//     Výchozí hodnota metody použité při přepínání snímků

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
)

func CreateImage(width int, height int, imageIndex int) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{150, 150, 150, 255},
		color.RGBA{250, 0, 0, 255},
	}

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	r := image.Rect(0, 0, width, height)
	draw.Draw(img, r, &image.Uniform{palette[0]}, image.ZP, draw.Src)

	r2 := image.Rect(imageIndex*60, 0, imageIndex*60+50, height)
	draw.Draw(img, r2, &image.Uniform{palette[1]}, image.ZP, draw.Src)

	return img
}

func main() {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 6; i++ {
		img := CreateImage(350, 50, i)
		images = append(images, img)
		delays = append(delays, 100)
	}

	outfile, err := os.Create("10.gif")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	gif.EncodeAll(outfile, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
