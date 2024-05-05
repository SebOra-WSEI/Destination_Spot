package handler

import (
	 "github.com/SebOra-WSEI/auth/internal/email"
	 "github.com/SebOra-WSEI/auth/internal/model"
	 "github.com/SebOra-WSEI/auth/internal/request"
	 "github.com/SebOra-WSEI/auth/internal/response"
	 "github.com/SebOra-WSEI/auth/startup"
	 "github.com/gin-gonic/gin"
	 "github.com/gin-gonic/gin/binding"
	 "net/http"
)

func SignUp(c *gin.Context) {
	 var body model.AuthBody
	 if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBody(
		  body.Email, body.Password, body.ConfirmPassword,
	 ) {
		  c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.EmptyFieldsErrorMsg))
		  return
	 }

	 if err := email.Verify(body.Email); err != nil {
		  c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
		  return
	 }

	 var user model.User
	 if err := startup.Db.Where("email = ?", body.Email).First(&user).Error; err == nil {
		  c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.UserAlreadyExistsErrorMsg))
		  return
	 }

}
