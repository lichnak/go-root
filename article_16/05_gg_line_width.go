// Seriál "Programovací jazyk Go"
//
// Šestnáctá část
//
// Demonstrační příklad číslo 5:
//     Nastavení šířky čáry.

package main

import "github.com/fogleman/gg"

func main() {
	const width = 320
	const height = 240

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 0, width, height)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()

	dc.SetRGBA(0.0, 0.0, 0.0, 1)

	for i := 0; i < 256; i += 16 {
		width := float64(i) / 20
		dc.SetLineWidth(width)

		x := float64(i + 32)
		dc.DrawLine(x, 20, x, height-20)

		dc.Stroke()
	}

	dc.SavePNG("05.png")
}
