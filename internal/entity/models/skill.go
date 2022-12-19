//go:generate easyjson -all skill.go
package models

//easyjson:json
type Skill struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	SkillName string `json:"skillSetName" gorm:"not null;"`
}
