package handler

import (
	"github.com/SebastianOraczek/auth/internal/response"
	"github.com/SebastianOraczek/auth/internal/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

const AuthHeader = "Authorization"

func AccessControl(c *gin.Context) {
	t, err := token.Verify(c.GetHeader(AuthHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(err.Error()))
		return
	}

	c.JSON(200, t)
}
