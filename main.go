package main

import (
	"encoding/base64"
	"gopkg.in/gographics/imagick.v2/imagick"
	"math"
	"strings"
)

const (
	ImageFitCover   = "cover"
	ImageFitContain = "contain"
)

var ProfileACESLinear, _ = base64.StdEncoding.DecodeString("AAAEMGxjbXMEMAAAbW50clJHQiBYWVogB+AABQABAA0AGQABYWNzcCpuaXgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPbWAAEAAAAA0y1sY21zAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMZGVzYwAAARQAAABGY3BydAAAAVwAAAE2d3RwdAAAApQAAAAUY2hhZAAAAqgAAAAsclhZWgAAAtQAAAAUYlhZWgAAAugAAAAUZ1hZWgAAAvwAAAAUclRSQwAAAxAAAAAQZ1RSQwAAAxAAAAAQYlRSQwAAAxAAAAAQY2hybQAAAyAAAAAkZG1uZAAAA0QAAADqbWx1YwAAAAAAAAABAAAADGVuVVMAAAAqAAAAHABBAEMARQBTAC0AZQBsAGwAZQAtAFYANAAtAGcAMQAwAC4AaQBjAGMAAAAAbWx1YwAAAAAAAAABAAAADGVuVVMAAAEaAAAAHABDAG8AcAB5AHIAaQBnAGgAdAAgADIAMAAxADYALAAgAEUAbABsAGUAIABTAHQAbwBuAGUAIAAoAGgAdAB0AHAAOgAvAC8AbgBpAG4AZQBkAGUAZwByAGUAZQBzAGIAZQBsAG8AdwAuAGMAbwBtAC8AKQAsACAAQwBDAC0AQgBZAC0AUwBBACAAMwAuADAAIABVAG4AcABvAHIAdABlAGQAIAAoAGgAdAB0AHAAcwA6AC8ALwBjAHIAZQBhAHQAaQB2AGUAYwBvAG0AbQBvAG4AcwAuAG8AcgBnAC8AbABpAGMAZQBuAHMAZQBzAC8AYgB5AC0AcwBhAC8AMwAuADAALwBsAGUAZwBhAGwAYwBvAGQAZQApAC4AAAAAWFlaIAAAAAAAAPbWAAEAAAAA0y1zZjMyAAAAAAABCL8AAARO///2aAAABYkAAP4D///8vv///jkAAALmAADQIlhZWiAAAAAAAAD9qwAAXKX///9OWFlaIAAAAAD///YJ///qZAAA0cJYWVogAAAAAAAAAyIAALj3AAACHXBhcmEAAAAAAAAAAAABAABjaHJtAAAAAAADAAAAALwWAABD6wAAAAAAAQAAAAAAB///7EltbHVjAAAAAAAAAAEAAAAMZW5VUwAAAM4AAAAcAEEAQwBFAFMAIABjAGgAcgBvAG0AYQB0AGkAYwBpAHQAaQBlAHMAIABmAHIAbwBtACAAVABCAC0AMgAwADEANAAtADAAMAA0ACwAIABoAHQAdABwADoALwAvAHcAdwB3AC4AbwBzAGMAYQByAHMALgBvAHIAZwAvAHMAYwBpAGUAbgBjAGUALQB0AGUAYwBoAG4AbwBsAG8AZwB5AC8AYQBjAGUAcwAvAGEAYwBlAHMALQBkAG8AYwB1AG0AZQBuAHQAYQB0AGkAbwBuAAAAAA==")
var ProfileSRGB, _ = base64.StdEncoding.DecodeString("AAAE6GxjbXMEMAAAbW50clJHQiBYWVogB+AABQABAA0AGQABYWNzcCpuaXgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPbWAAEAAAAA0y1sY21zAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMZGVzYwAAARQAAABOY3BydAAAAWQAAAE2d3RwdAAAApwAAAAUY2hhZAAAArAAAAAsclhZWgAAAtwAAAAUYlhZWgAAAvAAAAAUZ1hZWgAAAwQAAAAUclRSQwAAAxgAAAAgZ1RSQwAAAxgAAAAgYlRSQwAAAxgAAAAgY2hybQAAAzgAAAAkZG1uZAAAA1wAAAGMbWx1YwAAAAAAAAABAAAADGVuVVMAAAAyAAAAHABzAFIARwBCAC0AZQBsAGwAZQAtAFYANAAtAHMAcgBnAGIAdAByAGMALgBpAGMAYwAAAABtbHVjAAAAAAAAAAEAAAAMZW5VUwAAARoAAAAcAEMAbwBwAHkAcgBpAGcAaAB0ACAAMgAwADEANgAsACAARQBsAGwAZQAgAFMAdABvAG4AZQAgACgAaAB0AHQAcAA6AC8ALwBuAGkAbgBlAGQAZQBnAHIAZQBlAHMAYgBlAGwAbwB3AC4AYwBvAG0ALwApACwAIABDAEMALQBCAFkALQBTAEEAIAAzAC4AMAAgAFUAbgBwAG8AcgB0AGUAZAAgACgAaAB0AHQAcABzADoALwAvAGMAcgBlAGEAdABpAHYAZQBjAG8AbQBtAG8AbgBzAC4AbwByAGcALwBsAGkAYwBlAG4AcwBlAHMALwBiAHkALQBzAGEALwAzAC4AMAAvAGwAZQBnAGEAbABjAG8AZABlACkALgAAAABYWVogAAAAAAAA9tYAAQAAAADTLXNmMzIAAAAAAAEMQgAABd7///MlAAAHkwAA/ZD///uh///9ogAAA9wAAMBuWFlaIAAAAAAAAG+gAAA49QAAA5BYWVogAAAAAAAAJJ8AAA+EAAC2xFhZWiAAAAAAAABilwAAt4cAABjZcGFyYQAAAAAAAwAAAAJmZgAA8qcAAA1ZAAAT0AAACltjaHJtAAAAAAADAAAAAKPXAABUfAAATM0AAJmaAAAmZwAAD1xtbHVjAAAAAAAAAAEAAAAMZW5VUwAAAXAAAAAcAHMAUgBHAEIAIABjAGgAcgBvAG0AYQB0AGkAYwBpAHQAaQBlAHMAIABmAHIAbwBtACAAQQAgAFMAdABhAG4AZABhAHIAZAAgAEQAZQBmAGEAdQBsAHQAIABDAG8AbABvAHIAIABTAHAAYQBjAGUAIABmAG8AcgAgAHQAaABlACAASQBuAHQAZQByAG4AZQB0ACAALQAgAHMAUgBHAEIALAAgAGgAdAB0AHAAOgAvAC8AdwB3AHcALgB3ADMALgBvAHIAZwAvAEcAcgBhAHAAaABpAGMAcwAvAEMAbwBsAG8AcgAvAHMAUgBHAEIAOwAgAGEAbABzAG8AIABzAGUAZQAgAGgAdAB0AHAAOgAvAC8AdwB3AHcALgBjAG8AbABvAHIALgBvAHIAZwAvAHMAcABlAGMAaQBmAGkAYwBhAHQAaQBvAG4ALwBJAEMAQwAxAHYANAAzAF8AMgAwADEAMAAtADEAMgAuAHAAZABmAAA=")

type Size struct {
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
	Format string `json:"format"`
}

type Quality struct {
	CompressionQuality uint      `json:"compression_quality"`
	SamplingFactors    []float64 `json:"sampling_factors"`
}

type ColorProfile struct {
	data   []byte
	format string
}

type ImageMeta struct {
	wand     *imagick.MagickWand
	depth    uint
	profiles []ColorProfile
}

func NewImage(wand *imagick.MagickWand) (ImageMeta, error) {
	meta := ImageMeta{
		wand: wand,
	}

	if err := wand.AutoOrientImage(); err != nil {
		return meta, err
	}

	if len(wand.GetImageProfiles("i*")) == 0 {
		if err := wand.ProfileImage("icc", ProfileSRGB); err != nil {
			return meta, err
		}
	}

	for _, name := range wand.GetImageProfiles("*") {
		meta.profiles = append(meta.profiles, ColorProfile{
			data:   []byte(wand.GetImageProfile(name)),
			format: name,
		})
	}
	if err := wand.SetImageDepth(16); err != nil {
		return meta, err
	}
	if err := wand.ProfileImage("icc", ProfileACESLinear); err != nil {
		return meta, err
	}
	return meta, nil
}

func (image *ImageMeta) CloneImage() ImageMeta {
	return ImageMeta{
		image.wand.Clone(),
		image.depth,
		image.profiles,
	}
}

func (image *ImageMeta) SanitizeMetadata() error {
	var profiles []ColorProfile
	for _, profile := range image.profiles {
		if !strings.EqualFold("exif", profile.format) {
			profiles = append(profiles, profile)
		}
	}
	image.profiles = profiles
	image.wand.RemoveImageProfile("exif")

	if err := image.wand.SetOption("png:include-chunk", "bKGD,cHRM,iCCP"); err != nil {
		return err
	}
	if err := image.wand.SetOption("png:exclude-chunk", "EXIF,iTXt,tEXt,zTXt,date"); err != nil {
		return err
	}
	for _, key := range image.wand.GetImageProperties("png:*") {
		if err := image.wand.DeleteImageProperty(key); err != nil {
			return err
		}
	}

	return nil
}

func (image *ImageMeta) Crop(size Size) error {
	if size.Width == 0 || size.Height == 0 || size.Format != ImageFitCover {
		return nil
	}

	currentWidth := image.wand.GetImageWidth()
	currentHeight := image.wand.GetImageHeight()

	currentAspectRatio := float64(currentWidth) / float64(currentHeight)
	desiredAspectRatio := float64(size.Width) / float64(size.Height)

	if currentAspectRatio == desiredAspectRatio {
		return nil
	}

	var desiredWidth, desiredHeight uint
	if desiredAspectRatio > currentAspectRatio {
		desiredWidth = currentWidth
		desiredHeight = uint(math.Round(float64(currentWidth) / desiredAspectRatio))
	} else {
		desiredHeight = currentHeight
		desiredWidth = uint(math.Round(desiredAspectRatio * float64(currentHeight)))
	}

	offsetLeft := int((currentWidth - desiredWidth) / 2.0)
	offsetTop := int((currentHeight - desiredHeight) / 2.0)

	if err := image.wand.CropImage(desiredWidth, desiredHeight, offsetLeft, offsetTop); err != nil {
		return err
	}

	return nil
}

func determineDesiredSize(width uint, height uint, size Size) (uint, uint) {
	currentAspectRatio := float64(width) / float64(height)

	var desiredWidth, desiredHeight uint
	if size.Height != 0 && size.Width != 0 {
		if size.Format == ImageFitCover {
			var desiredAspectRatio = float64(size.Width) / float64(size.Height)
			var croppedWidth, croppedHeight uint
			if desiredAspectRatio > currentAspectRatio {
				croppedWidth = width
				croppedHeight = uint(math.Round(float64(width) / desiredAspectRatio))
			} else {
				croppedHeight = height
				croppedWidth = uint(math.Round(desiredAspectRatio * float64(height)))
			}

			desiredHeight = uint(math.Min(float64(size.Height), float64(croppedHeight)))
			desiredWidth = uint(math.Min(float64(size.Width), float64(croppedWidth)))
		} else if currentAspectRatio > 1 {
			desiredWidth = uint(math.Min(float64(size.Width), float64(width)))
			desiredHeight = uint(math.Round(float64(desiredWidth) / currentAspectRatio))
		} else {
			desiredHeight = uint(math.Min(float64(size.Height), float64(height)))
			desiredWidth = uint(math.Round(currentAspectRatio * float64(desiredHeight)))
		}
	} else if size.Height != 0 {
		desiredHeight = uint(math.Min(float64(size.Height), float64(height)))
		desiredWidth = uint(math.Round(currentAspectRatio * float64(desiredHeight)))
	} else if size.Width != 0 {
		desiredWidth = uint(math.Min(float64(size.Width), float64(width)))
		desiredHeight = uint(math.Round(float64(desiredWidth) / currentAspectRatio))
	} else {
		desiredWidth = width
		desiredHeight = height
	}

	return desiredWidth, desiredHeight
}

func (image *ImageMeta) Resize(size Size) error {
	if size.Width == 0 && size.Height == 0 {
		return nil
	}

	currentWidth := image.wand.GetImageWidth()
	currentHeight := image.wand.GetImageHeight()

	desiredWidth, desiredHeight := determineDesiredSize(currentWidth, currentHeight, size)

	if desiredWidth != currentWidth || desiredHeight != currentHeight {
		if err := image.wand.ResizeImage(desiredWidth, desiredHeight, imagick.FILTER_LANCZOS, 1); err != nil {
			return err
		}
	}

	return nil
}

func (image *ImageMeta) Write(quality Quality, target string) error {
	for _, profile := range image.profiles {
		if err := image.wand.ProfileImage(profile.format, profile.data); err != nil {
			return err
		}
	}
	if err := image.wand.SetImageDepth(image.depth); err != nil {
		return err
	}

	if quality.CompressionQuality != 0 {
		if err := image.wand.SetImageCompressionQuality(quality.CompressionQuality); err != nil {
			return err
		}
	}

	if len(quality.SamplingFactors) != 0 {
		if err := image.wand.SetSamplingFactors(quality.SamplingFactors); err != nil {
			return err
		}
	}

	if err := image.wand.WriteImage(target); err != nil {
		return err
	}

	return nil
}
