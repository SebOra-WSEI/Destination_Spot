package handler

import (
	"fmt"
	"github.com/SebastianOraczek/auth/database"
	"github.com/SebastianOraczek/auth/internal/model"
	"github.com/SebastianOraczek/auth/internal/password"
	"github.com/SebastianOraczek/auth/internal/request"
	"github.com/SebastianOraczek/auth/internal/response"
	"github.com/SebastianOraczek/auth/internal/token"
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
		c.JSON(http.StatusBadRequest, response.CreateError(err.Error()))
		return
	}

	var user model.User
	if err := database.Db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(response.UserNotFoundErrMsg))
		return
	}

	//claims := t.Claims.(jwt.MapClaims)
	//role, ok := claims["Role"]
	//if !ok {
	//	fmt.Println("Role can not be found in claims")
	//	c.JSON(http.StatusInternalServerError, response.CreateError(response.InternalServerErrMsg))
	//	return
	//}

	//if role != AdminRole {
	//	fmt.Println("Password must be changed by owner or admin")
	//	c.JSON(http.StatusInternalServerError, response.CreateError(response.ActionNotPermittedErrMsg))
	//	return
	//}

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 12)
	if err != nil {
		fmt.Println("Problem with hashing password", err.Error())
		c.JSON(http.StatusInternalServerError, response.CreateError(response.InternalServerErrMsg))
		return
	}

	user.Password = string(hashedPassword)

	if err := database.Db.Save(&user).Error; err != nil {
		fmt.Println("Problem saving user", err.Error())
		c.JSON(http.StatusInternalServerError, response.CreateError(response.ErrorWhileCreatingUSerErrMsg))
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
