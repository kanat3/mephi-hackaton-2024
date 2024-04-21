package models

type Emotion struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:20"`
}
