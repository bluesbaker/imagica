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

var (
	source string
	cols   int
	rows   int
)

// Initialize flags
func init() {
	flag.StringVar(&source, "source", "", "image source")
	flag.IntVar(&cols, "cols", 80, "columns")
	flag.IntVar(&rows, "rows", 40, "rows")

	flag.Usage = func() {
		fmt.Println("Example: imagica -source image.jpg -cols 200 -rows 100")
		flag.PrintDefaults()
	}
}

// Grayscale from RGB
func grayscale(colour color.Color) int {
	r, g, b, _ := colour.RGBA()
	grs := int(r+g+b) / 3
	return int(grs)
}

// Check errors
func check(heading string, err error) {
	if err != nil {
		fmt.Println(heading, err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	// checking arguments
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	} else if len(os.Args) == 2 {
		source = os.Args[1]
	}

	// open image
	imgFile, err := os.Open(source)
	check("Image error -", err)
	defer imgFile.Close()

	// decode image
	img, _, err := image.Decode(imgFile)
	check("Image error -", err)

	// reset EOF
	imgFile.Seek(0, 0)

	// decode config
	imgConfig, _, err := image.DecodeConfig(imgFile)
	check("Image error -", err)

	// options
	width := imgConfig.Width
	height := imgConfig.Height
	col := width / cols
	row := (height / rows) * 2
	ramp := "@%#+=-. "

	// print using ASCII
	for y := 0; y < height; y += row {
		for x := 0; x < width; x += col {
			grs := grayscale(img.At(x, y))
			fmt.Printf("%s", string(ramp[len(ramp)*grs/65536]))
		}
		fmt.Println()
	}
}
