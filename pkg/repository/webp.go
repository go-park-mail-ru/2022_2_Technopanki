package repository

import (
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image"
	"os"
	"strings"
)

func UploadWebpImage(path, name string, image *image.Image) (err error) {
	resultImage, createErr := os.Create(strings.Join([]string{path, name}, ""))
	if createErr != nil {
		return createErr
	}
	defer func(resultImage *os.File) {
		closeErr := resultImage.Close()
		if closeErr != nil {
			err = closeErr
		}
	}(resultImage)

	options, optionErr := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if optionErr != nil {
		return optionErr
	}

	if encodingErr := webp.Encode(resultImage, *image, options); err != nil {
		return encodingErr
	}
	return nil
}

func DeleteImage(path, name string) (err error) {
	removeErr := os.Remove(strings.Join([]string{path, name}, ""))
	if removeErr != nil {
		return removeErr
	}
	return nil
}
