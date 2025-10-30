package model

type UserOtpStatus string

const (
	UserOtpStatusPending         UserOtpStatus = "pending"
	UserOtpStatusVerified        UserOtpStatus = "verified"
	UserOtpStatusMaxAttempted    UserOtpStatus = "max_attempted"
	UserOtpStatusExpired         UserOtpStatus = "expired"
	UserOtpStatusUserInvalidated UserOtpStatus = "user_invalidated"
)
