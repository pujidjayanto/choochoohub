package model

type UserOTPChannel string

const (
	UserOTPChannelEmail UserOTPChannel = "email"
	UserOTPChannelSMS   UserOTPChannel = "sms"
)
