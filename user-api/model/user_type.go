package model

type UserType string

const (
	UserTypeUnverified UserType = "unverified"
	UserTypeBasic      UserType = "basic"
	UserTypePremium    UserType = "premium"
)
