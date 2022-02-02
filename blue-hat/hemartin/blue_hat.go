package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"
)

var asciiMapFilename = "ascii_maps"
var originalImageFilename = "red_hat.png"
var destinationImageFilename = "blue_hat.png"
var name = "hector"

const drawMarginBottom = 15
const maxChars = 11
const charSpacing = 1
const charPixelWidth = 5
const charPixelHeight = 7
const charScaleFactor = 6
const rectHeight = charPixelHeight * charScaleFactor

var rectWidth = (len(name) * charPixelWidth * charScaleFactor) + (len(name) * charSpacing * charScaleFactor)

type ColoredPixel struct {
	point image.Point
	color color.Color
}

func main() {
	imgReader, err := os.Open(originalImageFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer imgReader.Close()

	redHatImg, _, err := image.Decode(imgReader)
	if err != nil {
		fmt.Println(err)
		return
	}

	imgRectangle := redHatImg.Bounds()
	blueHatImg := image.NewRGBA(imgRectangle)
	for x := 0; x < imgRectangle.Max.X; x++ {
		for y := 0; y < imgRectangle.Max.Y; y++ {
			c := color.RGBAModel.Convert(redHatImg.At(x, y)).(color.RGBA)
			c.R, c.B = c.B, c.R
			blueHatImg.Set(x, y, c)
		}
	}

	asciiMap, err := createAsciiMap(asciiMapFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	nameImg := createNameImg(name, asciiMap)
	nameRect := nameImg.Bounds()

	drawOrigin := image.Point{blueHatImg.Rect.Max.X/2 - nameRect.Max.X/2, blueHatImg.Rect.Max.Y - nameRect.Max.Y - drawMarginBottom}
	drawRect := image.Rectangle{drawOrigin, drawOrigin.Add(nameImg.Bounds().Size())}
	draw.Draw(blueHatImg, drawRect, nameImg, nameRect.Min, draw.Src)

	destFile, err := os.Create(destinationImageFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destFile.Close()

	err = png.Encode(destFile, blueHatImg)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createNameImg(name string, asciiMap map[string]string) image.Image {
	nameRect := image.Rect(0, 0, rectWidth, rectHeight)
	nameImg := image.NewRGBA(nameRect)
	draw.Draw(nameImg, nameImg.Bounds(), image.White, image.Point{}, draw.Src)

	for charPosition, char := range name {
		encoding := asciiMap[string(char)]
		pixelOffset := charPosition * charScaleFactor

		for pixelPosition, pixelFlag := range encoding {
			hPosition := (pixelPosition % charPixelWidth) + pixelOffset
			vPosition := pixelPosition / charPixelWidth

			if string(pixelFlag) == "1" {
				// We need a 6x6 image to emulate an scaled pixel, so we create a new image here and paint it black
				pixelImg := image.NewRGBA(image.Rect(0, 0, charScaleFactor, charScaleFactor))
				draw.Draw(pixelImg, pixelImg.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

				drawOrigin := image.Point{hPosition * charScaleFactor, vPosition * charScaleFactor}
				drawRect := image.Rectangle{drawOrigin, drawOrigin.Add(pixelImg.Rect.Size())}
				draw.Draw(nameImg, drawRect, pixelImg, pixelImg.Rect.Min, draw.Src)
			}
		}
	}

	return nameImg
}

func createAsciiMap(asciiMapFilename string) (map[string]string, error) {
	asciiReader, err := os.Open(asciiMapFilename)
	if err != nil {
		return nil, err
	}
	defer asciiReader.Close()

	asciiMap := map[string]string{}
	asciiScanner := bufio.NewScanner(asciiReader)
	for asciiScanner.Scan() {
		line := asciiScanner.Text()
		elements := strings.SplitN(line, " ", 2)
		asciiMap[elements[0]] = elements[1]
	}

	if err := asciiScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return asciiMap, nil
}
