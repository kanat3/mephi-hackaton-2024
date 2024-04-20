package models

type Timepoint struct {
    ID        uuid.UUID gorm:"type:uuid;primaryKey"
    Time      time.Time gorm:"not null"

    UserID    uuid.UUID gorm:"type:uuid;not null"
    User      User      gorm:"foreignKey:UserID"

    MeetingID uuid.UUID gorm:"type:uuid;not null"
    Meeting   Meeting   gorm:"foreignKey:MeetingID"
	
    EmotionID uint      gorm:"not null"
    Emotion   Emotion   gorm:"foreignKey:EmotionID"
}