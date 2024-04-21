package models

import (
	"time"

	"github.com/google/uuid"
)

type Meeting struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Date     time.Time
	FilePath string `gorm:"size:225"`
}
