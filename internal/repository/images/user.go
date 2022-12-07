package images

import (
	"HeadHunter/configs"
	"HeadHunter/pkg/errorHandler"
	"fmt"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image"
	"image/color"
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

func Average(img image.Image) string {
	max := img.Bounds().Max
	min := img.Bounds().Min
	var sumR, sumG, sumB uint64
	count := uint64(max.Y-min.Y) * uint64(max.X-min.X)
	for y := min.Y; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			SR, SG, SB, _ := img.At(x, y).RGBA()
			r, g, b := uint8(SR), uint8(SG), uint8(SB)
			sumR += uint64(r)
			sumG += uint64(g)
			sumB += uint64(b)
		}
	}
	result := color.RGBA{
		R: uint8(sumR / count),
		G: uint8(sumG / count),
		B: uint8(sumB / count),
		A: 255,
	}
	resultStr := fmt.Sprintf("%d %d %d", result.R, result.G, result.B)
	return resultStr
}
