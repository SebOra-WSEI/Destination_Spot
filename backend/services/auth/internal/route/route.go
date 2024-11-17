package route

import (
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/handler/accessControl"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/handler/login"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/handler/register"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/handler/resetPassword"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.POST("/sign-up", register.SignUp)
	r.POST("/sign-in", login.SignIn)
	r.PUT("/reset-password/:id", resetPassword.ResetPassword)
	r.PUT("/access-control/:id", accessControl.GetAccessControl)
}
