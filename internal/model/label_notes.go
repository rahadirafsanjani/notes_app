package model

import "time"

type LabelNotes struct {
	ID        uint      `gorm:"primaryKey"`
	LabelID   uint      `gorm:"not null"`
	NoteID    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
