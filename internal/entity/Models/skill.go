package Models

type Skill struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	SkillName string `json:"skillSetName" gorm:"not null;"`
}
