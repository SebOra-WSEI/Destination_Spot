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

func SignIn(c *gin.Context) {
	var body model.LoginBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBodyFields(
		body.Email, body.Password,
	) {
		c.JSON(http.StatusBadRequest, response.CreateError(response.EmptyFieldsErrMsg))
		return
	}

	var user model.User
	if err := database.Db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		fmt.Println("User not found:", err.Error())
		c.JSON(http.StatusBadRequest, response.CreateError(response.InvalidLoginOrPasswordErrMsg))
		return
	}

	if err := password.Validate(body.Password, ""); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err.Error()))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		fmt.Println("Wrong password:", err.Error())
		c.JSON(http.StatusBadRequest, response.CreateError(response.InvalidLoginOrPasswordErrMsg))
		return
	}

	jwt, err := token.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(response.ErrorWhileCreatingTokenErrMsg))
		return
	}

	res := model.LoggedUserResponse{
		Token: jwt,
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
