package model

import "time"

type Notes struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:255;not null"`
	SubTitle    string    `gorm:"size:255;uniqueIndex;not null"`
	Description string    `gorm:"text;uniqueIndex;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
