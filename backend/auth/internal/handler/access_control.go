package handler

import (
	"github.com/SebOra-WSEI/auth/internal/response"
	"github.com/SebOra-WSEI/auth/internal/token"
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

	if !t.Valid {
		c.JSON(http.StatusBadRequest, response.CreateErrorResponse(response.TokenExpiredErrMsg))
		return
	}

	c.JSON(200, "abc")
}
