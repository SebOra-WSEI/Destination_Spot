package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccessControl(c *gin.Context) {
	c.JSON(http.StatusOK, "access-control")
}
