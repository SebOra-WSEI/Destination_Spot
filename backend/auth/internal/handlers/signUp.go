package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, "signUp")
}
