package models

import (
	"time"
)

type Extension struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:255;not null;" json:"title"`
	Description string    `gorm:"type:TEXT;" json:"description"`
	Content     string    `gorm:"type:TEXT;not null;" json:"content"`
	Tag         string    `gorm:"size:255;not null;" json:"tag"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	OwnerID     uint32    `gorm:"not null" json:"owner_id"`                          // Foreign key field
	Owner       User      `gorm:"foreignKey:OwnerID;references:ID;onDelete:CASCADE"` // Relationship
}
