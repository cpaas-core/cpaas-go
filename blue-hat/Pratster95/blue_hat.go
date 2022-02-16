package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	redHat, err := os.Open("red_hat.png")
	if err != nil {
		log.Fatal(err)
	}
	defer redHat.Close()

	imData, _, err := image.Decode(redHat)
	if err != nil {
		fmt.Println(err)
	}

	redHat_bounds := imData.Bounds()
	blueHat := image.NewRGBA(redHat_bounds)
	for y := redHat_bounds.Min.Y; y < redHat_bounds.Max.Y; y++ {
		for x := redHat_bounds.Min.X; x < redHat_bounds.Max.X; x++ {
			red_hat_color := imData.At(x, y)
			R, G, B, A := red_hat_color.RGBA()
			if R > G+B {
				blueHat.Set(x, y, color.RGBA{uint8(B), uint8(G), uint8(R), uint8(A)})
			} else {
				blueHat.Set(x, y, color.RGBA{uint8(R), uint8(G), uint8(B), uint8(A)})
			}

		}
	}

	out, err := os.Create("blue_hat.png")
	if err != nil {
		fmt.Println(err)
	}
	png.Encode(out, blueHat)
	fmt.Println("Image Saved....")
	out.Close()
}
