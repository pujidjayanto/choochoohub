package dto

type SignupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type (
	SigninRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	SigninResponse struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}
)
