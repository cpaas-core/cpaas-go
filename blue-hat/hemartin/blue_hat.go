package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("red_hat.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	imgRectangle := img.Bounds()
	newImg := image.NewRGBA(imgRectangle)
	for x := 0; x < imgRectangle.Max.X; x++ {
		for y := 0; y < imgRectangle.Max.Y; y++ {
			c := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			c.R, c.B = c.B, c.R
			newImg.Set(x, y, c)
		}
	}

	var resultImgBytes bytes.Buffer
	err = png.Encode(&resultImgBytes, newImg)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("blue_hat_hemartin.png", []byte(resultImgBytes.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
