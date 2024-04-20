package models

type Meeting struct {
    ID       uuid.UUID gorm:"type:uuid;primaryKey"
    Date     time.Time
    FilePath string    gorm:"size:225"
}