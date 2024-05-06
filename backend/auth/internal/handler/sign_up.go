package handler

import (
	"fmt"
	"github.com/SebastianOraczek/auth/internal/email"
	"github.com/SebastianOraczek/auth/internal/model"
	"github.com/SebastianOraczek/auth/internal/password"
	"github.com/SebastianOraczek/auth/internal/request"
	"github.com/SebastianOraczek/auth/internal/response"
	"github.com/SebastianOraczek/auth/startup"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const UserRole string = "user"

func SignUp(c *gin.Context) {
	var body model.AuthBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBody(
		body.Email, body.Password, body.ConfirmPassword,
	) {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.EmptyFieldsErrMsg))
		return
	}

	if err := email.Verify(body.Email); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
		return
	}

	if err := password.Validate(body.Password); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
		return
	}

	if body.Password != body.ConfirmPassword {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.PasswordNotTheSameErrMsg))
		return
	}

	var user model.User
	if err := startup.Db.Where("email = ?", body.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.UserAlreadyExistsErrMsg))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		fmt.Println("Problem with hashing password", err.Error())
		c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(response.InternalServerErrMsg))
		return
	}

	name, surname := email.CreateNameAndSurnameFromEmail(body.Email)

	newUser := model.User{
		Email:    body.Email,
		Password: string(hashedPassword),
		Name:     name,
		Surname:  surname,
		Role:     UserRole,
	}

	if err := startup.Db.Create(&newUser).Error; err != nil {
		fmt.Println("Problem creating new user", err.Error())
		c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(response.ProblemWhileRegistrationErrMsg))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": response.UserCreatedMsg})
}
