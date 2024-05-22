package model

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/auth/database"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/message"
)

type User struct {
	Id       uint
	Email    string
	Password string
	Role     string
	Name     string
	Surname  string
}

func (u User) FindByEmail(email string, user *User) error {
	if err := database.Db.Where("email = ?", email).First(&user).Error; err != nil {
		fmt.Println("User not found:", err.Error())
		return message.ErrInvalidLoginOrPassword
	}

	return nil
}

func (u User) FindById(id string, user *User) error {
	if err := database.Db.First(&user, id).Error; err != nil {
		fmt.Println("User not found:", err.Error())
		return message.ErrUserNotFound
	}

	return nil
}

func (u User) Update(user *User) error {
	if err := database.Db.Save(&user).Error; err != nil {
		fmt.Println("Problem saving user", err.Error())
		return message.ErrWhileUpdatingUser
	}

	return nil
}

func (u User) GetWithNoPassword() NoPasswordUser {
	return NoPasswordUser{
		ID:      u.Id,
		Email:   u.Email,
		Name:    u.Name,
		Surname: u.Surname,
		Role:    u.Role,
	}
}

type NoPasswordUser struct {
	ID      uint   `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Role    string `json:"role"`
}

type UserResponse struct {
	Message string         `json:"response"`
	User    NoPasswordUser `json:"user"`
}

type LoggedUserResponse struct {
	Token string         `json:"token"`
	User  NoPasswordUser `json:"user"`
}
