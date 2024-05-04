package handler

import (
	 "github.com/SebOra-WSEI/auth/internal/model"
	 "github.com/SebOra-WSEI/auth/internal/response"
	 "github.com/SebOra-WSEI/auth/start"
	 "github.com/gin-gonic/gin"
	 "github.com/gin-gonic/gin/binding"
	 "net/http"
)

func SignUp(c *gin.Context) {
	 var body model.AuthBody
	 if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		  c.JSON(http.StatusInternalServerError, response.CreateErrorResponse(err.Error()))
		  return
	 }

	 var user model.User
	 if err := start.Db.Where("name = ?", body.Name).First(&user).Error; err != nil {
		  c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
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
