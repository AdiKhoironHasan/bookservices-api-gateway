package service

import (
	"gorm.io/gorm"
)

// Repositories is a struct
type Repositories struct {
	DB *gorm.DB
}

// NewDBService is constructor
func NewDBService(db *gorm.DB) *Repositories {
	return &Repositories{
		DB: db,
	}
}
