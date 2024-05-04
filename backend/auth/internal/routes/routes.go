package routes

import (
	"github.com/SebOra-WSEI/auth/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/sign-up", handlers.SignUp)
	r.POST("/sign-in", handlers.SignIn)
	r.PUT("/access-control", handlers.AccessControl)
}
