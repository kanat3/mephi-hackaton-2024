package models

type User struct {
	ID    uuid.UUID gorm:"type:uuid;primaryKey"
    Login string    gorm:"type:varchar(16);unique"
}