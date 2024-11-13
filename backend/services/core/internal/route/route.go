package route

import (
	"github.com/SebOra-WSEI/Destination_spot/core/internal/handler"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	// Auth service endpoints
	r.POST("/sign-in", handler.SignIn)
	r.POST("/sign-up", handler.SignUp)
	r.PUT("/reset-password/:id", handler.ResetPassword)
	r.PUT("/access-control/:id", handler.AccessControl)

	// Users endpoints
	r.GET("/users", handler.GetAllUsers)
	r.GET("/users/:id", handler.GetUser)

	// Spots endpoints
	r.GET("/spots", handler.GetAllSpots)
	r.POST("/spots", handler.CreateSpot)
	r.GET("/spots/:id", handler.GetSpot)
	r.DELETE("/spots/:id", handler.DeleteSpot)
}
