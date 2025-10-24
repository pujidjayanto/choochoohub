package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	Email        string    `gorm:"not null;unique"`
	PasswordHash string    `gorm:"not null"`
	UserType     UserType  `gorm:"type:varchar(20);not null;default:'unverified'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Profile      *UserProfile `gorm:"constraint:OnDelete:CASCADE;"`
}

// https://gorm.io/docs/hooks.html#Hooks
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	if u.UserType == "" {
		u.UserType = "unverified"
	}
	return
}
