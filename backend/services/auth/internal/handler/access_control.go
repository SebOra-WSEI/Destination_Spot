package handler

import (
	"github.com/SebOra-WSEI/Destination_spot/internal/model"
	"github.com/SebOra-WSEI/Destination_spot/internal/password"
	"github.com/SebOra-WSEI/Destination_spot/internal/permission"
	"github.com/SebOra-WSEI/Destination_spot/internal/request"
	"github.com/SebOra-WSEI/Destination_spot/internal/response"
	"github.com/SebOra-WSEI/Destination_spot/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

const AuthHeader string = "Authorization"

func AccessControl(c *gin.Context) {
	id := c.Param("id")

	t, err := token.Verify(c.GetHeader(AuthHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	if code, err := permission.Admin(t.Claims.(jwt.MapClaims)); err != nil {
		c.JSON(code, response.CreateError(err))
		return
	}

	var user model.User
	if err := user.FindById(id, &user); err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(err))
		return
	}

	var body model.ActionControlBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBodyFields(
		body.NewPassword, body.ConfirmNewPassword,
	) {
		c.JSON(http.StatusBadRequest, response.CreateError(response.ErrEmptyFields))
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

	if err := user.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	res := model.UserResponse{
		Message: response.PasswordChangedMsg,
		User:    user.GetWithNoPassword(),
	}

	c.JSON(http.StatusOK, response.Create(res))
}
