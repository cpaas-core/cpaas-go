package main

import (
  "flag"
  "fmt"
  "os"
  "image"
  "image/png"
  "image/color"
)

func main() {
  infilePtr := flag.String("infile", "red_hat.png", "Input file name")
  outfilePtr := flag.String("outfile", "blue_hat.png", "Output file name")

  flag.Parse()

  fmt.Println("input file name:  ", *infilePtr)
  fmt.Println("output file name: ", *outfilePtr)

  infile, err := os.Open(*infilePtr)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to open input file: %v\n", err)
    os.Exit(1)
  }
  defer infile.Close()

  redHatImage, _, err := image.Decode(infile)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to decode input file: %v\n", err)
    os.Exit(1)
  }

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

  outfile, err := os.Create(*outfilePtr)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to create output file: %v\n", err)
    os.Exit(1)
  }
  defer outfile.Close()

  err = png.Encode(outfile, blueHatImage)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to encode output file: %v\n", err)
    os.Exit(1)
  }
}
