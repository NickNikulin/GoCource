package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var teal color.Color = color.RGBA{0, 255, 0, 255}
var red color.Color = color.RGBA{30, 30, 30, 30}
var reds color.Color = color.RGBA{200, 30, 30, 255}

func main() {
	file, err := os.Create("lines.png")

	if err != nil {
		fmt.Errorf("%s", err)
	}
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src)

	for x := 20; x < 380; x++ {
		y := x/3 + 15
		img.Set(x, y, red)
	}

	for x := 20; x < 380; x++ {
		y := x/50 + 15
		img.Set(x, y, red)
	}

	png.Encode(file, img)

}
