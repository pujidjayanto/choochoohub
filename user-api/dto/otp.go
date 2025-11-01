package dto

import (
	"time"

	"github.com/google/uuid"
)

type OtpRequest struct {
	UserId      uuid.UUID
	Channel     string
	Destination string
	Purpose     string
	ExpiredAt   time.Time
}

type VerifyOtpRequest struct {
	Destination string `json:"destination" validate:"required"`
	Purpose     string `json:"purpose" validate:"required"`
	OtpCode     string `json:"otpCode" validate:"required"`
}
