package handler

import (
	"github.com/SebOra-WSEI/Destination_spot/auth/database"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/password"
	"github.com/SebOra-WSEI/Destination_spot/shared/model"
	"github.com/SebOra-WSEI/Destination_spot/shared/permission"
	"github.com/SebOra-WSEI/Destination_spot/shared/request"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/SebOra-WSEI/Destination_spot/shared/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type ResetPasswordBody struct {
	CurrentPassword    string `json:"currentPassword"`
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}

func ResetPassword(c *gin.Context) {
	id := c.Param("id")
	t, err := token.Verify(c.GetHeader(AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var user model.User
	if err := user.FindById(database.Db, id, &user); err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(response.ErrUserNotFound))
		return
	}

	if code, err := permission.User(database.Db, user.ID, t.Claims.(jwt.MapClaims)); err != nil {
		c.JSON(code, response.CreateError(err))
		return
	}

	var body ResetPasswordBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBodyFields(
		body.CurrentPassword, body.NewPassword, body.ConfirmNewPassword,
	) {
		c.JSON(http.StatusBadRequest, response.CreateError(response.ErrEmptyFields))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.CurrentPassword)); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(response.ErrInvalidCurrentPassword))
		return
	}

	if body.CurrentPassword == body.NewPassword {
		c.JSON(http.StatusBadRequest, response.CreateError(response.ErrPasswordTheSame))
		return
	}

	if err := password.Validate(body.NewPassword, body.ConfirmNewPassword); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	newPassword, err := password.Generate(body.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	user.Password = newPassword

	if err := user.Update(database.Db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	res := model.UserResponse{
		Message: response.PasswordChangedMsg,
		User:    user.GetWithNoPassword(),
	}

	c.JSON(http.StatusOK, response.Create(res))
}
