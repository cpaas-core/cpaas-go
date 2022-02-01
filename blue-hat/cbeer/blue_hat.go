package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"
)

func main() {
	infilePtr := flag.String("infile", "red_hat.png", "Input file name")
	outfilePtr := flag.String("outfile", "blue_hat.png", "Output file name")
	userNamePtr := flag.String("userName", "Chris", "User name for the sticker")

	flag.Parse()

	fmt.Println("input file name: ", *infilePtr)
	fmt.Println("output file name:", *outfilePtr)
	fmt.Println("user name:       ", *userNamePtr)

	if len(*userNamePtr) > 11 {
		log.Fatal("User name too long, maximum of 11 characters")
	}

	infile, err := os.Open(*infilePtr)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to open input file: %v\n", err))
	}
	defer infile.Close()

	redHatImage, _, err := image.Decode(infile)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to decode input file: %v\n", err))
	}

	blueHatImage := ChangeImageColor(redHatImage)
	blueHatImageWithName := AddUserName(blueHatImage, *userNamePtr)

	outfile, err := os.Create(*outfilePtr)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create output file: %v\n", err))
	}
	defer outfile.Close()

	err = png.Encode(outfile, blueHatImageWithName)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to encode output file: %v\n", err))
	}
}

// Given an image, return an image with the red changed to blue
func ChangeImageColor(redHatImage image.Image) image.Image {
	blueHatImage := image.NewRGBA(redHatImage.Bounds())
	for y := redHatImage.Bounds().Min.Y; y < redHatImage.Bounds().Max.Y; y++ {
		for x := redHatImage.Bounds().Min.X; x < redHatImage.Bounds().Max.X; x++ {
			pixelColor := redHatImage.At(x, y)
			r, g, b, a := pixelColor.RGBA()
			if r > (g + b) {
				newColor := color.RGBA{uint8(b), uint8(g), uint8(r), uint8(a)}
				blueHatImage.Set(x, y, newColor)
			} else {
				blueHatImage.Set(x, y, pixelColor)
			}
		}
	}
	return blueHatImage
}

// Read the ascii_maps file and create the CharacterMap from it
func CreateCharacterMap() map[string]string {
	asciiMapsFile, err := os.Open("ascii_maps")
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to open file ascii_maps: %v\n", err))
	}

	characterMap := make(map[string]string)
	scanner := bufio.NewScanner(asciiMapsFile)
	for scanner.Scan() {
		breakdown := strings.Fields(scanner.Text())
		characterMap[breakdown[0]] = breakdown[1]
	}

	asciiMapsFile.Close()
	return characterMap
}

// Convert the characterRender string into a panel (map of lines)
func CreateCharacterPanel(characterRender string) [7]string {
	var characterPanel [7]string

	for i := 0; i < 7; i++ {
		firstChar := i * 5
		lastChar := firstChar + 5
		characterPanel[i] = characterRender[firstChar:lastChar]
	}

	return characterPanel
}

var characterColor = color.RGBA{0, 0, 0, 255}
var characterWidth = 6 // including space
var characterHeight = 7
var scaleFactor = 6
var letterTop = 330

// Add the user name to the image
func AddUserName(blueHatImage image.Image, userName string) image.Image {
	blueHatImageWithName := image.NewRGBA(blueHatImage.Bounds())
	draw.Draw(blueHatImageWithName, blueHatImage.Bounds(), blueHatImage, blueHatImage.Bounds().Min, draw.Src)
	characterMap := CreateCharacterMap()

	// figure out left offset for centering.
	imageWidth := blueHatImage.Bounds().Max.X
	nameWidth := len(userName) * characterWidth * scaleFactor
	leftOffset := (imageWidth - nameWidth) / 2

	for letterPosition, userNameRune := range userName {
		characterRender := characterMap[string(userNameRune)]
		characterPanel := CreateCharacterPanel(characterRender)
		for i := 0; i < characterHeight; i++ {
			for idx, e := range characterPanel[i] {
				if string(e) != "0" {
					rectX := leftOffset + (letterPosition * characterWidth * scaleFactor) + (idx * scaleFactor)
					rectY := letterTop + (i * scaleFactor)
					my_rect := image.Rect(rectX, rectY, rectX+scaleFactor, rectY+scaleFactor)
					draw.Draw(blueHatImageWithName, my_rect, &image.Uniform{characterColor}, image.ZP, draw.Src)
				}
			}
		}
	}
	return blueHatImageWithName
}
