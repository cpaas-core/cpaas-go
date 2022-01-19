package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	f, err := os.Open("red_hat.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	blue := image.NewRGBA(bounds)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := src.At(x, y)
			r, g, b, a := oldColor.RGBA()

			if r > g+b {
				blue.SetRGBA(x, y, color.RGBA{
					uint8(b),
					uint8(g),
					uint8(r),
					uint8(a),
				})
			} else {
				blue.SetRGBA(x, y, color.RGBA{
					uint8(r),
					uint8(g),
					uint8(b),
					uint8(a),
				})

			}
		}
	}

	outfile, err := os.Create("blue_hat.png")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	png.Encode(outfile, blue)
}
