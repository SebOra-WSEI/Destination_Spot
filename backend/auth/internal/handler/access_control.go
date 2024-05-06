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
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func AccessControl(c *gin.Context) {
	id := c.Param("id")

	t, err := token.Verify(c.GetHeader(AuthHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err.Error()))
		return
	}

	reqUserRole, ok := t.Claims.(jwt.MapClaims)["Role"]
	if !ok {
		fmt.Println("Role can not be found in claims")
		c.JSON(http.StatusBadRequest, response.CreateError(response.InternalServerErrMsg))
		return
	}

	if reqUserRole.(string) != AdminRole {
		fmt.Println("Action enabled only for admin")
		c.JSON(http.StatusBadRequest, response.CreateError(response.ActionNotPermittedErrMsg))
		return
	}

	var user model.User
	if err := database.Db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(response.UserNotFoundErrMsg))
		return
	}

	var body model.ActionControlBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBodyFields(
		body.NewPassword, body.ConfirmNewPassword,
	) {
		c.JSON(http.StatusBadRequest, response.CreateError(response.EmptyFieldsErrMsg))
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
