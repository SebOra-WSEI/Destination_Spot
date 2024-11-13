package model

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/jinzhu/gorm"
)

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

type UsersResponse struct {
	Users []NoPasswordUser `json:"users"`
}

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func (u User) FindByEmail(db *gorm.DB, email string, user *User) error {
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		fmt.Println("User not found:", err.Error())
		return response.ErrUserNotFound
	}

	return nil
}

func (u User) FindById(db *gorm.DB, id string, user *User) error {
	if err := db.First(&user, id).Error; err != nil {
		fmt.Println("User not found:", err.Error())
		return response.ErrUserNotFound
	}

	return nil
}

func (u User) Update(db *gorm.DB, user *User) error {
	if err := db.Save(&user).Error; err != nil {
		fmt.Println("Problem saving user", err.Error())
		return response.ErrWhileUpdatingUser
	}

	return nil
}

func (u User) GetWithNoPassword() NoPasswordUser {
	return NoPasswordUser{
		ID:      u.ID,
		Email:   u.Email,
		Name:    u.Name,
		Surname: u.Surname,
		Role:    u.Role,
	}
}
