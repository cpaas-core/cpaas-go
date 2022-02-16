package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const redHat = "red_hat.png"
const blueHat = "blue_hat.png"

func main() {
	inFile, err := os.Open(redHat)
	if err != nil {
		fmt.Println("Could not open file:" + redHat)
		panic(err)
	}
	defer inFile.Close()
	oldLogo, err := png.Decode(inFile)
	if err != nil {
		fmt.Println("Could not decode image")
		panic(err)
	}
	border := oldLogo.Bounds()

	newLogo := image.NewRGBA(border)
	for x := border.Min.X; x < border.Max.X; x++ {
		for y := border.Min.Y; y < border.Max.Y; y++ {
			r, g, b, a := oldLogo.At(x, y).RGBA()
			rgba := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			if rgba.R > rgba.G+rgba.B {
				swp := rgba.R
				rgba.R = rgba.B
				rgba.B = swp
			}
			newLogo.SetRGBA(x, y, rgba)
		}
	}

	outFile, err := os.Create(blueHat)
	if err != nil {
		fmt.Println("Could not create file:" + blueHat)
		panic(err)
	}
	defer outFile.Close()
	err = png.Encode(outFile, newLogo)
	if err != nil {
		fmt.Println("Could not encode the new logo")
		panic(err)
	}
	fmt.Println("New logo created at:" + blueHat)
}
