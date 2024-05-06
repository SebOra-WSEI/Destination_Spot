package handler

import (
	"fmt"
	"github.com/SebastianOraczek/auth/database"
	"github.com/SebastianOraczek/auth/internal/email"
	"github.com/SebastianOraczek/auth/internal/model"
	"github.com/SebastianOraczek/auth/internal/password"
	"github.com/SebastianOraczek/auth/internal/request"
	"github.com/SebastianOraczek/auth/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	UserRole  string = "user"
	AdminRole        = "admin"
)

func SignUp(c *gin.Context) {
	var body model.AuthBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBodyFields(
		body.Email, body.Password, body.ConfirmPassword,
	) {
		c.JSON(http.StatusBadRequest, response.CreateError(response.EmptyFieldsErrMsg))
		return
	}

	if err := email.Validate(body.Email); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err.Error()))
		return
	}

	if err := password.Validate(body.Password, body.ConfirmPassword); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err.Error()))
		return
	}

	var user model.User
	if err := database.Db.Where("email = ?", body.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, response.CreateError(response.UserAlreadyExistsErrMsg))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		fmt.Println("Problem with hashing password", err.Error())
		c.JSON(http.StatusInternalServerError, response.CreateError(response.InternalServerErrMsg))
		return
	}

	name, surname := email.GetNameAndSurname(body.Email)

	newUser := model.User{
		Email:    body.Email,
		Password: string(hashedPassword),
		Name:     name,
		Surname:  surname,
		Role:     UserRole,
	}

	if err := database.Db.Create(&newUser).Error; err != nil {
		fmt.Println("Problem creating new user", err.Error())
		c.JSON(http.StatusInternalServerError, response.CreateError(response.ProblemWhileRegistrationErrMsg))
		return
	}

	res := model.UserResponse{
		Message: response.UserCreatedMsg,
		User: model.NoPasswordUser{
			ID:      newUser.Id,
			Email:   newUser.Email,
			Name:    newUser.Name,
			Surname: newUser.Surname,
			Role:    newUser.Role,
		},
	}

	c.JSON(http.StatusCreated, response.Create(res))
}
