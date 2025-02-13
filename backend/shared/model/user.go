package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type NoPasswordUser struct {
	ID      uint   `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Role    string `json:"role"`
}

type UserResponseWithMessage struct {
	Message string         `json:"message"`
	User    NoPasswordUser `json:"user"`
}

type UserResponse struct {
	User NoPasswordUser `json:"user"`
}

type AllUsersResponse struct {
	Users []NoPasswordUser `json:"users"`
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

func (u User) FindByEmailSQL(db *sql.DB, c context.Context, email string, user *User) error {
	rows, err := db.QueryContext(c, "SELECT * FROM users WHERE email = ?", email)
	defer rows.Close()

	if err != nil {
		fmt.Println("User not found:", err.Error())
		return response.ErrUserNotFound
	}

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.Name, &user.Surname); err != nil {
			fmt.Println("Problem with scanning:", err.Error())
			return response.ErrInternalServer
		}
	}

	if user.ID == 0 {
		return response.ErrUserNotFound
	}

	return nil
}

func (u User) FindByIdSQL(db *sql.DB, c context.Context, id string, user *User) error {
	rows, err := db.QueryContext(c, "SELECT * FROM users WHERE id = ?", id)
	defer rows.Close()

	if err != nil {
		fmt.Println("User not found:", err.Error())
		return response.ErrUserNotFound
	}

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.Name, &user.Surname); err != nil {
			fmt.Println("Problem with scanning:", err.Error())
			return response.ErrInternalServer
		}
	}

	if user.ID == 0 {
		return response.ErrUserNotFound
	}

	return nil
}

func (u User) FindAllSQL(db *sql.DB, c context.Context, users *[]User) error {
	rows, err := db.QueryContext(c, "SELECT * FROM users")
	defer rows.Close()

	if err != nil {
		fmt.Println("User not found:", err.Error())
		return response.ErrUserNotFound
	}

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.Name, &user.Surname); err != nil {
			fmt.Println("Problem with scanning:", err.Error())
			return response.ErrInternalServer
		}

		*users = append(*users, user)
	}

	if len(*users) == 0 {
		return response.ErrUsersNotFound
	}

	return nil
}

func (u User) Update(db *gorm.DB, user *User) error {
	if err := db.Save(&user).Error; err != nil {
		fmt.Println("Problem while updating user", err.Error())
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
