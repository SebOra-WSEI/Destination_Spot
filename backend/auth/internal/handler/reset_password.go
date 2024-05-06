package handler

import (
	"fmt"
	"github.com/SebastianOraczek/auth/database"
	"github.com/SebastianOraczek/auth/internal/model"
	"github.com/SebastianOraczek/auth/internal/password"
	"github.com/SebastianOraczek/auth/internal/permission"
	"github.com/SebastianOraczek/auth/internal/request"
	"github.com/SebastianOraczek/auth/internal/response"
	"github.com/SebastianOraczek/auth/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func ResetPassword(c *gin.Context) {
	id := c.Param("id")

	t, err := token.Verify(c.GetHeader(AuthHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err.Error()))
		return
	}

	var user model.User
	if err := database.Db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(response.UserNotFoundErrMsg))
		return
	}

	if code, err := permission.User(user.Id, t.Claims.(jwt.MapClaims)); err != nil {
		c.JSON(code, response.CreateError(err.Error()))
		return
	}

	var body model.ResetPasswordBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBodyFields(
		body.CurrentPassword, body.NewPassword, body.ConfirmNewPassword,
	) {
		c.JSON(http.StatusBadRequest, response.CreateError(response.EmptyFieldsErrMsg))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.CurrentPassword)); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(response.InvalidCurrentPasswordErrMsg))
		return
	}

	if body.CurrentPassword == body.NewPassword {
		c.JSON(http.StatusBadRequest, response.CreateError(response.PasswordTheSameErrMsg))
		return
	}

	if err := password.Validate(body.NewPassword, body.ConfirmNewPassword); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err.Error()))
		return
	}

	newPassword, err := password.Generate(body.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err.Error()))
		return
	}

	user.Password = newPassword

	if err := database.Db.Save(&user).Error; err != nil {
		fmt.Println("Problem saving user", err.Error())
		c.JSON(http.StatusInternalServerError, response.CreateError(response.ErrorWhileCreatingUserErrMsg))
		return
	}

	res := model.UserResponse{
		Message: response.PasswordChangedMsg,
		User: model.NoPasswordUser{
			ID:      user.Id,
			Email:   user.Email,
			Name:    user.Name,
			Surname: user.Surname,
			Role:    user.Role,
		},
	}

	c.JSON(http.StatusOK, response.Create(res))
}
