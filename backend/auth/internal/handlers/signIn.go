package handlers

import (
	 "github.com/gin-gonic/gin"
	 "net/http"
)

func SignIn(c *gin.Context) {
	 c.JSON(http.StatusOK, "signIn")
}
