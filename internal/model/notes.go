package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notes struct {
	gorm.Model
	id          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
}
