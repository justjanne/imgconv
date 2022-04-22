package main

import (
	"flag"
	"github.com/justjanne/imgconv"
	"gopkg.in/gographics/imagick.v2/imagick"
	"os"
)

type arguments struct {
	Width   *uint
	Height  *uint
	Fit     *string
	Quality *uint
	Source  string
	Target  string
}

var args = arguments{
	Width: flag.Uint(
		"width",
		0,
		"Desired width of the image",
	),
	Height: flag.Uint(
		"height",
		0,
		"Desired height of the image",
	),
	Fit: flag.String(
		"fit",
		"contain",
		"Desired fit format for image. Allowed are cover and contain.",
	),
	Quality: flag.Uint(
		"quality",
		90,
		"Desired quality of output image",
	),
}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	imagick.Initialize()
	defer imagick.Terminate()

	source := flag.Arg(0)
	target := flag.Arg(1)

	if err := convert(source, target, imgconv.Quality{
		CompressionQuality: *args.Quality,
		SamplingFactors:    []float64{1.0, 1.0, 1.0, 1.0},
	}, imgconv.Size{
		Width:  *args.Width,
		Height: *args.Height,
		Format: *args.Fit,
	}); err != nil {
		panic(err)
	}
}

func convert(source string, target string, quality imgconv.Quality, size imgconv.Size) error {
	wand := imagick.NewMagickWand()
	defer wand.Destroy()

	var err error
	if err = wand.ReadImage(source); err != nil {
		return err
	}
	var image imgconv.ImageHandle
	if image, err = imgconv.NewImage(wand); err != nil {
		return err
	}
	if err := image.Crop(size); err != nil {
		return err
	}
	if err := image.Resize(size); err != nil {
		return err
	}
	if err := image.Write(quality, target); err != nil {
		return err
	}
	return nil
}
