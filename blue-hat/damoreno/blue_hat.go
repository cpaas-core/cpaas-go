package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const bannerScale = 6

func drawBanner(banner *Banner, img *image.RGBA, scale, startX, startY int) {
	for i := 0; i < banner.Height; i++ {
		for j := 0; j < banner.Width; j++ {
			if banner.Data[i][j] {
				drawSquare(img, scale, startX+j*scale, startY+i*scale)
			}
		}
	}
}

func drawSquare(img *image.RGBA, scale, startX, startY int) {
	rect := image.Rect(startX, startY, startX+scale, startY+scale)
	draw.Draw(img, rect, &image.Uniform{C: color.Black}, image.Point{}, draw.Src)
}

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
	if len(os.Args) != 2 {
		log.Fatal("you need to pass the name to print in the sticker")
	}

	img, err := loadImage("red_hat.png")
	if err != nil {
		log.Fatal(err)
	}

	recoloredImage := recolorImage(img)

	banner, err := NewBanner(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	drawBanner(
		banner,
		recoloredImage,
		bannerScale,
		(recoloredImage.Rect.Max.X-banner.Width*bannerScale)/2,
		recoloredImage.Rect.Max.Y-banner.Height*bannerScale-20,
	)

	err = saveImage("blue_hat.png", recoloredImage)
	if err != nil {
		log.Fatal(err)
	}
}
