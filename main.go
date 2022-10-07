package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"log"
	"os"
)

func createPartyAlpha(size image.Point, c color.RGBA) *image.RGBA {
	rgba := image.NewRGBA(image.Rectangle{ image.Point{ 0, 0 }, size })
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			rgba.Set(x, y, c)
		}
	}
	return rgba
}

func main() {
	reader, err := os.Open("./samples/sosukesuzuki.jpg")
	if err != nil {
		log.Fatal(err)	
	}
	defer reader.Close()

	img,  err := jpeg.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	size := img.Bounds().Size()

	partyColors := []color.RGBA{
		{ R: 255, G: 0, B: 100, A: 200 },
		{ R: 100, G: 255, B: 0, A: 200 },
		{ R: 0, G: 100, B: 255, A: 200 },
	}

	opts := &gif.GIF{}

	for _, partyColor := range partyColors {
		rgba := createPartyAlpha(size, partyColor)
		rectangle := image.Rectangle{ image.Point{ 0, 0 }, size }

		palettedImage := image.NewPaletted(rectangle, palette.WebSafe)
		draw.Draw(palettedImage, rectangle, img, image.Point{ 0, 0 }, draw.Src)
		draw.Draw(palettedImage, rectangle, rgba, image.Point{ 0, 0 }, draw.Over)

		opts.Image = append(opts.Image, palettedImage)
		opts.Delay = append(opts.Delay, 0)
	}


	f, _ := os.Create("./dist/sosukesuzuki.gif")
	defer f.Close()

	gif.EncodeAll(f, opts)
}
