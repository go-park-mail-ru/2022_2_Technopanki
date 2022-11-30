package images

import (
	"HeadHunter/configs"
	"HeadHunter/pkg/errorHandler"
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
		return errorHandler.ErrBadRequest
	}

	resultImage, createErr := os.Create(strings.Join([]string{fullPath, name}, ""))
	if createErr != nil {
		return createErr
	}

	defer func(resultImage *os.File) {
		errSync := resultImage.Sync()
		if errSync != nil {
			err = errSync
		}
		errClose := resultImage.Close()
		if errClose != nil {
			err = errSync
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
	//return nil
}

func DeleteUserAvatar(name string, cfg *configs.ImageConfig) error {
	fullPath := strings.Join([]string{cfg.Path, "avatar/"}, "")
	if name == "" {
		return errorHandler.ErrBadRequest
	}

	removeErr := os.Remove(strings.Join([]string{fullPath, name}, ""))
	if removeErr != nil {
		return removeErr
	}
	return nil
}
