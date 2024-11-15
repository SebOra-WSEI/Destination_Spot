package model

type LoggedUserResponse struct {
	Token string         `json:"token"`
	User  NoPasswordUser `json:"user"`
}
