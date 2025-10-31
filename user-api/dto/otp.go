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
