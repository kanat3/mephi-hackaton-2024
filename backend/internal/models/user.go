package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Login string    `gorm:"type:varchar(16);unique"`
}
