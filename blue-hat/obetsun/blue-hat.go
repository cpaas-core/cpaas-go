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

const POINT_SIZE = 6
const SYM_WIDTH = 5
const SYM_HEIGHT = 7

func Repaint(inputFile string, outputFile string, name string, mapsFile string) {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Unable to open input file: %v", err)
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("Unable to decode PNG: %v", err)
	}

	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	blue := image.NewRGBA(bounds)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := src.At(x, y)
			r, g, b, a := oldColor.RGBA()

			blue.SetRGBA(x, y, color.RGBA{
				uint8(b),
				uint8(g),
				uint8(r),
				uint8(a),
			})
		}
	}

	blue = DrawName(name, ReadAsciiMap(mapsFile), outputFile, blue)

	outfile, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Unable to create output file: %v", err)
	}
	defer outfile.Close()
	png.Encode(outfile, blue)

}

func ReadAsciiMap(mapsFile string) map[string]string {
	var letters map[string]string
	letters = make(map[string]string)

	f, err := os.Open(mapsFile)

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, " ")
		letters[parts[0]] = parts[1]
	}

	if err := scanner.Err(); err != nil {

		log.Fatal(err)
	}

	return letters
}

//Scale one pixel to square of pixels
func DrawPointSquare(col int, row int, img *image.RGBA) {

	for i := col * POINT_SIZE; i < col*POINT_SIZE+POINT_SIZE; i++ {
		for j := row * POINT_SIZE; j < row*POINT_SIZE+POINT_SIZE; j++ {
			img.Set(i, j, color.RGBA{0, 0, 0, 255})
		}
	}

}

func DrawName(name string, m map[string]string, outputFile string, blue *image.RGBA) *image.RGBA {

	//create canvas for name
	width := len(name)*SYM_WIDTH*POINT_SIZE + len(name)*POINT_SIZE + 1*POINT_SIZE
	height := (SYM_HEIGHT + 2) * POINT_SIZE

	r := image.Rect(0, 0, width, height)
	img := image.NewRGBA(r)

	//Fill name background with white color
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			img.Set(i, j, color.RGBA{255, 255, 255, 255})
		}
	}

	//One pixel border around the name
	offsetCol := 1
	offsetRow := 1

	//Draw scaled name
	for i, rn := range name {

		var letter = m[string(rn)]
		for j, p := range letter {
			if p == '1' {
				row := (j / SYM_WIDTH)
				col := (j - row*SYM_WIDTH) + offsetCol
				row += offsetRow
				DrawPointSquare(col, row, img)
			}
		}
		offsetCol = (SYM_WIDTH+1)*(i+1) + 1
	}

	//Copy name to blue hat canvas
	pt := image.Point{(400 - width) / 2, (400 / 5) * 4}
	bs := img.Bounds()
	draw.Draw(blue, image.Rectangle{pt, pt.Add(bs.Size())}, img, image.Point{0, 0}, draw.Src)

	return blue

}

func main() {

	if len(os.Args[1:]) < 4 {
		fmt.Println("Usage: \"blue-hat [name] [input file] [output file] [mapsFile]\"")
		os.Exit(1)

	}
	var name = os.Args[1]

	var inputFile = os.Args[2]
	var outputFile = os.Args[3]
	var mapsFile = os.Args[4]

	Repaint(inputFile, outputFile, name, mapsFile)

}
