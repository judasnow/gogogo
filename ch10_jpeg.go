package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func toJPEG(in io.Reader, out io.Writer) error {
    img, kind, err := image.Decode(in)
	if err != nil {
		return err
	} else {
		fmt.Fprintln(os.Stderr, "input format =", kind)
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	}
}

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}
