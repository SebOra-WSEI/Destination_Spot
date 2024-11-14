package handler

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/auth/database"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/email"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/password"
	"github.com/SebOra-WSEI/Destination_spot/shared/model"
	"github.com/SebOra-WSEI/Destination_spot/shared/request"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const UserRole string = "user"

type AuthBody struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func SignUp(c *gin.Context) {
	var body AuthBody
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

	res := model.UserResponseWithMessage{
		Message: response.UserCreatedMsg,
		User:    newUser.GetWithNoPassword(),
	}

	c.JSON(http.StatusCreated, response.Create(res))
}
