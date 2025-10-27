package userapi

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninResponse struct {
	Token string `json:"token"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
