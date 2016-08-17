package main

import (
    "fmt"
    "image"
    "image/jpeg"
    "image/png"
    "image/gif"
    "io"
    "os"
)

const VALID_FORMAT [3]string = [3]string{"gif", "png", "jpeg"}

func toJPEG(in io.Reader, out io.Writer) error {
    img, kind, err := image.Decode(in)
    if err != nil {
        return err
    } else {
        fmt.Fprintln(os.Stderr, "input format =", kind)
        return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
    }
}

// 转换到 gif 格式
func toGIF(in io.Reader, out io.Writer) error {
    img, kind, err := image.Decode(in)
    if err != nil {
        return err
    } else {
        fmt.Fprintln(os.Stderr, "input format =", kind)
        return gif.Encode(out, img, &gif.Options{Quality: 95})
    }
}

// 转换到 png 格式
func toPNG(in io.Reader, out io.Writer) error {
    img, kind, err := image.Decode(in)
    if err != nil {
        return err
    } else {
        fmt.Fprintln(os.Stderr, "input format =", kind)
        return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
    }
}

func main() {
    targetFormat := os.Args[1]
    fmt.Println("target format:", targetFormat)

    switch targetFormat {
    case "gif":
        err := toGIF(os.Stdin, os.Stdout)
        if err != nil {
            fmt.Fprintf(os.Stderr, "gif: %v\n", err)
            os.Exit(1)
        }
    case "jpeg":
        err := toJPEG(os.Stdin, os.Stdout)
        if err != nil {
            fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
        }
    case "png":
        err := toPNG(os.Stderr, os.Stdout)
        if err != nil {
            fmt.Fprintf(os.Stderr, "png: %v\n", err)
        }
    }
}
