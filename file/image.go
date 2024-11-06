package file

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	// We support gif for reading, but not writing
	_ "image/gif"

	"github.com/nfnt/resize"
)

// https://github.com/ajstarks/svgo

// Options represents image options
type Options struct {
	Path      string
	MaxHeight int64
	MaxWidth  int64
	Quality   int
	// Square bool
}

// SaveJpegRepresentations saves several image representation rescaled in proportion (always in proportion) using the specified max width, max height and quality
func SaveJpegRepresentations(r io.Reader, options []Options) error {

	if r == nil {
		return fmt.Errorf("assets: nil reader received in SaveJpegRepresentation")
	}

	// Read the image data, if we have a jpeg, convert to png?
	original, _, err := image.Decode(r)
	if err != nil {
		return fmt.Errorf("error decoding image:%s", err)
	}

	// For each option, save a file
	for _, o := range options {

		fmt.Printf("assets: saving image file - %v\n", o)

		// Resize this image given the params
		resized, err := resizeImage(original, o.MaxWidth, o.MaxHeight)
		if err != nil {
			return fmt.Errorf("error resizing file :%s", err)
		}

		// Write out to the desired file path
		w, err := os.Create(o.Path)
		if err != nil {
			return fmt.Errorf("error creating file :%s", err)
		}
		defer w.Close()
		err = jpeg.Encode(w, resized, &jpeg.Options{Quality: o.Quality})
		if err != nil {
			return fmt.Errorf("error encoding file :%s", err)
		}

	}

	return nil

}

// SavePNGRepresentations saves png representations according to options
func SavePNGRepresentations(r io.Reader, options []Options) error {

	if r == nil {
		return fmt.Errorf("assets: reader received in SaveJpegRepresentation")
	}

	// Read the image data, if we have a jpeg, convert to png?
	original, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	// For each option, save a file
	for _, o := range options {

		fmt.Printf("Saving image file - %v\n", o)

		// Resize this image given the params - this is always in proportion, NEVER stretched
		// If Square is true we crop to a square
		resized, err := resizeImage(original, o.MaxWidth, o.MaxHeight)
		if err != nil {
			return err
		}

		// Write out to the desired file path
		w, err := os.Create(o.Path)
		if err != nil {
			return err
		}
		defer w.Close()
		err = png.Encode(w, resized)
		if err != nil {
			return err
		}

	}

	return nil

}

// resizeImage resizes an image
func resizeImage(src image.Image, maxWidth int64, maxHeight int64) (image.Image, error) {

	srcSize := src.Bounds().Size()

	// Use the original image dimensions to keep it in pro
	ratio := float64(maxWidth) / float64(srcSize.X)
	yRatio := float64(maxHeight) / float64(srcSize.Y)
	if yRatio < ratio {
		ratio = yRatio
	}

	// Now adjust desired width and height according to ratio
	width := uint(float64(srcSize.X) * ratio)
	height := uint(float64(srcSize.Y) * ratio)

	return resize.Resize(width, height, src, resize.Bilinear), nil
}
