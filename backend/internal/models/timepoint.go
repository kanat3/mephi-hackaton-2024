package models

import (
	"time"

	"github.com/google/uuid"
)

type Timepoint struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Start_Time time.Time
	End_Time   time.Time

	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID"`

	MeetingID uuid.UUID `gorm:"type:uuid;not null"`
	Meeting   Meeting   `gorm:"foreignKey:MeetingID"`

	EmotionID uint    `gorm:"not null"`
	Emotion   Emotion `gorm:"foreignKey:EmotionID"`
}
