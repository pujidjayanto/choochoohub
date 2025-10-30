package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserOtp struct {
	ID           uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null;index:user_purpose_idx"`
	Channel      UserOtpChannel `gorm:"type:varchar(20);not null"`  // email | sms
	Destination  string         `gorm:"type:varchar(255);not null"` // email or phone
	OTPHash      string         `gorm:"type:text;not null"`         // hashed otp
	Purpose      string         `gorm:"type:varchar(50);not null"`  // signup | login | password_reset
	Status       UserOtpStatus  `gorm:"type:varchar(20);not null;default:'pending'"`
	SendAttempts int            `gorm:"not null;default:1"`
	ExpiresAt    time.Time      `gorm:"type:timestamptz;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

// BeforeCreate generates UUID if not set
func (o *UserOtp) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}
	return
}
