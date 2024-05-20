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
	"net/http"
)

const UserRole string = "user"

func SignUp(c *gin.Context) {
	var body model.AuthBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBodyFields(
		body.Email, body.Password, body.ConfirmPassword,
	) {
		c.JSON(http.StatusBadRequest, response.CreateError(response.ErrEmptyFields))
		return
	}

	if err := email.Validate(body.Email); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	if err := password.Validate(body.Password, body.ConfirmPassword); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var user model.User
	if err := database.Db.Where("email = ?", body.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, response.CreateError(response.ErrUserAlreadyExists))
		return
	}

	pass, err := password.Generate(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	name, surname := email.GetNameAndSurname(body.Email)

	newUser := model.User{
		Email:    body.Email,
		Password: pass,
		Name:     name,
		Surname:  surname,
		Role:     UserRole,
	}

	if err := database.Db.Create(&newUser).Error; err != nil {
		fmt.Println("Problem creating new user", err.Error())
		c.JSON(http.StatusInternalServerError, response.CreateError(response.ErrProblemWhileRegistration))
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
