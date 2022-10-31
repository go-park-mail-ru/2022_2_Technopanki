package images

import (
	"HeadHunter/configs"
	"HeadHunter/pkg/repository"
	"image"
	"strings"
)

func UploadUserAvatar(name string, image *image.Image, cfg *configs.ImageConfig) error {
	fullPath := strings.Join([]string{cfg.Path, "avatar/"}, "")
	return repository.UploadWebpImage(fullPath, name, image)
}
