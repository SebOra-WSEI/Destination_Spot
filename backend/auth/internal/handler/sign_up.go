package handler

import (
	"fmt"
	"github.com/SebOra-WSEI/auth/internal/email"
	"github.com/SebOra-WSEI/auth/internal/model"
	"github.com/SebOra-WSEI/auth/internal/password"
	"github.com/SebOra-WSEI/auth/internal/request"
	"github.com/SebOra-WSEI/auth/internal/response"
	"github.com/SebOra-WSEI/auth/startup"
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
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.EmptyFieldsErrorMsg))
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
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.PasswordNotTheSameErrorMsg))
		return
	}

	var user model.User
	if err := startup.Db.Where("email = ?", body.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.UserAlreadyExistsErrorMsg))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		fmt.Println("Problem with hashing password", err.Error())
		c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(response.InternalServerErrorMsg))
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

	res := model.ResponseUser{
		Email:   newUser.Email,
		Name:    newUser.Name,
		Surname: newUser.Surname,
		Role:    newUser.Role,
	}

	c.JSON(http.StatusCreated, gin.H{"user": res})
}
