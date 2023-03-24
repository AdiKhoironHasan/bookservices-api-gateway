package entity

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          int64          `gorm:"not null;uniqueIndex;primaryKe" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `gorm:"not null" json:"description"`
	CreatedAt   time.Time      `gorm:"not null;default:now()" json:"created_at,omitempty"`
	UpdatedAt   time.Time      `gorm:"not null;default:now()" json:"updated_at,omitempty"`
	DeleteAt    gorm.DeletedAt `gorm:"null" json:"delete_at,omitempty"`
}
