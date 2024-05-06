package model

type User struct {
	Id       uint
	Email    string
	Password string
	Role     string
	Name     string
	Surname  string
}

type NoPasswordUser struct {
	ID      uint   `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Role    string `json:"role"`
}

type UserResponse struct {
	Message string         `json:"message"`
	User    NoPasswordUser `json:"user"`
}

type LoggedUserResponse struct {
	Token string         `json:"token"`
	User  NoPasswordUser `json:"user"`
}
