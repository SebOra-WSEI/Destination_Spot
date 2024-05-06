package handler

import (
	"fmt"
	"github.com/SebastianOraczek/auth/internal/model"
	"github.com/SebastianOraczek/auth/internal/password"
	"github.com/SebastianOraczek/auth/internal/request"
	"github.com/SebastianOraczek/auth/internal/response"
	"github.com/SebastianOraczek/auth/internal/token"
	"github.com/SebastianOraczek/auth/startup"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const AuthHeader = "Authorization"

func AccessControl(c *gin.Context) {
	id := c.Param("id")

	_, err := token.Verify(c.GetHeader(AuthHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
		return
	}

	var user model.User
	if err := startup.Db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(response.UserNotFoundErrMsg))
		return
	}

	var body model.ResetPasswordBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBody(
		body.CurrentPassword, body.NewPassword, body.ConfirmNewPassword,
	) {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.EmptyFieldsErrMsg))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.CurrentPassword)); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.InvalidCurrentPasswordErrMsg))
		return
	}

	if body.CurrentPassword == body.NewPassword {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.PasswordTheSameErrMsg))
		return
	}

	if err := password.Validate(body.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
		return
	}

	if body.NewPassword != body.ConfirmNewPassword {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.PasswordNotTheSameErrMsg))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 12)
	if err != nil {
		fmt.Println("Problem with hashing password", err.Error())
		c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(response.InternalServerErrMsg))
		return
	}

	user.Password = string(hashedPassword)
	startup.Db.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": response.PasswordChangedMsg})
}
