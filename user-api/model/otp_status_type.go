package model

type UserOTPStatus string

const (
	UserOTPStatusPending         UserOTPStatus = "pending"
	UserOTPStatusVerified        UserOTPStatus = "verified"
	UserOTPStatusMaxAttempted    UserOTPStatus = "max_attempted"
	UserOTPStatusExpired         UserOTPStatus = "expired"
	UserOTPStatusUserInvalidated UserOTPStatus = "user_invalidated"
)
