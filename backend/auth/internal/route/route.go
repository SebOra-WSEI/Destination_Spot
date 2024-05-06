package route

import (
	"github.com/SebastianOraczek/auth/internal/handler"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.POST("/sign-up", handler.SignUp)
	r.POST("/sign-in", handler.SignIn)
	r.PUT("/reset-password/:id", handler.ResetPassword)
	r.PUT("/access-control/:id", handler.AccessControl)
}
