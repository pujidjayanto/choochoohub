package model

type UserType string

const (
	UserTypeStarter UserType = "starter" // signed up & OTP verified, minimal profile
	UserTypeBasic   UserType = "basic"   // profile completed
	UserTypePremium UserType = "premium" // high transaction user
)
