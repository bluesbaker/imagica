package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

var imagePath string

func init() {
	flag.StringVar(&imagePath, "source", "", "image source")
}

func grayscale(colour color.Color) int {
	r, g, b, _ := colour.RGBA()
	grs := (0.299 * float32(r)) + (0.587 * float32(g)) + (0.114 * float32(b))
	return int(grs)
}

func main() {
	flag.Parse()

	imgFile, err := os.Open(imagePath)
	if err != nil {
		fmt.Println(err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println(err)
	}

	imgFile.Seek(0, 0)

	imgConfig, _, err := image.DecodeConfig(imgFile)
	if err != nil {
		fmt.Println(err)
	}

	width := imgConfig.Width
	height := imgConfig.Height

	ramp := "@%#+=-. "
	for y := 0; y < height; y += 20 {
		for x := 0; x < width; x += 10 {
			grs := grayscale(img.At(x, y))
			fmt.Printf("%s", string(ramp[len(ramp)*grs/65536]))
		}
		fmt.Println()
	}
}
