package handler

import (
	"fmt"
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

	res := model.ResponseUser{
		ID:      user.ID,
		Email:   user.Email,
		Name:    user.Name,
		Surname: user.Surname,
		Role:    user.Role,
	}

	c.JSON(http.StatusOK, gin.H{"user": res})
}
