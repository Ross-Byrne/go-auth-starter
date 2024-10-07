package models

import (
	"time"
)

type Workspace struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;default:null"`
	Users     []User
	CreatedAt time.Time
	UpdatedAt time.Time
}
