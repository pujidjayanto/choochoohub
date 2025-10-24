package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfile struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	UserID         uuid.UUID `gorm:"not null;uniqueIndex"`
	Phone          string    `gorm:"type:varchar(20);unique"`
	Name           string    `gorm:"not null"`
	DOB            time.Time `gorm:"type:date;not null"`
	Gender         string    `gorm:"type:char(1);not null"`
	IdentityNumber string    `gorm:"type:varchar(50);not null"`
	IdentityType   string    `gorm:"type:varchar(10);not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

func (p *UserProfile) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
