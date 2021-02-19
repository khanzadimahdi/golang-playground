package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // register PNG decoder
	"io"
	"log"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		log.Fatalf("jpeg: %v\n", err)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format", kind)

	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

// how to run
// ./mandelbrot.png | ./jpeg >mandelbrot.jpg
