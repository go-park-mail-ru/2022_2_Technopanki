package entity

import "gorm.io/gorm"

type Resume struct {
	gorm.Model
	Description string
}
