package model

import "time"

type Labels struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Types     string    `gorm:"size:255;uniqueIndex;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
