package model

import (
	"fmt"
	"github.com/SebastianOraczek/auth/database"
	"github.com/SebastianOraczek/auth/internal/response"
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
		return response.ErrInvalidLoginOrPassword
	}

	return nil
}

func (u User) FindById(id string, user *User) error {
	if err := database.Db.First(&user, id).Error; err != nil {
		fmt.Println("User not found:", err.Error())
		return response.ErrUserNotFound
	}

	return nil
}

func (u User) Update(user *User) error {
	if err := database.Db.Save(&user).Error; err != nil {
		fmt.Println("Problem saving user", err.Error())
		return response.ErrWhileUpdatingUser
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
	Message string         `json:"message"`
	User    NoPasswordUser `json:"user"`
}

type LoggedUserResponse struct {
	Token string         `json:"token"`
	User  NoPasswordUser `json:"user"`
}
