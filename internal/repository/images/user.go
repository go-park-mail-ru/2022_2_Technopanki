package images

import (
	"HeadHunter/configs"
	"image"
	"strings"
)

func UploadUserAvatar(name string, image *image.Image, cfg *configs.ImageConfig) error {
	fullPath := strings.Join([]string{cfg.Path, "avatar/"}, "")
	return UploadWebpImage(fullPath, name, image)
}

func DeleteUserAvatar(name string, cfg *configs.ImageConfig) error {
	fullPath := strings.Join([]string{cfg.Path, "avatar/"}, "")
	return DeleteWebpImage(fullPath, name)
}
