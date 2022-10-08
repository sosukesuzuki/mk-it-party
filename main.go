package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"log"
	"math"
	"os"
)

func createPartyAlpha(size image.Point, c color.RGBA) *image.RGBA {
	rgba := image.NewRGBA(image.Rectangle{image.Point{0, 0}, size})
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			rgba.Set(x, y, c)
		}
	}
	return rgba
}

func changeInt(value int, rate float64) int {
	return int(math.Floor(float64(value) * rate))
}

func changeSize(size image.Point, rate float64) image.Point {
	return image.Point{
		changeInt(size.X, rate),
		changeInt(size.Y, rate),
	}
}

func main() {
	reader, err := os.Open("./samples/sosukesuzuki.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, err := jpeg.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	origSize := img.Bounds().Size()
	size := changeSize(origSize, 0.5)

	partyColors := []color.RGBA{
		{R: 255, G: 0, B: 100, A: 200},
		{R: 100, G: 255, B: 0, A: 200},
		{R: 0, G: 100, B: 255, A: 200},
		{R: 255, G: 0, B: 100, A: 200},
		{R: 100, G: 255, B: 0, A: 200},
		{R: 0, G: 100, B: 255, A: 200},
	}
	positions := []image.Point{
		{changeInt(origSize.X, 0.3), 0},
		{changeInt(origSize.X, 0.15), changeInt(origSize.Y, 0.15)},
		{changeInt(origSize.X, 0.1), changeInt(origSize.Y, 0.30)},
		{changeInt(origSize.X, 0.3), changeInt(origSize.Y, 0.15)},
		{changeInt(origSize.X, 0.3), changeInt(origSize.Y, 0.1)},
		{changeInt(origSize.X, 0.3), changeInt(origSize.Y, 0.05)},
	}

	opts := &gif.GIF{}

	for i, partyColor := range partyColors {
		position := positions[i]

		rgba := createPartyAlpha(origSize, partyColor)
		rectangle := image.Rectangle{image.Point{0, 0}, size}

		palettedImage := image.NewPaletted(rectangle, palette.WebSafe)
		draw.Draw(palettedImage, rectangle, img, position, draw.Src)
		draw.Draw(palettedImage, rectangle, rgba, position, draw.Over)

		opts.Image = append(opts.Image, palettedImage)
		opts.Delay = append(opts.Delay, 1)
	}

	f, _ := os.Create("./dist/sosukesuzuki.gif")
	defer f.Close()

	gif.EncodeAll(f, opts)
}
