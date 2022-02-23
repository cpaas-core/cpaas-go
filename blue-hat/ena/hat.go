package main

import (
	"image"

	"image/color"
	"image/png"

	// registration of the png format
	_ "image/png"
	"os"
)

func main() {

	const sourceFile = "red_hat.png"
	const outputFile = "modified_hat.png"

	fin, err := os.Open(sourceFile)
	defer func() { fin.Close() }()
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(fin)
	if err != nil {
		panic(err)
	}

	rect := img.Bounds()
	newImg := image.NewRGBA(rect)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			rgba := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			newImg.Set(x, y, color.RGBA{R: rgba.B, G: rgba.G, B: rgba.R, A: rgba.A})
		}
	}

	fout, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer func() { fout.Close() }()

	if err = png.Encode(fout, newImg); err != nil {
		panic(err)
	}

}
