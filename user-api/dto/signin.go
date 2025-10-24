package dto

type (
	SigninRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// token will be generated in api gateway
	SigninResponse struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}
)
