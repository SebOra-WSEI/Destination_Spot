package handler

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/message"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/model"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/password"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/token"
	"github.com/SebOra-WSEI/Destination_spot/shared/request"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignIn(c *gin.Context) {
	var body model.LoginBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBodyFields(
		body.Email, body.Password,
	) {
		c.JSON(http.StatusBadRequest, response.CreateError(message.ErrEmptyFields))
		return
	}

	var user model.User
	if err := user.FindByEmail(body.Email, &user); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	if err := password.Validate(body.Password, ""); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		fmt.Println("Wrong password:", err.Error())
		c.JSON(http.StatusBadRequest, response.CreateError(message.ErrInvalidLoginOrPassword))
		return
	}

	jwt, err := token.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(message.ErrWhileCreatingToken))
		return
	}

	res := model.LoggedUserResponse{
		Token: jwt,
		User:  user.GetWithNoPassword(),
	}

	c.JSON(http.StatusOK, response.Create(res))
}
