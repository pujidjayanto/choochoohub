package model

type UserOtpChannel string

const (
	UserOtpChannelEmail UserOtpChannel = "email"
	UserOtpChannelSms   UserOtpChannel = "sms"
)
