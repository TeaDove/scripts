package main

import (
	"github.com/pkg/errors"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func draw() error {
	width := 64
	height := 64

	// Create a new RGBA image
	// image.Rect(x0, y0, x1, y1) defines the image bounds
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill the image with a specific color (e.g., blue)
	blue := color.RGBA{R: 0, G: 0, B: 255, A: 255} // Red, Green, Blue, Alpha
	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}  // Red, Green, Blue, Alpha

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if math.Abs(math.Sin(float64(x))*32-float64(y)) < 10.0 {
				img.Set(x, y, red)
				continue
			}

			img.Set(x, y, blue)
		}
	}

	file, err := os.Create("output.png")
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func main() {
	err := draw()
	if err != nil {
		panic(err)
	}
}
