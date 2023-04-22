package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func main() {

	var format = flag.String("format", "jpeg", "the format of the output image")
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Println("Usage: ex10.1 [-format=jpeg|png] < in  > out")
		os.Exit(1)
	}

	img, kind, err := image.Decode(os.Stdin)

	if err != nil {
		fmt.Printf("error decoding image: %v\n", err)
		os.Exit(1)
	}

	if kind == *format {
		fmt.Println("Input and output formats are the same")
		os.Exit(1)
	}

	switch strings.ToLower(*format) {
	case "jpg", "jpeg":
		err = jpeg.Encode(os.Stdout, img, nil)
	case "png":
		err = png.Encode(os.Stdout, img)

	}

	if err != nil {
		fmt.Printf("error encoding image: %v\n", err)
		os.Exit(1)
	}
}
