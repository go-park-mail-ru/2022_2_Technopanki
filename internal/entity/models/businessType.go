//go:generate easyjson -all businessType.go
package models

//easyjson:json
type BusinessType struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null;"`
}
