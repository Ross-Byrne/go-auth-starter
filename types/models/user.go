package models

import "time"

type User struct {
	ID                uint   `gorm:"primaryKey"`
	FirstName         string `gorm:"not null;default:null"`
	LastName          string `gorm:"not null;default:null"`
	Email             string `gorm:"uniqueIndex;not null;default:null"`
	EncryptedPassword string `gorm:"not null;default:null"`
	WorkspaceID       uint   `gorm:"not null;default:null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
