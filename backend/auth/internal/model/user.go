package model

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type ResponseUser struct {
	ID      uint   `json:"id"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
