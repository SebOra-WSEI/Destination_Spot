package handler

import (
	"fmt"
	"github.com/SebastianOraczek/internal/model"
	"github.com/SebastianOraczek/internal/password"
	"github.com/SebastianOraczek/internal/request"
	"github.com/SebastianOraczek/internal/response"
	"github.com/SebastianOraczek/internal/token"
	"github.com/SebastianOraczek/startup"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignIn(c *gin.Context) {
	var body model.LoginBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBody(
		body.Email, body.Password,
	) {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.EmptyFieldsErrorMsg))
		return
	}

	var user model.User
	if err := startup.Db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		fmt.Println("User not found:", err.Error())
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.InvalidLoginOrPasswordErrMsg))
		return
	}

	if err := password.Validate(body.Password); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		fmt.Println("Wrong password:", err.Error())
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.InvalidLoginOrPasswordErrMsg))
		return
	}

	jwt, err := token.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(response.ErrorWhileCreatingTokenErrMsg))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwt})
}
