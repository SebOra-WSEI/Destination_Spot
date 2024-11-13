package model

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoggedUserResponse struct {
	Token string         `json:"token"`
	User  NoPasswordUser `json:"user"`
}
