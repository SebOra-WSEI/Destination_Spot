package model

type AuthBody struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
