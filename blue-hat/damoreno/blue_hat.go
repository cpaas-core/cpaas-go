package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// loadImage loads and returns an image from the given path.
func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	img, _, err := image.Decode(file)

	return img, err
}

// recolorImage recolors the given image, switching red anb blue components for each pixel's RGBA.
func recolorImage(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	newImage := image.NewRGBA(bounds)
	for i := 0; i < bounds.Max.X; i++ {
		for j := 0; j < bounds.Max.Y; j++ {
			r, g, b, a := img.At(i, j).RGBA()
			newImage.Set(i, j, color.RGBA{R: uint8(b), G: uint8(g), B: uint8(r), A: uint8(a)})
		}
	}

	return newImage
}

// saveImage saves an Image in the specified path.
func saveImage(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	return png.Encode(file, img)
}

func main() {
	img, err := loadImage("red_hat.png")
	if err != nil {
		log.Fatal(err)
	}

	err = saveImage("blue_hat.png", recolorImage(img))
	if err != nil {
		log.Fatal(err)
	}
}
