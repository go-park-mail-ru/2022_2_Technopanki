package images

import (
	"HeadHunter/configs"
	"HeadHunter/internal/errorHandler"
	"fmt"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
)

func UploadUserAvatar(name string, image *image.Image, cfg *configs.ImageConfig) (err error) {
	fullPath := strings.Join([]string{cfg.Path, "avatar/"}, "")
	if name == "" || fullPath == "" {
		fmt.Println("Error in file name/path (UploadUserAvatar)")
		return errorHandler.ErrBadRequest
	}

	resultImage, createErr := os.Create(strings.Join([]string{fullPath, name}, ""))
	if createErr != nil {
		return createErr
	}
	defer func(resultImage *os.File) {
		closeErr := resultImage.Close()
		if closeErr != nil {
			err = closeErr
		}
	}(resultImage)

	options, optionErr := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 15)
	if optionErr != nil {
		return optionErr
	}

	if encodingErr := webp.Encode(resultImage, *image, options); err != nil {
		return encodingErr
	}
	return nil
}

func DeleteUserAvatar(name string, cfg *configs.ImageConfig) error {
	fullPath := strings.Join([]string{cfg.Path, "avatar/"}, "")
	if fullPath == "" || name == "" {
		return errorHandler.ErrBadRequest
	}

	removeErr := os.Remove(strings.Join([]string{fullPath, name}, ""))
	if removeErr != nil {
		return removeErr
	}
	return nil
}
