package models

type Skill struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	SkillName string `json:"skillSetName" gorm:"not null;"`
}
